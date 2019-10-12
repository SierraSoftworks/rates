package rates

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestTicketBucket(t *testing.T) {
	conf := BucketConfig{
		MaxSize:  100,
		FillRate: 10,
	}

	bucket := NewBucket(&conf)
	require.NotNil(t, bucket, "it should return a bucket")

	normalBucket, ok := bucket.(*ticketBucket)
	require.True(t, ok, "it should actually be a *ticketBucket")

	assert.EqualValues(t, 100, normalBucket.Tickets, "it should start with 100 tickets")

	t.Run("Add()", func(t *testing.T) {
		bucket.Add(200)
		assert.EqualValues(t, 100, normalBucket.Tickets, "it should not allow you to add more tickets than the max")
		assert.True(t, bucket.Take(), "it should allow us to consume a ticket")
	})

	t.Run("Remove()", func(t *testing.T) {
		bucket.Add(200) // reset the bucket to full

		bucket.Remove(10)
		assert.EqualValues(t, 90, normalBucket.Tickets, "it should remove the required number of tickets from the bucket")
		assert.True(t, bucket.Take(), "it should allow us to consume a ticket")

		bucket.Remove(200)
		assert.EqualValues(t, 0, normalBucket.Tickets, "it should not allow you to remove more tickets than the minimum")
		assert.False(t, bucket.Take(), "it should not allow us to consume a ticket")
	})

	t.Run("Take()", func(t *testing.T) {
		bucket.Remove(200)
		assert.False(t, bucket.Take(), "it should not allow us to remove a token from an empty bucket")

		bucket.Add(1)
		assert.True(t, bucket.Take(), "it should allow us to remove a token once one is available")
		assert.EqualValues(t, 0, normalBucket.Tickets, "it should have no tickets left afterwards")
	})

	t.Run("TakeWhenAvailable()", func(t *testing.T) {
		bucket.Remove(200)

		start := time.Now()

		select {
		case <-bucket.TakeWhenAvailable():
			elapsed := time.Now().Sub(start).Seconds()
			assert.InEpsilon(t, 0.1, elapsed, 0.01, "it should have waited 100ms to hand out a token")
		case <-time.After(200 * time.Millisecond):
			t.Error("Timed out waiting for a token after 200ms")
		}
	})
}
