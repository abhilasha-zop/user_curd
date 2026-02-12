package handler

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"
	"user_crud/model"
	"user_crud/service"
)

type UserHandler struct {
	Service *service.UserService
}

func (h *UserHandler) Create(w http.ResponseWriter, r *http.Request) {
	var users []model.User
	json.NewDecoder(r.Body).Decode(&users)

	//err := h.Service.CreateUser(u)
	//if err != nil {
	//	http.Error(w, err.Error(), 400)
	//	return
	//}

	for _, u := range users {
		if err := h.Service.CreateUser(u); err != nil {
			http.Error(w, err.Error(), 400)
			return
		}

	}
	w.WriteHeader(http.StatusCreated)
}

func (h *UserHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	users, _ := h.Service.GetUsers()
	json.NewEncoder(w).Encode(users)
}

func (h *UserHandler) GetOne(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodGet {
		http.Error(w, "method not allowed", 405)
		return
	}

	parts := strings.Split(r.URL.Path, "/")
	if len(parts) < 3 {
		http.Error(w, "invalid path", 400)
		return
	}

	idStr := parts[2]

	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "invalid id", 400)
		return
	}

	user, err := h.Service.GetUser(id)
	if err != nil {
		http.Error(w, "user not found", 404)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}

func (h *UserHandler) Update(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodPut {
		http.Error(w, "method not allowed", 405)
		return
	}

	var u model.User

	err := json.NewDecoder(r.Body).Decode(&u)
	if err != nil {
		http.Error(w, "invalid body", 400)
		return
	}

	if u.ID == 0 {
		http.Error(w, "id required", 400)
		return
	}

	err = h.Service.UpdateUser(u)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(`{"message":"updated successfully"}`))
}

func (h *UserHandler) Delete(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodDelete {
		http.Error(w, "method not allowed", 405)
		return
	}

	idStr := strings.TrimPrefix(r.URL.Path, "/delete/")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "invalid id", 400)
		return
	}

	err = h.Service.DeleteUser(id)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	w.Write([]byte("deleted"))
}
