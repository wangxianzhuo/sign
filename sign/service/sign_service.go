package service

import (
	"github.com/wangxianzhuo/sign/env"
	"github.com/wangxianzhuo/sign/sign/model"
)

func Get(id string) (sign model.Sign, err error) {
	return model.Get(id, env.DB)
}

func GetByUserIDAndReferenceID(userID, referenceID string) (signs []model.Sign, err error) {
	return model.GetByUserIDAndReferenceID(userID, referenceID, env.DB)
}

func Delete(id string) error {
	return model.Delete(id, env.DB)
}
