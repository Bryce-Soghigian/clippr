package jobsv1

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
	return w.Metadata["videoUrl"]
}

// VideoHostPlatform will give us context into the type of video format we are recieving.
func (w *WorkItem) VideoHostPlatform() string {
	if w.Metadata == nil {
		return ""
	}

	return w.Metadata["host"]

}
