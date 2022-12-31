package jobsv1

import (
	"clippr/editor"
	"go.uber.org/zap"
	"k8s.io/client-go/kubernetes"
)

// Interface is the behaviors we expect from all runners
type Interface interface {
	Run(*WorkItem) error
}

// Runner is the type reponsible for running work items
type Runner struct {
	logger     *zap.Logger
	kubeClient kubernetes.Interface
	video      editor.Video
}

func (r *Runner) Run(w *WorkItem) error {
	action, err := GetActionFromFactory(w.Action)
	if err != nil {
		r.logger.Error("Failed to retrieve action from action factory")
		return err
	}
	err = action.Run(w)
	if err != nil {
		return err
	}
	return nil
}
