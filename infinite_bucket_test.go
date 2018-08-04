package rates

import (
	"testing"
	"time"

	. "github.com/smartystreets/goconvey/convey"
)

func TestInfiniteBucket(t *testing.T) {
	Convey("InfiniteBucket", t, func() {
		bucket := NewBucket(nil)

		Convey("Add", func() {
			bucket.Add(200)
			So(bucket.Take(), ShouldBeTrue)
		})

		Convey("Remove", func() {
			bucket.Remove(10)
			So(bucket.Take(), ShouldBeTrue)
			bucket.Remove(200)
			So(bucket.Take(), ShouldBeTrue)
		})

		Convey("Take", func() {
			bucket.Remove(200)
			So(bucket.Take(), ShouldBeTrue)

			bucket.Add(1)
			So(bucket.Take(), ShouldBeTrue)
		})

		Convey("TakeWhenAvailable", func() {
			start := time.Now()
			So(<-bucket.TakeWhenAvailable(), ShouldNotBeNil)
			So(time.Now().Sub(start).Seconds(), ShouldAlmostEqual, 0, 0.1)
		})
	})
}
