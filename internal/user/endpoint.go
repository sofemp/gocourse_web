// Creacion de enpoints para user

package user

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type (
	Controller func(w http.ResponseWriter, r *http.Request)

	Endpoints struct {
		Create Controller
		Get    Controller
		GetAll Controller
		Update Controller
		Delete Controller
	}
)

func MakeEndpoints() Endpoints {

	return Endpoints{
		Create: makeCreateEndpoint(),
		Get:    makeGetEndpoint(),
		GetAll: makeGetAllEndpoint(),
		Update: makeUpdateEndpoint(),
		Delete: makeDeleteEndpoint(),
	}

}

func makeCreateEndpoint() Controller {
	return func(w http.ResponseWriter, r *http.Request) {
		// Logic for creating a user
		fmt.Println("Create user endpoint called")
		json.NewEncoder(w).Encode(map[string]bool{"success": true})
	}
}

func makeGetEndpoint() Controller {
	return func(w http.ResponseWriter, r *http.Request) {
		// Logic for getting a user
		fmt.Println("Get user endpoint called")
		json.NewEncoder(w).Encode(map[string]bool{"success": true})
	}
}

func makeGetAllEndpoint() Controller {
	return func(w http.ResponseWriter, r *http.Request) {
		// Logic for getting all users
		fmt.Println("Get all users endpoint called")
		json.NewEncoder(w).Encode(map[string]bool{"success": true})
	}
}

func makeUpdateEndpoint() Controller {
	return func(w http.ResponseWriter, r *http.Request) {
		// Logic for updating a user
		fmt.Println("Update user endpoint called")
		json.NewEncoder(w).Encode(map[string]bool{"success": true})
	}
}

func makeDeleteEndpoint() Controller {
	return func(w http.ResponseWriter, r *http.Request) {
		// Logic for deleting a user
		fmt.Println("Delete user endpoint called")
		json.NewEncoder(w).Encode(map[string]bool{"success": true})
	}
}
