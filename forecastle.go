package main

import (
	"fmt"
	"net/http"

	"github.com/gobuffalo/packr"
	"github.com/gorilla/mux"
	"github.com/spf13/viper"
	"github.com/stakater/Forecastle/pkg/handlers"
	"github.com/stakater/Forecastle/pkg/log"
)

var (
	logger = log.New()
)

func init() {
	viper.SetConfigName("config")            // name of config file (without extension)
	viper.AddConfigPath("/etc/forecastle/")  // path to look for the config file in
	viper.AddConfigPath("$HOME/.forecastle") // call multiple times to add many search paths
	viper.AddConfigPath(".")                 // optionally look for config in the working directory
	viper.AutomaticEnv()
	err := viper.ReadInConfig() // Find and read the config file
	if err != nil {             // Handle errors reading the config file
		panic(fmt.Errorf("Fatal error config file: %s", err))
	}
}

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/apps", handlers.AppsHandler)
	router.HandleFunc("/config", handlers.ConfigHandler)

	// Serve static files using packr
	box := packr.NewBox("./static")
	router.PathPrefix("/").Handler(http.StripPrefix("/", http.FileServer(box)))

	logger.Info("Listening at port 3000")
	if err := http.ListenAndServe(":3000", router); err != nil {
		logger.Error(err)
	}
}
