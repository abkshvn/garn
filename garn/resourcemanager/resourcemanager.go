package resourcemanager

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type ResourceManager struct {
	Hostname string
	Port     int
}

func (r ResourceManager) query(path string, v interface{}) error {
	url := fmt.Sprintf("http://%s:%d/%s", r.Hostname, r.Port, path)
	response, err := http.Get(url)
	if err != nil {
		return fmt.Errorf("http endpoint unavailable: %v", err)
	}
	defer response.Body.Close()
	data, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return fmt.Errorf("could not read response body: %v", err)
	}

	err = json.Unmarshal(data, &v)
	if err != nil {
		return fmt.Errorf("could not unmarshal JSON response: %v", err)
	}
	return nil
}

func (r ResourceManager) ClusterInformation() (*ClusterInformation, error) {
	path := "ws/v1/cluster/info"
	responseWrapper := ClusterInformationWrapper{}
	err := r.query(path, &responseWrapper)
	if err != nil {
		return nil, fmt.Errorf("could not fetch cluster information: %v", err)
	}
	return &responseWrapper.ClusterInformation, nil
}

func (r ResourceManager) ClusterMetrics() (*ClusterMetrics, error) {
	path := "ws/v1/cluster/metrics"
	responseWrapper := ClusterMetricsWrapper{}
	err := r.query(path, &responseWrapper)
	if err != nil {
		return nil, fmt.Errorf("could not fetch cluster metrics: %v", err)
	}
	return &responseWrapper.ClusterMetrics, nil
}
