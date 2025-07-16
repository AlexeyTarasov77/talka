package sessions

import (
	"context"
	"encoding/json"
	"fmt"
)

//go:generate mockgen -source=manager.go -destination=./mocks_test.go -package=sessions_test
type Storage interface {
	Get(ctx context.Context, key string) (string, error)
	Set(ctx context.Context, key, value string) error
}

type Manager struct {
	sessionId string
	storage   Storage
}

func NewManager(sessionId string, storage Storage) *Manager {
	return &Manager{
		sessionId: sessionId,
		storage:   storage,
	}
}

func (m *Manager) StorageKey() string {
	return fmt.Sprintf("sessions:%s", m.sessionId)
}

func (m *Manager) GetSessionData(ctx context.Context) (map[string]any, error) {
	rawData, err := m.storage.Get(ctx, m.StorageKey())
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve session data: %w", err)
	}
	res := make(map[string]any)
	if err := json.Unmarshal([]byte(rawData), &res); err != nil {
		return nil, fmt.Errorf("failed to unmarshal session data: %w", err)
	}
	return res, nil
}

func (m *Manager) SetToSession(ctx context.Context, key, value string) error {
	jsonData, err := m.storage.Get(ctx, m.StorageKey())
	if err != nil {
		return fmt.Errorf("failed to retrieve session data: %w", err)
	}
	data := make(map[string]any)
	if err := json.Unmarshal([]byte(jsonData), &data); err != nil {
		return fmt.Errorf("failed to unmarshal session data: %w", err)
	}
	data[key] = value
	updatedJsonData, err := json.Marshal(data)
	if err != nil {
		return fmt.Errorf("failed to serialize updated session data: %w", err)
	}
	if err := m.storage.Set(ctx, m.StorageKey(), string(updatedJsonData)); err != nil {
		return fmt.Errorf("failed to set updated session data: %w", err)
	}
	return nil
}
