package main

import (
	"fmt"
	"log"
	"math"

	"github.com/rsmnarts/rsmnarts/services"
)

func main() {
	dataCorona := services.GrabDataCorona()
	if dataCorona == nil {
		log.Println("Data Corona Kosong")
	}

	dUT := dataCorona.Update.Total
	dUP := dataCorona.Update.Penambahan

	line := fmt.Sprintf("Tanggal data terakhir\t\t: %s\n", dUP.Created)
	line += fmt.Sprintf("😷 Terkonfirmasi\t: %s(+%s)\n", formatNumber(dUT.JumlahPositif), formatNumber(dUP.JumlahPositif))
	line += fmt.Sprintf("🤕 Dalam Perawat)an\t: %s(%.2f%% dari Terkonfirmasi)\n", formatNumber(dUT.JumlahDirawat), math.Round(100*float64(dUT.JumlahDirawat)/float64(dUT.JumlahPositif)))
	line += fmt.Sprintf("😊 Sembuh\t\t\t: %s(%.2f%% dari Terkonfirmasi)\n", formatNumber(dUT.JumlahSembuh), math.Round(100*float64(dUT.JumlahSembuh)/float64(dUT.JumlahPositif)))
	line += fmt.Sprintf("💀 Meninggal\t\t\t: %s(%.2f%% dari Terkonfirmasi)\n", formatNumber(dUT.JumlahMeninggal), math.Round(100*float64(dUT.JumlahMeninggal)/float64(dUT.JumlahPositif)))

	payloadGistCorona := map[string]interface{}{
		"description": "Data Corona Indonesia",
		"files": map[string]interface{}{
			"Data Corona Indonesia": map[string]interface{}{
				"content": line,
			},
		},
	}

	services.UpdateGist("94e1b3fe281a4daeefac41be87f48e8c", payloadGistCorona)
}
