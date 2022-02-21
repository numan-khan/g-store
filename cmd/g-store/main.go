package main

import (
	"flag"
	"log"
	"net/http"

	api_handlers "github.com/numan-khan/g-store/internal/app/g-store/api-handlers"
	"github.com/numan-khan/g-store/internal/pkg/store"
)

var (
	dbLocation = flag.String("db-location", "", "Path to the database")
	httpAddr   = flag.String("http-addr", "127.0.0.1:8080", "HTTP SERVER ADDRESS")
)

func parseCmdArgs() {
	flag.Parse()
	if *dbLocation == "" {
		log.Fatal("provide database location")
	}
}

func main() {

	parseCmdArgs()

	db, err := store.NewStore(*dbLocation)

	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	storeHandler := api_handlers.NewStoreHandler(db)

	http.HandleFunc("/get", storeHandler.GetKey)
	http.HandleFunc("/set", storeHandler.SetKey)

	log.Fatal(http.ListenAndServe(*httpAddr, nil))

}
