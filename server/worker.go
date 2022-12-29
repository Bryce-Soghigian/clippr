package server

func (s *ClipprServer) runWorkItemsProcessor() {
	// Process work items forever
	for s.processFrontWorkItem() {
	}
}

func (s *ClipprServer) processFrontWorkItem() bool {
	// TODO: Do something like this https://cs.github.com/kubernetes/kubernetes/blob/2bb77a13b1a709e0208c6ba9ba5323a5e54f79d6/pkg/controller/statefulset/stateful_set.go?q=workqueue.newNamedRate#L414
	return true
}
