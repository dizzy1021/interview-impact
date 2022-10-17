package model

import (
	"time"

	"gorm.io/gorm"
)

type UOM string

const (
	SHEET	UOM = "sheet"
	ROLL	UOM = "roll"
	PCS		UOM = "pcs"
)

type Product struct {
	ID			string			`gorm:"primaryKey"`
	Code		string			`gorm:"size:50;index:idx_product_code,unique;not null"`
	Name		string			`gorm:"size:125;not null"`
	Description	string			`gorm:"type:text;not null"`
	Price		uint32			`gorm:"not null"`
	UOM			UOM				`gorm:"type:uom;not null"`
	CreatedAt	time.Time
	UpdatedAt	time.Time		
	DeletedAt	gorm.DeletedAt	`gorm:"index"`
}