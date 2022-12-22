package server

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
	"time"

	"go.uber.org/zap"
	"k8s.io/apimachinery/pkg/util/wait"
	"k8s.io/client-go/util/workqueue"
	"k8s.io/sample-controller/pkg/signals"
)

type ClipprServer struct {
	workQueue workqueue.RateLimitingInterface
	logger    *zap.Logger
	// [TODO] Add S3 Client here
	// [TODO] Add KubeClient here
}

func NewServer(logger *zap.Logger) *ClipprServer {
	return &ClipprServer{
		workQueue: workqueue.NewNamedRateLimitingQueue(workqueue.DefaultControllerRateLimiter(), ""),
		logger:    logger,
	}
}

func (s *ClipprServer) Run() error {
	defer s.workQueue.ShutDown()
	stopSignal := signals.SetupSignalHandler()
	s.logger.Info("Starting workqueue worker")
	// Run all of our work in a separate goroutine
	go wait.Until(s.runWorkItemsProcessor, time.Second, stopSignal)
	http.HandleFunc("/recieveWorkitem", s.handleIngestWorkItem)
	server := &http.Server{
		Addr:              ":5555",
		ReadHeaderTimeout: time.Second,
	}
	go server.ListenAndServe()

	err := server.Shutdown(context.Background())
	return err
}

type WorkItem struct {
	actionName string
	metadata   map[string]string
}

func (s *ClipprServer) handleIngestWorkItem(response http.ResponseWriter, request *http.Request) {
	body, err := io.ReadAll(request.Body)
	if err != nil {
		http.Error(
			response,
			http.StatusText(http.StatusBadRequest),
			http.StatusBadRequest,
		)
		return
	}
	var workItems []*WorkItem
	err = json.Unmarshal(body, workItems)
	if err != nil {
		s.logger.Error(err.Error())
		return
	}

	for _, workitem := range workItems {
		s.workQueue.AddRateLimited(workitem)
	}

}

func (s *ClipprServer) runWorkItemsProcessor() {
	// Process work items forever
	for s.processFrontWorkItem() {
	}
}

func (s *ClipprServer) processFrontWorkItem() bool {
	// TODO: Do something like this https://cs.github.com/kubernetes/kubernetes/blob/2bb77a13b1a709e0208c6ba9ba5323a5e54f79d6/pkg/controller/statefulset/stateful_set.go?q=workqueue.newNamedRate#L414
	return true
}
