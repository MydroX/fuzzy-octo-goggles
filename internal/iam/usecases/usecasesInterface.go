package usecases

import "context"

type UsecasesInterface interface {
	Create(ctx context.Context, username, password, email, role string) error
	Get(ctx *context.Context)
	Update(ctx *context.Context)
	Delete(ctx *context.Context)
	Auth(ctx *context.Context)
}
