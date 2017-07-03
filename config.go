package worker

import "net/url"

type Config interface {
	GetWorkerName() string
	GetAmqpURL() url.URL
	GetExchangeName() string
	GetQueueName() string
}
type configImpl struct {
	workerName string
	amqpUrl url.URL
	exchangeName string
	queueName string
}
func (c *configImpl) GetWorkerName() string {
	return c.workerName
}
func (c *configImpl) GetAmqpURL() url.URL {
	return c.amqpUrl
}
func (c *configImpl) GetExchangeName() string {
	return c.exchangeName
}
func (c *configImpl) GetQueueName() string {
	return c.queueName
}

func NewConfig(workerName string, amqpUrl url.URL, exchangeName url.URL, queueName url.URL) Config {
	return &configImpl{
		workerName:workerName,
		amqpUrl:amqpUrl,
		exchangeName:exchangeName,
		queueName:queueName,
	}
}