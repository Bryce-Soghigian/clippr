package jobsv1

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetActionFactory(t *testing.T) {
	// Test with an invalid action name
	_, err := GetActionFromFactory("invalid_action_name")
	assert.Error(t, err)
}

func TestInit(t *testing.T) {
	// Test that all the expected actions are registered
}
