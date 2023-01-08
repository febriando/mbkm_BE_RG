package main

import (
	"fmt"
	"strconv"
	"strings"
)

// TODO: answer here

func DeliveryOrder(data []string, day string) map[string]float32 {

	var splitData []string
	var makeMap = make(map[string]float32)
	var biayaAdmin float32
	var harga float64
	var floatHarga float32

	lengthData := len(data)

	for i := 0; i < lengthData; i++ {
		splitData = strings.Split(data[i], ":")

		harga, _ = strconv.ParseFloat(string(splitData[2]), 32)
		floatHarga = float32(harga)

		switch day {
		case "senin":
			if string(splitData[3]) == "JKT" || string(splitData[3]) == "DPK" {
				biayaAdmin = floatHarga + (0.1 * floatHarga)
				makeMap[splitData[0]+"-"+splitData[1]] = biayaAdmin
			} else {
				continue
			}
		case "selasa":
			if string(splitData[3]) == "JKT" || string(splitData[3]) == "BKS" || string(splitData[3]) == "DPK" {
				biayaAdmin = floatHarga + (0.05 * floatHarga)
				makeMap[splitData[0]+"-"+splitData[1]] = biayaAdmin
			} else {
				continue
			}
		case "rabu":
			if string(splitData[3]) == "JKT" || string(splitData[3]) == "BDG" {
				biayaAdmin = floatHarga + (0.1 * floatHarga)
				makeMap[splitData[0]+"-"+splitData[1]] = biayaAdmin
			} else {
				continue
			}
		case "kamis":
			if string(splitData[3]) == "JKT" || string(splitData[3]) == "BDG" || string(splitData[3]) == "BKS" {
				biayaAdmin = floatHarga + (0.05 * floatHarga)
				makeMap[splitData[0]+"-"+splitData[1]] = biayaAdmin
			} else {
				continue
			}
		case "jumat":
			if string(splitData[3]) == "JKT" || string(splitData[3]) == "BKS" {
				biayaAdmin = floatHarga + (0.1 * floatHarga)
				makeMap[splitData[0]+"-"+splitData[1]] = biayaAdmin
			} else {
				continue
			}
		case "sabtu":
			if string(splitData[3]) == "JKT" || string(splitData[3]) == "BDG" {
				biayaAdmin = floatHarga + (0.05 * floatHarga)
				makeMap[splitData[0]+"-"+splitData[1]] = biayaAdmin
			} else {
				continue
			}
		case "minggu":
			return makeMap
		}

	}

	return makeMap // TODO: replace this
}

func main() {
	data := []string{"Budi:Gunawan:10000:JKT", "Andi:Sukirman:20000:JKT", "Budi:Sukirman:30000:BDG", "Andi:Gunawan:40000:BKS", "Budi:Gunawan:50000:DPK"}

	day := "minggu"

	deliveryData := DeliveryOrder(data, day)

	fmt.Println(deliveryData)
}
