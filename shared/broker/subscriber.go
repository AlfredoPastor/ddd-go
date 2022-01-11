package domain

type Subscriber chan interface{}

func NewSubscriber() chan interface{} {
	return make(chan interface{})
}
