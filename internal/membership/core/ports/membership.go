package ports

import "context"

type MembershipServicesAdapter interface {
	SubmitRegisterUser(ctx context.Context)
	GetUserInfo(ctx context.Context)
	SubmitLogin(ctx context.Context)
	SubmitLogout(ctx context.Context)
}
