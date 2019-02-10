package main

import (
	"encoding/json"
	"net"
	"net/http"
	"strings"
	"time"

	"github.com/hashicorp/raft"
	"github.com/justinas/alice"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/hlog"
)

type httpServer struct {
	address net.Addr
	node    *node
	logger  *zerolog.Logger
}

func (s *httpServer) Start() {
	s.logger.Info().Str("address", s.address.String()).Msg("Starting Server")
	c := alice.New()
	c = c.Append(hlog.AccessHandler(func(r *http.Request, status, size int, duration time.Duration) {
		hlog.FromRequest(r).Info().
			Str("req.method", r.Method).
			Str("req.url", r.URL.String()).
			Int("req.status", status).
			Int("req.size", size).
			Dur("req.duration", duration).
			Msg("")
	}))

	c = c.Append(hlog.RemoteAddrHandler("req.ip"))
	c = c.Append(hlog.UserAgentHandler("req.useragent"))
	c = c.Append(hlog.RefererHandler("req.referer"))
	c = c.Append(hlog.RequestIDHandler("req.id", "Request-Id"))
	handler := c.Then(s)

	if err := http.ListenAndServe(s.address.String(), handler); err != nil {
		s.logger.Fatal().Err(err).Msg("Error running HTTP s")
	}
}

func (s *httpServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	if strings.Contains(r.URL.Path, "/key") {
		s.handleRequest(w, r)
	} else if strings.Contains(r.URL.Path, "/join") {
		s.handleJoin(w, r)
	} else {
		w.WriteHeader(http.StatusBadRequest)
	}
}

func (s *httpServer) handleRequest(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		s.handleKeyPost(w, r)
	case http.MethodGet:
		s.handleKeyGet(w, r)
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func (s *httpServer) handleKeyPost(w http.ResponseWriter, r *http.Request) {
	request := struct {
		NewValue int `json:"newValue"`
	}{}

	defer r.Body.Close()
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		s.logger.Error().Err(err).Msg("Bad request")
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	event := &event{
		Type:  "set",
		Value: request.NewValue,
	}

	eventBytes, err := json.Marshal(event)
	if err != nil {
		s.logger.Error().Err(err).Msg("")
	}

	applyFuture := s.node.raftNode.Apply(eventBytes, 5*time.Second)
	if err := applyFuture.Error(); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (s *httpServer) handleKeyGet(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	response := struct {
		Value int `json:"value"`
	}{
		Value: s.node.fsm.stateValue,
	}
	responseBytes, err := json.Marshal(response)
	if err != nil {
		s.logger.Error().Err(err).Msg("")
	}
	w.Write(responseBytes)
}

func (s *httpServer) handleJoin(w http.ResponseWriter, r *http.Request) {
	peerAddress := r.Header.Get("Peer-Address")
	if peerAddress == "" {
		s.logger.Error().Msg("Peer-Address not set on request")
		w.WriteHeader(http.StatusBadRequest)
	}

	addPeerFuture := s.node.raftNode.AddVoter(raft.ServerID(peerAddress), raft.ServerAddress(peerAddress), 0, 0)
	if err := addPeerFuture.Error(); err != nil {
		s.logger.Error().
			Err(err).
			Str("peer.remoteaddr", peerAddress).
			Msg("Error joining peer to Raft")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	s.logger.Info().Str("peer.remoteaddr", peerAddress).Msg("Peer joined Raft")
	w.WriteHeader(http.StatusOK)
}
