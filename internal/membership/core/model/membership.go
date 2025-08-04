package model

// LoginInfo - to submit login payload
type LoginInfo struct {
	Username string
	Password string
}

// RegisterInfo - to submit register payload
type RegisterInfo struct {
	FullName string
	Status   string
	Username string
	Password string
}
