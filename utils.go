package tokengen

import (
	"github.com/martinlindhe/base36"
)

func convInt64ToBase36(ts uint64) string {
	return base36.Encode(ts)
}

func convBase36ToInt64(ts string) uint64 {
	return base36.Decode(ts)
}

func extractEvenElements(value string) string {
	v := ""
	for i, e := range value {
		if i%2 == 0 {
			v += string(e)
		}
	}
	return v
}
