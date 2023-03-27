package thezipcodes

type semaphore chan struct{}

func (s semaphore) push() {
	s <- struct{}{}
}
