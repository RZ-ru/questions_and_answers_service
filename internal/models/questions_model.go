package models

import "time"

type Question struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	Text      string    `json:"text"`
	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime"`
	Answers   []Answer  `json:"answers" gorm:"constraint:OnDelete:CASCADE;"`
}
