package encoding

import "encoding/json"

// MustMarshalJSON is a helper function that marshals JSON.
func MustMarshalJSON[T any](v T) []byte {
	b, err := json.Marshal(v)
	if err != nil {
		panic(err)
	}
	return b
}

// MustUnmarshalJSON is a helper function that unmarshals JSON.
func MustUnmarshalJSON[T any](b []byte) *T {
	v := new(T)
	err := json.Unmarshal(b, v)
	if err != nil {
		panic(err)
	}
	return v
}
