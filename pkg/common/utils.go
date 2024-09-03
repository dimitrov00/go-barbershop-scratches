package common

import (
	"barbershop/creativo/pkg/types"
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
	"time"
)

func NewAuditUpdater(userID types.UserID, now time.Time, forNewRecord bool) types.AuditChangeProcessorFn {
	return func(audit types.Audit) types.Audit {
		audit.UpdatedAt = now
		audit.UpdatedBy = userID

		if forNewRecord {
			audit.CreatedAt = now
			audit.CreatedBy = userID
		}

		return audit
	}
}

func GenerateRandomCode(length int) (string, error) {
	bytes := make([]byte, length)
	if _, err := rand.Read(bytes); err != nil {
		return "", err
	}

	return hex.EncodeToString(bytes)[:length], nil
}

func HashString(s string) (string, error) {
	hash := sha256.New()
	if _, err := hash.Write([]byte(s)); err != nil {
		return "", err
	}
	return hex.EncodeToString(hash.Sum(nil)), nil
}
