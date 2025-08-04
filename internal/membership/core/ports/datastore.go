package ports

import "context"

type CacheRepositoryAdapater interface {
	GetUserSessionFromCache(ctx context.Context)
	UpdateUserSessionIntoCache(ctx context.Context)
}

type DatabaseRepositoryAdapter interface {
	InsertUserIntoDb(ctx context.Context)
	GetUserFromDb(ctx context.Context)
	GetUserByUsername(ctx context.Context)
}
