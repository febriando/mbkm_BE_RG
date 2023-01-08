package main

import (
	"fmt"
	"net/http"
	"time"
)

func GetHandler() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		now := time.Now()
		hari := now.Weekday()
		tanggal := now.Day()
		bulan := now.Month()
		tahun := now.Year()
		result := fmt.Sprintf("%s, %02d %s %d", hari, tanggal, bulan, tahun)
		writer.Write([]byte(result))
	} // TODO: replace this
}

func main() {
	err := http.ListenAndServe("localhost:8080", GetHandler())
	if err != nil {
		fmt.Println(err)
	}
}
