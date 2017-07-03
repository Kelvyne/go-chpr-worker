package worker

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"net/url"
)

func TestNewConfig(t *testing.T) {
	config := NewConfig(
		"myWorker",
		"amqp://user:pwd@localhost:5672",
		"myExchange",
		"myQueue",
	)
	assert.Equal(t, "myWorker", config.GetWorkerName())
	assert.Equal(t, url.Parse("amqp://user:pwd@localhost:5672"), config.GetAmqpURL())
	assert.Equal(t, "myExchange", config.GetExchangeName())
	assert.Equal(t, "myQueue", config.GetQueueName())
}
