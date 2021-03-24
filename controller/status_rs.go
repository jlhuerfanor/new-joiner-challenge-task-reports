package controller

import (
	"bitbucket.org/new-joiners/taskreports/business"
	"encoding/json"
	"github.com/golobby/container"
	"log"
	"net/http"
)

func GetStatus(w http.ResponseWriter, _ *http.Request) {
	container.Make(func (statusSrv business.GetStatusBusiness) {
		var status = statusSrv.GetStatus()
		log.Printf("%+v", status)
		err := json.NewEncoder(w).Encode(status)
		log.Print(err)
	})
}