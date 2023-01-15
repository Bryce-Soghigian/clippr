package jobsv1

import (
	"clippr/editor"
	"errors"
	"fmt"
	"log"

	"go.uber.org/zap"
	"k8s.io/client-go/kubernetes"
)

var actionFactories = make(map[string]Factory)

type Interface interface {
	Run(w WorkItem) error
}

// Factory is the default shape we want all actions to have
type Factory func(
	kubeClient kubernetes.Interface,
	video editor.Editor,
	logger *zap.Logger,
	// todo] S3 Client

) (Interface, error)

// GetActionFromFactory Returns a factory object for a given action.
func GetActionFromFactory(name string) (Factory, error) {
	actionFactory, ok := actionFactories[name]
	if !ok {
		return nil, errors.New(
			fmt.Sprintf("invalid action %v", name),
		)
	}
	return actionFactory, nil
}

func RegisterAction(name string, factory Factory) {
	if factory == nil {
		log.Panicf("action factory %s does not exist", name)
	}

	_, registeredPreviously := actionFactories[name]
	if registeredPreviously {
		log.Panicf("action factory %s already registered", name)
	}
	actionFactories[name] = factory
}

// Init is a function that runs at compile time when we build the binary for our application
func init() {
	// Register Actions in the factory here
}
