package entity

type UpdateAge struct {
	NewAge int `json:"new_age"`
}

func NewUpdateAge() *UpdateAge {
	return new(UpdateAge)
}