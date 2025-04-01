package storage

type Person struct {
	ID          int    `json:"id,omitempty"`
	Name        string `json:"name,omitempty"`
	Surname     string `json:"surname,omitempty"`
	Patronymic  string `json:"patronymic,omitempty"`
	Age         int    `json:"age,omitempty"`
	Sex         string `json:"sex,omitempty"`
	Nationality string `json:"nationality,omitempty"`
}

type Age struct {
	Age int `json:"age"`
}

type Sex struct {
	Sex string `json:"gender"`
}

type Nationality struct {
	Nationality []Country `json:"country"`
}

type Country struct {
	Country string  `json:"country_id"`
	Prob    float64 `json:"probability"`
}
