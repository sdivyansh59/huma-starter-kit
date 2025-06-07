package snowflake

import (
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestGenerator_Next(t *testing.T) {
	g, err := NewGenerator(10)
	require.Nil(t, err)

	result := make(map[ID]bool)
	max := 500000

	for i := 0; i < max; i++ {
		id := g.Next()

		time.Sleep(time.Nanosecond)

		_, exists := result[id]
		require.False(t, exists)

		result[id] = true
	}
}
