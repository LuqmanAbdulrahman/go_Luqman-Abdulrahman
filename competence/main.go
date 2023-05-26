package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gorilla/mux"
)

// Item struct represents an electronic item
type Item struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Stock       int    `json:"stock"`
	Price       int    `json:"price"`
	Category    string `json:"category"`
}

// Sample data for demonstration
var items = []Item{
	{1, "Laptop", "Powerful laptop with high specifications", 5, 2000, "Electronics"},
	{2, "Smartphone", "Latest smartphone with advanced features", 10, 1000, "Electronics"},
}

// Custom JWT claims
type Claims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

// JWT secret key
var jwtSecret = []byte("your-secret-key")

// Logging middleware
func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%s %s %s", r.RemoteAddr, r.Method, r.URL)
		next.ServeHTTP(w, r)
	})
}

// Authentication middleware
func authenticationMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		tokenString := strings.Replace(authHeader, "Bearer ", "", 1)
		token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
			return jwtSecret, nil
		})

		if err != nil {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		claims, ok := token.Claims.(*Claims)
		if !ok || !token.Valid {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		log.Printf("Authenticated user: %s", claims.Username)
		next.ServeHTTP(w, r)
	})
}

// Handler to get all items
func getAllItems(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(items)
}

// Handler to get an item by ID
func getItemByID(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	itemID, err := strconv.Atoi(params["id"])
	if err != nil {
		http.Error(w, "Invalid item ID", http.StatusBadRequest)
		return
	}

	for _, item := range items {
		if item.ID == itemID {
			json.NewEncoder(w).Encode(item)
			return
		}
	}

	http.NotFound(w, r)
}

// Handler to add a new item
func addItem(w http.ResponseWriter, r *http.Request) {
	var item Item
	err := json.NewDecoder(r.Body).Decode(&item)
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	items = append(items, item)
	json.NewEncoder(w).Encode(item)
}

// Handler to update an item by ID
func updateItem(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	itemID, err := strconv.Atoi(params["id"])
	if err != nil {
		http.Error(w, "Invalid item ID", http.StatusBadRequest)
		return
	}

	var updatedItem Item
	err = json.NewDecoder(r.Body).Decode(&updatedItem)
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	for i := range items {
		if items[i].ID == itemID {
			items[i] = updatedItem
			json.NewEncoder(w).Encode(updatedItem)
			return
		}
	}

	http.NotFound(w, r)
}

// Handler to delete an item by ID
func deleteItem(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	itemID, err := strconv.Atoi(params["id"])
	if err != nil {
		http.Error(w, "Invalid item ID", http.StatusBadRequest)
		return
	}

	for i, item := range items {
		if item.ID == itemID {
			items = append(items[:i], items[i+1:]...)
			w.WriteHeader(http.StatusNoContent)
			return
		}
	}

	http.NotFound(w, r)
}

// Handler to get all items by category
func getItemsByCategory(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	categoryID := params["category_id"]

	var filteredItems []Item
	for _, item := range items {
		if item.Category == categoryID {
			filteredItems = append(filteredItems, item)
		}
	}

	json.NewEncoder(w).Encode(filteredItems)
}

// Handler to search items by keyword
func searchItemsByKeyword(w http.ResponseWriter, r *http.Request) {
	keyword := r.URL.Query().Get("keyword")

	var foundItems []Item
	for _, item := range items {
		if strings.Contains(strings.ToLower(item.Name), strings.ToLower(keyword)) {
			foundItems = append(foundItems, item)
		}
	}

	json.NewEncoder(w).Encode(foundItems)
}

// Handler to register a new user
func registerUser(w http.ResponseWriter, r *http.Request) {
	// In a real application, you would save the user data to a database or perform other registration-related operations.
	// This example just returns a success message.
	fmt.Fprint(w, "User registered successfully")
}

// Handler to authenticate a user and generate JWT token
func loginUser(w http.ResponseWriter, r *http.Request) {
	// In a real application, you would validate the user credentials and perform other authentication-related operations.
	// This example just generates a JWT token.
	username := r.FormValue("username")
	password := r.FormValue("password")

	// Perform authentication logic here, e.g., checking username and password against a database

	// Assume authentication is successful for demonstration purposes
	claims := &Claims{
		Username: username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtSecret)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	w.Write([]byte(tokenString))
}

func main() {
	r := mux.NewRouter()

	// Public routes
	r.HandleFunc("/register", registerUser).Methods("POST")
	r.HandleFunc("/login", loginUser).Methods("POST")

	// Protected routes
	api := r.PathPrefix("/api").Subrouter()
	api.Use(authenticationMiddleware)

	api.HandleFunc("/items", getAllItems).Methods("GET")
	api.HandleFunc("/items/{id}", getItemByID).Methods("GET")
	api.HandleFunc("/items", addItem).Methods("POST")
	api.HandleFunc("/items/{id}", updateItem).Methods("PUT")
	api.HandleFunc("/items/{id}", deleteItem).Methods("DELETE")
	api.HandleFunc("/items/category/{category_id}", getItemsByCategory).Methods("GET")
	api.HandleFunc("/items", searchItemsByKeyword).Methods("GET").Queries("keyword", "{keyword}")

	// Apply logging middleware to all routes
	r.Use(loggingMiddleware)

	log.Println("Server started on port 8000")
	log.Fatal(http.ListenAndServe(":8000", r))
}
