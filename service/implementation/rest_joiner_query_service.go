package implementation

import (
	"bitbucket.org/new-joiners/taskreports/model"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type RestJoinerQueryService struct {
	Configuration model.JoinerClientConfig
}

func (instance RestJoinerQueryService) GetByIdNumber(idNumber int64) (model.Joiner, error) {
	client := &http.Client{}
	req, err := http.NewRequest("GET", fmt.Sprintf("%v/joiner/%v", instance.Configuration.BaseUrl, idNumber), nil)

	if err != nil {
		return model.Joiner{}, err
	}
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")

	resp, err := client.Do(req)

	if err != nil {
		return model.Joiner{}, err
	} else {
		bodyBytes, err := ioutil.ReadAll(resp.Body)

		if err != nil {
			return model.Joiner{}, fmt.Errorf("%v | %v", err, resp.Body.Close())
		}

		var joiner  model.Joiner
		var jsonErr error

		if len(bodyBytes) > 0 {
			jsonErr = json.Unmarshal(bodyBytes, &joiner)
		} else {
			jsonErr = nil
		}

		if jsonErr != nil {
			return model.Joiner{}, fmt.Errorf("%v | %v", jsonErr, resp.Body.Close())
		}

		return joiner, resp.Body.Close()
	}
}
