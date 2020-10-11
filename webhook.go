package plugin

import (
	"context"
	"net/http"
)

// Config the plugin configuration.
type Config struct {
	URLS map[string]string `json:"webhookNotifier,omitempty"`
}

// CreateConfig creates the default plugin configuration.
func CreateConfig() *Config {
	return &Config{}
}

// Example a plugin.
type webhookNotifier struct {
	next http.Handler
	name string
	// ...
}

// New create a new webhook notifier plugin.
func New(ctx context.Context, next http.Handler, config *Config, name string) (http.Handler, error) {
	// ...
	return &webhookNotifier{
		// ...
	}, nil
}

func (e *webhookNotifier) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	// ...
	e.next.ServeHTTP(rw, req)
}
