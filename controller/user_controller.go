package controller

import (
	"encoding/json"
	"net/http"

	"github.com/gildemberg-santos/picpay-desafio-backend/model"
)

type Response struct {
	Message string
}

type User struct {
	Fullname      string
	TaxIdentifier string
	Email         string
	Password      string
}

func UserCreateController(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	user_request := User{}
	json.NewDecoder(r.Body).Decode(&user_request)

	var user model.User
	db, err := model.NewDB(&user)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(Response{Message: err.Error()})
		return
	}
	user = model.User{
		Fullname:      user_request.Fullname,
		TaxIdentifier: user_request.TaxIdentifier,
		Email:         user_request.Email,
		Password:      user_request.Password,
	}
	result := db.Create(&user)

	if result.Error != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(Response{Message: result.Error.Error()})
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(Response{Message: "User created successfully"})
}

func UserAllController(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	db, err := model.NewDB(&model.User{})
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(Response{Message: err.Error()})
		return
	}

	var users []model.User
	result := db.Find(&users)

	if result.Error != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(Response{Message: result.Error.Error()})
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(users)
}

func UserUpdateController(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	user_request := User{}
	json.NewDecoder(r.Body).Decode(&user_request)

	params := r.URL.Query()
	id := params.Get("id")

	var user model.User
	db, err := model.NewDB(&user)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(Response{Message: err.Error()})
		return
	}

	result := db.Find(&user, id)
	if result.Error != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(Response{Message: result.Error.Error()})
		return
	}

	user.Fullname = user_request.Fullname
	user.TaxIdentifier = user_request.TaxIdentifier
	user.Email = user_request.Email
	user.Password = user_request.Password

	result = db.Save(&user)
	if result.Error != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(Response{Message: result.Error.Error()})
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(Response{Message: "User updated successfully"})
}

func UserFindController(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := r.URL.Query()
	id := params.Get("id")

	var user model.User
	db, err := model.NewDB(&user)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(Response{Message: err.Error()})
		return
	}

	result := db.Find(&user, id)
	if result.Error != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(Response{Message: result.Error.Error()})
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(user)
}

func UserDeleteController(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := r.URL.Query()
	id := params.Get("id")

	var user model.User
	db, err := model.NewDB(&user)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(Response{Message: err.Error()})
		return
	}

	result := db.Find(&user, id)
	if result.Error != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(Response{Message: result.Error.Error()})
		return
	}

	result = db.Delete(&user)
	if result.Error != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(Response{Message: result.Error.Error()})
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(Response{Message: "User deleted successfully"})
}
