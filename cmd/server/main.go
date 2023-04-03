package main

import (
	"context"
	"log"
	"net/http"

	"github.com/bufbuild/connect-go"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"

	apiV1 "api/gen/api/v1"
	apiV1Connect "api/gen/api/v1/apiv1connect"
)

type InstallationService struct{}

func (s *InstallationService) RegisterInstallation(
    ctx context.Context,
    req *connect.Request[apiV1.RegisterArenaRequest],
) (*connect.Response[apiV1.EmptyResponse], error) {
    log.Println("Request: ", req.Msg)
    res := connect.NewResponse(&apiV1.EmptyResponse{})
    return res, nil
}

func (s *InstallationService) NotifyGameCreated(
    ctx context.Context,
    req *connect.Request[apiV1.GameEvent],
) (*connect.Response[apiV1.EmptyResponse], error) {
    log.Println("Request: ", req.Msg)
    res := connect.NewResponse(&apiV1.EmptyResponse{})
    return res, nil
}

func (s *InstallationService) NotifyGameStarted(
    ctx context.Context,
    req *connect.Request[apiV1.GameEvent],
) (*connect.Response[apiV1.EmptyResponse], error) {
    log.Println("Request: ", req.Msg)
    res := connect.NewResponse(&apiV1.EmptyResponse{})
    return res, nil
}

func (s *InstallationService) NotifyGameEnded(
    ctx context.Context,
    req *connect.Request[apiV1.GameEvent],
) (*connect.Response[apiV1.EmptyResponse], error) {
    log.Println("Request: ", req.Msg)
    res := connect.NewResponse(&apiV1.EmptyResponse{})
    return res, nil
}

func main() {
    installationService := &InstallationService{}
    mux := http.NewServeMux()
    path, handler := apiV1Connect.NewInstallationServiceHandler(installationService)
    mux.Handle(path, handler)
    log.Println("Server is running on localhost:8080");
    http.ListenAndServe(
        "localhost:8080",
        h2c.NewHandler(mux, &http2.Server{}),
    )
}
