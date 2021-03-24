package platform

import (
	"bitbucket.org/new-joiners/taskreports/platform/application"
	"bitbucket.org/new-joiners/taskreports/model"
)

func ConfigureContainer(config model.ApplicationConfig) {
	application.ConfigureStatusDomain()
}
