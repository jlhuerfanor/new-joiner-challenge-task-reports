package application

import (
	"github.com/golobby/container"
	"bitbucket.org/new-joiners/taskreports/service/api"
	"bitbucket.org/new-joiners/taskreports/service/implementation"
	"bitbucket.org/new-joiners/taskreports/business"
)

func ConfigureStatusDomain() {
	container.Singleton(func() api.StatusService {
		return &implementation.DefaultStatusService{}
	})
	container.Singleton(func(statusSrv api.StatusService) business.GetStatusBusiness {
		return business.GetStatusBusiness{StatusService: statusSrv}
	})
}
