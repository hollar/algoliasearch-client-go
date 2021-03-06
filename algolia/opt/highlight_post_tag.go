// Code generated by go generate. DO NOT EDIT.

package opt

import "encoding/json"

// HighlightPostTagOption is a wrapper for an HighlightPostTag option parameter. It holds
// the actual value of the option that can be accessed by calling Get.
type HighlightPostTagOption struct {
	value string
}

// HighlightPostTag wraps the given value into a HighlightPostTagOption.
func HighlightPostTag(v string) *HighlightPostTagOption {
	return &HighlightPostTagOption{v}
}

// Get retrieves the actual value of the option parameter.
func (o *HighlightPostTagOption) Get() string {
	if o == nil {
		return "</em>"
	}
	return o.value
}

// MarshalJSON implements the json.Marshaler interface for
// HighlightPostTagOption.
func (o HighlightPostTagOption) MarshalJSON() ([]byte, error) {
	return json.Marshal(o.value)
}

// UnmarshalJSON implements the json.Unmarshaler interface for
// HighlightPostTagOption.
func (o *HighlightPostTagOption) UnmarshalJSON(data []byte) error {
	if string(data) == "null" {
		o.value = "</em>"
		return nil
	}
	return json.Unmarshal(data, &o.value)
}

// Equal returns true if the given option is equal to the instance one. In case
// the given option is nil, we checked the instance one is set to the default
// value of the option.
func (o *HighlightPostTagOption) Equal(o2 *HighlightPostTagOption) bool {
	if o2 == nil {
		return o.value == "</em>"
	}
	return o.value == o2.value
}

// HighlightPostTagEqual returns true if the two options are equal.
// In case of one option being nil, the value of the other must be nil as well
// or be set to the default value of this option.
func HighlightPostTagEqual(o1, o2 *HighlightPostTagOption) bool {
	if o1 != nil {
		return o1.Equal(o2)
	}
	if o2 != nil {
		return o2.Equal(o1)
	}
	return true
}
