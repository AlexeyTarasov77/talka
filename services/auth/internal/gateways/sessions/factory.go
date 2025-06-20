package sessions

import "github.com/AlexeyTarasov77/messanger.users/internal/gateways"

type ManagerFactory struct {
	storage Storage
}

func NewManagerFactory(storage Storage) *ManagerFactory {
	return &ManagerFactory{storage: storage}
}

func (f *ManagerFactory) CreateSessionManager(sessionId string) gateways.SessionManager {
	return NewManager(sessionId, f.storage)
}
