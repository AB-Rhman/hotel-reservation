package types

import (
	"fmt"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
	"regexp"
)

const (
	bcryptCost      = 12
	minFirstNameLen = 2
	minLastNameLen  = 2
	minPasswordlen  = 7
)

type CreateUserParams struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Email     string `json:"email"`
	Password  string `json:"password"`
}

func (params CreateUserParams) Validate() []string {
	var errs []string

	if len(params.FirstName) < minFirstNameLen {
		errs = append(errs, fmt.Sprintf("firstName must be at least %d characters", minFirstNameLen))
	}
	if len(params.LastName) < minLastNameLen {
		errs = append(errs, fmt.Sprintf("lastName must be at least %d characters", minLastNameLen))
	}
	if len(params.Password) < minPasswordlen {
		errs = append(errs, fmt.Sprintf("password must be at least %d characters", minPasswordlen))
	}
	if !isEmailValid(params.Email) {
		errs = append(errs, fmt.Sprintf("email must be a valid email address"))
	}
	if len(errs) > 0 {
		return errs
	}
	return nil
}

func isEmailValid(email string) bool {
	emailRegex := regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`)
	return emailRegex.MatchString(email)
}

type User struct {
	ID                primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	FirstName         string             `bson:"firstName" json:"firstName"`
	LastName          string             `bson:"lastName" json:"lastName"`
	Email             string             `bson:"email" json:"email"`
	EncryptedPassword string             `bson:"encryptedPassword" json:"-"`
}

func NewUserFromParams(params CreateUserParams) (*User, error) {
	encPass, err := bcrypt.GenerateFromPassword([]byte(params.Password), bcryptCost)
	if err != nil {
		return nil, err
	}
	return &User{
		FirstName:         params.FirstName,
		LastName:          params.LastName,
		Email:             params.Email,
		EncryptedPassword: string(encPass),
	}, nil
}
