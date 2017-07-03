package worker

type Options struct {
	Heartbeat uint64
	TaskTimeout uint64
	ProcessExitTimeout uint64
	ChannelPrefetch uint64
	Logger Logger
}
func applyOptionsDefaultValues(options Options) Options {
	result := Options{
		Heartbeat:10,
		TaskTimeout:30000,
		ProcessExitTimeout:3000,
		ChannelPrefetch:100,
		Logger:NilLogger{},
	}
	if options.Heartbeat != 0 {
		result.Heartbeat = options.Heartbeat
	}
	if options.TaskTimeout != 0 {
		result.TaskTimeout = options.TaskTimeout
	}
	if options.ProcessExitTimeout != 0 {
		result.ProcessExitTimeout = options.ProcessExitTimeout
	}
	if options.ChannelPrefetch != 0 {
		result.ChannelPrefetch = options.ChannelPrefetch
	}
	if options.Logger != nil {
		result.Logger = options.Logger
	}
	return result
}
