package reporter

import (
	"encoding/json"
	"time"

	bosherr "github.com/cloudfoundry/bosh-utils/errors"
	bhttp "github.com/cloudfoundry/bosh-utils/http"
	bclient "github.com/cloudfoundry/bosh-utils/httpclient"
	boshlog "github.com/cloudfoundry/bosh-utils/logger"
)

type HTTP struct {
	endpoint string
	client   bclient.HTTPClient

	logTag string
	logger boshlog.Logger
}

func NewHTTP(endpoint string, logger boshlog.Logger) HTTP {
	rawClient := bclient.CreateDefaultClient(nil)
	retryClient := bhttp.NewNetworkSafeRetryClient(rawClient, 5, 500*time.Millisecond, logger)
	httpOpts := bclient.Opts{NoRedactUrlQuery: true}
	httpClient := bclient.NewHTTPClientOpts(retryClient, logger, httpOpts)
	return HTTP{endpoint, httpClient, "stats.reporter.HTTP", logger}
}

type statReq struct {
	Name  string            `json:"name"`
	Value string            `json:"value"`
	Tags  map[string]string `json:"tags,omitempty"`
}

func (r HTTP) Report(stats []Stat) error {
	var reqs []statReq

	for _, stat := range stats {
		req := statReq{
			Name:  stat.Name(),
			Value: stat.Value(),
			Tags:  stat.Tags(),
		}
		reqs = append(reqs, req)
	}

	// todo jsonl format
	body, err := json.Marshal(reqs)
	if err != nil {
		return bosherr.WrapError(err, "Marshaling stats for HTTP")
	}

	// todo add unique header
	_, err = r.client.Post(r.endpoint, body)
	if err != nil {
		return bosherr.WrapError(err, "Sending stats via HTTP")
	}

	return nil
}
