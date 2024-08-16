package domain

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

const (
	CollectionUser = "users"
)

type User struct {
	ID primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	UserName string `json:"userName" bson:"userName"`
	Email string `json:"email" bson:"email"`
	Password string `json:"password" bson:"password"`
	Verified bool `json:"verified" bson:"verified"`
	Date time.Time `json:"date" bson:"date"`
}

type UserRepository interface {
	Create(c context.Context, user *User) (primitive.ObjectID, error)
	Fetch(c context.Context) ([]*User, error)
	GetByEmail(c context.Context, email string) (*User, error)
	GetByID(c context.Context, id primitive.ObjectID) (*User, error)
	UpdateUserVerificationStatus(ctx context.Context, userID primitive.ObjectID) error
}