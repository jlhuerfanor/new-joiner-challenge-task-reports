package business

import (
	"testing"
	"bitbucket.org/new-joiners/taskreports/model"
)

var expectedResult = model.Status{Status: "Up"}
var calledGetStatusMock = false

type StatusServiceMock struct { }

func (mock StatusServiceMock) GetStatus() model.Status {
	calledGetStatusMock = true
	return expectedResult
}

func TestGetStatusBusiness_GetStatus(t *testing.T) {
	statusSrvMock := StatusServiceMock{}
	getStatusBusiness := GetStatusBusiness{StatusService: statusSrvMock}

	result := getStatusBusiness.GetStatus()

	if !calledGetStatusMock {
		t.Errorf("Service was not called. %v", calledGetStatusMock)
	}
	if result != expectedResult {
		t.Errorf("Result is not expected: %+v. expected %+v", result, expectedResult)
	}
}