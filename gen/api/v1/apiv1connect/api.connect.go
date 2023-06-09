// Code generated by protoc-gen-connect-go. DO NOT EDIT.
//
// Source: api/v1/api.proto

package apiv1connect

import (
	context "context"
	errors "errors"
	v1 "api/gen/api/v1"
	connect_go "github.com/bufbuild/connect-go"
	http "net/http"
	strings "strings"
)

// This is a compile-time assertion to ensure that this generated file and the connect package are
// compatible. If you get a compiler error that this constant is not defined, this code was
// generated with a version of connect newer than the one compiled into your binary. You can fix the
// problem by either regenerating this code with an older version of connect or updating the connect
// version compiled into your binary.
const _ = connect_go.IsAtLeastVersion0_1_0

const (
	// InstallationServiceName is the fully-qualified name of the InstallationService service.
	InstallationServiceName = "api.v1.InstallationService"
)

// InstallationServiceClient is a client for the api.v1.InstallationService service.
type InstallationServiceClient interface {
	RegisterInstallation(context.Context, *connect_go.Request[v1.RegisterArenaRequest]) (*connect_go.Response[v1.EmptyResponse], error)
	NotifyGameCreated(context.Context, *connect_go.Request[v1.GameEvent]) (*connect_go.Response[v1.EmptyResponse], error)
	NotifyGameStarted(context.Context, *connect_go.Request[v1.GameEvent]) (*connect_go.Response[v1.EmptyResponse], error)
	NotifyGameEnded(context.Context, *connect_go.Request[v1.GameEvent]) (*connect_go.Response[v1.EmptyResponse], error)
}

// NewInstallationServiceClient constructs a client for the api.v1.InstallationService service. By
// default, it uses the Connect protocol with the binary Protobuf Codec, asks for gzipped responses,
// and sends uncompressed requests. To use the gRPC or gRPC-Web protocols, supply the
// connect.WithGRPC() or connect.WithGRPCWeb() options.
//
// The URL supplied here should be the base URL for the Connect or gRPC server (for example,
// http://api.acme.com or https://acme.com/grpc).
func NewInstallationServiceClient(httpClient connect_go.HTTPClient, baseURL string, opts ...connect_go.ClientOption) InstallationServiceClient {
	baseURL = strings.TrimRight(baseURL, "/")
	return &installationServiceClient{
		registerInstallation: connect_go.NewClient[v1.RegisterArenaRequest, v1.EmptyResponse](
			httpClient,
			baseURL+"/api.v1.InstallationService/RegisterInstallation",
			opts...,
		),
		notifyGameCreated: connect_go.NewClient[v1.GameEvent, v1.EmptyResponse](
			httpClient,
			baseURL+"/api.v1.InstallationService/NotifyGameCreated",
			opts...,
		),
		notifyGameStarted: connect_go.NewClient[v1.GameEvent, v1.EmptyResponse](
			httpClient,
			baseURL+"/api.v1.InstallationService/NotifyGameStarted",
			opts...,
		),
		notifyGameEnded: connect_go.NewClient[v1.GameEvent, v1.EmptyResponse](
			httpClient,
			baseURL+"/api.v1.InstallationService/NotifyGameEnded",
			opts...,
		),
	}
}

// installationServiceClient implements InstallationServiceClient.
type installationServiceClient struct {
	registerInstallation *connect_go.Client[v1.RegisterArenaRequest, v1.EmptyResponse]
	notifyGameCreated    *connect_go.Client[v1.GameEvent, v1.EmptyResponse]
	notifyGameStarted    *connect_go.Client[v1.GameEvent, v1.EmptyResponse]
	notifyGameEnded      *connect_go.Client[v1.GameEvent, v1.EmptyResponse]
}

// RegisterInstallation calls api.v1.InstallationService.RegisterInstallation.
func (c *installationServiceClient) RegisterInstallation(ctx context.Context, req *connect_go.Request[v1.RegisterArenaRequest]) (*connect_go.Response[v1.EmptyResponse], error) {
	return c.registerInstallation.CallUnary(ctx, req)
}

