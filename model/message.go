package model

type EventMessageBlob[T any] struct {
	EventType string `json:"eventType"`
	Uri       string `json:"uri"`
	Data      T      `json:"data"`
}
