package storage

type Person struct {
	ID          int    `json:"id,omitempty" example:"1"`
	Name        string `json:"name,omitempty" example:"Bill"`
	Surname     string `json:"surname,omitempty" example:"Jay"`
	Patronymic  string `json:"patronymic,omitempty" example:"Bob"`
	Age         int    `json:"age,omitempty" example:"35"`
	Sex         string `json:"sex,omitempty" example:"male"`
	Nationality string `json:"nationality,omitempty" example:"US"`
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
