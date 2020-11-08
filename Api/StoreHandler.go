package api

import (
"bytes"
"database/sql"
"encoding/json"
"github.com/dunzoit/Store/Model"
"github.com/gorilla/mux"
"net/http"
"strconv"
)

type StoreProduct struct{
	Prod []int `json:"prod"`
}
func (a *App) handleStoreProduct(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid store ID")
		return
	}

	s := Model.Store{StoreId:id}

	if err := s.GetStoreProduct(a.DB); err != nil {
		switch err {
		case sql.ErrNoRows:
			respondWithError(w, http.StatusNotFound, "Product not found")
		default:
			respondWithError(w, http.StatusInternalServerError, err.Error())
		}
		return
	}

	respondWithJSON(w, http.StatusOK, s)
}


func (a *App) handleInsertProduct(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid Product ID")
		return
	}
	StoreToAdd := StoreProduct{}
	buf := new(bytes.Buffer)
	buf.ReadFrom(r.Body)
	JsonString := buf.String()
	var jsonStr = []byte(JsonString)
	err = json.Unmarshal(jsonStr, &StoreToAdd)
	if err !=nil {
		respondWithError(w, http.StatusInternalServerError,err.Error())
		return
	}
	ReferenceToStore := Model.Store{StoreId:id}
	Productlist := StoreToAdd.Prod
	err = ReferenceToStore.SetStoreProduct(a.DB,Productlist)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError,err.Error())
		return
	}
	respondWithJSON(w, http.StatusOK, err)
}


