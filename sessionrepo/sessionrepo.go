package sessionrepo

import (
	"errors"
)

var (
	ErrTokenNotFound = errors.New("token not found")
)

//go:generate mwgen -type=Repository -template=logging,metrics -outDir=.
type Repository interface {
	GetToken(userId int) (token string, err error)
	Delete(userId int) (err error)
}
