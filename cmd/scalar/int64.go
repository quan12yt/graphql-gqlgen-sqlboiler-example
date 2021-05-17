package scalar

import (
	"fmt"
	"io"
	"strconv"
)

// Int64 represents a custom GraphQL "Int64" scalar type
type Int64 string

// UnmarshalGQL unmarshals Int64, a custom scalar type
func (i *Int64) UnmarshalGQL(v interface{}) error {
	var err error
	switch input := v.(type) {
	case string:
		*i = Int64(input)
	case int64:
		*i = Int64(strconv.FormatInt(int64(input), 10))
	default:
		err = fmt.Errorf("int64 must be string")
	}
	return err
}

// MarshalGQL marshal Operator, required by gqlgen
func (i Int64) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(string(i)))
}

// ToGQLInt64 takes an int64 type and converts it to custom Int64 scalar type
func ToGQLInt64(i int64) Int64 {
	return Int64(strconv.FormatInt(int64(i), 10))
}
