// Copyright 2025 CloudWeGo Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by Kitex v0.13.1. DO NOT EDIT.

package streamingservice

import (
	"context"
	client "github.com/cloudwego/kitex/client"
	streamcall "github.com/cloudwego/kitex/client/callopt/streamcall"
	streamclient "github.com/cloudwego/kitex/client/streamclient"
	streaming "github.com/cloudwego/kitex/pkg/streaming"
	transport "github.com/cloudwego/kitex/transport"
	streaming0 "github.com/kitex-contrib/kitexcall/internal/test/server/kitex_gen/streaming"
)

// Client is designed to provide IDL-compatible methods with call-option parameter for kitex framework.
type Client interface {
}

// StreamClient is designed to provide Interface for Streaming APIs.
type StreamClient interface {
	BidirectionalStream(ctx context.Context, callOptions ...streamcall.Option) (stream StreamingService_BidirectionalStreamClient, err error)
	ServerStream(ctx context.Context, req *streaming0.Message, callOptions ...streamcall.Option) (stream StreamingService_ServerStreamClient, err error)
	ClientStream(ctx context.Context, callOptions ...streamcall.Option) (stream StreamingService_ClientStreamClient, err error)
}

type StreamingService_BidirectionalStreamClient interface {
	streaming.Stream
	Send(*streaming0.Message) error
	Recv() (*streaming0.StreamingResponse, error)
}

type StreamingService_ServerStreamClient interface {
	streaming.Stream
	Recv() (*streaming0.StreamingResponse, error)
}

type StreamingService_ClientStreamClient interface {
	streaming.Stream
	Send(*streaming0.Message) error
	CloseAndRecv() (*streaming0.StreamingResponse, error)
}

// NewClient creates a client for the service defined in IDL.
func NewClient(destService string, opts ...client.Option) (Client, error) {
	var options []client.Option
	options = append(options, client.WithDestService(destService))

	options = append(options, opts...)

	kc, err := client.NewClient(serviceInfoForClient(), options...)
	if err != nil {
		return nil, err
	}
	return &kStreamingServiceClient{
		kClient: newServiceClient(kc),
	}, nil
}

// MustNewClient creates a client for the service defined in IDL. It panics if any error occurs.
func MustNewClient(destService string, opts ...client.Option) Client {
	kc, err := NewClient(destService, opts...)
	if err != nil {
		panic(err)
	}
	return kc
}

type kStreamingServiceClient struct {
	*kClient
}

// NewStreamClient creates a stream client for the service's streaming APIs defined in IDL.
func NewStreamClient(destService string, opts ...streamclient.Option) (StreamClient, error) {
	var options []client.Option
	options = append(options, client.WithDestService(destService))
	options = append(options, client.WithTransportProtocol(transport.GRPC))
	options = append(options, streamclient.GetClientOptions(opts)...)

	kc, err := client.NewClient(serviceInfoForStreamClient(), options...)
	if err != nil {
		return nil, err
	}
	return &kStreamingServiceStreamClient{
		kClient: newServiceClient(kc),
	}, nil
}

// MustNewStreamClient creates a stream client for the service's streaming APIs defined in IDL.
// It panics if any error occurs.
func MustNewStreamClient(destService string, opts ...streamclient.Option) StreamClient {
	kc, err := NewStreamClient(destService, opts...)
	if err != nil {
		panic(err)
	}
	return kc
}

type kStreamingServiceStreamClient struct {
	*kClient
}

func (p *kStreamingServiceStreamClient) BidirectionalStream(ctx context.Context, callOptions ...streamcall.Option) (stream StreamingService_BidirectionalStreamClient, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, streamcall.GetCallOptions(callOptions))
	return p.kClient.BidirectionalStream(ctx)
}

func (p *kStreamingServiceStreamClient) ServerStream(ctx context.Context, req *streaming0.Message, callOptions ...streamcall.Option) (stream StreamingService_ServerStreamClient, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, streamcall.GetCallOptions(callOptions))
	return p.kClient.ServerStream(ctx, req)
}

func (p *kStreamingServiceStreamClient) ClientStream(ctx context.Context, callOptions ...streamcall.Option) (stream StreamingService_ClientStreamClient, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, streamcall.GetCallOptions(callOptions))
	return p.kClient.ClientStream(ctx)
}
