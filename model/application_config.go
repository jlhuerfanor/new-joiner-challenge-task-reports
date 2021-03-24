package model

type ApplicationConfig struct {
	Port int `json:"port"`
	ContextPath string `json:"contextPath"`
	TaskClient TaskClientConfig `json:"taskClient"`
}
