package tokengen

import (
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestHashData(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name   string
		setUp  func(*testing.T)
		expect func(*testing.T)
	}{
		{
			name:  "success",
			setUp: func(_ *testing.T) {},
			expect: func(t *testing.T) {
				value := "foo"
				res := hashData(value)
				require.NotEmpty(t, res)
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.setUp(t)
			tt.expect(t)
		})
	}
}
func TestHashAndEncryptWithSecretKey(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name   string
		setUp  func(*testing.T)
		expect func(*testing.T)
	}{
		{
			name:  "success",
			setUp: func(_ *testing.T) {},
			expect: func(t *testing.T) {
				value := "foo"
				secretKey := "bar"
				res := hashAndEncryptWithSecretKey(value, []byte(secretKey))
				require.NotEmpty(t, res)
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.setUp(t)
			tt.expect(t)
		})
	}
}

func TestSaltedHmac(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name   string
		setUp  func(*testing.T)
		expect func(*testing.T)
	}{
		{
			name:  "success",
			setUp: func(_ *testing.T) {},
			expect: func(t *testing.T) {
				value := "foo"
				secretKey := "bar"
				salt := "foobar"
				res := saltedHmac(salt, value, secretKey)
				require.NotEmpty(t, res)
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.setUp(t)
			tt.expect(t)
		})
	}
}

func TestSaltedHmacHex(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name   string
		setUp  func(*testing.T)
		expect func(*testing.T)
	}{
		{
			name:  "success",
			setUp: func(_ *testing.T) {},
			expect: func(t *testing.T) {
				value := "foo"
				secretKey := "bar"
				salt := "foobar"
				res := saltedHmacHex(salt, value, secretKey)
				require.NotEmpty(t, res)
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.setUp(t)
			tt.expect(t)
		})
	}
}

func TestGetTimestamp(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name   string
		setUp  func(*testing.T)
		expect func(*testing.T)
	}{
		{
			name:  "location unavailable",
			setUp: func(_ *testing.T) {},
			expect: func(t *testing.T) {
				n := time.Now()
				res := getTimestamp(n, nil)
				require.NotEmpty(t, res)
			},
		},
		{
			name:  "location available",
			setUp: func(_ *testing.T) {},
			expect: func(t *testing.T) {
				n := time.Now()
				res := getTimestamp(n, time.Local)
				require.NotEmpty(t, res)
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.setUp(t)
			tt.expect(t)
		})
	}
}
func TestGetNow(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name   string
		setUp  func(*testing.T)
		expect func(*testing.T)
	}{
		{
			name:  "success",
			setUp: func(_ *testing.T) {},
			expect: func(t *testing.T) {
				res := getNow()
				require.NotEmpty(t, res)
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.setUp(t)
			tt.expect(t)
		})
	}
}

func Test_makeTokenWithTs(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name   string
		setUp  func(*testing.T)
		expect func(*testing.T)
	}{
		{
			name:  "success",
			setUp: func(_ *testing.T) {},
			expect: func(t *testing.T) {
				value := "foo"
				secretKey := "bar"
				salt := "foobar"
				res := makeTokenWithTs(salt, value, secretKey, getTimestamp(getNow(), nil))
				require.NotEmpty(t, res)

				l := strings.Split(res, "-")
				require.Len(t, l, 2)
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.setUp(t)
			tt.expect(t)
		})
	}
}

func TestMakeTokenWithSalt(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name   string
		setUp  func(*testing.T)
		expect func(*testing.T)
	}{
		{
			name:  "success",
			setUp: func(_ *testing.T) {},
			expect: func(t *testing.T) {
				value := "foo"
				secretKey := "bar"
				salt := "foobar"
				res := MakeTokenWithSalt(salt, value, secretKey)
				require.NotEmpty(t, res)

				l := strings.Split(res, "-")
				require.Len(t, l, 2)
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.setUp(t)
			tt.expect(t)
		})
	}
}

func TestMakeToken(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name   string
		setUp  func(*testing.T)
		expect func(*testing.T)
	}{
		{
			name:  "success",
			setUp: func(_ *testing.T) {},
			expect: func(t *testing.T) {
				value := "foo"
				secretKey := "bar"
				res := MakeToken(value, secretKey)
				require.NotEmpty(t, res)

				l := strings.Split(res, "-")
				require.Len(t, l, 2)
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.setUp(t)
			tt.expect(t)
		})
	}
}

func TestValidateTokenWithKeySaltAndGetNow(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name   string
		setUp  func(*testing.T)
		expect func(*testing.T)
	}{
		{
			name:  "wrong format",
			setUp: func(_ *testing.T) {},
			expect: func(t *testing.T) {
				value := "foo"
				secretKey := "bar"
				salt := "foobar"
				timeout, _ := time.ParseDuration("60s")
				tk := "foobar"
				res := validateTokenWithKeySaltAndGetNow(salt, value, secretKey, tk, timeout, getNow)
				require.Error(t, res)
			},
		},
		{
			name:  "decode error",
			setUp: func(_ *testing.T) {},
			expect: func(t *testing.T) {
				value := "foo"
				secretKey := "bar"
				salt := "foobar"
				timeout, _ := time.ParseDuration("60s")
				tk := "foobar-foo"
				res := validateTokenWithKeySaltAndGetNow(salt, value, secretKey, tk, timeout, getNow)
				require.Errorf(t, res, "encoding/hex")
			},
		},
		{
			name:  "timeout exceeded",
			setUp: func(_ *testing.T) {},
			expect: func(t *testing.T) {
				value := "foo"
				secretKey := "bar"
				salt := "foobar"
				timeout, _ := time.ParseDuration("60s")
				now := time.Now()
				tk := makeTokenWithSaltAndGetNow(salt, value, secretKey, func() time.Time {
					return now
				})
				res := validateTokenWithKeySaltAndGetNow(salt, value, secretKey, tk, timeout, func() time.Time {
					return now.Add(timeout).Add(time.Hour)
				})
				require.Errorf(t, res, "timeout exceeded")
			},
		},
		{
			name:  "success",
			setUp: func(_ *testing.T) {},
			expect: func(t *testing.T) {
				value := "foo"
				secretKey := "bar"
				salt := "foobar"
				timeout, _ := time.ParseDuration("60s")
				n := time.Now()
				tk := makeTokenWithSaltAndGetNow(salt, value, secretKey, func() time.Time {
					return n
				})
				res := validateTokenWithKeySaltAndGetNow(salt, value, secretKey, tk, timeout, func() time.Time {
					return n
				})
				require.Nil(t, res)
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.setUp(t)
			tt.expect(t)
		})
	}
}

func Fuzz_makeTokenWithTs(f *testing.F) {
	testCases := []string{
		"foo", " ", "123!",
	}
	for _, tc := range testCases {
		f.Add(tc)
	}
	salt := "foobar"
	secretKey := "bar"
	f.Fuzz(func(t *testing.T, value string) {
		nw := getNow()
		tk := makeTokenWithTs(salt, value, secretKey, getTimestamp(nw, nil))
		res := validateTokenWithKeySaltAndGetNow(salt, value, secretKey, tk, time.Since(nw)+time.Hour, func() time.Time {
			return nw
		})
		if res != nil {
			t.Errorf("err: %q, salt: %q, value: %q, secretkey: %q", res, salt, value, secretKey)
		}
	})
}
