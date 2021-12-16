package ast

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestValue(t *testing.T) {

	// make sure that an empty list (*not* null) when evaluated as
	// an literal argument does not get turned into a nil list
	//
	// e.g., if we have a field with arguments such as:
	//     foo(bar: [String!])
	//
	// we want to be able to distinguish between an invocation such as
	//     foo
	// and
	//     foo(bar: [])

	t.Run("empty list value", func(t *testing.T) {
		value := &Value{
			Raw:      "[]",
			Children: ChildValueList{},
			Kind:     ListValue,
		}
		lst, err := value.Value(nil)
		require.Nil(t, err)
		require.Equal(t, len(lst.([]interface{})), 0)
		require.NotNil(t, lst)
	})

	t.Run("null list value", func(t *testing.T) {
		value := &Value{
			Raw:  "null",
			Kind: ListValue,
		}
		lst, err := value.Value(nil)
		require.Nil(t, err)
		require.Equal(t, len(lst.([]interface{})), 0)
		require.Nil(t, lst)
	})
}
