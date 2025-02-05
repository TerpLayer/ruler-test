package rulertest

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestAdd(t *testing.T) {
	result := Add(1, 2)
	require.Equal(t, 3, result)

}

func TestSub(t *testing.T) {
	result := Sub(2, 1)
	require.Equal(t, 1, result)

	result = Sub(3, 1)
	require.Equal(t, 2, result)

	result = Sub(3, 1)
	require.Equal(t, 1, result)
}
