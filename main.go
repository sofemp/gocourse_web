package main

import (
	"/internal/user" // Importing the user package for endpoints
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter() // Create a new router using Gorilla Mux

	userEnd := user.MakeEndpoints() // Create user endpoints using the user package

	// Ya no se usará http.HandleFunc, sino que se usará el router. Por ello se elimina el puerto 3535 de http.HandleFunc

	//port := ":3535"
	router.HandleFunc("/users", userEnd.Create).Methods("POST")   // Handle POST requests for /users
	router.HandleFunc("/users", userEnd.GetAll).Methods("GET")    // Handle GET requests for /users
	router.HandleFunc("/users", userEnd.Update).Methods("PATCH")  // Handle PATH requests for /users
	router.HandleFunc("/users", userEnd.Delete).Methods("DELETE") // Handle DELETE requests for /users

	router.HandleFunc("/courses", getCourses).Methods("GET") // Handle GET requests for /courses

	//err := http.ListenAndServe(port, nil)
	/*
		if err != nil {
			fmt.Println("Error starting server:", err)
			return
		}
	*/
	// Start the server using the router
	server := &http.Server{
		//Handler:      http.TimeoutHandler(router, time.Second*3, "Request timed out"), // Set a timeout handler for the router
		Handler:      router,          // Use the Gorilla Mux router as the handler
		Addr:         ":8000",         // Cambiado a puerto 8000
		ReadTimeout:  5 * time.Second, // Set a read timeout of 5 seconds
		WriteTimeout: 5 * time.Second, // Set a write timeout of 5 seconds
	}

	err := server.ListenAndServe() // Start the server
	if err != nil {
		log.Fatalf("Error starting server: %v", err) // Log any errors that occur while starting the server
	}
	//log.Fatal(server.ListenAndServe()) // Start the server and log any errors

}

func getUsers(w http.ResponseWriter, r *http.Request) {
	// This function will handle the logic to get users
	//fmt.Fprintln(w, "got /users")
	//io.WriteString(w, "This is my user endpoint.\n")
	time.Sleep(10 * time.Second)                // Simulate a delay of 10 seconds
	fmt.Println("Processing /users request...") // Log the processing of the request

	json.NewEncoder(w).Encode(map[string]bool{"success": true})
}

func getCourses(w http.ResponseWriter, r *http.Request) {
	// This function will handle the logic to get courses
	//fmt.Fprintln(w, "got /courses")
	//io.WriteString(w, "This is my course endpoint.\n")
	json.NewEncoder(w).Encode(map[string]bool{"success": true})
}
