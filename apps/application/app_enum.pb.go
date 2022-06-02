// Code generated by github.com/infraboard/mcube
// DO NOT EDIT

package application

import (
	"bytes"
	"fmt"
	"strings"
)

// ParseTypeFromString Parse Type from string
func ParseTypeFromString(str string) (Type, error) {
	key := strings.Trim(string(str), `"`)
	v, ok := Type_value[strings.ToUpper(key)]
	if !ok {
		return 0, fmt.Errorf("unknown Type: %s", str)
	}

	return Type(v), nil
}

// Equal type compare
func (t Type) Equal(target Type) bool {
	return t == target
}

// IsIn todo
func (t Type) IsIn(targets ...Type) bool {
	for _, target := range targets {
		if t.Equal(target) {
			return true
		}
	}

	return false
}

// MarshalJSON todo
func (t Type) MarshalJSON() ([]byte, error) {
	b := bytes.NewBufferString(`"`)
	b.WriteString(strings.ToUpper(t.String()))
	b.WriteString(`"`)
	return b.Bytes(), nil
}

// UnmarshalJSON todo
func (t *Type) UnmarshalJSON(b []byte) error {
	ins, err := ParseTypeFromString(string(b))
	if err != nil {
		return err
	}
	*t = ins
	return nil
}

// ParseSCM_PROVIDERFromString Parse SCM_PROVIDER from string
func ParseSCM_PROVIDERFromString(str string) (SCM_PROVIDER, error) {
	key := strings.Trim(string(str), `"`)
	v, ok := SCM_PROVIDER_value[strings.ToUpper(key)]
	if !ok {
		return 0, fmt.Errorf("unknown SCM_PROVIDER: %s", str)
	}

	return SCM_PROVIDER(v), nil
}

// Equal type compare
func (t SCM_PROVIDER) Equal(target SCM_PROVIDER) bool {
	return t == target
}

// IsIn todo
func (t SCM_PROVIDER) IsIn(targets ...SCM_PROVIDER) bool {
	for _, target := range targets {
		if t.Equal(target) {
			return true
		}
	}

	return false
}

// MarshalJSON todo
func (t SCM_PROVIDER) MarshalJSON() ([]byte, error) {
	b := bytes.NewBufferString(`"`)
	b.WriteString(strings.ToUpper(t.String()))
	b.WriteString(`"`)
	return b.Bytes(), nil
}

// UnmarshalJSON todo
func (t *SCM_PROVIDER) UnmarshalJSON(b []byte) error {
	ins, err := ParseSCM_PROVIDERFromString(string(b))
	if err != nil {
		return err
	}
	*t = ins
	return nil
}

// ParseLANGUAGEFromString Parse LANGUAGE from string
func ParseLANGUAGEFromString(str string) (LANGUAGE, error) {
	key := strings.Trim(string(str), `"`)
	v, ok := LANGUAGE_value[strings.ToUpper(key)]
	if !ok {
		return 0, fmt.Errorf("unknown LANGUAGE: %s", str)
	}

	return LANGUAGE(v), nil
}

// Equal type compare
func (t LANGUAGE) Equal(target LANGUAGE) bool {
	return t == target
}

// IsIn todo
func (t LANGUAGE) IsIn(targets ...LANGUAGE) bool {
	for _, target := range targets {
		if t.Equal(target) {
			return true
		}
	}

	return false
}

// MarshalJSON todo
func (t LANGUAGE) MarshalJSON() ([]byte, error) {
	b := bytes.NewBufferString(`"`)
	b.WriteString(strings.ToUpper(t.String()))
	b.WriteString(`"`)
	return b.Bytes(), nil
}

// UnmarshalJSON todo
func (t *LANGUAGE) UnmarshalJSON(b []byte) error {
	ins, err := ParseLANGUAGEFromString(string(b))
	if err != nil {
		return err
	}
	*t = ins
	return nil
}

// ParseDescribeByFromString Parse DescribeBy from string
func ParseDescribeByFromString(str string) (DescribeBy, error) {
	key := strings.Trim(string(str), `"`)
	v, ok := DescribeBy_value[strings.ToUpper(key)]
	if !ok {
		return 0, fmt.Errorf("unknown DescribeBy: %s", str)
	}

	return DescribeBy(v), nil
}

// Equal type compare
func (t DescribeBy) Equal(target DescribeBy) bool {
	return t == target
}

// IsIn todo
func (t DescribeBy) IsIn(targets ...DescribeBy) bool {
	for _, target := range targets {
		if t.Equal(target) {
			return true
		}
	}

	return false
}

// MarshalJSON todo
func (t DescribeBy) MarshalJSON() ([]byte, error) {
	b := bytes.NewBufferString(`"`)
	b.WriteString(strings.ToUpper(t.String()))
	b.WriteString(`"`)
	return b.Bytes(), nil
}

// UnmarshalJSON todo
func (t *DescribeBy) UnmarshalJSON(b []byte) error {
	ins, err := ParseDescribeByFromString(string(b))
	if err != nil {
		return err
	}
	*t = ins
	return nil
}
