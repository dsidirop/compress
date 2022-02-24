typedef string Timestamp # in golang we can grab the utc representation of time as    time.Now().UTC().MarshalText()

struct THFooItem {
	1: string          ID,
	2: string          Email,
	3: list<string>    Roles,
	4: string          APIKey,
	5: THFooProfile    Profile,
	6: string          Username,
	7: Timestamp       CreatedAt,
	8: Timestamp       UpdatedAt
}

struct THFooProfile {
	1: string          Dob,
	2: string          Name,
	3: string          About,
	4: string          Address,
	5: string          Company,
	6: THFooLocation   Location
}

struct THFooLocation {
	1: double   Lat,
	2: double   Long
}
