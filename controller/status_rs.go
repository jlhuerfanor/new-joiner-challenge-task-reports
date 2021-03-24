package controller

import (
	"bitbucket.org/new-joiners/taskreports/business"
	"github.com/golobby/container"
	"net/http"
	"encoding/json"
	"log"
)

func GetStatus(w http.ResponseWriter, r *http.Request) {
	container.Make(func (statusSrv business.GetStatusBusiness) {
		var status = statusSrv.GetStatus()
		log.Printf("%+v", status)
		json.NewEncoder(w).Encode(status)
	})
}