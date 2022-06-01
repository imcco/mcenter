// Code generated by github.com/infraboard/mcube
// DO NOT EDIT

package instance

import (
	"bytes"
	"fmt"
	"strings"
)

// ParseCommandTypeFromString Parse CommandType from string
func ParseCommandTypeFromString(str string) (CommandType, error) {
	key := strings.Trim(string(str), `"`)
	v, ok := CommandType_value[strings.ToUpper(key)]
	if !ok {
		return 0, fmt.Errorf("unknown CommandType: %s", str)
	}

	return CommandType(v), nil
}

// Equal type compare
func (t CommandType) Equal(target CommandType) bool {
	return t == target
}

// IsIn todo
func (t CommandType) IsIn(targets ...CommandType) bool {
	for _, target := range targets {
		if t.Equal(target) {
			return true
		}
	}

	return false
}

// MarshalJSON todo
func (t CommandType) MarshalJSON() ([]byte, error) {
	b := bytes.NewBufferString(`"`)
	b.WriteString(strings.ToUpper(t.String()))
	b.WriteString(`"`)
	return b.Bytes(), nil
}

// UnmarshalJSON todo
func (t *CommandType) UnmarshalJSON(b []byte) error {
	ins, err := ParseCommandTypeFromString(string(b))
	if err != nil {
		return err
	}
	*t = ins
	return nil
}
