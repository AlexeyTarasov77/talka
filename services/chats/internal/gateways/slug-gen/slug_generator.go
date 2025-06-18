package slug_gen

import (
	"crypto/rand"
	"encoding/base64"
)

type SlugGenerator struct {
	length int
}

func New(length int) *SlugGenerator {
	return &SlugGenerator{length: length}
}

func (s *SlugGenerator) GenerateRandomSlug() (string, error) {
	buf := make([]byte, s.length*6/8)
	_, err := rand.Read(buf)
	if err != nil {
		return "", err
	}
	return base64.RawURLEncoding.EncodeToString(buf), nil
}
