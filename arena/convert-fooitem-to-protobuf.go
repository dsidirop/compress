package arena

import (
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
)

func ConvertFooItemToPBFooItem(x FooItem) *PBFooItem {
	return &PBFooItem{
		ID:        x.ID,
		Email:     x.Email,
		Roles:     x.Roles,
		APIKey:    x.APIKey,
		Username:  x.Username,
		CreatedAt: timestamppb.New(x.CreatedAt),
		UpdatedAt: timestamppb.New(x.UpdatedAt),

		Profile: &PBFooProfile{
			Dob:     x.Profile.Dob,
			Name:    x.Profile.Name,
			About:   x.Profile.About,
			Address: x.Profile.Address,
			Company: x.Profile.Company,

			Location: &PBFooLocation{
				Lat:  x.Profile.Location.Lat,
				Long: x.Profile.Location.Long,
			},
		},
	}
}
