package gapi

import (
	"fmt"

	db "github.com/afushimi-source/simplebank/db/sqlc"
	"github.com/afushimi-source/simplebank/pb"
	"github.com/afushimi-source/simplebank/token"
	"github.com/afushimi-source/simplebank/util"
	"github.com/gin-gonic/gin"
)

type Server struct {
	pb.UnimplementedSimpleBankServer
	config     util.Config
	store      db.Store
	tokenMaker token.Maker
	router     *gin.Engine
}

func NewServer(config util.Config, store db.Store) (*Server, error) {
	tokenMaker, err := token.NewPasetoMaker(config.TokenSymmetricKey)
	if err != nil {
		return nil, fmt.Errorf("cannnot create token maker: %w", err)
	}
	server := &Server{
		config:     config,
		store:      store,
		tokenMaker: tokenMaker,
	}

	return server, nil
}
