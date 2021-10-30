package entity

type AddFriend struct {
	SourceId string `json:"source_id"`
	TargetId string `json:"target_id"`
}

func NewAddFriend() *AddFriend {
	return new(AddFriend)
}