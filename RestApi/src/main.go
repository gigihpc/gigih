package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Mahasiswa struct {
	ID      string   `json:"id,omitempty"`
	Name    string   `json:"name,omitempty"`
	Old     string   `json:"old,omitempty"`
	Address *Address `json:"address,omitempty"`
}
type Address struct {
	City string `json:"city,omitempty"`
}

var mhsw []Mahasiswa

func GetMahasiswa(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(mhsw)
}
func GetMahasiswas(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	for _, item := range mhsw {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(&Mahasiswa{})
	log.Fatalln(mhsw)
}
func CreateMahasiswa(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var mhs Mahasiswa
	_ = json.NewDecoder(r.Body).Decode(&mhs)
	mhs.ID = params["id"]
	mhsw = append(mhsw, mhs)
	json.NewEncoder(w).Encode(mhsw)
	log.Fatalln(mhsw)
}
func DeleteMahasiswa(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	for index, item := range mhsw {
		if item.ID == params["id"] {
			mhsw = append(mhsw[:index], mhsw[index+1])
			break
		}
		json.NewEncoder(w).Encode(mhsw)
	}
	log.Fatalln(mhsw)
}

//our main function
func main() {
	//input data
	mhsw = append(mhsw, Mahasiswa{ID: "1", Name: "Haidar", Old: "12", Address: &Address{City: "Malang"}})
	mhsw = append(mhsw, Mahasiswa{ID: "2", Name: "Gigih", Old: "27", Address: &Address{City: "Malang"}})
	mhsw = append(mhsw, Mahasiswa{ID: "3", Name: "Kartika", Old: "22", Address: &Address{City: "Malang"}})
	mhsw = append(mhsw, Mahasiswa{ID: "3", Name: "Saya", Old: "22", Address: &Address{City: "Malang"}})
	//sesion, err := mgo.Dial("127.0.0.1:27017")
	//if err != nil {
	//	panic(err)
	//}
	//defer sesion.Close()
	//sesion.SetMode(mgo.Monotonic, true)

	//create route
	router := mux.NewRouter()
	router.HandleFunc("/api/mhsw", GetMahasiswa).Methods("GET")
	router.HandleFunc("/api/mhsw/{id}", GetMahasiswas).Methods("GET")
	router.HandleFunc("/api/mhsw/{id}", CreateMahasiswa).Methods("POST")
	router.HandleFunc("/api/mhsw/{id}", DeleteMahasiswa).Methods("DELETE")
	port := "8000"
	println("server on " + port)
	log.Fatal(http.ListenAndServe(":"+port, router))
}
