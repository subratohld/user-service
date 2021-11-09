package service

import (
	logger "github.com/subratohld/logger"
	"github.com/subratohld/user-service/internal/model"
	"github.com/subratohld/user-service/internal/repository"
	"go.uber.org/zap"
)

type User interface {
	Save(user *model.User) (*model.User, error)
}

func NewUserService(repo repository.User, logger logger.Logger) User {
	return &user{
		repo:   repo,
		logger: logger,
	}
}

type user struct {
	logger logger.Logger
	repo   repository.User
}

func (usr *user) Save(user *model.User) (*model.User, error) {
	id, err := usr.repo.Save(user)
	if err != nil {
		usr.logger.Error("cound not save user", zap.Error(err))
		return nil, err
	}
	user.Id = id
	return user, nil
}
