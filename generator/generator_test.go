package generator_test

import (
	"strings"
	"testing"
	"time"

	"github.com/hyuti/pwdTokenGenerator/generator"
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
				res := generator.HashData(value)
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
func TestHashDataWithSecretKey(t *testing.T) {
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
				res := generator.HashDataWithSecretKey(value, []byte(secretKey))
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
				res := generator.SaltedHmac(salt, value, secretKey)
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
				res := generator.SaltedHmacHex(salt, value, secretKey)
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
				res := generator.GetTimestamp(n, nil)
				require.NotEmpty(t, res)
			},
		},
		{
			name:  "location available",
			setUp: func(_ *testing.T) {},
			expect: func(t *testing.T) {
				n := time.Now()
				res := generator.GetTimestamp(n, time.Local)
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
				res := generator.GetNow()
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

func TestMakeTokenWithTs(t *testing.T) {
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
				res := generator.MakeTokenWithTs(salt, value, secretKey, generator.GetTimestamp(generator.GetNow(), nil))
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
				res := generator.MakeTokenWithSalt(salt, value, secretKey)
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
				res := generator.MakeToken(value, secretKey)
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
				res := generator.ValidateTokenWithKeySaltAndGetNow(salt, value, secretKey, tk, timeout, generator.GetNow)
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
				res := generator.ValidateTokenWithKeySaltAndGetNow(salt, value, secretKey, tk, timeout, generator.GetNow)
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
				tk := generator.MakeTokenWithSaltAndGetNow(salt, value, secretKey, func() time.Time {
					return now
				})
				res := generator.ValidateTokenWithKeySaltAndGetNow(salt, value, secretKey, tk, timeout, func() time.Time {
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
				tk := generator.MakeTokenWithSaltAndGetNow(salt, value, secretKey, func() time.Time {
					return n
				})
				res := generator.ValidateTokenWithKeySaltAndGetNow(salt, value, secretKey, tk, timeout, func() time.Time {
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
