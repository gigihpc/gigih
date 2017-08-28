package main 

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type Propinsi struct {
	Data []struct {
		KodeWilayah    string `json:"kode_wilayah"`
		MstKodeWilayah string `json:"mst_kode_wilayah"`
		Nama           string `json:"nama"`
	} `json:"data"`
}

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

func RequestCode(url string, c chan string){
	
	res, err := http.Get(url)
	if err != nil {
		log.Fatal("res1 error: ", err)
	}
	
	body,readErr :=ioutil.ReadAll(res.Body)
	if readErr != nil{
		log.Fatal("body err: ", readErr)
	}
	prop1 := new(Propinsi)
	body = bytes.TrimPrefix(body, []byte("\xef\xbb\xbf"))
	jsonErr := json.Unmarshal(body, &prop1)
	if jsonErr != nil {
		log.Fatal("jsonErr: ", jsonErr)
	}
	
	for i:=0;i<len(prop1.Data);i++{
		c <- prop1.Data[i].KodeWilayah
	}
} 

func makeRequest(c chan string){
	for{
		url := "http://jendela.data.kemdikbud.go.id/api/index.php/CcariMuseum/searchGET?kode_prop="
		kode_prop := <-c
		url += kode_prop
		
		fmt.Println(url)
	
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
		fmt.Println(data1.Data)
	}
}

func (data *Museum) saveCsv(){
	data.Data[0]
}

func main() {
	ch := make(chan string)
	url := "http://jendela.data.kemdikbud.go.id/api/index.php/CWilayah/wilayahGET"
	go RequestCode(url, ch)
	go makeRequest(ch)
	
	
//	fmt.Println(len(data1.Data))
//	fmt.Println(data1.Data[0].Kecamatan)
//	fmt.Println(data1.Data[0].Nama)
	
}
