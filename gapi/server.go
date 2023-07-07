package gapi

import (
	"fmt"

	db "github.com/hitesh25kumar/db/sqlc"
	"github.com/hitesh25kumar/db/util"
	"github.com/hitesh25kumar/pb"
	"github.com/hitesh25kumar/token"
)

type Server struct {
	pb.UnimplementedSimpleBankServer
	config     util.Config
	store      db.Store
	tokenMaker token.Maker
}

func NewServer(config util.Config, store db.Store) (*Server, error) {
	tokenMaker, err := token.NewPasetoMaker(config.TokenSymmetrickKey)
	if err != nil {
		return nil, fmt.Errorf("cannot create token maker: %w", err)
	}
	server := &Server{
		config:     config,
		store:      store,
		tokenMaker: tokenMaker,
	}

	return server, nil
}

// func (server *Server) setupRouter() {
// 	router := gin.Default()

// 	router.POST("/users", server.createUser)
// 	router.POST("/users/login", server.loginUser)

// 	authRoutes := router.Group("/").Use(authMiddleware(server.tokenMaker))

// 	authRoutes.POST("/accounts", server.createAccount)
// 	authRoutes.GET("/accounts/:id", server.getAccount)
// 	authRoutes.GET("/accounts", server.listAccount)

// 	authRoutes.POST("/transfers", server.createTransfer)

// 	server.router = router
// }

// func (server *Server) Start(address string) error {
// 	return server.router.Run(address)
// }

// func errorResponse(err error) gin.H {
// 	return gin.H{"error": err.Error()}
// }
