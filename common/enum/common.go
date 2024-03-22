package enum

import (
	"cmp"
	"slices"
)

func getKeys[TK cmp.Ordered, TV any](m map[TK]TV) []TK {
	var keys []TK
	for k := range m {
		keys = append(keys, k)
	}
	slices.Sort(keys)
	return keys
}
