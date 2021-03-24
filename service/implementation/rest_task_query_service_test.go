package implementation


import (
	"bitbucket.org/new-joiners/taskreports/model"
	"fmt"
	"testing"
)

func TestRestTaskQueryService_GetTaskIdList(t *testing.T) {
	taskQueryService := RestTaskQueryService{
		Configuration: model.TaskClientConfig { BaseUrl: "https://joiner-tasks.herokuapp.com/wap/tasks" } }

	list, err := taskQueryService.GetTaskIdList()

	if err != nil {
		t.Errorf("Error requesting to server. %v", err)
		return
	}

	if len(list) == 0 {
		t.Error("Expected results, but get nothing")
	}
}

func TestRestTaskQueryService_GetTask(t *testing.T) {
	taskQueryService := RestTaskQueryService{
		Configuration: model.TaskClientConfig { BaseUrl: "https://joiner-tasks.herokuapp.com/wap/tasks" } }

	list, err := taskQueryService.GetTaskIdList()

	if err == nil {
		var tasks = make([]model.Task, len(list))

		for i, id := range list {
			taskI, err := taskQueryService.GetTask(id)
			if err == nil {
				tasks[i] = taskI
				fmt.Printf("Task[%v] = %+v", i, tasks[i])
				fmt.Println()

				if tasks[i].Id !=  id {
					t.Errorf("Unexpected task id. Current: %v. Expected: %v.", tasks[i].Id, id)
				}
 			} else {
				t.Errorf("Could not request task details. Task Id: %v", id)
				return
			}
		}
	} else {
		t.Error("Could not request task ids.")
		return
	}
}