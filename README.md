# Rates
**Simple rate limiting primitives for Go**

This package provides a couple of basic rate limiting primitives which can
be used in Go applications to restrict the rate at which certain tasks are
performed.

```go
package main

import (
    "log"
    "time"
    
    "github.com/SierraSoftworks/rates"
)

func main() {

    // Buckets are a useful primitive which allow a constant rate of
    // execution while still allowing well behaved clients to burst
    // past that rate for short periods.
    bucket := rates.NewBucket(&rates.BucketConfig{
        MaxSize: 10,
        FillRate: 1,
    })

    for i := 0; i < 20; i++ {
        if bucket.Take() {
            log.Println("Executed task", i)
        } else {
            log.Println("Rate limited task", i)
        }
    }

    select {
        case <-bucket.TakeWhenAvailable():
            log.Println("Waited until we had an available task before running this")
        case <-time.After(time.Second):
            log.Println("We decided not to wait any longer")
    }
}
```