// NotifyGameCreated calls api.v1.InstallationService.NotifyGameCreated.
func (c *installationServiceClient) NotifyGameCreated(ctx context.Context, req *connect_go.Request[v1.GameEvent]) (*connect_go.Response[v1.EmptyResponse], error) {
	return c.notifyGameCreated.CallUnary(ctx, req)
}

// NotifyGameStarted calls api.v1.InstallationService.NotifyGameStarted.
func (c *installationServiceClient) NotifyGameStarted(ctx context.Context, req *connect_go.Request[v1.GameEvent]) (*connect_go.Response[v1.EmptyResponse], error) {
	return c.notifyGameStarted.CallUnary(ctx, req)
}

// NotifyGameEnded calls api.v1.InstallationService.NotifyGameEnded.
func (c *installationServiceClient) NotifyGameEnded(ctx context.Context, req *connect_go.Request[v1.GameEvent]) (*connect_go.Response[v1.EmptyResponse], error) {
	return c.notifyGameEnded.CallUnary(ctx, req)
}

// InstallationServiceHandler is an implementation of the api.v1.InstallationService service.
type InstallationServiceHandler interface {
	RegisterInstallation(context.Context, *connect_go.Request[v1.RegisterArenaRequest]) (*connect_go.Response[v1.EmptyResponse], error)
	NotifyGameCreated(context.Context, *connect_go.Request[v1.GameEvent]) (*connect_go.Response[v1.EmptyResponse], error)
	NotifyGameStarted(context.Context, *connect_go.Request[v1.GameEvent]) (*connect_go.Response[v1.EmptyResponse], error)
	NotifyGameEnded(context.Context, *connect_go.Request[v1.GameEvent]) (*connect_go.Response[v1.EmptyResponse], error)
}

// NewInstallationServiceHandler builds an HTTP handler from the service implementation. It returns
// the path on which to mount the handler and the handler itself.
//
// By default, handlers support the Connect, gRPC, and gRPC-Web protocols with the binary Protobuf
// and JSON codecs. They also support gzip compression.
func NewInstallationServiceHandler(svc InstallationServiceHandler, opts ...connect_go.HandlerOption) (string, http.Handler) {
	mux := http.NewServeMux()
	mux.Handle("/api.v1.InstallationService/RegisterInstallation", connect_go.NewUnaryHandler(
		"/api.v1.InstallationService/RegisterInstallation",
		svc.RegisterInstallation,
		opts...,
	))
	mux.Handle("/api.v1.InstallationService/NotifyGameCreated", connect_go.NewUnaryHandler(
		"/api.v1.InstallationService/NotifyGameCreated",
		svc.NotifyGameCreated,
		opts...,
	))
	mux.Handle("/api.v1.InstallationService/NotifyGameStarted", connect_go.NewUnaryHandler(
		"/api.v1.InstallationService/NotifyGameStarted",
		svc.NotifyGameStarted,
		opts...,
	))
	mux.Handle("/api.v1.InstallationService/NotifyGameEnded", connect_go.NewUnaryHandler(
		"/api.v1.InstallationService/NotifyGameEnded",
		svc.NotifyGameEnded,
		opts...,
	))
	return "/api.v1.InstallationService/", mux
}

// UnimplementedInstallationServiceHandler returns CodeUnimplemented from all methods.
type UnimplementedInstallationServiceHandler struct{}

func (UnimplementedInstallationServiceHandler) RegisterInstallation(context.Context, *connect_go.Request[v1.RegisterArenaRequest]) (*connect_go.Response[v1.EmptyResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("api.v1.InstallationService.RegisterInstallation is not implemented"))
}

func (UnimplementedInstallationServiceHandler) NotifyGameCreated(context.Context, *connect_go.Request[v1.GameEvent]) (*connect_go.Response[v1.EmptyResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("api.v1.InstallationService.NotifyGameCreated is not implemented"))
}

func (UnimplementedInstallationServiceHandler) NotifyGameStarted(context.Context, *connect_go.Request[v1.GameEvent]) (*connect_go.Response[v1.EmptyResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("api.v1.InstallationService.NotifyGameStarted is not implemented"))
}

func (UnimplementedInstallationServiceHandler) NotifyGameEnded(context.Context, *connect_go.Request[v1.GameEvent]) (*connect_go.Response[v1.EmptyResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("api.v1.InstallationService.NotifyGameEnded is not implemented"))
}
