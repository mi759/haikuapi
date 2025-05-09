package domain

import "context"

type User struct {
	UserID        int32
	DisplayUserID string
	UserName      string
	ProfileText   string
}
// リポジトリインターフェース
type UserRepository interface {
	GetByID(ctx context.Context, id int32) (*User, error)
	Create(ctx context.Context, user *User) error
}