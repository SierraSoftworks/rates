package rates

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestBucket(t *testing.T) {
	Convey("Bucket", t, func() {
		Convey("NewBucket()", func() {
			So(NewBucket(nil), ShouldHaveSameTypeAs, &infiniteBucket{})
			So(NewBucket(&BucketConfig{}), ShouldHaveSameTypeAs, &ticketBucket{})
		})
	})
}
