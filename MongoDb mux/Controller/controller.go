package controller

import (
	"encoding/json"
	model "go-rest-api-db/Model"
	service "go-rest-api-db/Service"
	"net/http"

	"github.com/gorilla/mux"
)

type NetflixController interface {
	GetAll(w http.ResponseWriter, r *http.Request)
	Get(w http.ResponseWriter, r *http.Request)
	Create(w http.ResponseWriter, r *http.Request)
	Update(w http.ResponseWriter, r *http.Request)
	Delete(w http.ResponseWriter, r *http.Request)
	DeleteAll(w http.ResponseWriter, r *http.Request)
}

type netflixController struct {
	service service.NetflixService
}

func NewController(s service.NetflixService) NetflixController {
	return &netflixController{
		service: s,
	}
}

//controller

var (
	s service.NetflixService = service.NewNetflixService()
)

func (c *netflixController) GetAll(w http.ResponseWriter, r *http.Request) {
	//w.Header().Set("Content-Type", "application/x-www-form-urlencode")
	w.Header().Set("Content-Type", "application/json")
	allmovies := s.GetAllMovie()
	json.NewEncoder(w).Encode(allmovies)
}
func (c *netflixController) Get(w http.ResponseWriter, r *http.Request) {
	//w.Header().Set("Content-Type", "application/x-www-form-urlencode")
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	allmovies := s.GetMovie(params["id"])
	json.NewEncoder(w).Encode(allmovies)
}
func (c *netflixController) Create(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencode")
	w.Header().Set("Allow-Control-Allow-Methods", "POST")
	var movie model.Netflix
	_ = json.NewDecoder(r.Body).Decode(&movie)
	s.Insert(movie)
	json.NewEncoder(w).Encode(movie)
}
func (c *netflixController) Update(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencode")
	w.Header().Set("Allow-Control-Allow-Methods", "POST")
	params := mux.Vars(r)
	s.Update(params["id"])
}
func (c *netflixController) DeleteAll(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencode")
	w.Header().Set("Allow-Control-Allow-Methods", "POST")
	count := s.DeleteAll()
	json.NewEncoder(w).Encode(count)
}
func (c *netflixController) Delete(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencode")
	w.Header().Set("Allow-Control-Allow-Methods", "POST")
	params := mux.Vars(r)
	s.Delete(params["id"])
	json.NewEncoder(w).Encode(params["id"])
}
