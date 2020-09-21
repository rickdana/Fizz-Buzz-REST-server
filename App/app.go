package App

import (
	"crypto/tls"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/rickdana/fizzbuzzApi/Config"
	"github.com/rickdana/fizzbuzzApi/Controller"
	"github.com/rickdana/fizzbuzzApi/Logger"
	basicauthmiddleware "github.com/rickdana/fizzbuzzApi/Middlewares/Basicauthmiddleware"
	"github.com/rickdana/fizzbuzzApi/Repository"
	"github.com/rickdana/fizzbuzzApi/Service"
	"log"
	"net"
	"net/http"
	"strconv"
	"time"
)

type App struct {
	Router                *mux.Router
	Config                *Config.Config
	Logger                *Logger.Logger
	em                    *Repository.EntityManager
	FizzBuzzController    *Controller.FizzBuzzController
	HealthCheckController *Controller.HealthCheckController
	StatisticsController  *Controller.StatisticsController
}

var excludedPath = []string{"/api/v1/health-check"}

func (app *App) Initialize() {

	basicAuthMiddleware := basicauthmiddleware.NewBasicAuthMiddleware(app.Config.Auth.Username, app.Config.Auth.Password, excludedPath)
	app.Router = mux.NewRouter().PathPrefix("/api/v1").Subrouter().StrictSlash(true)
	app.Router.Use(basicAuthMiddleware.Middleware)

	//Entity Manager
	fileRepository := Repository.NewFileRepository(Repository.NewFileRepositoryConfig(app.Config.DataSource.FileStore))

	// Services
	fizzBuzzService := Service.NewFizzBuzzService(app.Logger)
	statisticsService := Service.NewStatisticService(app.Logger, fileRepository)

	//controller
	app.FizzBuzzController = Controller.NewFizzBuzzController(fizzBuzzService, statisticsService)
	app.HealthCheckController = &Controller.HealthCheckController{}
	app.StatisticsController = Controller.NewStatisticsController(statisticsService)

	app.setRouters()
}

func (app *App) setRouters() {
	app.Router.HandleFunc("/fizz-buzz", app.FizzBuzzController.PrintFizzBuzzer).Methods("POST")
	app.Router.HandleFunc("/health-check", app.HealthCheckController.GetState).Methods("GET")
	app.Router.HandleFunc("/statistics", app.StatisticsController.GetStatistics).Methods("GET")
}

func (app *App) Run(host string, port uint16) {
	srv := &http.Server{
		Addr: net.JoinHostPort(host, strconv.Itoa(int(port))),
		// Good practice to set timeouts to avoid Slowloris attacks.
		WriteTimeout: time.Second * 15,
		ReadTimeout:  time.Second * 15,
		IdleTimeout:  time.Second * 60,
		Handler:      app.Router,
		TLSConfig: &tls.Config{
			NextProtos: []string{"h2", "http/1.1"},
		},
	}

	fmt.Printf("Server listening on %s", srv.Addr)
	if err := srv.ListenAndServeTLS("certs/localhost.crt", "certs/localhost.key"); err != nil {
		log.Fatal(err)
	}
}
