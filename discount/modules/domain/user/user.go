package user

import (
	"github.com/google/uuid"
	"github.com/sergio-vaz-abreu/discount/infrastructure/time_seed"
	"time"
)

func NewUser(dateOfBirth time.Time) User {
	return User{Id: uuid.New().String(), DateOfBirth: dateOfBirth}
}

type User struct {
	Id          string
	DateOfBirth time.Time
}

func (u User) IsBirthDay() bool {
	return time_seed.IsSameDayOfTheYear(time_seed.Now(), u.DateOfBirth)
}
