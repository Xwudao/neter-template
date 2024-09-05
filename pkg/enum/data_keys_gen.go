// Code generated by go-enum DO NOT EDIT.
// Version:
// Revision:
// Build Date:
// Built By:

package enum

import (
	"fmt"
	"strings"
)

const (
	// DataKeyFriendLink is a DataKey of type friend_link.
	DataKeyFriendLink DataKey = "friend_link"
	// DataKeyTopLink is a DataKey of type top_link.
	DataKeyTopLink DataKey = "top_link"
)

var ErrInvalidDataKey = fmt.Errorf("not a valid DataKey, try [%s]", strings.Join(_DataKeyNames, ", "))

var _DataKeyNames = []string{
	string(DataKeyFriendLink),
	string(DataKeyTopLink),
}

// DataKeyNames returns a list of possible string values of DataKey.
func DataKeyNames() []string {
	tmp := make([]string, len(_DataKeyNames))
	copy(tmp, _DataKeyNames)
	return tmp
}

// DataKeyValues returns a list of the values for DataKey
func DataKeyValues() []DataKey {
	return []DataKey{
		DataKeyFriendLink,
		DataKeyTopLink,
	}
}

// String implements the Stringer interface.
func (x DataKey) String() string {
	return string(x)
}

// IsValid provides a quick way to determine if the typed value is
// part of the allowed enumerated values
func (x DataKey) IsValid() bool {
	_, err := ParseDataKey(string(x))
	return err == nil
}

var _DataKeyValue = map[string]DataKey{
	"friend_link": DataKeyFriendLink,
	"top_link":    DataKeyTopLink,
}

// ParseDataKey attempts to convert a string to a DataKey.
func ParseDataKey(name string) (DataKey, error) {
	if x, ok := _DataKeyValue[name]; ok {
		return x, nil
	}
	return DataKey(""), fmt.Errorf("%s is %w", name, ErrInvalidDataKey)
}

// MarshalText implements the text marshaller method.
func (x DataKey) MarshalText() ([]byte, error) {
	return []byte(string(x)), nil
}

// UnmarshalText implements the text unmarshaller method.
func (x *DataKey) UnmarshalText(text []byte) error {
	tmp, err := ParseDataKey(string(text))
	if err != nil {
		return err
	}
	*x = tmp
	return nil
}