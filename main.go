package main

import (
	"CryptoLAB/controllers/stdhttp"
	"CryptoLAB/gate/psg"
	"log"
	"net/http"
)

func main() {

	db, err := psg.NewPsg("localhost:5432", "postgres", "1234")
	if err != nil {
		log.Fatal("Error creating database connection:", err)
	}

	controller := stdhttp.NewController("localhost:8080", db)

	http.HandleFunc("/record/add", controller.RecordAdd)
	http.HandleFunc("/records/get", controller.RecordsGet)
	http.HandleFunc("/record/update", controller.RecordUpdate)
	http.HandleFunc("/record/delete", controller.RecordDeleteByPhone)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
