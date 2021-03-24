package implementation

import "bitbucket.org/new-joiners/taskreports/model"

type DefaultStatusService struct { }

func (m DefaultStatusService) GetStatus() model.Status {
	return model.Status{ Status: "Up"}
}
