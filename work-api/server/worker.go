package server

import (
	"clippr/editor"
	"clippr/jobsv1"
	"errors"
)

func (s *ClipprServer) runWorkItemsProcessor() {
	// Process work items forever
	for s.processFrontWorkItem() {
	}
}

func (s *ClipprServer) processFrontWorkItem() bool {
	// TODO: Do something like this https://cs.github.com/kubernetes/kubernetes/blob/2bb77a13b1a709e0208c6ba9ba5323a5e54f79d6/pkg/controller/statefulset/stateful_set.go?q=workqueue.newNamedRate#L414
	item, shutdown := s.workQueue.Get()
	if shutdown {
		return false
	}
	err := s.processActionFromWorkItem(item)
	if err != nil {
		s.logger.Error(err.Error())
	}
	return true
}

func (s *ClipprServer) processActionFromWorkItem(item interface{}) error {
	defer s.workQueue.Done(item)
	var workItem jobsv1.WorkItem
	var ok bool
	workItem, ok = item.(jobsv1.WorkItem)
	if !ok {
		s.workQueue.Forget(item)
		return errors.New("Failed to convert item to workitem")
	}

	err := s.TriggerRunner(workItem)
	if err != nil {
		return err
	}
	return nil

}

func (s *ClipprServer) TriggerRunner(w jobsv1.WorkItem) error {
	// Retrieve Video from S3 and pass it to the runner here.
	runner := jobsv1.NewRunner(s.kubeClient, editor.Video{}, s.logger)
	return runner.Run(w)

}
