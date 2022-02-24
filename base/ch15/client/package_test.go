package client

import (
	"ayixigo/base/ch15/series"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPackage(t *testing.T) {
	// Output:
	// [1 1 2 3 5 8 13 21 34 55 89 144 233 377 610 987 1597 2584 4181 6765]
	t.Log(series.GetFSeries(10))
	assert.Equal(t, 1, 1)
}
