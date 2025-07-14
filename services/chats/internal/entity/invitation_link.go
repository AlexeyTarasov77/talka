package entity

import "time"

type (
	InvitationLink struct {
		ID int
		// Url is unique link code
		Url     string
		GroupId int `json:"group_id"`
		// Name is optional link label visible to admins only
		Name string
		// RequiresAdminApproval indicates whether admin should approve user's join request by this link before become participant. Default: false
		RequiresAdminApproval bool `json:"requires_admin_approval"`
		// ActivationsLimit indicates how many users can use this link. Optional
		ActivationsLimit int `json:"activations_limit"`
		// ExpiresAt  indicates when link expires and can't be used anymore. Optional
		ExpiresAt time.Time `json:"expires_at"`
	}
)

func (l *InvitationLink) IsExpired() bool {
	return !l.ExpiresAt.IsZero() && l.ExpiresAt.Before(time.Now())
}
