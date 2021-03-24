package api

import "bitbucket.org/new-joiners/taskreports/model"

type TaskQueryService interface {
	GetTaskIdList() ([]int, error)
	GetTask(id int) (model.Task, error)
}