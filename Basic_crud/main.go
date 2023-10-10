package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// making a struct that has fields(details of an album) needed to be stored as data
type Album struct {
	ID     string `json:"id"`
	Title  string `json:"title"`
	Artist string `json:"artist"`
	Price  string `json:"price"`
}

// making a slice that already have some data in it to work with
var albums = []Album{
	{ID: "1", Title: "MMLP", Artist: "Eminem", Price: "$300"},
	{ID: "2", Title: "MMLP2", Artist: "Eminem", Price: "$400"},
	{ID: "3", Title: "Kamikaze", Artist: "Eminem", Price: "$500"},
	{ID: "4", Title: "Music to be murdered by", Artist: "Eminem", Price: "$600"},
}

var Port = 3000

func main() {

	fmt.Println("Server is running at port", Port)

	r := mux.NewRouter()
	r.HandleFunc("/albums", getAlbums).Methods("GET")
	r.HandleFunc("/albums/{id}", getAlbumById).Methods("GET")
	r.HandleFunc("/albums", postAlbums).Methods("POST")
	r.HandleFunc("/albums/{id}", deleteAlbumById).Methods("DELETE")
	r.HandleFunc("/albums/{id}", updateAlbum).Methods("PUT")

	http.ListenAndServe(":3000", r)
}

func getAlbums(w http.ResponseWriter, r *http.Request) {
	//setting the content type to json
	w.Header().Set("Content-Type", "application/json")

	//encoding the list of albums as json and sending it as the response
	json.NewEncoder(w).Encode(albums)
	// if err := json.NewEncoder(w).Encode(&albums); err != nil {
	//  http.Error(w, err.Error(), http.StatusInternalServerError)
	//  return
	// }
}

func getAlbumById(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r) //created a variable to store id, here r is request object which retrieves route parameters with the help of m (in this case it is id)
	for _, album := range albums {
		if album.ID == params["id"] {
			json.NewEncoder(w).Encode(album)
			return
		}

	}
}

func postAlbums(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var new_album Album
	//decode is used to create entries(sending data to server)
	_ = json.NewDecoder(r.Body).Decode(&new_album)
	new_album.ID = strconv.Itoa(rand.Intn(1000000000))

	albums = append(albums, new_album)
	//encode is used to get data from server
	json.NewEncoder(w).Encode(albums)

}

func deleteAlbumById(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var found bool
	//first i need the id of album i wanna delete
	params := mux.Vars(r)
	//then iterate through the albums slice to match the id
	for index, album := range albums {
		if album.ID == params["id"] {
			//delete the album by replacing
			albums = append(albums[:index], albums[index+1:]...)
			found = true
			break
		}

	}
	if found {
		// If the album was found and deleted, send a success message.
		response := map[string]string{"message": "Album deleted successfully"}
		json.NewEncoder(w).Encode(response)
		fmt.Println("Album deleted successfully!!")
	} else {
		// If the album was not found, send an error message.
		w.WriteHeader(http.StatusNotFound)
		response := map[string]string{"error": "Album not found"}
		json.NewEncoder(w).Encode(response)
		fmt.Println("Album not found")
	}

}

func updateAlbum(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	//first get the album you wanna delete by id
	params := mux.Vars(r)
	for index, album := range albums {
		if album.ID == params["id"] {
			//first deleted the album and now its place is empty
			albums = append(albums[:index], albums[index+1:]...)
			//now add the updated album
			var album Album
			_ = json.NewDecoder(r.Body).Decode(&album)
			album.ID = params["id"]
			albums = append(albums, album)
			json.NewEncoder(w).Encode(album)
			return

		}

	}
	json.NewEncoder(w).Encode(albums)
	// response := map[string]string{"message": "Album updted successfully"}
	// json.NewEncoder(w).Encode(response)
	// fmt.Println("Album updated successfully!!")

}
