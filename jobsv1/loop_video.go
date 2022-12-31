package jobsv1

import (
	"clippr/editor"

	"go.uber.org/zap"
	"k8s.io/client-go/kubernetes"
)

var VIDEO_LOOPER = "VideoLooperAction"

type VideoLooper struct {
	logger     *zap.Logger
	kubeClient kubernetes.Interface
	video      editor.Video
}

func NewVideoLooper(logger *zap.Logger, client kubernetes.Interface, video editor.Video) *VideoLooper {
	return &VideoLooper{
		logger:     logger,
		kubeClient: client,
		video:      video,
	}
}

func (v *VideoLooper) Run(w WorkItem) error {

	return nil
}
