package newrelic

import (
	"context"
	"errors"
	"os"

	opencensus "github.com/devopsfaith/krakend-opencensus"
	"github.com/newrelic/newrelic-opencensus-exporter-go/nrcensus"
)

func init() {
	opencensus.RegisterExporterFactories(func(ctx context.Context, cfg opencensus.Config) (interface{}, error) {
		return Exporter(ctx, cfg)
	})
}

func Exporter(ctx context.Context, cfg opencensus.Config) (*nrcensus.Exporter, error) {
	if cfg.Exporters.NewRelic == nil {
		return nil, errDisabled
	}
	key := os.Getenv(cfg.Exporters.NewRelic.ApiKeyVar)
	if key == "" {
		return nil, errApiKey
	}
	e, err := nrcensus.NewExporter(cfg.Exporters.NewRelic.Service, key)
	if err != nil {
		return e, err
	}

	return e, nil

}

var errDisabled = errors.New("opencensus New Relic exporter disabled")
var errApiKey = errors.New("no New Relic API key provided")
