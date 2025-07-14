package entity

import "time"

type (
	InvitationLink struct {
		// Url serves purpose of unique link identifier
		Url     string
		GroupId int
		// Name is optional link label visible to admins only
		Name string
		// RequestAdminApproval indicates whether admin should approve user's join request by this link before become participant. Default: false
		RequestAdminApproval bool
		// ActivationsLimit indicates how many users can use this link. Optional
		ActivationsLimit int
		// ExpiresAt  indicates when link expires and can't be used anymore. Optional
		ExpiresAt time.Time
	}
	ChatJoinReq struct {
		GroupChatId int
		FromUserId  int
	}
)

func (l *InvitationLink) IsExpired() bool {
	return !l.ExpiresAt.IsZero() && l.ExpiresAt.Before(time.Now())
}
