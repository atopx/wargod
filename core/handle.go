package core

import (
	"fmt"
	"strings"
)

type handle struct {
	uri         string
	name        string
	subscribe   []byte
	unsubscribe []byte
	prefix      []byte
	handler     func(data []byte) error
}

func newHandler(uri string, handler func(data []byte) error) *handle {
	name := strings.ReplaceAll(strings.TrimPrefix(uri, "/"), "/", "_")
	return &handle{
		uri:         uri,
		name:        name,
		subscribe:   []byte(fmt.Sprintf(`[5, "OnJsonApiEvent_%s"]`, name)),
		unsubscribe: []byte(fmt.Sprintf(`[6, "OnJsonApiEvent_%s"]`, name)),
		prefix:      []byte(fmt.Sprintf(`[8,"OnJsonApiEvent_%s"`, name)),
		handler:     handler,
	}
}
