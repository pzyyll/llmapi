package repository

import (
	"llmapi/src/internal/model"

	"gorm.io/gorm"
)

type APIKeyRepo interface {
	Create(key *model.APIKeyRecord) error
	GetByLookupHash(lookupHash string) (*model.APIKeyRecord, error)
	GetByUserID(userID int64) ([]*model.APIKeyRecord, error)
	Update(key *model.APIKeyRecord) error
	Delete(key *model.APIKeyRecord) error
	DeleteByLookupKey(lookupKey string) error
}

type apiKeyRepo struct {
	db *gorm.DB
}

func NewAPIKeyRepo(db *gorm.DB) APIKeyRepo {
	return &apiKeyRepo{
		db: db,
	}
}

func (r *apiKeyRepo) Create(key *model.APIKeyRecord) error {
	return r.db.Create(key).Error
}

func (r *apiKeyRepo) GetByLookupHash(lookupHash string) (*model.APIKeyRecord, error) {
	var key model.APIKeyRecord
	if err := r.db.Where(&model.APIKeyRecord{LookupKey: lookupHash}).First(&key).Error; err != nil {
		return nil, err
	}
	return &key, nil
}

func (r *apiKeyRepo) Update(key *model.APIKeyRecord) error {
	return r.db.Model(key).Select("Name", "Scopes", "ExpiresAt").Updates(key).Error
}

func (r *apiKeyRepo) Delete(key *model.APIKeyRecord) error {
	return r.db.Delete(key).Error
}

func (r *apiKeyRepo) GetByUserID(userID int64) ([]*model.APIKeyRecord, error) {
	var keys []*model.APIKeyRecord
	if err := r.db.Where(&model.APIKeyRecord{UserID: userID}).Find(&keys).Error; err != nil {
		return nil, err
	}
	return keys, nil
}

func (r *apiKeyRepo) DeleteByLookupKey(lookupKey string) error {
	var key model.APIKeyRecord
	if err := r.db.Where(&model.APIKeyRecord{LookupKey: lookupKey}).First(&key).Error; err != nil {
		return err
	}
	return r.db.Delete(&key).Error
}
