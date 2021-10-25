package entity

type AddFriend struct {
	SourceId int `json:"source_id"`
	TargetId int `json:"target_id"`
}

func NewAddFriend() *AddFriend {
	return new(AddFriend)
}