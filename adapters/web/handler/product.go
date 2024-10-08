package handler

import (
	"encoding/json"
	"net/http"

	"github.com/felipeamendes/arq-hexagonal/adapters/dto"
	"github.com/felipeamendes/arq-hexagonal/application"
	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
)

func MakeProductHandlers(r *mux.Router, n *negroni.Negroni, service application.ProductServiceInterface) {
	r.Handle("/product/{id}/enable", n.With(
		negroni.Wrap(enableProduct(service)),
	)).Methods("PUT", "OPTIONS")
	r.Handle("/product/{id}/disable", n.With(
		negroni.Wrap(disableProduct(service)),
	)).Methods("PUT", "OPTIONS")
	r.Handle("/product/{id}", n.With(
		negroni.Wrap(getProduct(service)),
	)).Methods("GET", "OPTIONS")
	r.Handle("/product", n.With(
		negroni.Wrap(createProduct(service)),
	)).Methods("POST", "OPTIONS")
}

func getProduct(service application.ProductServiceInterface) http.Handler {
	return http.HandlerFunc(func(response http.ResponseWriter, request *http.Request) {
		response.Header().Set("Content-Type", "application/json")
		vars := mux.Vars(request)
		id := vars["id"]
		product, err := service.Get(id)
		if err != nil {
			response.WriteHeader(http.StatusNotFound)
			return
		}
		err = json.NewEncoder(response).Encode(product)
		if err != nil {
			response.WriteHeader(http.StatusInternalServerError)
			return
		}
	})
}

func createProduct(service application.ProductServiceInterface) http.Handler {
	return http.HandlerFunc(func(response http.ResponseWriter, request *http.Request) {
		response.Header().Set("Content-Type", "application/json")
		var productDto dto.Product
		err := json.NewDecoder(request.Body).Decode(&productDto)
		if err != nil {
			response.WriteHeader(http.StatusInternalServerError)
			response.Write(jsonError(err.Error()))
			return
		}
		product, err := service.Create(productDto.Name, productDto.Price)
		if err != nil {
			response.WriteHeader(http.StatusInternalServerError)
			response.Write(jsonError(err.Error()))
			return
		}
		err = json.NewEncoder(response).Encode(product)
		if err != nil {
			response.WriteHeader(http.StatusInternalServerError)
			response.Write(jsonError(err.Error()))
			return
		}
	})
}

func enableProduct(service application.ProductServiceInterface) http.Handler {
	return http.HandlerFunc(func(response http.ResponseWriter, request *http.Request) {
		response.Header().Set("Content-Type", "application/json")
		vars := mux.Vars(request)
		id := vars["id"]
		product, err := service.Get(id)
		if err != nil {
			response.WriteHeader(http.StatusNotFound)
			return
		}
		result, err := service.Enable(product)
		if err != nil {
			response.WriteHeader(http.StatusInternalServerError)
			response.Write(jsonError(err.Error()))
			return
		}
		err = json.NewEncoder(response).Encode(result)
		if err != nil {
			response.WriteHeader(http.StatusInternalServerError)
			return
		}
	})
}

func disableProduct(service application.ProductServiceInterface) http.Handler {
	return http.HandlerFunc(func(response http.ResponseWriter, request *http.Request) {
		response.Header().Set("Content-Type", "application/json")
		vars := mux.Vars(request)
		id := vars["id"]
		product, err := service.Get(id)
		if err != nil {
			response.WriteHeader(http.StatusNotFound)
			return
		}

		var productDto dto.Product
		err = json.NewDecoder(request.Body).Decode(&productDto)
		if err != nil {
			response.WriteHeader(http.StatusInternalServerError)
			response.Write(jsonError(err.Error()))
			return
		}
		err = product.ChangePrice(productDto.Price)
		if err != nil {
			response.WriteHeader(http.StatusInternalServerError)
			response.Write(jsonError(err.Error()))
			return
		}

		result, err := service.Disable(product)
		if err != nil {
			response.WriteHeader(http.StatusInternalServerError)
			response.Write(jsonError(err.Error()))
			return
		}
		err = json.NewEncoder(response).Encode(result)
		if err != nil {
			response.WriteHeader(http.StatusInternalServerError)
			return
		}
	})
}
