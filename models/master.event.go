package models

import (
	"time"

	"gorm.io/gorm"
)

type MasterEvent struct {
	gorm.Model
	EventOrder  []EventOrder `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	EventName   string       `json:"title"`
	Cover       string       `json:"cover"`
	Description string       `json:"description"`
	Date        time.Time    `json:"date"`
	StartTime   time.Time    `json:"started_at"`
	FinishTime  time.Time    `json:"finished_at"`
	Price       int          `json:"Price"`
	Speaker     string       `json:"speaker"`
	SpeakerRole string       `json:"speaker_role"`
	SpeakerComp string       `json:"speaker_com"`
}
