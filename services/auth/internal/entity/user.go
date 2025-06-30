package entity

import "time"

type User struct {
	ID int
	// Username is optional depending on signup method and oauth provider (if used)
	Username  string
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	// Password can be nil, if user signed up with oauth provider
	Password   []byte
	AvatarUrl  string    `json:"avatar_url"`
	LastSeenAt time.Time `json:"last_seen_at"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
	IsActive   bool      `json:"is_active"`
	BirthDate  time.Time `json:"birth_date"`
	// ProviderId is optional, present only if user signed up with oauth. Indicates which oauth provider was used
	OAuthProviderId OAuthProviderId `json:"provider_id"`
	// OAuthAccID is optional, present only if user signed up with oauth. Indicates unique id of user's oauth acc (e.g google acc)
	OAuthAccID string `json:"oauth_acc_id"`
}
