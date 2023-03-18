package main

import (
	"fmt"
	"log"

	"fake.com/medium/config"
	"fake.com/medium/db"
	"fake.com/medium/middleware"
	"fake.com/medium/routes"
	"fake.com/medium/tokenauth"
	"fake.com/medium/wired"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func main() {
	config, err := config.LoadConfig(".")
	if err != nil {
		log.Fatal("cannot load config:", err)
	}

	db := db.ConnectDataBase(config.DBDriver, config.DBSource)

	server, err := NewServer(config, db)
	if err != nil {
		log.Fatal("cannot create server:", err)
	}

	err = server.Start(config.ServerAddress)
	if err != nil {
		log.Fatal("cannot start server:", err)
	}

}

type Server struct {
	config     config.Config
	tokenMaker tokenauth.Maker
	router     *gin.Engine
	db         *gorm.DB
}

// NewServer creates a new HTTP server and set up routing.
func NewServer(config config.Config, DB *gorm.DB) (*Server, error) {
	tokenMaker, err := tokenauth.NewPasetoMaker(config.TokenSymmetricKey)

	if err != nil {
		return nil, fmt.Errorf("cannot create token maker: %w", err)
	}

	server := &Server{
		config:     config,
		tokenMaker: tokenMaker,
		db:         DB,
	}

	server.setupRouter()
	return server, nil
}

func (server *Server) setupRouter() {

	authorized := wired.InitAuthorized(server.db)
	router := gin.Default()

	router.Use(middleware.Cors2())
	authorized2 := router.Group("/")
	//inicializa servicios usuario
	apiUser := wired.InitUserAPI(server.db, server.tokenMaker, server.config.AccessTokenDuration)
	authorized2.Use(middleware.AuthMiddleware(server.tokenMaker))
	routes.UsersHandlers(router, authorized2, apiUser)

	routes.AuthHandlers(router, apiUser)
	router.Use(middleware.AuthMiddleware(server.tokenMaker))
	router.Use(middleware.AuthorizedMiddleware(authorized))

	apiRol := wired.InitRolAPI(server.db)
	apiCategoria := wired.InitCategoriaAPI(server.db)
	apiLugar := wired.InitLugarAPI(server.db)
	apiTipoDocumentos := wired.InitTipoDocumentoAPI(server.db)

	/*without authentication*/
	routes.RolHandlers(router, apiRol)
	routes.CategoriaHandlers(router, apiCategoria)
	routes.TipoTipoDocumentoHandlers(router, apiTipoDocumentos)
	routes.LugaresHandlers(router, apiLugar)

	/*with authentication */

	server.router = router

}

// Start runs the HTTP server on a specific address.
func (server *Server) Start(address string) error {
	return server.router.Run(address)
}
