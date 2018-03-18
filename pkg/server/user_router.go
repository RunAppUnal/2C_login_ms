package server

import (
	"encoding/json"
	"errors"
	"go_rest_api/pkg"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type userRouter struct {
	userService root.UserService
}

func NewUserRouter(u root.UserService, router *mux.Router) *mux.Router {
	userRouter := userRouter{u}

	router.HandleFunc("/register", userRouter.createUserHandler).Methods("POST")
  router.HandleFunc("/{username}", userRouter.getUserHandler).Methods("GET")
	router.HandleFunc("/{username}", userRouter.deleteUserHandler).Methods("DELETE")
  router.HandleFunc("/login", userRouter.loginHandler).Methods("POST")
	return router
}

func (ur *userRouter) createUserHandler(w http.ResponseWriter, r *http.Request) {
	user, err := decodeUser(r)
	if err != nil {
		Error(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
  err = ur.userService.CreateUser(&user)
	if err != nil {
		Error(w, http.StatusInternalServerError, err.Error())
		return
	}

	Json(w, http.StatusOK, err)
}

func (ur *userRouter) deleteUserHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	log.Println(vars)
	username := vars["username"]

	err := ur.userService.DeleteByUsername(username)
	if err != nil {
		Error(w, http.StatusNotFound, err.Error())
		return
	}

	Json(w, http.StatusOK, err)
}

func (ur *userRouter) getUserHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	log.Println(vars)
	username := vars["username"]

	user, err := ur.userService.GetByUsername(username)
	if err != nil {
		Error(w, http.StatusNotFound, err.Error())
		return
	}

	Json(w, http.StatusOK, user)
}

func(ur* userRouter) loginHandler(w http.ResponseWriter, r *http.Request) {
  log.Println("loginHandler")
  err, credentials := decodeCredentials(r)
  if err != nil {
    Error(w, http.StatusBadRequest, "Invalid request payload")
    return
  }

  var user root.User
  err, user = ur.userService.Login(credentials)
	log.Println(user)
  if err == nil {
    Json(w, http.StatusOK, user)
  } else {
    Error(w, http.StatusInternalServerError, "Incorrect password")
  }
}

func decodeUser(r *http.Request) (root.User, error) {
  var u root.User
  if r.Body == nil {
    return u, errors.New("no request body")
  }
  decoder := json.NewDecoder(r.Body)
  err := decoder.Decode(&u)
  return u, err
}

func decodeCredentials(r *http.Request) (error,root.Credentials) {
  var c root.Credentials
  if r.Body == nil {
    return errors.New("no request body"), c
  }
  decoder := json.NewDecoder(r.Body)
  err := decoder.Decode(&c)
  return err, c
}