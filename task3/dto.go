package main

type Message[T any] interface {
	GetContent() T
}

type HttpMessage[T any] struct {
	content T
}

type WebsocketMessage[T any] struct {
	content T
}

func (hm HttpMessage[T]) GetContent() T {
	return hm.content
}

func (wm WebsocketMessage[T]) GetContent() T {
	return wm.content
}
