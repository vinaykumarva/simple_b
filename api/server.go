package api

import (
	"github.com/gin-gonic/gin"
	db "github.com/techschool/simplebank/db/sqlc"
)

//This Server will serves all HTTP requests for our banking service
type Server struct {
	//It will allow us to interact with the database when processing API requests from clients.
	store *db.Store
	//gin.Engine. This router will help us send each API request to the correct handler for processing.
	router *gin.Engine
}

//First, we create a new Server object with the input store
func NewServer(store *db.Store) *Server {
	server := &Server{store: store, router: gin.Default()}
	//router := gin.Default()

	// TODO: add routes to router

	server.router.POST("/accounts", server.createAccount)

	//server.router = router
	return server
}

func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
