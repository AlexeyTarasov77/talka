package security

type Option func(*SecurityProvider)

func TokenEntropy(n int) Option {
	return func(s *SecurityProvider) {
		s.tokenEntropy = n
	}
}
