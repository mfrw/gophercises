package kitoc

import "context"

type Service interface {
	Hello(ctx context.Context, firstName, lastName string) (greeting string, err error)
}
