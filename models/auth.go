package models

import "time"

type UserRole struct {
	ID        int64      `name:"Id" json:"id" db:"id"`
	CreatedAt time.Time  `name:"CreatedAt" json:"created_at" db:"created_at"`
	UpdatedAt time.Time  `name:"UpdatedAt" json:"updated_at" db:"updated_at"`
	DeletedAt *time.Time `name:"DeletedAt" json:"deleted_at" db:"deleted_at"`
	FirstName string     `name:"FirstName" json:"first_name" db:"first_name"`
	LastName  string     `name:"LastName" json:"last_name" db:"last_name"`
	Email     string     `name:"Email" json:"email" db:"email"`
	Phone     string     `name:"Phone" json:"phone" db:"phone"`
	ImageUrl  *string    `name:"ImageUrl" json:"image_url" db:"image_url"`
	RoleID    int64      `name:"RoleID" json:"role_id" db:"role_id"`
	RoleName  string     `name:"RoleName" json:"role_name" db:"role_name"`
}

type Session struct {
	UserID    int64
	ExpiresAt int64
}

type VerifiedUser struct {
	ID        int64     `db:"id"`
	UserId    int64     `db:"user_id"`
	Code      string    `db:"code"`
	CreatedAt time.Time `db:"created_at"`
}

func (u VerifiedUser) GetTableNameAndAlias() (string, string) {
	return "verified_user", "verified_user"
}

type CreatePasswordRequest struct {
	OldPassword  string `json:"old_password" validate:"required,min=8" valerr:"password minimal 8 karakter"`
	NewPassword  string `json:"new_password" validate:"required,min=8" valerr:"password minimal 8 karakter"`
	NewPassword2 string `json:"new_password2" validate:"required,min=8" valerr:"password minimal 8 karakter"`
}

type AuthRequest struct {
	Email    string `json:"email" validate:"required,email" valerr:"Email wajib diisi"`
	Password string `json:"password" validate:"required" valerr:"Password wajib diisi"`
}

type AuthResponse struct {
	Token     string `json:"token"`
	User      User   `json:"user"`
	Role      Role   `json:"role"`
	ExpiredAt string `json:"expired_at"`
}
