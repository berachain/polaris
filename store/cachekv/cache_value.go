package cachekv

// cValue represents a cached value.
// If dirty is true, it indicates the cached value is different from the underlying value.
type cValue struct {
	value []byte
	dirty bool
}

//nolint:revive
func NewCValue(v []byte, d bool) *cValue {
	return &cValue{
		value: v,
		dirty: d,
	}
}

func (cv *cValue) Dirty() bool {
	return cv.dirty
}

func (cv *cValue) Value() []byte {
	return cv.value
}

// deepCopy creates a new cValue object with the same value and dirty flag as the original
// cValue object. This function is used to create a deep copy of the prev field in
// DeleteCacheValue and SetCacheValue objects, so that modifications to the original prev value do
// not affect the cloned DeleteCacheValue or SetCacheValue object.
func (cv *cValue) deepCopy() *cValue {
	// Return a new cValue with the same value and dirty flag
	return &cValue{
		value: append([]byte(nil), cv.value...),
		dirty: cv.dirty,
	}
}
