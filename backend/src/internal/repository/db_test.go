package repository

import (
	"os"
	"testing"

	"llmapi/src/internal/model"
	"llmapi/src/pkg/logger"

	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

func GetPostgresDB() (*gorm.DB, error) {
	logger.SetLevelString("debug")
	return CreateDB(&Options{
		DSN:        "postgres://postgres:postgres@localhost:15432/default",
		DBLogLevel: 4,
	})
}

func TestCreateSqliteDB(t *testing.T) {
	db, err := CreateDB(&Options{
		DSN:        "sqlite://llmapi.db",
		DBLogLevel: 1,
	})
	assert.NoError(t, err, "expected no error while creating database")
	assert.NotNil(t, db, "expected a valid database connection, got nil")
}

func TestCreatePostgresDB(t *testing.T) {
	db, err := GetPostgresDB()
	assert.NoError(t, err, "expected no error while creating database")
	assert.NotNil(t, db, "expected a valid database connection, got nil")
}

func TestCreateMysqlDB(t *testing.T) {
	db, err := CreateDB(&Options{
		DSN:        "mysql://root:llmapi@tcp(localhost:13306)/llmapi?charset=utf8mb4&parseTime=True&loc=Local",
		DBLogLevel: 1,
	})
	assert.NoError(t, err, "expected no error while creating database")
	assert.NotNil(t, db, "expected a valid database connection, got nil")
}

func TestCreateSqlServerDB(t *testing.T) {
	dsn := os.Getenv("SQL_SERVER_URL")
	assert.NotEmpty(t, dsn, "expected non-empty SQL_SERVER_URL environment variable")
	db, err := CreateDB(&Options{
		DSN:        dsn,
		DBLogLevel: 2,
	})
	assert.NoError(t, err, "expected no error while creating database")
	assert.NotNil(t, db, "expected a valid database connection, got nil")
}

func TestModel(t *testing.T) {
	db, err := GetPostgresDB()
	assert.NoError(t, err, "expected no error while creating database")
	assert.NotNil(t, db, "expected a valid database connection, got nil")

	// Drop tables to ensure fresh migration with correct constraints
	// if db.Migrator().HasTable(&model.Model{}) {
	// 	db.Migrator().DropTable(&model.Model{})
	// }
	// if db.Migrator().HasTable(&model.ModelChannel{}) {
	// 	db.Migrator().DropTable(&model.ModelChannel{})
	// }
	// if db.Migrator().HasTable(&model.ModelAlias{}) {
	// 	db.Migrator().DropTable(&model.ModelAlias{})
	// }

	// Migrate ModelAlias first as Model might depend on it for FK creation via Model.Alias referencing ModelAlias.Alias
	err = db.AutoMigrate(&model.ModelAlias{})
	assert.NoError(t, err, "expected no error while auto-migrating ModelAlias")

	// Then migrate other models that might depend on ModelAlias
	err = db.AutoMigrate(&model.Model{}, &model.ModelChannel{})
	assert.NoError(t, err, "expected no error while auto-migrating Model and ModelChannel")

	tx := db.Begin()
	defer tx.Rollback()

	modelChannel := &model.ModelChannel{
		Name: "[Test] Test Channel",
	}

	tx.Create(modelChannel)
	assert.NoError(t, tx.Error, "expected no error while creating model channel")
	assert.NotEqual(t, modelChannel.ID, 0, "expected model channel ID to be set")
	assert.Equal(t, len(modelChannel.Models), 0, "expected model channel to have models")

	var modelChannelCount int64
	tx.Model(&model.ModelChannel{}).Where("id = ?", modelChannel.ID).Count(&modelChannelCount)
	assert.Equal(t, int64(1), modelChannelCount, "expected model channel count to be 1")

	testModel := &model.Model{
		ModelName:      "[Test] Test Model",
		ModelChannelID: modelChannel.ID,
	}

	tx.Create(testModel)
	assert.NoError(t, tx.Error, "expected no error while creating model")
	assert.NotEmpty(t, testModel.ID, "expected model ID to be set")
	assert.Equal(t, len(modelChannel.Models), 0, "expected model channel to have models")

	// preload modes to model channel
	tx.Preload("Models").First(modelChannel)
	assert.NoError(t, tx.Error, "expected no error while preloading models")
	assert.Equal(t, len(modelChannel.Models), 1, "expected model channel to have 1 model")

	var modelCount int64
	tx.Model(&model.Model{}).Where("id = ?", testModel.ID).Count(&modelCount)
	assert.Equal(t, int64(1), modelCount, "expected model count to be 1")

	tx.Delete(modelChannel)
	assert.NoError(t, tx.Error, "expected no error while deleting model channel")
	var modelChannelCountAfterDelete int64
	tx.Model(&model.ModelChannel{}).Where("id = ?", modelChannel.ID).Count(&modelChannelCountAfterDelete)
	assert.Equal(t, int64(0), modelChannelCountAfterDelete, "expected model channel count to be 0 after delete")

	var modelCountAfterDelete int64
	tx.Model(&model.Model{}).Where("model_channel_id = ?", modelChannel.ID).Count(&modelCountAfterDelete)
	assert.Equal(t, int64(0), modelCountAfterDelete, "expected model count to be 0 after delete")

	var modelAliasCount int64
	tx.Model(&model.ModelAlias{}).Where("alias = ?", testModel.Alias).Count(&modelAliasCount)
	assert.Equal(t, int64(0), modelAliasCount, "expected model alias count to be 0 after delete")

	// Test ModelChannel2
	modelChannel2 := &model.ModelChannel{
		Name: "[Test] Test Channel 2",
		Models: []model.Model{
			{ModelName: "[Test] Test Model 2"},
			{ModelName: "[Test] Test Model 3"},
		},
	}

	tx.Create(modelChannel2)
	// tx.Save(modelChannel2)
	assert.NoError(t, tx.Error, "expected no error while creating model channel 2")
	assert.NotEqual(t, tx.RowsAffected, 0, "expected model channel 2 to be created")
	assert.NotEqual(t, modelChannel2.ID, 0, "expected model channel 2 ID to be set")
	assert.Equal(t, len(modelChannel2.Models), 2, "expected model channel 2 to have 2 models")

	var modelChannel2Count int64
	tx.Model(&model.ModelChannel{}).Where("id = ?", modelChannel2.ID).Count(&modelChannel2Count)
	assert.Equal(t, int64(1), modelChannel2Count, "expected model channel 2 count to be 1")
	var modelCount2 int64
	tx.Model(&model.Model{}).Where("model_channel_id = ?", modelChannel2.ID).Count(&modelCount2)
	assert.Equal(t, int64(2), modelCount2, "expected model count to be 2")

	// Check Update Model Alias
	testModel = &modelChannel2.Models[0]
	testModelOldAlias := testModel.Alias
	// tx.Model(&model.Model{}).Save(testModel)
	tx.Model(testModel).Update("alias", "[Test] Test Model 2 Updated")
	assert.NoError(t, tx.Error, "expected no error while updating model alias")

	var modelAliasCount2 int64
	tx.Model(&model.ModelAlias{}).Where("alias = ?", testModel.Alias).Count(&modelAliasCount2)
	assert.Equal(t, int64(1), modelAliasCount2, "expected model alias count to be 1 after update")
	var modelAliasCountOld int64
	tx.Model(&model.ModelAlias{}).Where("alias = ?", testModelOldAlias).Count(&modelAliasCountOld)
	assert.Equal(t, int64(0), modelAliasCountOld, "expected model alias count to be 0 after update")

	oldAlias2 := testModel.Alias
	testModel.Alias = "[Test] Test Model 2 Updated Again"
	tx.Model(testModel).Updates(testModel)
	assert.NoError(t, tx.Error, "expected no error while updating model alias again")
	var modelAliasCount2AfterUpdate int64
	tx.Model(&model.ModelAlias{}).Where("alias = ?", testModel.Alias).Count(&modelAliasCount2AfterUpdate)
	assert.Equal(t, int64(1), modelAliasCount2AfterUpdate, "expected model alias count to be 1 after update again")
	var modelAliasCountOld2 int64
	tx.Model(&model.ModelAlias{}).Where("alias = ?", oldAlias2).Count(&modelAliasCountOld2)
	assert.Equal(t, int64(0), modelAliasCountOld2, "expected model alias count to be 0 after update again")

	oldAlias3 := testModel.Alias
	testModel.Alias = "[Test] Test Model 2 Updated Again Again"
	tx.Save(testModel)
	assert.NoError(t, tx.Error, "expected no error while saving model alias again again")
	var modelAliasCount3AfterSave int64
	tx.Model(&model.ModelAlias{}).Where("alias = ?", testModel.Alias).Count(&modelAliasCount3AfterSave)
	assert.Equal(t, int64(1), modelAliasCount3AfterSave, "expected model alias count to be 1 after save again")
	var modelAliasCountOld3 int64
	tx.Model(&model.ModelAlias{}).Where("alias = ?", oldAlias3).Count(&modelAliasCountOld3)
	assert.Equal(t, int64(0), modelAliasCountOld3, "expected model alias count to be 0 after save again")

	oldAlias4 := testModel.Alias
	testModel.Alias = "[Test] Test Model 2 Updated Again Again Again"
	tx.Model(&model.Model{}).Where("id = ?", testModel.ID).Updates(testModel)
	assert.NoError(t, tx.Error, "expected no error while updating model alias again again")
	var modelAliasCount4AfterUpdate int64
	tx.Model(&model.ModelAlias{}).Where("alias = ?", testModel.Alias).Count(&modelAliasCount4AfterUpdate)
	assert.Equal(t, int64(1), modelAliasCount4AfterUpdate, "expected model alias count to be 1 after update again again")
	var modelAliasCountOld4 int64
	tx.Model(&model.ModelAlias{}).Where("alias = ?", oldAlias4).Count(&modelAliasCountOld4)
	assert.Equal(t, int64(0), modelAliasCountOld4, "expected model alias count to be 0 after update again again")

	tx.Delete(modelChannel2)
	assert.NoError(t, tx.Error, "expected no error while deleting model channel 2")
	var modelChannel2CountAfterDelete int64
	tx.Model(&model.ModelChannel{}).Where("id = ?", modelChannel2.ID).Count(&modelChannel2CountAfterDelete)
	assert.Equal(t, int64(0), modelChannel2CountAfterDelete, "expected model channel 2 count to be 0 after delete")
}

func TestModelRouting(t *testing.T) {
	db, err := GetPostgresDB()
	assert.NoError(t, err, "expected no error while creating database")
	assert.NotNil(t, db, "expected a valid database connection, got nil")

	tx := db.Begin()
	defer tx.Rollback()

	// tx.Migrator().DropTable(&model.ModelRouting{})
	// tx.Migrator().DropTable(&model.ModelRoutingTarget{})
	// tx.Migrator().DropTable(&model.ModelEndpoint{})
	// tx.Migrator().DropTable(&model.ModelProvider{})
	// tx.Migrator().DropTable(&model.Channel{})

	err = tx.AutoMigrate(
		&model.Channel{},
		&model.ModelProvider{},
	)
	assert.NoError(t, err, "expected no error while auto-migrating Channel and ModelProvider")

	err = tx.AutoMigrate(
		&model.ModelRouting{},
	)
	assert.NoError(t, err, "expected no error while auto-migrating ModelRouting")

	err = tx.AutoMigrate(
		&model.ModelRoutingTarget{})
	assert.NoError(t, err, "expected no error while auto-migrating ModelRoutingTarget")

	err = tx.AutoMigrate(
		&model.ModelEndpoint{},
	)
	assert.NoError(t, err, "expected no error while auto-migrating ModelEndpoint")

	channelA := &model.Channel{
		Name: "[Test] Test Channel A",
	}
	assert.NoError(t, tx.Create(channelA).Error, "expected no error while creating channel A")

	modelProviderA := &model.ModelProvider{
		Name:    "[Test] Test Model Provider A",
		Channel: *channelA,
	}
	assert.NoError(t, tx.Create(modelProviderA).Error, "expected no error while creating model provider A")
	assert.NotEmpty(t, modelProviderA.ChannelID, "expected ChannelID to be set")

	modelEndpointA := &model.ModelEndpoint{
		Name:         "[Test] Test Model Endpoint A",
		ModelRouting: &model.ModelRouting{Strategy: 1},
	}
	assert.NoError(t, tx.Create(modelEndpointA).Error, "expected no error while creating model endpoint A")
	assert.NotEmpty(t, modelEndpointA.ID, "expected ModelEndpoint ID to be set")
	assert.NotEmpty(t, modelEndpointA.ModelRouting.ID, "expected ModelRouting ID to be set")

	modelRoutingTargetA := &model.ModelRoutingTarget{
		ModelRoutingID: modelEndpointA.ModelRouting.ID,
	}
	assert.NoError(t, tx.Create(modelRoutingTargetA).Error, "expected no error while creating model routing target A")
	assert.NotEmpty(t, modelRoutingTargetA.ModelRoutingID, "expected ModelRoutingID to be set")

	assert.Empty(t, modelEndpointA.ModelRouting.ModelRoutingTargets, "expected ModelRouting ID to be empty before saving")
	tx.Model(&modelEndpointA.ModelRouting).Preload("ModelRoutingTargets").First(&modelEndpointA.ModelRouting)
	assert.NoError(t, tx.Error, "expected no error while preloading ModelRoutingTargets")
	assert.NotEmpty(t, modelEndpointA.ModelRouting.ModelRoutingTargets, "expected ModelRouting ID to be set after saving")

	modelRoutingTargetA2 := model.ModelRoutingTarget{
		Weight: 1,
	}
	modelEndpointA.ModelRouting.ModelRoutingTargets = append(modelEndpointA.ModelRouting.ModelRoutingTargets, &modelRoutingTargetA2)
	assert.NoError(t, tx.Save(modelEndpointA).Error, "expected no error while saving model endpoint A")
	assert.NotEmpty(t, modelRoutingTargetA2.ModelRoutingID, "expected ModelRoutingID to be set after saving")

	var routingTargetCount int64
	assert.NoError(t, tx.Model(&model.ModelRoutingTarget{}).Count(&routingTargetCount).Error, "expected count to be retrieved without error")
	assert.Equal(t, int64(2), routingTargetCount, "expected ModelRoutingTarget count to be 2 after saving")

	assert.Empty(t, modelRoutingTargetA.TargetID, "expected TargetID to be empty after saving")
	modelEndpointA.ModelRoutingTargets = append(modelEndpointA.ModelRoutingTargets, modelRoutingTargetA)
	assert.NoError(t, tx.Save(modelEndpointA).Error, "expected no error while saving model endpoint A")
	assert.NotEmpty(t, modelRoutingTargetA.TargetID, "expected TargetID to be set after saving")
	assert.Equal(t, modelRoutingTargetA.TargetID, modelEndpointA.ID, "expected TargetID to match ModelEndpoint ID")

	modelEndpoint2 := &model.ModelEndpoint{
		Name:         "[Test] Test Model Endpoint 2",
		ModelRouting: &model.ModelRouting{Strategy: 2},
	}
	assert.NoError(t, tx.Create(modelEndpoint2).Error, "expected no error while creating model endpoint 2")
	assert.NotEmpty(t, modelEndpoint2.ID, "expected ModelEndpoint ID to be set")
	assert.NotEmpty(t, modelEndpoint2.ModelRouting.ID, "expected ModelRouting ID to be set")

	modelRoutingTarget2 := &model.ModelRoutingTarget{
		ModelRoutingID: modelEndpoint2.ModelRouting.ID,
		Weight:         1,
		Priority:       1,
	}
	assert.NoError(t, tx.Create(modelRoutingTarget2).Error, "expected no error while creating model routing target 2")
	assert.NotEmpty(t, modelRoutingTarget2.ID, "expected ModelRoutingTarget ID to be set")

	modelProviderA.ModelRoutingTargets = append(modelProviderA.ModelRoutingTargets, modelRoutingTarget2)
	assert.NoError(t, tx.Model(modelProviderA).Updates(modelProviderA).Error, "expected no error while updating model provider A")
	assert.NotEmpty(t, modelRoutingTarget2.TargetID, "expected TargetID to be set after saving")
	var routingTargetCount2 int64
	assert.NoError(t, tx.Model(&model.ModelRoutingTarget{}).Count(&routingTargetCount2).Error, "expected count to be retrieved without error")
	assert.Equal(t, int64(3), routingTargetCount2, "expected ModelRoutingTarget count to be 3 after saving")

	switch modelRoutingTargetA.TargetType {
	case (model.ModelEndpoint{}).TableName():
		var modelEndpoint model.ModelEndpoint
		tx.Table(modelRoutingTargetA.TargetType).Where("id = ?", modelRoutingTargetA.TargetID).First(&modelEndpoint)
		modelRoutingTargetA.Target = &modelEndpoint
		assert.NoError(t, tx.Error, "expected no error while retrieving model endpoint")
		assert.Equal(t, modelEndpoint.ID, modelRoutingTargetA.TargetID, "expected ModelEndpoint ID to match")
	case (model.ModelProvider{}).TableName():
		var modelProvider model.ModelProvider
		tx.Table(modelRoutingTargetA.TargetType).Where("id = ?", modelRoutingTargetA.TargetID).First(&modelProvider)
		modelRoutingTargetA.Target = &modelProvider
		assert.NoError(t, tx.Error, "expected no error while retrieving model provider")
		assert.Equal(t, modelProvider.ID, modelRoutingTargetA.TargetID, "expected ModelProvider ID to match")
	default:
		t.Errorf("unexpected target type: %s", modelRoutingTargetA.TargetType)
	}
}
