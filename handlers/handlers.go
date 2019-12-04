package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/vm/models"
)

type Pagination struct {
	Pagesize int    `json:"pagesize,omitempty"`
	Offset   int    `json:"offset,omitempty"`
	OrderBy  string `json:"sort_by,omitempty"`
}

const (
	DefaultPagesize = 4
	DefaultOffset   = 0
	DefaultOrderBy  = "instance_id asc"
)

//CreateInstance method create instance in database
func CreateInstance(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	var instance models.Instance
	err := json.NewDecoder(r.Body).Decode(&instance)
	if err != nil {
		w.WriteHeader(400)
		return
	}
	resp, err := models.CreateInstance(&instance)
	if err != nil {
		w.WriteHeader(404)
		return
	}
	json.NewEncoder(w).Encode(resp)
}

//GetAllInstance method returns instance details from database based on query parameters
func GetAllInstances(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	queryParamMap := r.URL.Query()
	pagesize := DefaultPagesize
	offset := DefaultOffset
	order_by := DefaultOrderBy

	value, _ := queryParamMap["pagesize"]

	key, _ := queryParamMap["offset"]

	order, _ := queryParamMap["order_by"]

	if value != nil {
		Pagesize, err := strconv.Atoi(value[0])
		if err != nil {
			err.Error()
		}

		if Pagesize != 0 {
			pagesize = Pagesize
		}
	}

	if key != nil {
		Offset, err := strconv.Atoi(key[0])
		if err != nil {
			err.Error()
		}

		if Offset != 0 {
			offset = Offset
		}
	}

	if order != nil {
		Order_by := order[0]

		if Order_by != "" {
			order_by = Order_by
		}
	}

	instances, err := models.GetAllInstances(pagesize, offset, order_by)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(instances)
}

//GetInstance method returns an instance detail from database
func GetInstance(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var instance models.Instance
	var err error
	param := mux.Vars(r)
	instanceid := param["instanceid"]
	resp, err := models.GetInstance(&instanceid, &instance)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(resp)
}

//UpdateInstance method update specific instance in database
func UpdateInstance(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var instance models.Instance
	var err error
	param := mux.Vars(r)
	instanceid := param["instance_id"]
	_ = json.NewDecoder(r.Body).Decode(&instance)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	models.UpdateInstance(&instanceid, &instance)
	json.NewEncoder(w).Encode(instance)
}

////DeleteInstance method delete specific instance from database
func DeleteInstance(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var instance models.Instance
	var err error
	param := mux.Vars(r)
	instanceid := param["instanceid"]
	resp, err := models.DeleteInstance(&instanceid, &instance)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(resp)
}
