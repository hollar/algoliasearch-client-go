// Code generated by go generate. DO NOT EDIT.

package opt

import (
	"encoding/json"
	"reflect"
)

// ExtraHeadersOption is a wrapper for an ExtraHeaders option parameter. It holds
// the actual value of the option that can be accessed by calling Get.
type ExtraHeadersOption struct {
	value map[string]string
}

// ExtraHeaders wraps the given value into a ExtraHeadersOption.
func ExtraHeaders(v map[string]string) *ExtraHeadersOption {
	return &ExtraHeadersOption{v}
}

// Get retrieves the actual value of the option parameter.
func (o *ExtraHeadersOption) Get() map[string]string {
	if o == nil {
		return map[string]string{}
	}
	return o.value
}

// MarshalJSON implements the json.Marshaler interface for
// ExtraHeadersOption.
func (o ExtraHeadersOption) MarshalJSON() ([]byte, error) {
	return json.Marshal(o.value)
}

// UnmarshalJSON implements the json.Unmarshaler interface for
// ExtraHeadersOption.
func (o *ExtraHeadersOption) UnmarshalJSON(data []byte) error {
	if string(data) == "null" {
		o.value = map[string]string{}
		return nil
	}
	return json.Unmarshal(data, &o.value)
}

// Equal returns true if the given option is equal to the instance one. In case
// the given option is nil, we checked the instance one is set to the default
// value of the option.
func (o *ExtraHeadersOption) Equal(o2 *ExtraHeadersOption) bool {
	if o2 == nil {
		return reflect.DeepEqual(o.value, map[string]string{})
	}
	return reflect.DeepEqual(o.value, o2.value)
}

// ExtraHeadersEqual returns true if the two options are equal.
// In case of one option being nil, the value of the other must be nil as well
// or be set to the default value of this option.
func ExtraHeadersEqual(o1, o2 *ExtraHeadersOption) bool {
	if o1 != nil {
		return o1.Equal(o2)
	}
	if o2 != nil {
		return o2.Equal(o1)
	}
	return true
}
