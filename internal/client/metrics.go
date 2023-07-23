package client

import (
	"GSS/internal/metrics"
	"log"

	"github.com/google/uuid"
)

type UserMetricStorage struct {
	UUID          uuid.UUID             `json:"uuid"`
	MetricStorage metrics.MetricStorage `json:"metrics"`
}

func (user *User) sendUserMetric(url string, request UserMetricStorage) (err error) {
	agent := user.Client.Post(url)
	err = agent.JSON(request).Parse()
	if err != nil {
		return
	}

	statusCode, _, errs := agent.String()
	log.Println("Send JSON:", statusCode)
	if errs != nil {
		return errs[0]
	}
	return
}

func (user *User) SendMetricStorage() (err error) {
	mStorage, err := metrics.Get()
	if err != nil {
		return
	}

	request := UserMetricStorage{
		UUID:          user.UUID,
		MetricStorage: mStorage,
	}
	err = user.sendUserMetric("http://"+user.Config.ServerAddr, request)
	if err != nil {
		return
	}

	log.Printf("Sending metrics...")
	return
}

func (user *User) StreamingMetrics() (err error) {
	mStorage, err := metrics.Get()
	if err != nil {
		return
	}

	timestamp := mStorage.Timestamp
	for metricKey, metricVal := range mStorage.MetricMapFloat64 {
		request := UserMetricStorage{
			UUID: user.UUID,
			MetricStorage: metrics.MetricStorage{
				Timestamp: timestamp,
				MetricMapFloat64: metrics.MetricMapFloat64{
					metricKey: metricVal,
				},
				MetricMapUint32: metrics.MetricMapUint32{},
				MetricMapUint64: metrics.MetricMapUint64{},
			},
		}
		err := user.sendUserMetric("http://"+user.Config.ServerAddr, request)
		if err != nil {
			return err
		}
		log.Printf("Sending metric '%s'...", metricKey)
	}

	for metricKey, metricVal := range mStorage.MetricMapUint32 {
		request := UserMetricStorage{
			UUID: user.UUID,
			MetricStorage: metrics.MetricStorage{
				Timestamp:        timestamp,
				MetricMapFloat64: metrics.MetricMapFloat64{},
				MetricMapUint32: metrics.MetricMapUint32{
					metricKey: metricVal,
				},
				MetricMapUint64: metrics.MetricMapUint64{},
			},
		}
		err := user.sendUserMetric("http://"+user.Config.ServerAddr, request)
		if err != nil {
			return err
		}
		log.Printf("Sending metric '%s'...", metricKey)
	}

	for metricKey, metricVal := range mStorage.MetricMapUint64 {
		request := UserMetricStorage{
			UUID: user.UUID,
			MetricStorage: metrics.MetricStorage{
				Timestamp:        timestamp,
				MetricMapFloat64: metrics.MetricMapFloat64{},
				MetricMapUint32:  metrics.MetricMapUint32{},
				MetricMapUint64: metrics.MetricMapUint64{
					metricKey: metricVal,
				},
			},
		}
		err := user.sendUserMetric("http://"+user.Config.ServerAddr, request)
		if err != nil {
			return err
		}
		log.Printf("Sending metric '%s'...", metricKey)
	}

	return
}
