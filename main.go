package main

import (
	"context"
	"errors"
	"fmt"
	"microgen/sessionrepo"
)

// @microgen middleware,logging,metrics,http,error-logging
type Service interface {
	// @http-method GET
	// @http-path jrss/api/session/validate
	Validate(ctx context.Context, userId int, token string) (err error)

	// @http-method DELETE
	// @http-path jrss/api/session
	Delete(ctx context.Context, userId int, token string) (err error)
}
type service struct {
	sessionRepo sessionrepo.Repository
}

func NewService(sessionRepo sessionrepo.Repository) *service {
	return &service{
		sessionRepo: sessionRepo,
	}
}

func (svc *service) Validate(ctx context.Context, userId int, token string) (err error) {
	if userId < 1 || len(token) < 1 {
		return err
	}

	srcToken, err := svc.sessionRepo.GetToken(userId)
	switch {
	case err == nil:
	case errors.Is(err, sessionrepo.ErrTokenNotFound):
		return err
	default:
		return err
	}

	if srcToken != token {
		return err
	}

	return nil
}

func (svc *service) Delete(ctx context.Context, userId int, token string) (err error) {
	err = svc.Validate(ctx, userId, token)
	if err != nil {
		return
	}
	return svc.sessionRepo.Delete(userId)
}
func main() {
	fmt.Println("Hello Megha")
}

// microgen -file=./main.go -out=. -package=microgen/main
