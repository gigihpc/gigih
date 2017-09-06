package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// Mahasiswa Struct
type Mahasiswa struct {
	Name    string `json:"name"`
	Address string `json:"address"`
	Old     string `json:"old"`
}

var mhsw = map[string]*Mahasiswa{
	"1": &Mahasiswa{Name: "Gigih", Address: "Malang", Old: "27"},
	"2": &Mahasiswa{Name: "Kartika", Address: "Malang", Old: "27"},
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/api/mhsw", handleMhsw).Methods("GET")
	router.HandleFunc("/api/mhsw/{id}", handleMahasiswa).Methods("GET", "DELETE", "POST")
	port := "8001"
	println("open port: " + port)
	http.ListenAndServe(":"+port, router)
}

func handleMahasiswa(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(req)
	id := vars["id"]

	switch req.Method {
	case "GET":
		mhsw, ok := mhsw[id]
		if !ok {
			res.WriteHeader(http.StatusNotFound)
			fmt.Fprint(res, string("Student Not Found"))
		}
		outgoingJSON, error := json.Marshal(mhsw)
		if error != nil {
			log.Println(error.Error())
			http.Error(res, error.Error(), http.StatusInternalServerError)
			return
		}
		fmt.Fprint(res, string(outgoingJSON))
	case "DELETE":
		delete(mhsw, id)
		res.WriteHeader(http.StatusNoContent)
	case "POST":
		mhsw1 := new(Mahasiswa)
		decoder := json.NewDecoder(req.Body)
		error := decoder.Decode(&mhsw1)
		if error != nil {
			log.Println(error.Error())
			http.Error(res, error.Error(), http.StatusInternalServerError)
			return
		}
		mhsw[id] = mhsw1
		outgoingJSON, err := json.Marshal(mhsw1)
		if err != nil {
			log.Println(error.Error())
			http.Error(res, err.Error(), http.StatusInternalServerError)
			return
		}
		res.WriteHeader(http.StatusCreated)
		fmt.Fprint(res, string(outgoingJSON))
	}
}

func handleMhsw(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "application/json")
	outgoingJSON, error := json.Marshal(mhsw)
	if error != nil {
		log.Println(error.Error())
		http.Error(res, error.Error(), http.StatusInternalServerError)
		return
	}
	fmt.Fprint(res, string(outgoingJSON))
}
