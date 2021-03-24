package implementation

import (
	"testing"
)

func TestDefaultStatusService_GetStatus(t *testing.T) {
	statusSrv := DefaultStatusService{}
	status := statusSrv.GetStatus()

	if status.Status != "Up" {
		t.Errorf("Status is not equals to Up. Status: %v", status.Status)
	}
}
