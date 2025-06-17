package users_api

import "context"

type UsersAPI struct{}

func New() *UsersAPI {
	return &UsersAPI{}
}

func (api *UsersAPI) CheckExists(ctx context.Context, id int) (bool, error) {
	return false, nil
}
