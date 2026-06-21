package storage

import "sc/go13/types"

type Storage interface {
	Get(int) *types.User
}
