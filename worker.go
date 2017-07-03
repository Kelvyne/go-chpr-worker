package worker

import "net/url"

const DEFAULT_EXCHANGE_TYPE = "topic"

type Logger interface {
	Debug()
	Info()
	Warn()
	Error()
}

type NilLogger struct {}
func (L NilLogger) Debug() {}
func (L NilLogger) Info() {}
func (L NilLogger) Warn() {}
func (L NilLogger) Error() {}

type Handler interface {
	Handle()
	Validate()
	RoutingKey() string
}

type Worker struct {
	// mandatory
	Handlers []Handler
	WorkerName string
	AmqpURL url.URL
	ExchangeName string
	QueueName string
	// optional
	Heartbeat int
	TaskTimeout int
	ProcessExitTimeout int
	ChannelPrefetch int
	Logger *Logger
}

func (W *Worker) GetRoutingKeys() []string {
	keys := make([]string, len(W.Handlers))
	for index, handler := range W.Handlers {
		keys[index] = handler.RoutingKey()
	}
	return keys
}
func (W *Worker) Listen() {

}
func (W *Worker) Close() {

}

/*
CreateWorkers creates a worker instance based on configuration.
 */
func CreateWorker(handlers []Handler, config Config, options Options) *Worker {
	options = applyOptionsDefaultValues(options)
	config = applyOptionsDefaultValues(options)
	w := &Worker{
		Handlers:handlers,
	}
	return w
}

