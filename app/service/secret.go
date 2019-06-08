package service

import (
	"github.com/gin-gonic/gin"
	"secret-server/app/crerrors"
	"secret-server/app/models"
	"secret-server/app/repository"
)

type ISecretService interface {
	CreateSecret(ctx *gin.Context, secret *models.Secret) crerrors.IError
	GetSecret(ctx *gin.Context, id string) (crerrors.IError, *models.Secret)
}

type SecretService struct {
	repo repository.ISecretsRepo
}

func NewSecretService(repo repository.ISecretsRepo) *SecretService {
	return &SecretService{repo: repo}
}

func (s *SecretService) CreateSecret(ctx *gin.Context, secret *models.Secret) crerrors.IError {
	return s.repo.Create(ctx, secret)
}

func (s *SecretService) GetSecret(ctx *gin.Context, id string) (crerrors.IError, *models.Secret) {
	secret := &models.Secret{}
	err := s.repo.FindByPkIdString(ctx, secret, id)
	if err != nil {
		return err, nil
	}
	return nil, secret
}
