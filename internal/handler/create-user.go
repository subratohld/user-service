package handler

import (
	"encoding/json"
	"net/http"

	"github.com/subratohld/user-service/internal/model"
	"github.com/subratohld/user-service/internal/service"
	"go.uber.org/multierr"
)

func NewCreateUserHandler(userSvc service.User) http.Handler {
	return &createUserHandler{
		svc: userSvc,
	}
}

type createUserHandler struct {
	svc service.User
}

func (h createUserHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	resp := new(jsonResp)
	if r.Body == nil {
		resp.Errors("empty request body").Write(w, http.StatusBadRequest)
		return
	}

	defer func() {
		r.Body.Close()
	}()

	var user model.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		resp.Errors("could not process request").Write(w, http.StatusUnprocessableEntity)
		return
	}

	err := user.Validate()
	if len(multierr.Errors(err)) > 0 {
		resp.Errors("could not process request").Write(w, http.StatusBadRequest)
		return
	}

	dbUser, err := h.svc.Save(&user)
	if err != nil {
		resp.Errors("could not process request").Write(w, http.StatusInternalServerError)
		return
	}

	resp.Body(dbUser).Message("User inserted successfully!").Write(w, http.StatusOK)
}
