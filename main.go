package main

import (
	"adventureWorks-api/infra/database"
	"adventureWorks-api/infra/routes"
	"fmt"
	"log"

	"github.com/kardianos/service"
)

var (
	logger service.Logger
	debug  *bool
)

type program struct{}

// Start function initializes the program as a service
func (p *program) Start(s service.Service) error {
	go p.run()
	return nil
}

// Run function initializes the routes that API will serves
func (p *program) run() {

	routes.InitRoutes()
}

// Stop function stop athe program
func (p *program) Stop(s service.Service) error {
	return nil
}

func main() {
	err := database.Connect()
	if err != nil {
		fmt.Println(err)
		return
	}

	defer database.DB.Close()

	svcConfig := &service.Config{
		Name:        "Adventure Works API",
		DisplayName: "Adventure Works API",
		Description: "API running on port 8700",
	}

	prg := &program{}
	s, err := service.New(prg, svcConfig)
	if err != nil {
		log.Fatal(err)
	}
	logger, err = s.Logger(nil)
	if err != nil {
		log.Fatal(err)
	}
	err = s.Run()
	if err != nil {
		logger.Error(err)
	}

}
