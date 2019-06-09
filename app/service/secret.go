package service

import (
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"secret-server/app/crerrors"
	"secret-server/app/models"
	"secret-server/app/repository"
	"secret-server/app/utils"
	"time"
)

type ISecretService interface {
	CreateSecret(ctx *gin.Context, secret *models.SecretRequest) (*models.Secret, crerrors.IError)
	GetSecret(ctx *gin.Context, id string) (*models.Secret, crerrors.IError)
}

type SecretService struct {
	repo repository.ISecretsRepo
}

func NewSecretService(repo repository.ISecretsRepo) *SecretService {
	return &SecretService{repo: repo}
}

func (s *SecretService) CreateSecret(ctx *gin.Context, secretReq *models.SecretRequest) (*models.Secret, crerrors.IError) {
	now := time.Now()
	secret := &models.Secret{
		ID:         utils.GenerateUUID(),
		SecretText: secretReq.SecretText,
		CreatedAt:  now,
		ExpiresAt:  now.Add(time.Duration(secretReq.ExpireAfter) * time.Minute),
		MaxViews:   secretReq.ExpireAfterViews,
		RemViews:   secretReq.ExpireAfterViews,
	}
	iError := s.repo.Create(ctx, secret)
	return secret, iError
}

func (s *SecretService) GetSecret(ctx *gin.Context, id string) (*models.Secret, crerrors.IError) {
	secret := &models.Secret{}
	err := s.repo.FindByPkIdString(ctx, secret, id)
	if err != nil {
		return nil, err
	}
	var iError crerrors.IError
	if time.Now().After(secret.ExpiresAt) {
		iError = crerrors.NewCrError(ctx, crerrors.SecretExpired, errors.New("The requested secret has been expired"))
		return nil, iError
	}
	if secret.RemViews == 0 {
		iError = crerrors.NewCrError(ctx, crerrors.SecretLimitReached, errors.New("The requested secret view limit has been reached"))
		return nil, iError
	}
	values := make(map[string]interface{})
	values["rem_views"] = secret.RemViews - 1
	values["expires_at"] = secret.ExpiresAt
	iError = s.repo.Update(ctx, secret, values)
	return secret, iError
}
