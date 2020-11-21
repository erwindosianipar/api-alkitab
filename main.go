package main

import (
	"fmt"
	"net/http"
	"os"

	rootHandler "api-alkitab/root/handler"

	passageHandler "api-alkitab/passage/handler"
	passageRepository "api-alkitab/passage/repository"
	passageService "api-alkitab/passage/service"

	listHandler "api-alkitab/list/handler"
	listRepository "api-alkitab/list/repository"
	listService "api-alkitab/list/service"

	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func init() {
	viper.SetConfigFile("config.json")
	err := viper.ReadInConfig()
	if err != nil {
		logrus.Fatal(err)
	}
}

func setupServerAddress() (string, error) {
	port := os.Getenv("PORT")
	if port == "" {
		port := ":9000"
		return port, fmt.Errorf("$PORT not set defaulting to " + port)
	}
	return ":" + port, nil
}

func main() {
	port, err := setupServerAddress()
	if err != nil {
		logrus.Error("[main][Error][setupServerAddress] ", err)
	}

	router := mux.NewRouter().StrictSlash(true)

	rootHandler.SetupHandler(router)

	passageBaseURL := viper.GetString("url.passage")
	passageRepository := passageRepository.CreatePassageRepository(passageBaseURL)
	passageService := passageService.CreatePassageService(passageRepository)
	passageHandler.SetupHandler(router, passageService)

	listCSVFile := viper.GetString("file.list_passage_csv")
	listRepository := listRepository.CreateListRepository(listCSVFile)
	listService := listService.CreateListService(listRepository)
	listHandler.SetupHandler(router, listService)

	logrus.Info("Starting web server at ", port)
	logrus.Fatal(http.ListenAndServe(port, router))
}
