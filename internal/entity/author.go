package entity

type Author struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
	Bio  string `json:"bio"`
}
