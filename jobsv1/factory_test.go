package action

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"
	"k8s.io/client-go/kubernetes"
)

func TestGetActionFactory(t *testing.T) {
	// Test with an invalid action name
	_, err = GetActionFromFactory("invalid_action_name")
	assert.Error(t, err)
}

func TestRegister(t *testing.T) {
	// Test with a valid action name and factory
	name := "test_action"
	factory := func(
		kubeClient kubernetes.Interface,
		logger *zap.Logger,
	) (Interface, error) {
		return nil, nil
	}
	RegisterAction(name, factory)

	// Test that the action was registered correctly
	registeredFactory, err := GetActionFromFactory(name)
	assert.NoError(t, err)
	assert.Equal(t, factory, registeredFactory)

	// Test registering the same action name again
	defer func() {
		// Recover from panic and check error message
		r := recover()
		assert.Equal(t, "action factory for test_action already registered", r)
	}()
	RegisterAction(name, factory)
}

func TestInit(t *testing.T) {
	// Test that all the expected actions are registered
}
