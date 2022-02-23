package arena

import (
	"time"
)

type FooItem struct {
	ID        string     `json:"id"              `
	Email     string     `json:"email"           `
	Roles     []string   `json:"roles"           `
	APIKey    string     `json:"apiKey"          `
	Profile   FooProfile `json:"profile"         `
	Username  string     `json:"username"        `
	CreatedAt time.Time  `json:"createdAt"       `
	UpdatedAt time.Time  `json:"updatedAt"       `
}

type FooProfile struct {
	Dob      string      `json:"dob"             `
	Name     string      `json:"name"            `
	About    string      `json:"about"           `
	Address  string      `json:"address"         `
	Company  string      `json:"company"         `
	Location FooLocation `json:"location"        `
}

type FooLocation struct {
	Lat  float64 `json:"lat"     `
	Long float64 `json:"long"    `
}
