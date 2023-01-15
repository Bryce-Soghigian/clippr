package jobsv1

import (
	"clippr/editor"
	"go.uber.org/zap"
	"k8s.io/client-go/kubernetes"
)

// Runner is the type reponsible for running work items
type Runner struct {
	logger     *zap.Logger
	kubeClient kubernetes.Interface
	video      editor.Editor
}

func NewRunner(kubeClient kubernetes.Interface, video editor.Editor, logger *zap.Logger) *Runner {
	return &Runner{
		logger:     logger,
		kubeClient: kubeClient,
		video:      video,
	}
}

func (r *Runner) Run(w WorkItem) error {
	newAction, err := GetActionFromFactory(w.Action)
	if err != nil {
		r.logger.Error("Failed to retrieve action from action factory")
		return err
	}
	action, err := newAction(r.kubeClient, r.video, r.logger)
	if err != nil {
		return err
	}
	err = action.Run(w)
	if err != nil {
		return err
	}
	return nil
}
