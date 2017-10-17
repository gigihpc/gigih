package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Numverify struct {
	museum_id          string `json:"museum_id"`
	kode_pengelolaan   string `json:"kode_pengelolaan"`
	nama               string `json:"nama"`
	sdm                string `json:"sdm"`
	alamat_jalan       string `json:"alamat_jalan"`
	desa_kelurahan     string `json:"desa_kelurahan"`
	kecamatan          string `json:"kecamatan"`
	kabupaten_kota     string `json:"kabupaten_kota"`
	propinsi           string `json:"propinsi"`
	lintang            string `json:"lintang"`
	bujur              string `json:"bujur"`
	koleksi            string `json:"koleksi"`
	sumber_dana        string `json:"sumber_dana"`
	pengelola          string `json:"pengelola"`
	tipe               string `json:"tipe"`
	standar            string `json:"standar"`
	tahun_berdiri      string `json:"tahun_berdiri"`
	bangunan           string `json:"bangunan"`
	luas_tanah         string `json:"luas_tanah"`
	status_kepemilikan string `json:"status_kepemilikan"`
}

func main() {
	//phone := "010000"
	// QueryEscape escapes the phone string so
	// it can be safely placed inside a URL query
	//safePhone := url.QueryEscape(phone)

	url := "http://jendela.data.kemdikbud.go.id/api/index.php/CcariMuseum/searchGET?nama=Museum Gunung Api Batur/Vulkanologi"

	// Build the request
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatal("NewRequest: ", err)
		return
	}

	// For control over HTTP client headers,
	// redirect policy, and other settings,
	// create a Client
	// A Client is an HTTP client
	client := &http.Client{}

	// Send the request via a client
	// Do sends an HTTP request and
	// returns an HTTP response
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal("Do: ", err)
		return
	}

	// Callers should close resp.Body
	// when done reading from it
	// Defer the closing of the body
	defer resp.Body.Close()

	// Fill the record with the data from the JSON
	var record Numverify

	// Use json.Decode for reading streams of JSON data
	if err := json.NewDecoder(resp.Body).Decode(&record); err != nil {
		log.Println(err)
	}
	fmt.Println(resp.Body)
	fmt.Println("kabupaten_kota = ", record.kabupaten_kota)
	fmt.Println("koleksi   = ", record.koleksi)
	fmt.Println("alamat jalan  = ", record.alamat_jalan)
	fmt.Println("status kepemilikan   = ", record.status_kepemilikan)
	fmt.Println("bangunan  = ", record.bangunan)

}
