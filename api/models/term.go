package models

import (
	"time"

	"github.com/jinzhu/gorm"
	uuid "github.com/satori/go.uuid"
)

// Term - table name: terms
type Term struct {
	ID           uuid.UUID     `gorm:"type:uuid;primary_key;" json:"id"`
	Name         string        `gorm:"size:255;not null;unique" json:"name"`
	Pronuciation string        `gorm:"size:255;" json:"pronuciation"`
	Definition   string        `gorm:"type:text;" json:"definition"`
	Synonyms     string        `gorm:"type:text;" json:"synonyms"`
	RelatedTerms []RelatedTerm `json:"related_terms"`
	CreatedByID  uuid.UUID     `gorm:"type:uuid;not null" json:"created_by_id"`
	CreatedAt    time.Time     `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedByID  uuid.UUID     `gorm:"type:uuid;not null" json:"updated_by_id"`
	UpdatedAt    time.Time     `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}

// BeforeCreate - create new uuid
func (term *Term) BeforeCreate(tx *gorm.DB) (err error) {
	term.ID = uuid.NewV4()
	term.UpdatedByID = term.CreatedByID
	if len(term.RelatedTerms) > 0 {
		for index := range term.RelatedTerms {
			term.RelatedTerms[index].ID = uuid.NewV4()
			term.RelatedTerms[index].TermID = term.ID
		}
	}
	return
}

// BeforeSave = update updatedAt
func (term *Term) BeforeSave(tx *gorm.DB) (err error) {
	term.UpdatedAt = time.Now()
	return
}
