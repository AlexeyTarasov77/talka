package dto

type (
	CreatePersonalChat struct {
		CurrentUserId int `json:"current_user_id"`
		WithUserId    int `json:"with_user_id"`
	}
	CreateGroupChat struct {
		CurrentUserId int    `json:"current_user_id"`
		ImageURL      string `json:"image_url"`
		MembersIds    []int  `json:"members_ids"`
		Name          string
		Description   string
		Slug          string
		IsPublic      bool `json:"is_public"`
	}
)
