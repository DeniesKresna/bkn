package models

import "time"

type Role struct {
	ID        int64      `json:"id" db:"id"`
	CreatedBy string     `json:"created_by" db:"created_by"`
	CreatedAt time.Time  `json:"created_at" db:"created_at"`
	UpdatedBy string     `json:"updated_by" db:"updated_by"`
	UpdatedAt time.Time  `json:"updated_at" db:"updated_at"`
	DeletedBy *string    `json:"deleted_by" db:"deleted_by"`
	DeletedAt *time.Time `json:"deleted_at" db:"deleted_at"`
	Name      string     `json:"name" db:"name"`
}

type CreateRoleRequest struct {
	Name string `validate:"required"`
}

func (u Role) GetTableName() string {
	return "roles"
}
