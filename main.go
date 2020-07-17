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
	line += fmt.Sprintf("😷 Terkonfirmasi\t: %d(+%d)\n", dUT.JumlahPositif, dUP.JumlahPositif)
	line += fmt.Sprintf("🤕 Dalam Perawatan\t: %d(%.2f%% dari Terkonfirmasi)\n", dUT.JumlahDirawat, math.Round(100*float64(dUT.JumlahDirawat)/float64(dUT.JumlahPositif)))
	line += fmt.Sprintf("😊 Sembuh\t\t\t: %d(%.2f%% dari Terkonfirmasi)\n", dUT.JumlahSembuh, math.Round(100*float64(dUT.JumlahSembuh)/float64(dUT.JumlahPositif)))
	line += fmt.Sprintf("💀 Meninggal\t\t\t: %d(%.2f%% dari Terkonfirmasi)\n", dUT.JumlahMeninggal, math.Round(100*float64(dUT.JumlahMeninggal)/float64(dUT.JumlahPositif)))
	fmt.Print(line)
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
