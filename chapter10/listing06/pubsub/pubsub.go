package main

type PubSub struct {
	host string
}

func New(host string) *PubSub {
	ps := PubSub{
		host: host,
	}

	return &ps
}

func (ps *PubSub) Publish(key string, v interface{}) error {

	return nil
}

func (ps *PubSub) Subscribe(key string) error {

	return nil
}
