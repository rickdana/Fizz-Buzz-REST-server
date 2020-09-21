package main

import (
	"flag"
	"fmt"
	"github.com/rickdana/fizzbuzzApi/App"
	"github.com/rickdana/fizzbuzzApi/Config"
	"github.com/rickdana/fizzbuzzApi/Logger"
	"github.com/rickdana/fizzbuzzApi/Repository"
	"gopkg.in/yaml.v3"
	"log"
	"os"
)

var (
	warningLogger *log.Logger
	infoLogger    *log.Logger
	errorLogger   *log.Logger
)
var defaultConfigPath = "./config.yml"

func main() {

	configPathPtr := flag.String("configPath", defaultConfigPath, "Absolute path of the config file (yml)")

	flag.Parse()

	var cfg Config.Config

	readFile(*configPathPtr, &cfg)

	app := &App.App{
		Config: &cfg,
		Logger: initLogger(cfg.Logger.FilePath),
	}
	app.Initialize()
	app.Run(app.Config.Server.Host, app.Config.Server.Port)
}

func initLogger(logFilePath string) *Logger.Logger {
	Repository.CreateFile(logFilePath)
	file, err := os.OpenFile(logFilePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal(err)
	}

	infoLogger = log.New(file, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
	warningLogger = log.New(file, "WARNING: ", log.Ldate|log.Ltime|log.Lshortfile)
	errorLogger = log.New(file, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)
	return Logger.NewFZLogger(infoLogger, warningLogger, errorLogger)
}

func readFile(configPath string, cfg *Config.Config) {
	f, err := os.Open(configPath)
	if err != nil {
		processError(err)
	}
	defer f.Close()

	decoder := yaml.NewDecoder(f)
	err = decoder.Decode(cfg)
	if err != nil {
		processError(err)
	}
}

func processError(err error) {
	fmt.Println(err)
	os.Exit(2)
}
