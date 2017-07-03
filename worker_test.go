package worker

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

const amqpURL = "amqp://guest:guest@localhost:5672"

type VoidHandler struct {
	routingKey string
}
func (H *VoidHandler) Handle() {}
func (H *VoidHandler) Validate() {}
func (H *VoidHandler) RoutingKey() string {
	return H.routingKey
}


func TestWorker_GetRoutingKeys(t *testing.T) {
	w := &Worker{
		Handlers:[]Handler{
			&VoidHandler{routingKey:"key-1"},
			&VoidHandler{routingKey:"key-2"},
		},
	}
	assert.Equal(
		t,
		[]string{"key-1", "key-2"},
		w.GetRoutingKeys(),
	)
}

