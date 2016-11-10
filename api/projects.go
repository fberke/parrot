package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/anthonynsimon/parrot/errors"
	"github.com/anthonynsimon/parrot/model"
	"github.com/anthonynsimon/parrot/render"
	"github.com/pressly/chi"
)

func createProject(w http.ResponseWriter, r *http.Request) {
	project := model.Project{}
	errs := decodeAndValidate(r.Body, &project)
	if errs != nil {
		render.Error(w, http.StatusUnprocessableEntity, errs)
		return
	}
	userID, err := getUserIDFromContext(r.Context())
	if err != nil {
		handleError(w, errors.ErrInternal)
		return
	}

	result, err := store.CreateProject(&project)
	if err != nil {
		handleError(w, errors.ErrInternal)
		return
	}
	pu := model.ProjectUser{ProjectID: result.ID, UserID: userID, Role: "admin"}
	err = store.AssignProjectUser(pu)
	if err != nil {
		handleError(w, errors.ErrInternal)
		return
	}

	render.JSON(w, http.StatusCreated, result)
}

func updateProject(w http.ResponseWriter, r *http.Request) {
	projectID, err := strconv.Atoi(chi.URLParam(r, "projectID"))
	if err != nil {
		handleError(w, errors.ErrBadRequest)
		return
	}

	project := model.Project{}
	if err := json.NewDecoder(r.Body).Decode(&project); err != nil {
		handleError(w, errors.ErrBadRequest)
		return
	}
	project.ID = projectID
	project.SanitizeKeys()

	err = store.UpdateProject(&project)
	if err != nil {
		handleError(w, errors.ErrInternal)
		return
	}

	render.JSON(w, http.StatusOK, project)
}

func showProject(w http.ResponseWriter, r *http.Request) {
	projectID, err := strconv.Atoi(chi.URLParam(r, "projectID"))
	if err != nil {
		handleError(w, errors.ErrBadRequest)
		return
	}

	project, err := store.GetProject(projectID)
	if err != nil {
		handleError(w, err)
		return
	}

	render.JSON(w, http.StatusOK, project)
}

func deleteProject(w http.ResponseWriter, r *http.Request) {
	projectID, err := strconv.Atoi(chi.URLParam(r, "projectID"))
	if err != nil {
		handleError(w, errors.ErrBadRequest)
		return
	}

	resultID, err := store.DeleteProject(projectID)
	if err != nil {
		handleError(w, errors.ErrInternal)
		return
	}

	render.JSON(w, http.StatusOK, map[string]interface{}{
		"message": fmt.Sprintf("deleted project with id %d and all related locales", resultID),
	})
}
