package blog

import (
	"Project2-v3/pkg/utils"
	"encoding/json"
	"net/http"
)

type BlogHandler struct {
	service *BlogService
}

func NewBlogHandler(service *BlogService) *BlogHandler {
	return &BlogHandler{service: service}
}

func (bh *BlogHandler) CreateBlog(w http.ResponseWriter, r *http.Request) {
	var req CreateBlogRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err.Error())
		return
	}

	err = bh.service.Create(req)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err.Error())
		return
	}

	utils.WriteJSON(w, http.StatusCreated, map[string]string{"message": "Blog created successfully"})
}

func (bh *BlogHandler) GetAllBlogs(w http.ResponseWriter, r *http.Request) {
	blogs, err := bh.service.GetAll()
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err.Error())
		return
	}

	utils.WriteJSON(w, http.StatusOK, blogs)
}

func (bh *BlogHandler) GetBlogById(w http.ResponseWriter, r *http.Request) {
	id, ok := utils.ParseID(w, r)
	if !ok {
		return
	}

	blog, err := bh.service.GetById(id)
	if err != nil {
		utils.WriteError(w, http.StatusNotFound, "blog not found")
		return
	}

	utils.WriteJSON(w, http.StatusOK, blog)
}

func (bh *BlogHandler) UpdateBlog(w http.ResponseWriter, r *http.Request) {
	id, ok := utils.ParseID(w, r)
	if !ok {
		return
	}

	var req UpdateBlogRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err.Error())
		return
	}

	err = bh.service.Update(id, req)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err.Error())
		return
	}

	utils.WriteJSON(w, http.StatusOK, map[string]string{"message": "Blog updated successfully"})
}

func (bh *BlogHandler) DeleteBlog(w http.ResponseWriter, r *http.Request) {
	id, ok := utils.ParseID(w, r)
	if !ok {
		return
	}
	err := bh.service.Delete(id)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err.Error())
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
