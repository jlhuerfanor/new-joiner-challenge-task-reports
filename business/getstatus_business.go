package business

import (
	"bitbucket.org/new-joiners/taskreports/model"
	"bitbucket.org/new-joiners/taskreports/service/api"
)

type GetStatusBusiness struct {
	StatusService api.StatusService
}

func (instance GetStatusBusiness) GetStatus() model.Status {
	return instance.StatusService.GetStatus()
}
