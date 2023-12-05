// Code generated by github.com/infraboard/mcube/v2
// DO NOT EDIT

package user

import (
	"bytes"
	"fmt"
	"strings"
)

// ParsePROVIDERFromString Parse PROVIDER from string
func ParsePROVIDERFromString(str string) (PROVIDER, error) {
	key := strings.Trim(string(str), `"`)
	v, ok := PROVIDER_value[strings.ToUpper(key)]
	if !ok {
		return 0, fmt.Errorf("unknown PROVIDER: %s", str)
	}

	return PROVIDER(v), nil
}

// Equal type compare
func (t PROVIDER) Equal(target PROVIDER) bool {
	return t == target
}

// IsIn todo
func (t PROVIDER) IsIn(targets ...PROVIDER) bool {
	for _, target := range targets {
		if t.Equal(target) {
			return true
		}
	}

	return false
}

// MarshalJSON todo
func (t PROVIDER) MarshalJSON() ([]byte, error) {
	b := bytes.NewBufferString(`"`)
	b.WriteString(strings.ToUpper(t.String()))
	b.WriteString(`"`)
	return b.Bytes(), nil
}

// UnmarshalJSON todo
func (t *PROVIDER) UnmarshalJSON(b []byte) error {
	ins, err := ParsePROVIDERFromString(string(b))
	if err != nil {
		return err
	}
	*t = ins
	return nil
}

// ParseTYPEFromString Parse TYPE from string
func ParseTYPEFromString(str string) (TYPE, error) {
	key := strings.Trim(string(str), `"`)
	v, ok := TYPE_value[strings.ToUpper(key)]
	if !ok {
		return 0, fmt.Errorf("unknown TYPE: %s", str)
	}

	return TYPE(v), nil
}

// Equal type compare
func (t TYPE) Equal(target TYPE) bool {
	return t == target
}

// IsIn todo
func (t TYPE) IsIn(targets ...TYPE) bool {
	for _, target := range targets {
		if t.Equal(target) {
			return true
		}
	}

	return false
}

// MarshalJSON todo
func (t TYPE) MarshalJSON() ([]byte, error) {
	b := bytes.NewBufferString(`"`)
	b.WriteString(strings.ToUpper(t.String()))
	b.WriteString(`"`)
	return b.Bytes(), nil
}

// UnmarshalJSON todo
func (t *TYPE) UnmarshalJSON(b []byte) error {
	ins, err := ParseTYPEFromString(string(b))
	if err != nil {
		return err
	}
	*t = ins
	return nil
}

// ParseGenderFromString Parse Gender from string
func ParseGenderFromString(str string) (Gender, error) {
	key := strings.Trim(string(str), `"`)
	v, ok := Gender_value[strings.ToUpper(key)]
	if !ok {
		return 0, fmt.Errorf("unknown Gender: %s", str)
	}

	return Gender(v), nil
}

// Equal type compare
func (t Gender) Equal(target Gender) bool {
	return t == target
}

// IsIn todo
func (t Gender) IsIn(targets ...Gender) bool {
	for _, target := range targets {
		if t.Equal(target) {
			return true
		}
	}

	return false
}

// MarshalJSON todo
func (t Gender) MarshalJSON() ([]byte, error) {
	b := bytes.NewBufferString(`"`)
	b.WriteString(strings.ToUpper(t.String()))
	b.WriteString(`"`)
	return b.Bytes(), nil
}

// UnmarshalJSON todo
func (t *Gender) UnmarshalJSON(b []byte) error {
	ins, err := ParseGenderFromString(string(b))
	if err != nil {
		return err
	}
	*t = ins
	return nil
}

// ParseCREATE_FROMFromString Parse CREATE_FROM from string
func ParseCREATE_FROMFromString(str string) (CREATE_FROM, error) {
	key := strings.Trim(string(str), `"`)
	v, ok := CREATE_FROM_value[strings.ToUpper(key)]
	if !ok {
		return 0, fmt.Errorf("unknown CREATE_FROM: %s", str)
	}

	return CREATE_FROM(v), nil
}

// Equal type compare
func (t CREATE_FROM) Equal(target CREATE_FROM) bool {
	return t == target
}

// IsIn todo
func (t CREATE_FROM) IsIn(targets ...CREATE_FROM) bool {
	for _, target := range targets {
		if t.Equal(target) {
			return true
		}
	}

	return false
}

// MarshalJSON todo
func (t CREATE_FROM) MarshalJSON() ([]byte, error) {
	b := bytes.NewBufferString(`"`)
	b.WriteString(strings.ToUpper(t.String()))
	b.WriteString(`"`)
	return b.Bytes(), nil
}

// UnmarshalJSON todo
func (t *CREATE_FROM) UnmarshalJSON(b []byte) error {
	ins, err := ParseCREATE_FROMFromString(string(b))
	if err != nil {
		return err
	}
	*t = ins
	return nil
}

// ParseDESCRIBE_BYFromString Parse DESCRIBE_BY from string
func ParseDESCRIBE_BYFromString(str string) (DESCRIBE_BY, error) {
	key := strings.Trim(string(str), `"`)
	v, ok := DESCRIBE_BY_value[strings.ToUpper(key)]
	if !ok {
		return 0, fmt.Errorf("unknown DESCRIBE_BY: %s", str)
	}

	return DESCRIBE_BY(v), nil
}

// Equal type compare
func (t DESCRIBE_BY) Equal(target DESCRIBE_BY) bool {
	return t == target
}

// IsIn todo
func (t DESCRIBE_BY) IsIn(targets ...DESCRIBE_BY) bool {
	for _, target := range targets {
		if t.Equal(target) {
			return true
		}
	}

	return false
}

// MarshalJSON todo
func (t DESCRIBE_BY) MarshalJSON() ([]byte, error) {
	b := bytes.NewBufferString(`"`)
	b.WriteString(strings.ToUpper(t.String()))
	b.WriteString(`"`)
	return b.Bytes(), nil
}

// UnmarshalJSON todo
func (t *DESCRIBE_BY) UnmarshalJSON(b []byte) error {
	ins, err := ParseDESCRIBE_BYFromString(string(b))
	if err != nil {
		return err
	}
	*t = ins
	return nil
}
