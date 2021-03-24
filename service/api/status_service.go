package api

import "bitbucket.org/new-joiners/taskreports/model"

type StatusService interface {
	GetStatus() model.Status
}
