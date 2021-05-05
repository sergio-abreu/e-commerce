package user

import (
	. "github.com/onsi/gomega"
	"github.com/sergio-vaz-abreu/discount/infrastructure/time_seed"
	"testing"
	"time"
)

func TestUser(t *testing.T) {
	g := NewGomegaWithT(t)
	t.Run("A day before user's date of birth is not its birthday", func(t *testing.T) {
		time_seed.ChangeTimeSeed(func() time.Time {
			return time.Date(2021, time.May, 03, 0, 0, 0, 0, time.UTC)
		})
		dateOfBirth := time.Date(2000, time.May, 04, 0, 0, 0, 0, time.UTC)
		sut := NewUser(dateOfBirth)

		isBirthDay := sut.IsBirthDay()

		g.Expect(isBirthDay).Should(
			BeFalse())
	})
	t.Run("A day after user's date of birth is not its birthday", func(t *testing.T) {
		time_seed.ChangeTimeSeed(func() time.Time {
			return time.Date(2021, time.May, 05, 0, 0, 0, 0, time.UTC)
		})
		dateOfBirth := time.Date(2000, time.May, 04, 0, 0, 0, 0, time.UTC)
		sut := NewUser(dateOfBirth)

		isBirthDay := sut.IsBirthDay()

		g.Expect(isBirthDay).Should(
			BeFalse())
	})
	t.Run("The same day as user's date of birth is its birthday", func(t *testing.T) {
		time_seed.ChangeTimeSeed(func() time.Time {
			return time.Date(2021, time.May, 04, 0, 0, 0, 0, time.UTC)
		})
		dateOfBirth := time.Date(2000, time.May, 04, 0, 0, 0, 0, time.UTC)
		sut := NewUser(dateOfBirth)

		isBirthDay := sut.IsBirthDay()

		g.Expect(isBirthDay).Should(
			BeTrue())
	})
}
