package shared

import "context"

type ContractReader interface {
	Read(ctx context.Context, req, dest any) error
}
