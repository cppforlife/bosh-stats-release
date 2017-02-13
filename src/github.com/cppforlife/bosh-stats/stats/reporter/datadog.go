package reporter

import (
	"time"

	bosherr "github.com/cloudfoundry/bosh-utils/errors"
	boshlog "github.com/cloudfoundry/bosh-utils/logger"
	datadog "github.com/zorkian/go-datadog-api"
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
	event := &datadog.Event{
		Title: "",
		Text:  "",
		Time:  int(time.Now().Unix()),

		Priority:  "normal",
		AlertType: "info",

		Host:        "",
		Aggregation: "",
		SourceType:  "",

		Tags:     nil,
		Resource: "",
	}

	_, err := r.client.PostEvent(event)

	return err
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
