package service

import "time"

type Service struct{}

func New() *Service {
	return &Service{}
}

func (svc *Service) DaysFromNow() int {
	year := 2025
	month := time.January
	day := 1

	then := time.Date(year, month, day, 0, 0, 0, 0, time.UTC)

	days := time.Until(then).Hours() / 24

	return int(days)
}
