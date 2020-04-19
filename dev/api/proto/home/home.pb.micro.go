// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: github.com/micro/services/dev/api/proto/home/home.proto

package go_micro_api_dev

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

import (
	context "context"
	api "github.com/micro/go-micro/v2/api"
	client "github.com/micro/go-micro/v2/client"
	server "github.com/micro/go-micro/v2/server"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

// Reference imports to suppress errors if they are not otherwise used.
var _ api.Endpoint
var _ context.Context
var _ client.Option
var _ server.Option

// Api Endpoints for Home service

func NewHomeEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for Home service

type HomeService interface {
	ReadUser(ctx context.Context, in *ReadUserRequest, opts ...client.CallOption) (*ReadUserResponse, error)
	ListApps(ctx context.Context, in *ListAppsRequest, opts ...client.CallOption) (*ListAppsResponse, error)
}

type homeService struct {
	c    client.Client
	name string
}

func NewHomeService(name string, c client.Client) HomeService {
	return &homeService{
		c:    c,
		name: name,
	}
}

func (c *homeService) ReadUser(ctx context.Context, in *ReadUserRequest, opts ...client.CallOption) (*ReadUserResponse, error) {
	req := c.c.NewRequest(c.name, "Home.ReadUser", in)
	out := new(ReadUserResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *homeService) ListApps(ctx context.Context, in *ListAppsRequest, opts ...client.CallOption) (*ListAppsResponse, error) {
	req := c.c.NewRequest(c.name, "Home.ListApps", in)
	out := new(ListAppsResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Home service

type HomeHandler interface {
	ReadUser(context.Context, *ReadUserRequest, *ReadUserResponse) error
	ListApps(context.Context, *ListAppsRequest, *ListAppsResponse) error
}

func RegisterHomeHandler(s server.Server, hdlr HomeHandler, opts ...server.HandlerOption) error {
	type home interface {
		ReadUser(ctx context.Context, in *ReadUserRequest, out *ReadUserResponse) error
		ListApps(ctx context.Context, in *ListAppsRequest, out *ListAppsResponse) error
	}
	type Home struct {
		home
	}
	h := &homeHandler{hdlr}
	return s.Handle(s.NewHandler(&Home{h}, opts...))
}

type homeHandler struct {
	HomeHandler
}

func (h *homeHandler) ReadUser(ctx context.Context, in *ReadUserRequest, out *ReadUserResponse) error {
	return h.HomeHandler.ReadUser(ctx, in, out)
}

func (h *homeHandler) ListApps(ctx context.Context, in *ListAppsRequest, out *ListAppsResponse) error {
	return h.HomeHandler.ListApps(ctx, in, out)
}