package repository

import (
	"llmapi/src/internal/model"

	"gorm.io/gorm"
)

type ModelRepo interface {
	CreateModelChannel(modelChannel *model.ModelChannel) error
	CreateModel(model *model.Model) error
	UpdateModelAlias(modelAlias *model.ModelAlias, updates map[string]any) error
	DeleteModelChannel(modelChannel *model.ModelChannel) error
	DeleteModel(model *model.Model) error
}

type modelRepo struct {
	db *gorm.DB
}

func NewModelRepo(db *gorm.DB) ModelRepo {
	return &modelRepo{
		db: db,
	}
}

func (r *modelRepo) CreateModelChannel(modelChannel *model.ModelChannel) error {
	return r.db.Create(modelChannel).Error
}

func (r *modelRepo) CreateModel(model *model.Model) error {
	return r.db.Create(model).Error
}

func (r *modelRepo) UpdateModelAlias(modelAlias *model.ModelAlias, updates map[string]any) error {
	return r.db.Model(modelAlias).Updates(updates).Error
}

func (r *modelRepo) DeleteModelChannel(modelChannel *model.ModelChannel) error {
	return r.db.Delete(modelChannel).Error
}

func (r *modelRepo) DeleteModel(model *model.Model) error {
	return r.db.Delete(model).Error
}
