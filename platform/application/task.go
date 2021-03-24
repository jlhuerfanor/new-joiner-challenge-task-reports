package application

import (
	"bitbucket.org/new-joiners/taskreports/model"
	"bitbucket.org/new-joiners/taskreports/service/api"
	"bitbucket.org/new-joiners/taskreports/service/implementation"
	"github.com/golobby/container"
)

func ConfigureTaskDomain(config model.ApplicationConfig) {
	container.Singleton(func() api.TaskQueryService {
		return &implementation.RestTaskQueryService{ Configuration: config.TaskClient}
	})
}
