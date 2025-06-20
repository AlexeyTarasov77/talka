package security

import (
	"crypto/rand"
	"encoding/base64"
)

var _defaultTokenEntropy = 64

type SecurityProvider struct {
	tokenEntropy int
}

func New(opts ...Option) *SecurityProvider {
	securityProvider := &SecurityProvider{tokenEntropy: _defaultTokenEntropy}
	for _, opt := range opts {
		opt(securityProvider)
	}
	return securityProvider
}

func (s *SecurityProvider) GenerateSecureUrlSafeToken() string {
	buf := make([]byte, s.tokenEntropy)
	_, err := rand.Read(buf)
	if err != nil {
		panic(err)
	}
	return base64.URLEncoding.EncodeToString(buf)
}
