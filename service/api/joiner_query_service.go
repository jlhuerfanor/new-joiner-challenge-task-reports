package api

import "bitbucket.org/new-joiners/taskreports/model"

type JoinerQueryService interface {
	GetByIdNumber(idNumber int64) (model.Joiner, error)
}
