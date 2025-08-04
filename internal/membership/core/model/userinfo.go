package model

import "time"

// UserInfo - used by secondary ports
type UserInfo struct {
	AccountNumber string    `json:"account_number" db:"account_number"`
	Username      string    `json:"username" db:"username"`
	Hash          string    `json:"hash" db:"hash"`
	LastLogin     time.Time `json:"last_login" db:"last_login"`
	CreatedAt     time.Time `json:"created_at" db:"created_at"`
}

// UserProfileInfo - used by secondary ports
// This struct is used to represent the user profile information
type UserProfileInfo struct {
	AccountNumber   string    `json:"account_number" db:"account_number"`
	Email           string    `json:"email" db:"email"`
	FullName        string    `json:"full_name" db:"full_name"`
	StatusUser      string    `json:"status_user" db:"status_user"`
	Address         string    `json:"address" db:"address"`
	PostCode        string    `json:"post_code" db:"post_code"`
	BirthDate       string    `json:"birth_date" db:"birth_date"`
	Phone           string    `json:"phone" db:"phone"`
	WhatsApp        string    `json:"whatsapp" db:"whatsapp"`
	Npwp            string    `json:"npwp" db:"npwp"`
	Ktp             string    `json:"ktp" db:"ktp"`
	Avatar          string    `json:"avatar" db:"avatar"`
	BeneficiaryName string    `json:"beneficiary_name" db:"beneficiary_name"`
	LastLogin       time.Time `json:"last_login" db:"last_login"`
	CreatedAt       time.Time `json:"created_at" db:"created_at"`
}
