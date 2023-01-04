package jobsv1

import (
	"clippr/editor"
	"errors"

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
	if w.LoopCount() == 0 {
		return errors.New("Missing Loop Count from workitem metadata")
	}
	newVideo, err := v.video.Loop(w.LoopCount(), v.logger)
	if err != nil {
		return err
	}
	v.video = *newVideo
	// todo Use S3 Client to upload the video somewhere
	return nil
}
