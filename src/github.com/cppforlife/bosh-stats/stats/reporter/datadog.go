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

// name="director.version" value="260.5.0 (00000000)" tags="map[string]string(nil)"
// name="director.uuid" value="69f9a704-676b-4aaf-81db-5e2c3f1b87ce" tags="map[string]string(nil)"
// name="director.auth.type" value="basic" tags="map[string]string(nil)"
// name="director.cpi" value="warden_cpi" tags="map[string]string(nil)"
// name="azs.names" value="z1,z2,z3" tags="map[string]string(nil)"

// name="release" value="true" tags="map[string]string{"name":"bosh-stats", "version":"0+dev.1492881605"}"
// name="stemcell" value="true" tags="map[string]string{"version":"3363.19", "name":"bosh-warden-boshlite-ubuntu-trusty-go_agent"}"

// name="azs.count" value="3" tags="map[string]string(nil)"
// name="disk_types.count" value="1" tags="map[string]string(nil)"
// name="vm_types.count" value="1" tags="map[string]string(nil)"
// name="networks.count" value="1" tags="map[string]string(nil)"
// name="networks.manual.count" value="1" tags="map[string]string(nil)"
// name="networks.dynamic.count" value="0" tags="map[string]string(nil)"
// name="networks.vip.count" value="0" tags="map[string]string(nil)"
// name="compilation.workers" value="5" tags="map[string]string(nil)"
// name="releases.count" value="8" tags="map[string]string(nil)"
// name="stemcells.count" value="1" tags="map[string]string(nil)"
// name="deployments.count" value="2" tags="map[string]string(nil)"
// name="instances.count" value="6" tags="map[string]string(nil)"

// name="deployment.instances.count" value="1" tags="map[string]string{"deployment":"stats"}"
// name="deployment.instances.count" value="5" tags="map[string]string{"deployment":"zookeeper"}"

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
			Metric: fmt.Sprintf("bosh.stats.%s", stat.Name()),
			Points: []datadog.DataPoint{{float64(time.Now().Unix()), float64(value)}},
			Type:   "gauge",
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
