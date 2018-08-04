package rates

import (
	"testing"
	"time"

	. "github.com/smartystreets/goconvey/convey"
)

func TestTicketBucket(t *testing.T) {
	Convey("TicketBucket", t, func() {
		conf := BucketConfig{
			MaxSize:  100,
			FillRate: 10,
		}

		bucket := NewBucket(&conf)
		normalBucket := bucket.(*ticketBucket)
		So(bucket, ShouldNotBeNil)
		So(normalBucket.Config, ShouldResemble, &conf)

		Convey("Initial Size", func() {
			So(normalBucket.Tickets, ShouldEqual, 100)
		})

		Convey("Add", func() {
			bucket.Add(200)
			So(normalBucket.Tickets, ShouldEqual, 100)
			So(bucket.Take(), ShouldBeTrue)
		})

		Convey("Remove", func() {
			bucket.Remove(10)
			So(normalBucket.Tickets, ShouldEqual, 90)
			So(bucket.Take(), ShouldBeTrue)
			bucket.Remove(200)
			So(normalBucket.Tickets, ShouldEqual, 0)
			So(bucket.Take(), ShouldBeFalse)
		})

		Convey("Take", func() {
			bucket.Remove(200)
			So(bucket.Take(), ShouldBeFalse)

			bucket.Add(1)
			So(bucket.Take(), ShouldBeTrue)
			So(normalBucket.Tickets, ShouldAlmostEqual, 0, 0.1)
		})

		Convey("TakeWhenAvailable", func() {
			bucket.Remove(200)

			start := time.Now()
			So(<-bucket.TakeWhenAvailable(), ShouldNotBeNil)
			So(time.Now().Sub(start).Seconds(), ShouldAlmostEqual, 0.1, 0.01)
		})
	})
}
