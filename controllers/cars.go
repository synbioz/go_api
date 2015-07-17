package controllers

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/synbioz/go_api/models"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
)

func CarsIndex(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json;charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(models.AllCars())
}

func CarsCreate(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json;charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	body, err := ioutil.ReadAll(r.Body)

	if err != nil {
		log.Fatal(err)
	}

	var car models.Car

	err = json.Unmarshal(body, &car)

	if err != nil {
		log.Fatal(err)
	}

	models.NewCar(&car)

	json.NewEncoder(w).Encode(car)
}

func CarsShow(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json;charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])

	if err != nil {
		log.Fatal(err)
	}

	car := models.FindCarById(id)

	json.NewEncoder(w).Encode(car)
}
func CarsUpdate(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json;charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])

	if err != nil {
		log.Fatal(err)
	}

	body, err := ioutil.ReadAll(r.Body)

	if err != nil {
		log.Fatal(err)
	}

	car := models.FindCarById(id)

	err = json.Unmarshal(body, &car)

	models.UpdateCar(car)

	json.NewEncoder(w).Encode(car)
}
func CarsDelete(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json;charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	vars := mux.Vars(r)

	// strconv.Atoi is shorthand for ParseInt
	id, err := strconv.Atoi(vars["id"])

	if err != nil {
		log.Fatal(err)
	}

	err = models.DeleteCarById(id)
}
