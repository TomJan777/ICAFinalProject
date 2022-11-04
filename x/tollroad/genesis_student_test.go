package tollroad_test

import (
	"testing"

	"github.com/b9lab/toll-road/x/tollroad/types"
	"github.com/stretchr/testify/require"
)

func TestExpectedDefaultGenesis(t *testing.T) {
	require.Equal(t, uint64(1), *&types.DefaultGenesis().SystemInfo.NextOperatorId)
}
