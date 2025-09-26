package models

import (
	"time"
)

type Driver struct {
	ID               uint       `json:"id" gorm:"primaryKey"`
	UserID           uint       `json:"user_id"`
	User             User       `json:"user" gorm:"foreignKey:UserID"`
	CurrentLatitude  *float64   `json:"current_latitude"`
	CurrentLongitude *float64   `json:"current_longitude"`
	IsOnShift        bool       `json:"is_on_shift" gorm:"default:false"`
	ShiftStartedAt   *time.Time `json:"shift_started_at"`
	CreatedAt        time.Time  `json:"created_at"`
	UpdatedAt        time.Time  `json:"updated_at"`
}