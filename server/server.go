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
