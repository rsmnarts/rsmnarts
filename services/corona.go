package services

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

type resCovid19GoId struct {
	Data struct {
		ID                   int `json:"id"`
		JumlahOdp            int `json:"jumlah_odp"`
		JumlahPdp            int `json:"jumlah_pdp"`
		TotalSpesimen        int `json:"total_spesimen"`
		TotalSpesimenNegatif int `json:"total_spesimen_negatif"`
	} `json:"data"`
	Update struct {
		Penambahan struct {
			JumlahPositif   int    `json:"jumlah_positif"`
			JumlahMeninggal int    `json:"jumlah_meninggal"`
			JumlahSembuh    int    `json:"jumlah_sembuh"`
			JumlahDirawat   int    `json:"jumlah_dirawat"`
			Tanggal         string `json:"tanggal"`
			Created         string `json:"created"`
		} `json:"penambahan"`
		Harian []struct {
			KeyAsString     time.Time `json:"key_as_string"`
			Key             int64     `json:"key"`
			DocCount        int       `json:"doc_count"`
			JumlahMeninggal struct {
				Value int `json:"value"`
			} `json:"jumlah_meninggal"`
			JumlahSembuh struct {
				Value int `json:"value"`
			} `json:"jumlah_sembuh"`
			JumlahPositif struct {
				Value int `json:"value"`
			} `json:"jumlah_positif"`
			JumlahDirawat struct {
				Value int `json:"value"`
			} `json:"jumlah_dirawat"`
			JumlahPositifKum struct {
				Value int `json:"value"`
			} `json:"jumlah_positif_kum"`
			JumlahSembuhKum struct {
				Value int `json:"value"`
			} `json:"jumlah_sembuh_kum"`
			JumlahMeninggalKum struct {
				Value int `json:"value"`
			} `json:"jumlah_meninggal_kum"`
			JumlahDirawatKum struct {
				Value int `json:"value"`
			} `json:"jumlah_dirawat_kum"`
		} `json:"harian"`
		Total struct {
			JumlahPositif   int `json:"jumlah_positif"`
			JumlahDirawat   int `json:"jumlah_dirawat"`
			JumlahSembuh    int `json:"jumlah_sembuh"`
			JumlahMeninggal int `json:"jumlah_meninggal"`
		} `json:"total"`
	} `json:"update"`
}

func GrabDataCorona() *resCovid19GoId {
	resp, err := http.Get("https://data.covid19.go.id/public/api/update.json")
	if err != nil {
		log.Println(err)
		return nil
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
		return nil
	}

	var jsonResCovid19GoId resCovid19GoId
	if err = json.Unmarshal(body, &jsonResCovid19GoId); err != nil {
		log.Println(err)
		return nil
	}

	return &jsonResCovid19GoId
}
