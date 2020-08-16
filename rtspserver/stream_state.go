package rtspserver

import "github.com/jhondev/dorsvr/livemedia"

type StreamServerState struct {
	subsession  livemedia.IServerMediaSubsession
	streamToken *livemedia.StreamState
}
