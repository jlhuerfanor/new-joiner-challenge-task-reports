package implementation

import (
	"bitbucket.org/new-joiners/taskreports/model"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type RestTaskQueryService struct {
	Configuration model.TaskClientConfig
}

func (instance RestTaskQueryService) GetTaskIdList() ([]int, error) {
	client := &http.Client{}
	req, err := http.NewRequest("GET", fmt.Sprintf("%v/task", instance.Configuration.BaseUrl), nil)

	if err != nil {
		return nil, err
	}
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")

	resp, err := client.Do(req)

	if err != nil {
		return nil, err
	} else {
		bodyBytes, err := ioutil.ReadAll(resp.Body)

		if err != nil {
			return nil, fmt.Errorf("%v | %v", err, resp.Body.Close())
		}

		ids := make([]int, 0)
		jsonErr := json.Unmarshal(bodyBytes, &ids)

		if jsonErr != nil {
			return nil, fmt.Errorf("%v | %v", jsonErr, resp.Body.Close())
		}

		return ids, resp.Body.Close()
	}
}

func (instance RestTaskQueryService) GetTask(id int) (model.Task, error) {
	client := &http.Client{}
	req, err := http.NewRequest("GET", fmt.Sprintf("%v/task/%v", instance.Configuration.BaseUrl, id), nil)

	if err != nil {
		return model.Task{}, err
	}
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")

	resp, err := client.Do(req)

	if err != nil {
		return model.Task{}, err
	} else {
		bodyBytes, err := ioutil.ReadAll(resp.Body)

		if err != nil {
			return model.Task{}, fmt.Errorf("%v | %v", err, resp.Body.Close())
		}

		var task model.Task
		jsonErr := json.Unmarshal(bodyBytes, &task)

		if jsonErr != nil {
			return model.Task{}, fmt.Errorf("%v | %v", jsonErr, resp.Body.Close())
		}

		return task, resp.Body.Close()
	}
}