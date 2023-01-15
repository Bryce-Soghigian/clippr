package editor

import (
	ffmpeg_go "github.com/u2takey/ffmpeg-go"
	"go.uber.org/zap"
)

type Editor struct {
	Video *ffmpeg_go.Stream
}

func NewEditor(Video *ffmpeg_go.Stream) *Editor {
	return &Editor{Video: Video}
}

func (e Editor) CopyContent() *ffmpeg_go.Stream {
	cp := e
	return cp.Video
}

func (e *Editor) Loop(n int64, logger *zap.Logger) (*Editor, error) {
	finalCut := []*ffmpeg_go.Stream{}
	times := int(n)
	for i := 0; i < times; i++ {
		finalCut = append(finalCut, e.CopyContent())
	}
	Video := ffmpeg_go.Concat(finalCut)
	e.Video = Video
	return e, nil
}

func (e *Editor) WriteToFile(path string) error {
	err := e.Video.Output(path).Run()
	return err
}
