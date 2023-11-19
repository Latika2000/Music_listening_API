package main

import (
	"crypto/rand"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

/********** GENERATE SECRET CODE **************/

func generateSecretCode() string {
	randomBytes := make([]byte, 16)
	rand.Read(randomBytes)
	return hex.EncodeToString(randomBytes)
}

type User struct {
	ID         string     `json:"id"`
	Name       string     `json:"name"`
	Email      string     `json:"email"`
	Secretcode string     `json:"secretcode"`
	Playlists  []Playlist `json:"playlists"`
}

type Playlist struct {
	ID    string `json:"id"`
	Title string `json:"title"`
	Songs []Song `json:"songs"`
}

type Song struct {
	Title string `json:"title"`
}

var users []User

/***************************************************** TO LOGIN ************************************************/

func login(f http.ResponseWriter, r *http.Request) {

	f.Header().Set("Content-Type", "application/json")
	var userCreds struct {
		Secretcode string `json:"secret_code"`
	}

	err := json.NewDecoder(r.Body).Decode(&userCreds)
	if err != nil {
		f.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(f, "Invalid request payload")
		return
	}

	for _, user := range users {
		if user.Secretcode == userCreds.Secretcode {
			json.NewEncoder(f).Encode(user)
			return
		}
	}

	f.WriteHeader(http.StatusNotFound)
	fmt.Fprintf(f, "User not found")

}

/******************************************* TO REGISTER*****************************************************/

func register(f http.ResponseWriter, r *http.Request) {
	f.Header().Set("Content-Type", "application/json")
	var newUser User

	err := json.NewDecoder(r.Body).Decode(&newUser)
	if err != nil {
		f.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(f, "Invalid request payload")
		return
	}

	newUser.Secretcode = generateSecretCode()
	json.NewEncoder(f).Encode(newUser)

}

/********************************************************* TO VIEW PROFILE *************************************/

func viewProfile(f http.ResponseWriter, r *http.Request) {

	f.Header().Set("Content-Type", "application/json")
	var currentUser User
	userID := User.Secretcode

	for _, user := range users {
		if user.ID == userID {
			currentUser = user
			break
		}
	}

	json.NewEncoder(f).Encode(currentUser)

}

func addSongToPlaylist(f http.ResponseWriter, r *http.Request) {

}

func getAllSongsOfPlaylist(f http.ResponseWriter, r *http.Request) {

}

func createPlaylist(f http.ResponseWriter, r *http.Request) {
	f.Header().Set("Content-Type", "application/json")

}

func deletePlaylist(f http.ResponseWriter, r *http.Request) {

}

func getSongDetail(f http.ResponseWriter, r *http.Request) {

}

func deleteSongFromPlaylist(f http.ResponseWriter, r *http.Request) {

}

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/login", login).Methods("POST")
	router.HandleFunc("/register", register).Methods("POST")
	router.HandleFunc("/viewProfile", viewProfile).Methods("GET")
	router.HandleFunc("/getAllSongsOfPlaylist/{playlistID}", getAllSongsOfPlaylist).Methods("GET")
	router.HandleFunc("/createPlaylist", createPlaylist).Methods("POST")
	router.HandleFunc("/addSongToPlaylist/{playlistID}", addSongToPlaylist).Methods("POST")
	router.HandleFunc("/deletePlaylist/{playlistID}", deletePlaylist).Methods("DELETE")
	router.HandleFunc("/deleteSongFromPlaylist/{playlistID}/{songID}", deleteSongFromPlaylist).Methods("DELETE")
	router.HandleFunc("/getSongDetail/{playlistID}/{songID}", getSongDetail).Methods("GET")

	http.ListenAndServe(":6010", router)

}
