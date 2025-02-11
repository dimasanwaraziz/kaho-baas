package models

import "time"

type HashOptions struct {
	Type       string `json:"type"`
	MemoryCost int    `json:"memoryCost"`
	TimeCost   int    `json:"timeCost"`
	Threads    int    `json:"threads"`
}

type Target struct {
	ID           string    `json:"$id"`
	CreatedAt    time.Time `json:"$createdAt"`
	UpdatedAt    time.Time `json:"$updatedAt"`
	Name         string    `json:"name"`
	UserID       string    `json:"userId"`
	ProviderID   string    `json:"providerId"`
	ProviderType string    `json:"providerType"`
	Identifier   string    `json:"identifier"`
	Expired      bool      `json:"expired"`
}

type User struct {
	ID                string      `json:"$id"`
	CreatedAt         time.Time   `json:"$createdAt"`
	UpdatedAt         time.Time   `json:"$updatedAt"`
	Name              string      `json:"name"`
	Password          string      `json:"password"`
	Hash              string      `json:"hash"`
	HashOptions       HashOptions `json:"hashOptions"`
	Registration      time.Time   `json:"registration"`
	Status            bool        `json:"status"`
	Labels            []string    `json:"labels"`
	PasswordUpdate    time.Time   `json:"passwordUpdate"`
	Email             string      `json:"email"`
	Phone             string      `json:"phone"`
	EmailVerification bool        `json:"emailVerification"`
	PhoneVerification bool        `json:"phoneVerification"`
	MFA               bool        `json:"mfa"`
	Prefs             interface{} `json:"prefs"`
	Targets           []Target    `json:"targets"`
	AccessedAt        time.Time   `json:"accessedAt"`
}
