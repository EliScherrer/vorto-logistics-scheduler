package models

import (
	"time"
)

type Load struct {
	ID                  uint       `json:"id" gorm:"primaryKey"`
	PickupLatitude      float64    `json:"pickup_latitude"`
	PickupLongitude     float64    `json:"pickup_longitude"`
	DropoffLatitude     float64    `json:"dropoff_latitude"`
	DropoffLongitude    float64    `json:"dropoff_longitude"`
	Status              string     `json:"status" gorm:"default:awaiting_driver"`
	DriverStatus        string     `json:"driver_status" gorm:"default:unassigned"`
	AssignedDriverID    *uint      `json:"assigned_driver_id"`
	AssignedDriver      *Driver    `json:"assigned_driver,omitempty" gorm:"foreignKey:AssignedDriverID"`
	CreatedBy           uint       `json:"created_by"`
	Creator             User       `json:"creator" gorm:"foreignKey:CreatedBy"`
	PickupCompletedAt   *time.Time `json:"pickup_completed_at"`
	DropoffCompletedAt  *time.Time `json:"dropoff_completed_at"`
	CreatedAt           time.Time  `json:"created_at"`
	UpdatedAt           time.Time  `json:"updated_at"`
}