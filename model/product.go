package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type UOM string

const (
	SHEET	UOM = "sheet"
	ROLL	UOM = "roll"
	PCS		UOM = "pcs"
)

type Product struct {
	ID			string			`gorm:"primaryKey;type:uuid;column:id" json:"product_id"`
	Code		string			`gorm:"size:50;index:idx_product_code,unique;not null;column:code" json:"product_code" binding:"required,min=3,max=50"`
	Name		string			`gorm:"size:125;not null;column:name" json:"product_name" binding:"required,min=3,max=125"`
	Description	string			`gorm:"type:text;not null;column:description" json:"product_desc" binding:"required,min=10"`
	Price		uint32			`gorm:"not null;column:price" json:"product_price" binding:"required,min=0,max=4294967295"`
	UOM			UOM				`gorm:"type:uom;not null;column:uom" json:"product_uom" binding:"required,oneof='sheet' 'roll' 'pcs'"`
	CreatedAt	time.Time		`json:"created_at"`
	UpdatedAt	time.Time		`json:"updated_at"`
	DeletedAt	gorm.DeletedAt	`gorm:"index"`
}

func (m *Product) BeforeCreate(tx *gorm.DB) (err error) {
	m.ID = uuid.NewString()
	return
  }
