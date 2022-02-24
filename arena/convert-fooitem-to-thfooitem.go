package arena

import (
	"time"

	"github.com/klauspost/compress/arena/thfooitem"
)

func ConvertFooItemToTHFooItem(x FooItem) thfooitem.THFooItem {
	nowBytes, _ := time.Now().UTC().MarshalText()
	nowTimestampString := thfooitem.Timestamp(nowBytes)

	y := thfooitem.NewTHFooItem()
	y.ID = x.ID
	y.Email = x.Email
	y.Roles = x.Roles
	y.APIKey = x.APIKey
	y.Username = x.Username
	y.CreatedAt = nowTimestampString
	y.UpdatedAt = nowTimestampString

	y.Profile = thfooitem.NewTHFooProfile()
	y.Profile.Dob = x.Profile.Dob
	y.Profile.Name = x.Profile.Name
	y.Profile.About = x.Profile.About
	y.Profile.Address = x.Profile.Address
	y.Profile.Company = x.Profile.Company

	y.Profile.Location = thfooitem.NewTHFooLocation()
	y.Profile.Location.Lat = x.Profile.Location.Lat
	y.Profile.Location.Long = x.Profile.Location.Long

	return *y
}
