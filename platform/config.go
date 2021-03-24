package platform

import (
	"bitbucket.org/new-joiners/taskreports/model"
	"os"
	"strconv"
)

func GetConfiguration() model.ApplicationConfig {
	port, portError := strconv.Atoi(os.Getenv("PORT"))

	if portError != nil {
		port = 10000
	}

	return model.ApplicationConfig {
		Port: port,
		ContextPath: "/wap/task-reports",
		TaskClient: model.TaskClientConfig{ BaseUrl: "https://joiner-tasks.herokuapp.com/wap/tasks" },
		JoinerClient: model.JoinerClientConfig{ BaseUrl: "https://new-joiners.herokuapp.com/wap/new-joiners" } }
}