package time_seed

import "time"

var defaultTime = time.Now

func Now() time.Time {
	return defaultTime()
}

type Seed func() time.Time

func ChangeTimeSeed(seed Seed) {
	defaultTime = seed
}

func IsSameDayOfTheYear(first time.Time, second time.Time) bool {
	first = first.UTC()
	second = second.UTC()
	return first.Day() == second.Day() && first.Month() == second.Month()
}
