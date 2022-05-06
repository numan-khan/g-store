package main

import (
	"errors"
	"flag"
	"fmt"
	"log"
	"net/http"

	api_handlers "github.com/numan-khan/g-store/internal/app/g-store/api-handlers"

	"github.com/numan-khan/g-store/internal/pkg/models"
	"github.com/numan-khan/g-store/internal/pkg/store"
	"github.com/spf13/viper"
)

var (
	dbLocation = flag.String("db-location", "", "Path to the database")
	httpAddr   = flag.String("http-addr", "127.0.0.1:8080", "HTTP SERVER ADDRESS")
	shard      = "shard1"
)

func parseCmdArgs() {
	flag.Parse()
	if *dbLocation == "" {
		log.Fatal("provide database location")
	}
}

func parseConfig() models.Configuration {

	// Set the file name of the configurations file
	viper.SetConfigName("config")

	// Set the path to look for the configurations file
	viper.AddConfigPath("../../configs")

	// Enable VIPER to read Environment Variables
	viper.AutomaticEnv()

	viper.SetConfigType("yml")
	var configuration models.Configuration

	if err := viper.ReadInConfig(); err != nil {
		fmt.Printf("Error reading config file, %s", err)
	}

	err := viper.Unmarshal(&configuration)
	if err != nil {
		fmt.Printf("Unable to decode into struct, %v", err)
	}

	return configuration
}

func main() {

	parseCmdArgs()

	config := parseConfig()

	myShard, err := config.GetShardByName(shard)

	if err != nil {
		errors.New("Error loadig configuration ...")
	}

	db, err := store.NewStore(*dbLocation, myShard, len(config.Shards) )

	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	storeHandler := api_handlers.NewStoreHandler(db)

	http.HandleFunc("/get", storeHandler.GetKey)
	http.HandleFunc("/set", storeHandler.SetKey)

	log.Fatal(http.ListenAndServe(*httpAddr, nil))

}
