package rates

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestInfiniteBucket(t *testing.T) {
	bucket := NewBucket(nil)
	require.NotNil(t, bucket, "it should return a bucket")

	t.Run("Add()", func(t *testing.T) {
		bucket.Add(200)
		assert.True(t, bucket.Take())
	})

	t.Run("Remove()", func(t *testing.T) {
		bucket.Remove(10)
		assert.True(t, bucket.Take())

		bucket.Remove(200)
		assert.True(t, bucket.Take())
	})

	t.Run("Take()", func(t *testing.T) {
		bucket.Remove(200)
		assert.True(t, bucket.Take())

		bucket.Add(1)
		assert.True(t, bucket.Take())
	})

	t.Run("TakeWhenAvailable()", func(t *testing.T) {
		select {
		case <-bucket.TakeWhenAvailable():
		case <-time.After(100 * time.Millisecond):
			t.Error("Timed out waiting for a token")
		}
	})
}
