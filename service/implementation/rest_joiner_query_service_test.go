package implementation

import (
	"bitbucket.org/new-joiners/taskreports/model"
	"testing"
)

func TestRestJoinerQueryService_GetByIdNumber(t *testing.T) {
	taskQueryService := RestJoinerQueryService{
		Configuration: model.JoinerClientConfig { BaseUrl: "https://new-joiners.herokuapp.com/wap/new-joiners" } }

	profile, err := taskQueryService.GetByIdNumber(1023456789)

	if err != nil {
		t.Errorf("Error requesting to server. %v", err)
		return
	}

	if len(profile.IdNumber) == 0 {
		t.Error("Expected results, but got nothing")
	}
}

func TestRestJoinerQueryService_GetByIdNumber_Empty(t *testing.T) {
	taskQueryService := RestJoinerQueryService{
		Configuration: model.JoinerClientConfig { BaseUrl: "https://new-joiners.herokuapp.com/wap/new-joiners" } }

	profile, err := taskQueryService.GetByIdNumber(25)

	if err != nil {
		t.Errorf("Error requesting to server. %v", err)
		return
	}

	if len(profile.IdNumber) > 0 {
		t.Error("Expected nothing, but got results")
	}
}
