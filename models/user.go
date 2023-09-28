package models

import "time"

type User struct {
	ID        int64      `json:"id" db:"id"`
	CreatedBy string     `json:"created_by" db:"created_by"`
	CreatedAt time.Time  `json:"created_at" db:"created_at"`
	UpdatedBy string     `json:"updated_by" db:"updated_by"`
	UpdatedAt time.Time  `json:"updated_at" db:"updated_at"`
	DeletedBy *string    `json:"deleted_by" db:"deleted_by"`
	DeletedAt *time.Time `json:"deleted_at" db:"deleted_at"`
	FirstName string     `json:"first_name" db:"first_name"`
	LastName  string     `json:"last_name" db:"last_name"`
	Email     string     `json:"email" db:"email"`
	Phone     string     `json:"phone" db:"phone"`
	ImageUrl  *string    `json:"image_url" db:"image_url"`
	Password  string     `json:"-" db:"password"`
	Active    int        `db:"active"`
	RoleId    int64      `json:"role_id" db:"role_id"`
}

func (u User) GetTableName() string {
	return "users"
}

type UserTableData struct {
	ID       int64   `json:"id" db:"id"`
	ImageUrl *string `json:"image" db:"image_url"`
	Name     string  `json:"name" db:"name"`
	Email    string  `json:"email" db:"email"`
	RoleId   int64   `json:"role_id" db:"role_id"`
	Phone    string  `json:"phone" db:"phone"`
}

type CreateUserRequest struct {
	FirstName string  `json:"first_name" validate:"required" valerr:"Nama awal wajib diisi"`
	LastName  string  `json:"last_name" validate:"required" valerr:"Nama akhir wajib diisi"`
	Email     string  `json:"email" validate:"required,email" valerr:"Email wajib diisi"`
	Phone     string  `json:"phone" validate:"numeric,min=11,max=15" valerr:"Telepon harus angka"`
	ImageUrl  *string `json:"image" validate:"-"`
	Password  string  `json:"password" validate:"required,min=8" valerr:"Password minimal 8 karakter"`
	RoleId    int64   `json:"role_id" validate:"required,numeric" valerr:"Peran wajib diisi"`
	Active    int
}

type UpdateUserRequest struct {
	FirstName string  `json:"first_name" validate:"required" valerr:"Nama awal wajib diisi"`
	LastName  string  `json:"last_name" validate:"required" valerr:"Nama akhir wajib diisi"`
	Email     string  `json:"email" validate:"required,email" valerr:"Email wajib diisi"`
	Phone     string  `json:"phone" validate:"numeric,min=11,max=15" valerr:"Telepon harus angka"`
	ImageUrl  *string `json:"image" validate:"-"`
}

type CreateImgRequest struct {
	Image string `json:"image" validate:"required"`
}

type UserRegister struct {
	FirstName       string `json:"first_name" validate:"required" valerr:"Nama awal wajib diisi"`
	LastName        string `json:"last_name" validate:"required" valerr:"Nama akhir wajib diisi"`
	Email           string `json:"email" validate:"required,email" valerr:"Email wajib diisi"`
	Phone           string `json:"phone" validate:"numeric,min=11,max=15" valerr:"Telepon harus angka"`
	Password        string `json:"password" validate:"required,min=8" valerr:"Password minimal 8 karakter"`
	ConfirmPassword string `json:"confirm_password" validate:"required,min=8" valerr:"Password minimal 8 karakter"`
}

type UserSearch struct {
	Email  string `json:"email"`
	Phone  string `json:"phone"`
	Search string `json:"search"`
}

type UserProfileDetail struct {
	ID         int64      `json:"id" db:"id"`
	CreatedBy  string     `json:"created_by" db:"created_by"`
	CreatedAt  time.Time  `json:"created_at" db:"created_at"`
	UpdatedBy  string     `json:"updated_by" db:"updated_by"`
	UpdatedAt  time.Time  `json:"updated_at" db:"updated_at"`
	DeletedBy  *string    `json:"deleted_by" db:"deleted_by"`
	DeletedAt  *time.Time `json:"deleted_at" db:"deleted_at"`
	FirstName  string     `json:"first_name" db:"first_name"`
	LastName   string     `json:"last_name" db:"last_name"`
	Email      string     `json:"email" db:"email"`
	Phone      string     `json:"phone" db:"phone"`
	ImageUrl   *string    `json:"image_url" db:"image_url"`
	Password   string     `json:"-" db:"password"`
	Active     int        `db:"active"`
	RoleId     int64      `json:"role_id" db:"role_id"`
	Profession *string    `json:"profession" db:"profession"`
	Company    *string    `json:"company" db:"company"`
	Domicile   *string    `json:"domicile" db:"domicile"`
}

// =========================================
// ================ PROFILES ===============
// =========================================

type Profile struct {
	ID         int64      `json:"id" db:"id"`
	Profession string     `json:"profession" db:"profession"`
	Company    string     `json:"company" db:"company"`
	Domicile   string     `json:"domicile" db:"domicile"`
	CreatedBy  string     `json:"created_by" db:"created_by"`
	CreatedAt  time.Time  `json:"created_at" db:"created_at"`
	UpdatedBy  string     `json:"updated_by" db:"updated_by"`
	UpdatedAt  time.Time  `json:"updated_at" db:"updated_at"`
	DeletedBy  *string    `json:"deleted_by" db:"deleted_by"`
	DeletedAt  *time.Time `json:"deleted_at" db:"deleted_at"`
	UserID     int64      `json:"user_id" db:"user_id"`
}

type ProfileRequest struct {
	Profession string `json:"profession"`
	Company    string `json:"company"`
	Domicile   string `json:"domicile"`
}

type ProfileDetail struct {
	Profession string     `json:"profession" db:"profession"`
	Company    string     `json:"company" db:"company"`
	Domicile   string     `json:"domicile" db:"domicile"`
}

func (p Profile) GetTableName() string {
	return "profiles"
}
