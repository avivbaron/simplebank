package gapi

import (
	"fmt"

	db "github.com/avivbaron/simplebank/db/sqlc"
	"github.com/avivbaron/simplebank/pb"
	"github.com/avivbaron/simplebank/token"
	"github.com/avivbaron/simplebank/util"
)

// Servert serves gRPC requrests fir out banking service
type Server struct {
	pb.UnimplementedSimpleBankServer
	config     util.Config
	store      db.Store
	tokenMaker token.Maker
}


func NewServer(cfg util.Config, store db.Store) (*Server, error) {
	tokenMaker, err := token.NewPasetoMaker(cfg.TokenSymmetricKey)
	if err != nil {
		return nil, fmt.Errorf("cannot create token maker: %w", err)
	}

	server := &Server{
		store:      store,
		config:     cfg,
		tokenMaker: tokenMaker,
	}

	return server, nil
}
