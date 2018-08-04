package rates

type infiniteBucket struct {
}

func (b *infiniteBucket) Add(tickets float64) {

}

func (b *infiniteBucket) Remove(tickets float64) {

}

func (b *infiniteBucket) Take() bool {
	return true
}

func (b *infiniteBucket) Refilled() <-chan struct{} {
	c := make(chan struct{}, 1)
	c <- struct{}{}
	close(c)
	return c
}

func (b *infiniteBucket) TakeWhenAvailable() <-chan struct{} {
	c := make(chan struct{}, 1)
	c <- struct{}{}
	close(c)
	return c
}
