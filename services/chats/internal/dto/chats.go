package dto

type (
	CreatePersonalChat struct {
		FromUserId int `json:"from_user_id"`
		ToUserId   int `json:"to_user_id"`
	}
	CreateGroupChat struct {
		OwnerId     int    `json:"owner_id"`
		ImageURL    string `json:"image_url"`
		MembersIds  []int  `json:"members_ids"`
		Name        string
		Description string
		Slug        string
		IsPublic    bool `json:"is_public"`
	}
)
