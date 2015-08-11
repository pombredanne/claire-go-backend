package main

import (
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type MakeController struct {
	BaseController
	makeRepo IMakeRepository
}

func (mc MakeController) Index(rw http.ResponseWriter, request *http.Request) {
	makes, err := mc.makeRepo.GetAll(2)

	err = mc.Render(rw, makes, err)
	if err != nil {
		// TODO Implement as error method on base controller
		rw.Write([]byte("Internal server error"))
	}
}

func (mc MakeController) View(rw http.ResponseWriter, request *http.Request) {
	params := mux.Vars(request)
	id, _ := strconv.Atoi(params["id"])

	make, err := mc.makeRepo.Get(id)

	mc.Render(rw, make, err)
}
