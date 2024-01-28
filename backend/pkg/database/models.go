package database

type Omit bool

// Model Shadows gorm.Model, overriding the json tags.
type Model struct {
	ID        Omit `json:"ID,omitempty"`
	CreatedAt Omit `json:"CreatedAt,omitempty"`
	UpdatedAt Omit `json:"UpdatedAt,omitempty"`
	DeletedAt Omit `json:"DeletedAt,omitempty"`
}
