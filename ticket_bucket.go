package rates

import (
	"sync"
	"time"
)

type ticketBucket struct {
	Tickets float64       `json:"tickets"`
	Config  *BucketConfig `json:"config"`

	lastUpdate time.Time
	mutex      sync.Mutex
}

func (b *ticketBucket) updateTickets() {
	now := time.Now()

	b.mutex.Lock()
	seconds := now.Sub(b.lastUpdate).Seconds()
	b.lastUpdate = now
	b.mutex.Unlock()

	b.Add(seconds * b.Config.FillRate)
}

func (b *ticketBucket) Add(tickets float64) {
	b.mutex.Lock()
	defer b.mutex.Unlock()

	b.Tickets = b.Tickets + tickets
	if b.Tickets > b.Config.MaxSize {
		b.Tickets = b.Config.MaxSize
	}
}

func (b *ticketBucket) Remove(tickets float64) {
	b.mutex.Lock()
	defer b.mutex.Unlock()

	b.Tickets = b.Tickets - tickets
	if b.Tickets < 0 {
		b.Tickets = 0
	}
}

func (b *ticketBucket) Take() bool {
	b.updateTickets()

	b.mutex.Lock()
	defer b.mutex.Unlock()
	if b.Tickets < 1 {
		return false
	}

	b.Tickets = b.Tickets - 1
	return true
}

func (b *ticketBucket) Refilled() <-chan struct{} {
	c := make(chan struct{})

	time.AfterFunc(
		time.Duration(float64(time.Second)/b.Config.FillRate),
		func() {
			c <- struct{}{}
			close(c)
		})

	return c
}

func (b *ticketBucket) TakeWhenAvailable() <-chan struct{} {
	c := make(chan struct{})

	go func() {
		for !b.Take() {
			time.Sleep(time.Duration(float64(time.Second) / b.Config.FillRate))
		}

		c <- struct{}{}
		close(c)
	}()

	return c
}
