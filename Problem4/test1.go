package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

// type museum struct {
// 	museum_id          string `json:"museum_id"`
// 	kode_pengelolaan   string `json:"kode_pengelolaan"`
// 	nama               string `json:"nama"`
// 	sdm                string `json:"sdm"`
// 	alamat_jalan       string `json:"alamat_jalan"`
// 	desa_kelurahan     string `json:"desa_kelurahan"`
// 	kecamatan          string `json:"kecamatan"`
// 	kabupaten_kota     string `json:"kabupaten_kota"`
// 	propinsi           string `json:"propinsi"`
// 	lintang            string `json:"lintang"`
// 	bujur              string `json:"bujur"`
// 	koleksi            string `json:"koleksi"`
// 	sumber_dana        string `json:"sumber_dana"`
// 	pengelola          string `json:"pengelola"`
// 	tipe               string `json:"tipe"`
// 	standar            string `json:"standar"`
// 	tahun_berdiri      string `json:"tahun_berdiri"`
// 	bangunan           string `json:"bangunan"`
// 	luas_tanah         string `json:"luas_tanah"`
// 	status_kepemilikan interface{} `json:"status_kepemilikan"`
// }
type Museum struct {
	Data []struct {
		AlamatJalan       string      `json:"alamat_jalan"`
		Bangunan          string      `json:"bangunan"`
		Bujur             string      `json:"bujur"`
		DesaKelurahan     string      `json:"desa_kelurahan"`
		KabupatenKota     string      `json:"kabupaten_kota"`
		Kecamatan         string      `json:"kecamatan"`
		KodePengelolaan   string      `json:"kode_pengelolaan"`
		Koleksi           string      `json:"koleksi"`
		Lintang           string      `json:"lintang"`
		LuasTanah         string      `json:"luas_tanah"`
		MuseumID          string      `json:"museum_id"`
		Nama              string      `json:"nama"`
		Pengelola         string      `json:"pengelola"`
		Propinsi          string      `json:"propinsi"`
		Sdm               string      `json:"sdm"`
		Standar           string      `json:"standar"`
		StatusKepemilikan interface{} `json:"status_kepemilikan"`
		SumberDana        string      `json:"sumber_dana"`
		TahunBerdiri      string      `json:"tahun_berdiri"`
		Tipe              string      `json:"tipe"`
	} `json:"data"`
}

func main() {

	url := "http://jendela.data.kemdikbud.go.id/api/index.php/CcariMuseum/searchGET?kode_kec=016001"

	res, err := http.Get(url)
	if err != nil {
		log.Fatal("res error: ", err)
	}

	body, readErr := ioutil.ReadAll(res.Body)
	if readErr != nil {
		log.Fatal("body err: ", readErr)
	}

	// text := `{"data":[{"museum_id":"B16CC26B-D535-4CB5-BDBD-0186C8ED74B5","kode_pengelolaan":"MS000346","nama":" Museum Gunung Api Batur\/Vulkanologi","sdm":"0","alamat_jalan":"Jl. Raya Pengelokan","desa_kelurahan":"-","kecamatan":"Kintamani","kabupaten_kota":"Kab. Bangli","propinsi":"Prop. Bali","lintang":"-8.2838530","bujur":"115.3644660","koleksi":"0","sumber_dana":"-","pengelola":"Diresmikan 10 Mei 2010. Kementerian ESDM.","tipe":"Khusus","standar":"C","tahun_berdiri":"0","bangunan":"-","luas_tanah":"0","status_kepemilikan":"-"}]}`
	// bytes := []byte(text)
	data1 := new(Museum)
	//fmt.Println(string(body))
	body = bytes.TrimPrefix(body, []byte("\xef\xbb\xbf"))
	jsonErr := json.Unmarshal(body, &data1)
	if jsonErr != nil {
		log.Fatal("jsonErr: ", jsonErr)
	}
	fmt.Println(data1.Data[0].Kecamatan)
	fmt.Println(data1.Data[0].Nama)
}
