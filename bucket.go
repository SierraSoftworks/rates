package rates

import (
	"time"
)

type BucketConfig struct {
	MaxSize  float64 `json:"maxSize"`
	FillRate float64 `json:"fillRate"`
}

type Bucket interface {
	// Add will add the provided number of tickets to this
	// bucket, capping at the configured maximum bucket size.
	// This method is thread safe.
	Add(tickets float64)

	// Remove will remove the provided number of tickets
	// from this bucket, never falling below 0.
	// This method is thread safe.
	Remove(tickets float64)

	// Take will attempt to retrieve a ticket from this
	// bucket, returning false if there were no tickets
	// available.
	// This method is thread safe.
	Take() bool

	// TakeWhenAvailable acts similarly to Take, however it
	// will return a channel which will emit a single value
	// when a ticket has been successfully retrieved from
	// the bucket.
	TakeWhenAvailable() <-chan struct{}

	// Refilled informs you with a message on the channel when
	// a new ticket will have been added to the bucket. This
	// ticket might not be available for your use if another
	// goroutine claims it first, however it is guaranteed to
	// exist for someone.
	Refilled() <-chan struct{}
}

// NewBucket will create a rate bucket for the given
// config
func NewBucket(config *BucketConfig) Bucket {
	if config == nil {
		return &infiniteBucket{}
	}

	return &ticketBucket{
		Tickets: config.MaxSize,
		Config:  config,

		lastUpdate: time.Now(),
	}
}
