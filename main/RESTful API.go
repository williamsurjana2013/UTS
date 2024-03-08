package main

import (
	"encoding/json"
	"net/http"
)

// User merupakan definisi dari resource user pada layanan yang disediakan
type User struct {
	ID   int    `json:"id"`
	Nama string `json:"nama"`
}

func handleUser() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			handleAllUser(w, r)
		case http.MethodPost:
			// tambahkan user baru:
		default:
			w.WriteHeader(http.StatusMethodNotAllowed)
		}
	}
}

func handleAllUser(w http.ResponseWriter, r *http.Request) {
	var users []User
	// Ambil data dari pangkalan data atau sumber data lain
	// users = service.GetAllUsers()
	// Sebagai contoh disini hanya menggunakan data in-memory
	users = []User{
		User{ID: 1, Nama: "Aconk"},
		User{ID: 2, Nama: "Mamat"},
	}
	enc := json.NewEncoder(w)
	enc.SetIndent("", "    ")
	err := enc.Encode(users)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	return
}

func main() {
	http.Handle("/user", handleUser())
	http.ListenAndServe(":8000", http.DefaultServeMux)
}

// Output dari CURL
// curl localhost:8000/user
// [
//     {
//         "id": 1,
//         "nama": "Aconk"
//     },
//     {
//        "id": 2,
//        "nama": "Mamat"
//     }
// ]
