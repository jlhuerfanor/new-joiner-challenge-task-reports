package application

import (
	"bitbucket.org/new-joiners/taskreports/model"
	"bitbucket.org/new-joiners/taskreports/service/api"
	"bitbucket.org/new-joiners/taskreports/service/implementation"
	"github.com/golobby/container"
)

func ConfigureJoinerDomain(config model.ApplicationConfig) {
	container.Singleton(func() api.JoinerQueryService {
		return &implementation.RestJoinerQueryService{ Configuration: config.JoinerClient}
	})
}
