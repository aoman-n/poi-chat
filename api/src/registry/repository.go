package registry

import (
	"github.com/laster18/poi/api/src/domain"
	"github.com/laster18/poi/api/src/infra/redis"
	"github.com/laster18/poi/api/src/repository"
	"gorm.io/gorm"
)

// handlerはregistryを持ち、registryからrepositoryを取得することができる
// usecaseはrepositoryを持つ。registryを持たない

// interface化することでテスト時には必要なものだけ実装すればよくなる
type Repository interface {
	NewGlobalUser() domain.GlobalUserRepo
}

type repositoryImpl struct {
	redis *redis.Client
	db    *gorm.DB
}

func NewRepository() Repository {
	return &repositoryImpl{}
}

func (r *repositoryImpl) NewGlobalUser() domain.GlobalUserRepo {
	return repository.NewGlobalUserRepo(r.redis)
}
