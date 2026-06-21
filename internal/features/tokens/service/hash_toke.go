package tokens_service

import (
	"crypto/sha256"
	"encoding/hex"
)

func (s *TokensServise) hashToken(token string) string {
	hash := sha256.Sum256([]byte(token))
	return hex.EncodeToString(hash[:])
}
