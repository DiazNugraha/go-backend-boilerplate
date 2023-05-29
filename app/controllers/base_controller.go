package controllers

import (
	"backend/app/config"
	"backend/app/models"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/rs/cors"

	"github.com/urfave/cli"

	"gorm.io/gorm"
)

type Server struct {
	DB        *gorm.DB
	Router    *mux.Router
	AppConfig *AppConfig
}

type AppConfig struct {
	AppName string
	AppEnv  string
	AppPort string
	AppURL  string
}

func (s *Server) Initialize(config AppConfig) {
	s.Router = mux.NewRouter()
  s.InitializeAppConfig(config)
  s.InitializeDB()
  s.InitializeRoutes()

}

func (s *Server) InitializeCommand(config AppConfig) {
	s.InitializeAppConfig(config)
	s.InitializeDB()

	cmdApp := cli.NewApp()

	cmdApp.Commands = []cli.Command{
		{
			Name: "db:migrate",
			Action: func(cli *cli.Context) error {
				s.MigrateDB()
				return nil
			},
		},
	}

	err := cmdApp.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}

func (s *Server) InitializeAppConfig(config AppConfig) {
	s.AppConfig = &config
}

func (s *Server) InitializeDB() {
	conn := config.DBConn()
	var err error
	s.DB, err = gorm.Open(conn, &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
}

func (s *Server) MigrateDB() {
	for _, model := range models.RegisterModel() {
		err := s.DB.AutoMigrate(model.Model)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("========================Migrated Successfully========================")
	}
}

func (s *Server) Run() {
	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE"},
		AllowCredentials: true,
		AllowedHeaders:   []string{"*"},
	})
	handler := c.Handler(s.Router)
	log.Fatal(http.ListenAndServe(":"+s.AppConfig.AppPort, handler))
}
