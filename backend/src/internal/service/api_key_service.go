package service

import (
	"fmt"
	"strings"
	"time"

	"llmapi/src/internal/model"
	"llmapi/src/internal/repository"
	"llmapi/src/internal/utils"
)

const (
	DefaultAPIKeyPrefix     = "sk-llm-v1-"
	SecretPartLengthBytes   = 32 // 32 bytes = 256 bits of entropy
	Argon2DefaultTime       = 1
	Argon2DefaultMemory     = 64 * 1024 // 64MB
	Argon2DefaultThreads    = 4
	Argon2DefaultKeyLength  = 32
	Argon2DefaultSaltLength = 16
	APIKeyLogTableName      = "api_key_usage_logs" // 假设的日志表名
)

var (
	ErrorInvalidTokenFormat       = fmt.Errorf("invalid token format")
	ErrorFailedToDecodeSecretPart = fmt.Errorf("failed to decode secret part")
	ErrorFailedToGetAPIKeyRecord  = fmt.Errorf("failed to get API key record")
	ErrorFailedToDecodeSalt       = fmt.Errorf("failed to decode salt")
	ErrorInvalidAPIKey            = fmt.Errorf("invalid API key")
)

type APIKeyService interface {
	CreateAPIKey(user *model.User, name string, scopes int64, expire int64) (string, *model.APIKeyRecord, error)
	GetAPIKeys(user *model.User) ([]*model.APIKeyRecord, error)
	GetAPIKeyRecordByToken(token string) (*model.APIKeyRecord, error)
	ValidateAPIKey(token string) (*model.APIKeyRecord, *model.User, error)
	GetAPIKeyRecordByLookupKey(lookupKey string) (*model.APIKeyRecord, error)
	DeleteAPIKeyRecordByLookupKey(lookupKey string) error
	DeleteAPIKeyRecord(token string) error
	UpdateAPIKeyRecord(apiKey *model.APIKeyRecord) error
}

type apiKeyService struct {
	userRepo   repository.UserRepo
	apiKeyRepo repository.APIKeyRepo
}

func NewAPIKeyService(userRepo repository.UserRepo, apiKeyRepo repository.APIKeyRepo) *apiKeyService {
	return &apiKeyService{
		userRepo:   userRepo,
		apiKeyRepo: apiKeyRepo,
	}
}

func (s *apiKeyService) CreateAPIKey(user *model.User, name string, scopes int64, expire int64) (string, *model.APIKeyRecord, error) {
	secretRawKey, err := utils.RandBase64(SecretPartLengthBytes)
	if err != nil {
		return "", nil, err
	}
	prefix := DefaultAPIKeyPrefix

	fullAPIKey := prefix + secretRawKey

	rawSalt, encodedSalt, err := utils.RandSaltWithBase64(Argon2DefaultSaltLength)
	if err != nil {
		return "", nil, err
	}

	encodeKeyHash := utils.EncodeArgon2([]byte(secretRawKey), rawSalt, Argon2DefaultTime, Argon2DefaultMemory, Argon2DefaultThreads, Argon2DefaultKeyLength)

	// Used for querying
	lookupKey := utils.EncodeSha256([]byte(secretRawKey))

	apiKey := &model.APIKeyRecord{
		UserID:      user.UserID,
		Name:        name,
		Prefix:      prefix,
		LookupKey:   lookupKey,
		SecretHash:  encodeKeyHash,
		Salt:        encodedSalt,
		SecretBrief: utils.KeyBrief(fullAPIKey, prefix),
		Scopes:      scopes,
		ExpiresAt:   nil,
	}

	if expire > 0 {
		expireTime := time.Now().Add(time.Duration(expire) * time.Second)
		apiKey.ExpiresAt = &expireTime
	}

	if err := s.apiKeyRepo.Create(apiKey); err != nil {
		return "", nil, err
	}

	return fullAPIKey, apiKey, nil
}

func (s *apiKeyService) GetAPIKeys(user *model.User) ([]*model.APIKeyRecord, error) {
	apiKeys, err := s.apiKeyRepo.GetByUserID(user.UserID)
	if err != nil {
		return nil, err
	}
	return apiKeys, nil
}

func (s *apiKeyService) GetAPIKeyRecordByToken(token string) (*model.APIKeyRecord, error) {
	if !strings.HasPrefix(token, DefaultAPIKeyPrefix) {
		return nil, ErrorInvalidTokenFormat
	}
	secretPart := strings.TrimPrefix(token, DefaultAPIKeyPrefix)

	lookupKey := utils.EncodeSha256([]byte(secretPart))
	apiKeyRecord, err := s.apiKeyRepo.GetByLookupHash(lookupKey)
	if err != nil {
		return nil, ErrorFailedToGetAPIKeyRecord
	}

	decodedSalt, err := utils.DecodeBase64(apiKeyRecord.Salt)
	if err != nil {
		return nil, ErrorFailedToDecodeSalt
	}

	encodedKeyHash := utils.EncodeArgon2([]byte(secretPart), decodedSalt, Argon2DefaultTime, Argon2DefaultMemory, Argon2DefaultThreads, Argon2DefaultKeyLength)
	if encodedKeyHash != apiKeyRecord.SecretHash {
		return nil, ErrorInvalidAPIKey
	}

	return apiKeyRecord, nil
}

func (s *apiKeyService) DeleteAPIKeyRecord(token string) error {
	apiKey, err := s.GetAPIKeyRecordByToken(token)
	if err != nil {
		return err
	}
	return s.apiKeyRepo.Delete(apiKey)
}

func (s *apiKeyService) GetAPIKeyRecordByLookupKey(lookupKey string) (*model.APIKeyRecord, error) {
	return s.apiKeyRepo.GetByLookupHash(lookupKey)
}

func (s *apiKeyService) DeleteAPIKeyRecordByLookupKey(lookupKey string) error {
	return s.apiKeyRepo.DeleteByLookupKey(lookupKey)
}

func (s *apiKeyService) ValidateAPIKey(token string) (*model.APIKeyRecord, *model.User, error) {
	apiKeyRecord, err := s.GetAPIKeyRecordByToken(token)
	if err != nil {
		return nil, nil, err
	}

	user, err := s.userRepo.GetUserByUserID(apiKeyRecord.UserID)
	if err != nil {
		return nil, nil, err
	}

	if user == nil {
		return nil, nil, fmt.Errorf("user not found")
	}

	return apiKeyRecord, user, nil
}

func (s *apiKeyService) UpdateAPIKeyRecord(apiKey *model.APIKeyRecord) error {
	return s.apiKeyRepo.Update(apiKey)
}