package models

import (
	uuid "github.com/satori/go.uuid"
)

// RelatedTerm - table name: related_terms
type RelatedTerm struct {
	ID        uuid.UUID `gorm:"type:uuid;primary_key;" json:"id"`
	Term      Term      `json:"-"`
	TermID    uuid.UUID `gorm:"type:uuid;not null" json:"term_id"`
	Related   Term      `json:"-"`
	RelatedID uuid.UUID `gorm:"type:uuid;not null" json:"related_id"`
}
