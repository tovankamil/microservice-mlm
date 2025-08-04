package repository

import (
	"e-wallet-microservices/lib/config"

	"github.com/go-redis/redis/v8"
)

type MembershipRepository struct {
	config *config.Config
	client *redis.Client
}

// func NewRepository(cfg *config.Config, client *redis.Client) ports.MembershipServicesAdapter {
// 	return &MembershipRepository{
// 		config: cfg,
// 		client: client,
// 	}
// }
