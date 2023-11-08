package domain

import "time"

type Certificate struct {
	Participant, Course string
	DeliveredAt         time.Time
}
