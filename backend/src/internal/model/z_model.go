package model

import (
	"fmt"

	"gorm.io/datatypes"
	"gorm.io/gorm"
)

const oldAliasKey = "gorm:old_alias"

type Model struct {
	gorm.Model
	ModelName      string `gorm:"not null;index"`
	Alias          string `gorm:"not null;index"`
	Enabled        bool   `gorm:"default:true"`
	Status         uint8  `gorm:"default:0"`
	ExtraParams    datatypes.JSONType[any]
	ModelChannelID uint         `gorm:"index"`
	ModelChannel   ModelChannel `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

func (Model) TableName() string {
	return "models"
}

func (m *Model) BeforeCreate(tx *gorm.DB) (err error) {
	if m.ModelName == "" {
		return fmt.Errorf("model name cannot be empty")
	}
	if m.Alias == "" {
		m.Alias = m.ModelName
	}
	// Ensure ModelAlias exists before creating Model
	if err := tx.Model(&ModelAlias{}).Where("alias = ?", m.Alias).FirstOrCreate(&ModelAlias{Alias: m.Alias}).Error; err != nil {
		return fmt.Errorf("failed to ensure ModelAlias for Alias '%s' before creating Model: %w", m.Alias, err)
	}
	return nil
}

func (m *Model) AfterSave(tx *gorm.DB) (err error) {
	// This hook is now primarily for handling alias changes during updates.
	// Creation of ModelAlias is handled in BeforeCreate.

	if oldAliasValue, ok := tx.Statement.InstanceGet(oldAliasKey); ok {
		oldAlias := oldAliasValue.(string)
		// Only proceed if oldAlias is different from current m.Alias,
		// indicating an update where the alias actually changed.
		if oldAlias != m.Alias {
			// Check if the new alias exists, if not, create it.
			// This handles cases where an alias is updated to a brand new one.
			// if err := tx.FirstOrCreate(&ModelAlias{Alias: m.Alias}).Error; err != nil {
			// 	return fmt.Errorf("failed to ensure ModelAlias for new Alias '%s' after update: %w", m.Alias, err)
			// }

			// Check if the old alias is still in use by other models
			var count int64
			// Ensure we are looking for models other than the current one if it's an update
			if err = tx.Model(&Model{}).Where("alias = ? AND id != ?", oldAlias, m.ID).Count(&count).Error; err != nil {
				return fmt.Errorf("failed to count models for old alias '%s': %w", oldAlias, err)
			}

			if count == 0 {
				// If no other Model records use the old Alias, delete the old ModelAlias record
				if err = tx.Where("alias = ?", oldAlias).Delete(&ModelAlias{}).Error; err != nil {
					// Log or handle error, but don't necessarily fail the whole transaction
					// if the primary operation (model save) was successful.
					// Consider logging this as a warning.
					tx.Logger.Warn(tx.Statement.Context, "failed to delete unused old ModelAlias for '%s': %v", oldAlias, err)
				}
			}
		}
	}
	return nil
}

func (m *Model) BeforeDelete(tx *gorm.DB) (err error) {
	if m.ID == 0 || m.Alias == "" {
		return nil
	}
	var count int64
	// 检查是否还有 Model 记录使用这个 Alias
	// 注意：此时当前 Model (m.ID) 的记录已被删除（或在删除事务中即将完成删除）
	if err = tx.Model(&Model{}).Where(&Model{Alias: m.Alias}).Where("id != ?", m.ID).Count(&count).Error; err != nil {
		return fmt.Errorf("failed to count models for alias '%s': %w", m.Alias, err)
	}
	if count == 0 {
		if err = tx.Where("alias = ?", m.Alias).Delete(&ModelAlias{}).Error; err != nil {
			return fmt.Errorf("failed to delete ModelAlias for '%s': %w", m.Alias, err)
		}
	}
	return nil
}

func (m *Model) checkAlias(tx *gorm.DB) (err error) {
	// 仅当在 Model 中能找到 id 时才能生效
	// 如果通过以下方式：tx.Model(&model.Model{}).Where("id = ?", testModel.ID).Update("alias", testModel.Alias)
	// 可能会导致无法找到 id
	var oldAlias string
	var destAlias string

	id := m.ID

	if destMap, ok := tx.Statement.Dest.(*Model); ok {
		destAlias = destMap.Alias
		if id == 0 && destMap.ID != 0 {
			id = destMap.ID
		}
	} else if destMap, ok := tx.Statement.Dest.(map[string]interface{}); ok {
		if alias, ok := destMap["alias"]; ok {
			destAlias = alias.(string)
		}
	} else {
		return fmt.Errorf("unsupported destination type: %T", tx.Statement.Dest)
	}

	if destAlias == "" {
		return fmt.Errorf("alias cannot be empty")
	}

	if m.Alias != "" && m.Alias != destAlias {
		oldAlias = m.Alias
	} else {
		// 当前可能无法判断是否没有改变，用户提供用于查询的m对象可能只包含的ID,需要查询数据库获取旧值
		if err = tx.Model(&Model{}).Where("id = ?", id).Select("alias").Scan(&oldAlias).Error; err != nil {
			return fmt.Errorf("failed to get old alias for model ID %d: %w", id, err)
		}
	}

	if oldAlias != destAlias {
		if err := tx.Model(&ModelAlias{}).Where("alias = ?", destAlias).FirstOrCreate(&ModelAlias{Alias: destAlias}).Error; err != nil {
			return fmt.Errorf("failed to ensure ModelAlias for new Alias '%s' after update: %w", destAlias, err)
		}

		tx.Statement.InstanceSet(oldAliasKey, oldAlias)
	}
	return nil
}

// Hook for Model: BeforeUpdate
// 当 Model 的 Alias 字段发生变更时，在更新前捕获其旧的 Alias 值。
func (m *Model) BeforeUpdate(tx *gorm.DB) (err error) {
	if err := m.checkAlias(tx); err != nil {
		return err
	}
	return nil
}
