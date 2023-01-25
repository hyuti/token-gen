package generator

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"fmt"
	"strings"
	"time"
)

const (
	KeySalt = "github.com/hyuti/pwdTokenGenerator"
)

func HashData(value string) []byte {
	h := sha256.New()
	h.Write([]byte(value))
	return h.Sum(nil)
}
func HashDataWithSecretKey(value string, secretKey []byte) []byte {
	h := hmac.New(sha256.New, secretKey)
	h.Write([]byte(value))
	return h.Sum(nil)
}
func SaltedHmac(keySalt, value, secretKey string) []byte {
	key := HashData(fmt.Sprintf("%s%s", keySalt, secretKey))
	return HashDataWithSecretKey(value, key)
}

func SaltedHmacHex(keySalt, value, secretKey string) string {
	tk := SaltedHmac(keySalt, value, secretKey)
	return hex.EncodeToString(tk)
}

func GetTimestamp(t time.Time, loc *time.Location) uint64 {
	if loc == nil {
		loc = time.UTC
	}
	d := t.Sub(time.Date(2001, time.January, 1, 0, 0, 0, 0, loc))
	return uint64(d.Seconds())
}

func GetNow() time.Time {
	return time.Now()
}

func MakeTokenWithTs(keySalt, value, secretKey string, ts uint64) string {
	tk := SaltedHmacHex(keySalt, value, secretKey)
	tk = extractEvenElements(tk)
	ts36 := convInt64ToBase36(ts)
	return fmt.Sprintf("%s-%s", ts36, tk)
}
func MakeTokenWithSalt(keySalt, value, secretKey string) string {
	return MakeTokenWithSaltAndGetNow(keySalt, value, secretKey, GetNow)
}
func MakeTokenWithSaltAndGetNow(keySalt, value, secretKey string, getNow func() time.Time) string {
	return MakeTokenWithTs(keySalt, value, secretKey, GetTimestamp(getNow(), nil))
}
func MakeToken(value, secretKey string) string {
	return MakeTokenWithSalt(KeySalt, value, secretKey)
}

func validateToken(expectedTk, actualTk string) error {
	lAcTk := strings.Split(actualTk, "-")
	if len(lAcTk) != 2 {
		return errors.New("invalid format")
	}
	acTk := lAcTk[1]
	acTkB, err := hex.DecodeString(acTk)
	if err != nil {
		return err
	}
	lExTk := strings.Split(expectedTk, "-")
	if len(lExTk) != 2 {
		return errors.New("invalid format")
	}
	exTk := lExTk[1]
	exTkB, err := hex.DecodeString(exTk)
	if err != nil {
		return err
	}
	if !hmac.Equal(exTkB, acTkB) {
		return errors.New("invalid token")
	}
	return nil
}

func ValidateTokenWithKeySalt(keySalt, value, secretKey, actualTk string, timeout time.Duration) error {
	return ValidateTokenWithKeySaltAndGetNow(keySalt, value, secretKey, actualTk, timeout, GetNow)
}
func ValidateTokenWithKeySaltAndGetNow(keySalt, value, secretKey, actualTk string, timeout time.Duration, getNow func() time.Time) error {
	lAcTk := strings.Split(actualTk, "-")
	if len(lAcTk) != 2 {
		return errors.New("invalid format")
	}
	tsStr := lAcTk[0]
	ts := convBase36ToInt64(tsStr)

	if err := validateToken(MakeTokenWithTs(keySalt, value, secretKey, ts), actualTk); err != nil {
		return err
	}

	if (GetTimestamp(getNow(), nil) - ts) > uint64(timeout.Seconds()) {
		return errors.New("timeout exceeded")
	}
	return nil
}

func ValidateToken(value, secretKey, actualTk string, timeout time.Duration) error {
	return ValidateTokenWithKeySalt(KeySalt, value, secretKey, actualTk, timeout)
}
