package arena

import (
	"time"
)

type FooItem struct {
	ID        string     `   json:"id"           msg:"id"          `
	Email     string     `   json:"email"        msg:"email"       `
	Roles     []string   `   json:"roles"        msg:"roles"       `
	APIKey    string     `   json:"apiKey"       msg:"apiKey"      `
	Profile   FooProfile `   json:"profile"      msg:"profile"     `
	Username  string     `   json:"username"     msg:"username"    `
	CreatedAt time.Time  `   json:"createdAt"    msg:"createdAt"   `
	UpdatedAt time.Time  `   json:"updatedAt"    msg:"updatedAt"   `
}

type FooProfile struct {
	Dob      string      `   json:"dob"          msg:"dob"         `
	Name     string      `   json:"name"         msg:"name"        `
	About    string      `   json:"about"        msg:"about"       `
	Address  string      `   json:"address"      msg:"address"     `
	Company  string      `   json:"company"      msg:"company"     `
	Location FooLocation `   json:"location"     msg:"location"    `
}

type FooLocation struct {
	Lat  float64 `           json:"lat"          msg:"lat"         `
	Long float64 `           json:"long"         msg:"long"        `
}
