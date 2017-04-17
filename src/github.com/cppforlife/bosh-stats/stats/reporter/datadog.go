package reporter

import (
	"fmt"
	"strconv"
	"time"

	bosherr "github.com/cloudfoundry/bosh-utils/errors"
	boshlog "github.com/cloudfoundry/bosh-utils/logger"
	datadog "gopkg.in/zorkian/go-datadog-api.v1"
)

type DatadogConfig struct {
	APIKey string
	AppKey string
}

type Datadog struct {
	client *datadog.Client

	logTag string
	logger boshlog.Logger
}

func NewDatadog(config DatadogConfig, logger boshlog.Logger) Datadog {
	return Datadog{
		client: datadog.NewClient(config.APIKey, config.AppKey),

		logTag: "stats.reporter.Datadog",
		logger: logger,
	}
}

func (r Datadog) Report(stats []Stat) error {
	metrics := []datadog.Metric{}
	for _, stat := range stats {
		tags := []string{}
		for key, value := range stat.Tags() {
			tags = append(tags, fmt.Sprintf("%s:%s", key, value))
		}
		value, err := strconv.Atoi(stat.Value())
		if err != nil {
			r.logger.Error(r.logTag, fmt.Sprintf("Could not convert value to integer: %v", err.Error()))
		}

		metric := datadog.Metric{
			Metric: stat.Name(),
			Points: []datadog.DataPoint{{float64(time.Now().Unix()), float64(value)}},
			Type:   "info",
			Host:   "director",
			Tags:   tags,
		}
		metrics = append(metrics, metric)
	}

	return r.client.PostMetrics(metrics)
}

func (c DatadogConfig) Required() bool { return len(c.AppKey) > 0 }

func (c DatadogConfig) Validate() error {
	if !c.Required() {
		return nil
	}

	if len(c.APIKey) == 0 {
		return bosherr.Error("Missing 'APIKey'")
	}

	return nil
}
