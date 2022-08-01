package domain

import "github.com/eznd-otus-msa/hw3/app/pkg/types"

type UserId int64

func (t UserId) Validate() error {
	if t <= 0 {
		return ErrInvalidUserId
	}
	return nil
}

type User struct {
	Id        UserId
	Username  string
	FirstName string
	LastName  string
	Email     string
	Phone     string
}

type UserPartialData = types.Kv
