package abi

import (
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
)

type (
	Arguments = abi.Arguments
	Event     = abi.Event
)

var (
	MakeTopics  = abi.MakeTopics
	ToCamelCase = abi.ToCamelCase
)

// `ToMixedCase` converts a under_score formatted string to mixedCase format (camelCase with the
// first letter lowercase). This function is inspired by the geth abi.ToCamelCase function.
func ToMixedCase(input string) string {
	parts := strings.Split(input, "_")
	for i, s := range parts {
		if i > 0 && len(s) > 0 {
			parts[i] = strings.ToUpper(s[:1]) + s[1:]
		}
	}
	return strings.Join(parts, "")
}
