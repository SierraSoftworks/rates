package rates

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewBucket(t *testing.T) {
	assert.IsType(t, &infiniteBucket{}, NewBucket(nil), "it should return an infinite bucket if no config is provided")
	assert.IsType(t, &ticketBucket{}, NewBucket(&BucketConfig{}), "it should return a ticket bucket if a config is provided")
}
