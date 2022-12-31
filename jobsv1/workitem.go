package jobsv1

import (
	"log"
	"strconv"
)

const (
	LoopCountKey = "loopCount"
	VideoUrlKey  = "videoUrl"
	HostKey      = "host"
)

type WorkItem struct {
	Action   string
	ID       string
	Metadata map[string]string
}

// VideoUrl is a link to the orignal source content
func (w *WorkItem) VideoUrl() string {
	if w.Metadata == nil {
		return ""
	}
	return w.Metadata[VideoUrlKey]
}

// VideoHostPlatform will give us context into the type of video format we are recieving.
func (w *WorkItem) VideoHostPlatform() string {
	if w.Metadata == nil {
		return ""
	}

	return w.Metadata[HostKey]
}

// LoopCount is used to determine the number of times you would
// like to loop a video
func (w *WorkItem) LoopCount() int64 {
	if w.Metadata == nil {
		return 0
	}
	if _, ok := w.Metadata[LoopCountKey]; !ok {
		return 0
	}
	base, bitSize := 10, 32
	count, err := strconv.ParseInt(
		w.Metadata[LoopCountKey], base, bitSize,
	)
	if err != nil {
		log.Print(err.Error())
		return 0
	}
	return count
}
