package worker

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

type CustomLogger struct {}
func (L CustomLogger) Debug() {}
func (L CustomLogger) Info() {}
func (L CustomLogger) Warn() {}
func (L CustomLogger) Error() {}

func TestApplyOptionsDefaultValues_Default(t *testing.T)  {
	options := applyOptionsDefaultValues(Options{})

	assert.Equal(t, uint64(10), options.Heartbeat)
	assert.Equal(t, uint64(30000), options.TaskTimeout)
	assert.Equal(t, uint64(3000), options.ProcessExitTimeout)
	assert.Equal(t, uint64(100), options.ChannelPrefetch)
	assert.IsType(t, NilLogger{}, options.Logger)
}

func TestApplyOptionsDefaultValues_Specified(t *testing.T)  {
	options := applyOptionsDefaultValues(Options{
		Heartbeat:1,
		TaskTimeout:1,
		ProcessExitTimeout:1,
		ChannelPrefetch:1,
		Logger:CustomLogger{},
	})

	assert.Equal(t, uint64(1), options.Heartbeat)
	assert.Equal(t, uint64(1), options.TaskTimeout)
	assert.Equal(t, uint64(1), options.ProcessExitTimeout)
	assert.Equal(t, uint64(1), options.ChannelPrefetch)
	assert.IsType(t, CustomLogger{}, options.Logger)
}
