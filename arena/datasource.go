package arena

import (
	"time"
)

const TimeFormat = "2006-01-02T15:04:05.000Z"

func parseTimeString(input string) time.Time {

	t, err := time.Parse(TimeFormat, input)
	if err != nil {
		return time.Time{}
	}

	return t
}

var Datasource = []FooItem{
	{
		ID:    "620e7d1f89c0231fc95854d8",
		Email: "latasha_hanson@digirang.mf",
		Roles: []string{
			"owner",
			"member",
		},
		APIKey: "790c6327-f9a6-487b-8cad-6a2ede14c4e5",
		Profile: FooProfile{
			Dob:     "1988-01-12",
			Name:    "Latasha Hanson",
			About:   "Ad aliqua ullamco nulla officia laborum do nulla et laboris nisi duis nisi consectetur. Do occaecat labore quis nulla pariatur non.",
			Address: "70 McKibben Street, Omar, Connecticut",
			Company: "Digirang",
			Location: FooLocation{
				Lat:  22.344456,
				Long: -79.955604,
			},
		},
		Username:  "latasha88",
		CreatedAt: parseTimeString("2013-11-22T01:00:35.839Z"),
		UpdatedAt: parseTimeString("2013-11-23T01:00:35.839Z"),
	},
	{
		ID:    "620e7d1f4af1d347144b3461",
		Email: "gregory_lowery@hairport.amsterdam",
		Roles: []string{
			"admin",
			"owner",
		},
		APIKey: "5af7b181-b6fc-45e0-97a6-31614e5f2228",
		Profile: FooProfile{
			Dob:     "1992-10-13",
			Name:    "Gregory Lowery",
			About:   "Ut culpa cupidatat minim sint quis fugiat ex deserunt. Reprehenderit incididunt consectetur sunt dolor ad nulla incididunt ad ea amet Lorem.",
			Address: "24 Arlington Avenue, Bendon, Kentucky",
			Company: "Hairport",
			Location: FooLocation{
				Lat:  -79.090737,
				Long: -118.633953,
			},
		},
		Username:  "gregory92",
		CreatedAt: parseTimeString("2010-04-25T07:40:29.588Z"),
		UpdatedAt: parseTimeString("2010-04-26T07:40:29.588Z"),
	},
	{
		ID:    "620e7d1f92c8e4e23eb7b1f7",
		Email: "natasha_newman@orbean.org",
		Roles: []string{
			"member",
			"guest",
		},
		APIKey: "bb6b9343-4433-427f-badf-12ee6fc40b54",
		Profile: FooProfile{
			Dob:     "1993-06-10",
			Name:    "Natasha Newman",
			About:   "Incididunt in aliquip ad dolore adipisicing excepteur ea non minim ipsum aliquip exercitation. Enim labore cillum excepteur aliquip ipsum.",
			Address: "6 Himrod Street, Irwin, Illinois",
			Company: "Orbean",
			Location: FooLocation{
				Lat:  -2.844543,
				Long: -152.699558,
			},
		},
		Username:  "natasha93",
		CreatedAt: parseTimeString("2011-06-26T14:49:09.267Z"),
		UpdatedAt: parseTimeString("2011-06-27T14:49:09.267Z"),
	},
	{
		ID:    "620e7d1f8e4ec925b1a3550b",
		Email: "olive_fulton@telequiet.cn",
		Roles: []string{
			"admin",
			"owner",
		},
		APIKey: "e866792d-44ef-4e10-b4fa-bd3282eac283",
		Profile: FooProfile{
			Dob:     "1992-10-24",
			Name:    "Olive Fulton",
			About:   "Exercitation minim irure esse occaecat aute magna dolor et dolor mollit incididunt Lorem aliqua exercitation. Aliqua ex velit qui ipsum id irure consectetur eiusmod nostrud esse mollit occaecat exercitation.",
			Address: "74 Montauk Avenue, Wolcott, Arizona",
			Company: "Telequiet",
			Location: FooLocation{
				Lat:  34.275255,
				Long: -111.547445,
			},
		},
		Username:  "olive92",
		CreatedAt: parseTimeString("2014-03-23T09:15:55.839Z"),
		UpdatedAt: parseTimeString("2014-03-24T09:15:55.839Z"),
	},
	{
		ID:    "620e7d1f7120fe86785d32cb",
		Email: "smith_chen@balooba.training",
		Roles: []string{
			"owner",
		},
		APIKey: "e013e9d3-8d6a-4e39-bc1e-63e052775371",
		Profile: FooProfile{
			Dob:     "1992-10-30",
			Name:    "Smith Chen",
			About:   "Aliquip minim nisi duis culpa non veniam enim. Sunt do dolore ea ipsum.",
			Address: "3 Russell Street, Indio, North Carolina",
			Company: "Balooba",
			Location: FooLocation{
				Lat:  -53.661435,
				Long: 17.718808,
			},
		},
		Username:  "smith92",
		CreatedAt: parseTimeString("2014-04-22T02:13:27.414Z"),
		UpdatedAt: parseTimeString("2014-04-23T02:13:27.414Z"),
	},
	{
		ID:    "620e7d1f671fad80323c1b60",
		Email: "davenport_tate@xumonk.xxx",
		Roles: []string{
			"guest",
			"admin",
		},
		APIKey: "4a693735-1e83-4bfa-a385-0f4dafdcf53f",
		Profile: FooProfile{
			Dob:     "1992-09-23",
			Name:    "Davenport Tate",
			About:   "Amet excepteur in excepteur magna sint laborum incididunt irure do ipsum voluptate exercitation dolore. Nisi elit ad eiusmod eu.",
			Address: "93 Kansas Place, Machias, Iowa",
			Company: "Xumonk",
			Location: FooLocation{
				Lat:  61.146843,
				Long: -63.557771,
			},
		},
		Username:  "davenport92",
		CreatedAt: parseTimeString("2012-08-27T09:24:14.467Z"),
		UpdatedAt: parseTimeString("2012-08-28T09:24:14.467Z"),
	},
	{
		ID:    "620e7d1f9a01dedb6a92d947",
		Email: "tate_wade@decratex.shriram",
		Roles: []string{
			"guest",
			"admin",
		},
		APIKey: "2ced64d1-5039-4663-a39b-b67abca4eb57",
		Profile: FooProfile{
			Dob:     "1989-04-04",
			Name:    "Tate Wade",
			About:   "Ad duis nisi velit excepteur sint aute esse eiusmod. Qui commodo duis laborum amet laboris pariatur eu aliquip pariatur laboris.",
			Address: "33 Kiely Place, Loretto, North Dakota",
			Company: "Decratex",
			Location: FooLocation{
				Lat:  51.779527,
				Long: 2.504407,
			},
		},
		Username:  "tate89",
		CreatedAt: parseTimeString("2012-09-24T05:23:37.269Z"),
		UpdatedAt: parseTimeString("2012-09-25T05:23:37.269Z"),
	},
	{
		ID:    "620e7d1f2ab41022f7fc88a0",
		Email: "dolly_stout@otherway.mini",
		Roles: []string{
			"owner",
			"guest",
		},
		APIKey: "5fc9bb84-b536-462c-96c6-eacfb6f105ce",
		Profile: FooProfile{
			Dob:     "1991-06-19",
			Name:    "Dolly Stout",
			About:   "Sunt voluptate ad sint proident ipsum. Est Lorem eu adipisicing adipisicing do cupidatat veniam minim magna ea do nulla fugiat.",
			Address: "16 Dean Street, Greer, Vermont",
			Company: "Otherway",
			Location: FooLocation{
				Lat:  -27.874049,
				Long: 135.526772,
			},
		},
		Username:  "dolly91",
		CreatedAt: parseTimeString("2013-02-02T04:24:59.233Z"),
		UpdatedAt: parseTimeString("2013-02-03T04:24:59.233Z"),
	},
}
