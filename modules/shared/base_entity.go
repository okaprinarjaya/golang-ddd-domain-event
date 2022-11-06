package mod_shared

import "time"

const (
	NONE     = "NONE"
	NEW      = "NEW"
	MODIFIED = "MODIFIED"
	DELETED  = "DELETED"
)

type BaseEntity struct {
	Id                string
	CreatedAt         time.Time
	UpdatedAt         time.Time
	CreatedBy         string
	CreatedByName     string
	UpdatedBy         string
	UpdatedByName     string
	PersistenceStatus string
}
