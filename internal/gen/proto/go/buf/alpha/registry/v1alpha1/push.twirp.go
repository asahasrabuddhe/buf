// Copyright 2020-2021 Buf Technologies, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by protoc-gen-twirp v8.1.0, DO NOT EDIT.
// source: buf/alpha/registry/v1alpha1/push.proto

package registryv1alpha1

import context "context"
import fmt "fmt"
import http "net/http"
import ioutil "io/ioutil"
import json "encoding/json"
import strconv "strconv"
import strings "strings"

import protojson "google.golang.org/protobuf/encoding/protojson"
import proto "google.golang.org/protobuf/proto"
import twirp "github.com/twitchtv/twirp"
import ctxsetters "github.com/twitchtv/twirp/ctxsetters"

// Version compatibility assertion.
// If the constant is not defined in the package, that likely means
// the package needs to be updated to work with this generated code.
// See https://twitchtv.github.io/twirp/docs/version_matrix.html
const _ = twirp.TwirpPackageMinVersion_8_1_0

// =====================
// PushService Interface
// =====================

// PushService is the Push service.
type PushService interface {
	// Push pushes.
	Push(context.Context, *PushRequest) (*PushResponse, error)
}

// ===========================
// PushService Protobuf Client
// ===========================

type pushServiceProtobufClient struct {
	client      HTTPClient
	urls        [1]string
	interceptor twirp.Interceptor
	opts        twirp.ClientOptions
}

// NewPushServiceProtobufClient creates a Protobuf client that implements the PushService interface.
// It communicates using Protobuf and can be configured with a custom HTTPClient.
func NewPushServiceProtobufClient(baseURL string, client HTTPClient, opts ...twirp.ClientOption) PushService {
	if c, ok := client.(*http.Client); ok {
		client = withoutRedirects(c)
	}

	clientOpts := twirp.ClientOptions{}
	for _, o := range opts {
		o(&clientOpts)
	}

	// Using ReadOpt allows backwards and forwads compatibility with new options in the future
	literalURLs := false
	_ = clientOpts.ReadOpt("literalURLs", &literalURLs)
	var pathPrefix string
	if ok := clientOpts.ReadOpt("pathPrefix", &pathPrefix); !ok {
		pathPrefix = "/twirp" // default prefix
	}

	// Build method URLs: <baseURL>[<prefix>]/<package>.<Service>/<Method>
	serviceURL := sanitizeBaseURL(baseURL)
	serviceURL += baseServicePath(pathPrefix, "buf.alpha.registry.v1alpha1", "PushService")
	urls := [1]string{
		serviceURL + "Push",
	}

	return &pushServiceProtobufClient{
		client:      client,
		urls:        urls,
		interceptor: twirp.ChainInterceptors(clientOpts.Interceptors...),
		opts:        clientOpts,
	}
}

func (c *pushServiceProtobufClient) Push(ctx context.Context, in *PushRequest) (*PushResponse, error) {
	ctx = ctxsetters.WithPackageName(ctx, "buf.alpha.registry.v1alpha1")
	ctx = ctxsetters.WithServiceName(ctx, "PushService")
	ctx = ctxsetters.WithMethodName(ctx, "Push")
	caller := c.callPush
	if c.interceptor != nil {
		caller = func(ctx context.Context, req *PushRequest) (*PushResponse, error) {
			resp, err := c.interceptor(
				func(ctx context.Context, req interface{}) (interface{}, error) {
					typedReq, ok := req.(*PushRequest)
					if !ok {
						return nil, twirp.InternalError("failed type assertion req.(*PushRequest) when calling interceptor")
					}
					return c.callPush(ctx, typedReq)
				},
			)(ctx, req)
			if resp != nil {
				typedResp, ok := resp.(*PushResponse)
				if !ok {
					return nil, twirp.InternalError("failed type assertion resp.(*PushResponse) when calling interceptor")
				}
				return typedResp, err
			}
			return nil, err
		}
	}
	return caller(ctx, in)
}

func (c *pushServiceProtobufClient) callPush(ctx context.Context, in *PushRequest) (*PushResponse, error) {
	out := new(PushResponse)
	ctx, err := doProtobufRequest(ctx, c.client, c.opts.Hooks, c.urls[0], in, out)
	if err != nil {
		twerr, ok := err.(twirp.Error)
		if !ok {
			twerr = twirp.InternalErrorWith(err)
		}
		callClientError(ctx, c.opts.Hooks, twerr)
		return nil, err
	}

	callClientResponseReceived(ctx, c.opts.Hooks)

	return out, nil
}

// =======================
// PushService JSON Client
// =======================

type pushServiceJSONClient struct {
	client      HTTPClient
	urls        [1]string
	interceptor twirp.Interceptor
	opts        twirp.ClientOptions
}

// NewPushServiceJSONClient creates a JSON client that implements the PushService interface.
// It communicates using JSON and can be configured with a custom HTTPClient.
func NewPushServiceJSONClient(baseURL string, client HTTPClient, opts ...twirp.ClientOption) PushService {
	if c, ok := client.(*http.Client); ok {
		client = withoutRedirects(c)
	}

	clientOpts := twirp.ClientOptions{}
	for _, o := range opts {
		o(&clientOpts)
	}

	// Using ReadOpt allows backwards and forwads compatibility with new options in the future
	literalURLs := false
	_ = clientOpts.ReadOpt("literalURLs", &literalURLs)
	var pathPrefix string
	if ok := clientOpts.ReadOpt("pathPrefix", &pathPrefix); !ok {
		pathPrefix = "/twirp" // default prefix
	}

	// Build method URLs: <baseURL>[<prefix>]/<package>.<Service>/<Method>
	serviceURL := sanitizeBaseURL(baseURL)
	serviceURL += baseServicePath(pathPrefix, "buf.alpha.registry.v1alpha1", "PushService")
	urls := [1]string{
		serviceURL + "Push",
	}

	return &pushServiceJSONClient{
		client:      client,
		urls:        urls,
		interceptor: twirp.ChainInterceptors(clientOpts.Interceptors...),
		opts:        clientOpts,
	}
}

func (c *pushServiceJSONClient) Push(ctx context.Context, in *PushRequest) (*PushResponse, error) {
	ctx = ctxsetters.WithPackageName(ctx, "buf.alpha.registry.v1alpha1")
	ctx = ctxsetters.WithServiceName(ctx, "PushService")
	ctx = ctxsetters.WithMethodName(ctx, "Push")
	caller := c.callPush
	if c.interceptor != nil {
		caller = func(ctx context.Context, req *PushRequest) (*PushResponse, error) {
			resp, err := c.interceptor(
				func(ctx context.Context, req interface{}) (interface{}, error) {
					typedReq, ok := req.(*PushRequest)
					if !ok {
						return nil, twirp.InternalError("failed type assertion req.(*PushRequest) when calling interceptor")
					}
					return c.callPush(ctx, typedReq)
				},
			)(ctx, req)
			if resp != nil {
				typedResp, ok := resp.(*PushResponse)
				if !ok {
					return nil, twirp.InternalError("failed type assertion resp.(*PushResponse) when calling interceptor")
				}
				return typedResp, err
			}
			return nil, err
		}
	}
	return caller(ctx, in)
}

func (c *pushServiceJSONClient) callPush(ctx context.Context, in *PushRequest) (*PushResponse, error) {
	out := new(PushResponse)
	ctx, err := doJSONRequest(ctx, c.client, c.opts.Hooks, c.urls[0], in, out)
	if err != nil {
		twerr, ok := err.(twirp.Error)
		if !ok {
			twerr = twirp.InternalErrorWith(err)
		}
		callClientError(ctx, c.opts.Hooks, twerr)
		return nil, err
	}

	callClientResponseReceived(ctx, c.opts.Hooks)

	return out, nil
}

// ==========================
// PushService Server Handler
// ==========================

type pushServiceServer struct {
	PushService
	interceptor      twirp.Interceptor
	hooks            *twirp.ServerHooks
	pathPrefix       string // prefix for routing
	jsonSkipDefaults bool   // do not include unpopulated fields (default values) in the response
	jsonCamelCase    bool   // JSON fields are serialized as lowerCamelCase rather than keeping the original proto names
}

// NewPushServiceServer builds a TwirpServer that can be used as an http.Handler to handle
// HTTP requests that are routed to the right method in the provided svc implementation.
// The opts are twirp.ServerOption modifiers, for example twirp.WithServerHooks(hooks).
func NewPushServiceServer(svc PushService, opts ...interface{}) TwirpServer {
	serverOpts := newServerOpts(opts)

	// Using ReadOpt allows backwards and forwads compatibility with new options in the future
	jsonSkipDefaults := false
	_ = serverOpts.ReadOpt("jsonSkipDefaults", &jsonSkipDefaults)
	jsonCamelCase := false
	_ = serverOpts.ReadOpt("jsonCamelCase", &jsonCamelCase)
	var pathPrefix string
	if ok := serverOpts.ReadOpt("pathPrefix", &pathPrefix); !ok {
		pathPrefix = "/twirp" // default prefix
	}

	return &pushServiceServer{
		PushService:      svc,
		hooks:            serverOpts.Hooks,
		interceptor:      twirp.ChainInterceptors(serverOpts.Interceptors...),
		pathPrefix:       pathPrefix,
		jsonSkipDefaults: jsonSkipDefaults,
		jsonCamelCase:    jsonCamelCase,
	}
}

// writeError writes an HTTP response with a valid Twirp error format, and triggers hooks.
// If err is not a twirp.Error, it will get wrapped with twirp.InternalErrorWith(err)
func (s *pushServiceServer) writeError(ctx context.Context, resp http.ResponseWriter, err error) {
	writeError(ctx, resp, err, s.hooks)
}

// handleRequestBodyError is used to handle error when the twirp server cannot read request
func (s *pushServiceServer) handleRequestBodyError(ctx context.Context, resp http.ResponseWriter, msg string, err error) {
	if context.Canceled == ctx.Err() {
		s.writeError(ctx, resp, twirp.NewError(twirp.Canceled, "failed to read request: context canceled"))
		return
	}
	if context.DeadlineExceeded == ctx.Err() {
		s.writeError(ctx, resp, twirp.NewError(twirp.DeadlineExceeded, "failed to read request: deadline exceeded"))
		return
	}
	s.writeError(ctx, resp, twirp.WrapError(malformedRequestError(msg), err))
}

// PushServicePathPrefix is a convenience constant that may identify URL paths.
// Should be used with caution, it only matches routes generated by Twirp Go clients,
// with the default "/twirp" prefix and default CamelCase service and method names.
// More info: https://twitchtv.github.io/twirp/docs/routing.html
const PushServicePathPrefix = "/twirp/buf.alpha.registry.v1alpha1.PushService/"

func (s *pushServiceServer) ServeHTTP(resp http.ResponseWriter, req *http.Request) {
	ctx := req.Context()
	ctx = ctxsetters.WithPackageName(ctx, "buf.alpha.registry.v1alpha1")
	ctx = ctxsetters.WithServiceName(ctx, "PushService")
	ctx = ctxsetters.WithResponseWriter(ctx, resp)

	var err error
	ctx, err = callRequestReceived(ctx, s.hooks)
	if err != nil {
		s.writeError(ctx, resp, err)
		return
	}

	if req.Method != "POST" {
		msg := fmt.Sprintf("unsupported method %q (only POST is allowed)", req.Method)
		s.writeError(ctx, resp, badRouteError(msg, req.Method, req.URL.Path))
		return
	}

	// Verify path format: [<prefix>]/<package>.<Service>/<Method>
	prefix, pkgService, method := parseTwirpPath(req.URL.Path)
	if pkgService != "buf.alpha.registry.v1alpha1.PushService" {
		msg := fmt.Sprintf("no handler for path %q", req.URL.Path)
		s.writeError(ctx, resp, badRouteError(msg, req.Method, req.URL.Path))
		return
	}
	if prefix != s.pathPrefix {
		msg := fmt.Sprintf("invalid path prefix %q, expected %q, on path %q", prefix, s.pathPrefix, req.URL.Path)
		s.writeError(ctx, resp, badRouteError(msg, req.Method, req.URL.Path))
		return
	}

	switch method {
	case "Push":
		s.servePush(ctx, resp, req)
		return
	default:
		msg := fmt.Sprintf("no handler for path %q", req.URL.Path)
		s.writeError(ctx, resp, badRouteError(msg, req.Method, req.URL.Path))
		return
	}
}

func (s *pushServiceServer) servePush(ctx context.Context, resp http.ResponseWriter, req *http.Request) {
	header := req.Header.Get("Content-Type")
	i := strings.Index(header, ";")
	if i == -1 {
		i = len(header)
	}
	switch strings.TrimSpace(strings.ToLower(header[:i])) {
	case "application/json":
		s.servePushJSON(ctx, resp, req)
	case "application/protobuf":
		s.servePushProtobuf(ctx, resp, req)
	default:
		msg := fmt.Sprintf("unexpected Content-Type: %q", req.Header.Get("Content-Type"))
		twerr := badRouteError(msg, req.Method, req.URL.Path)
		s.writeError(ctx, resp, twerr)
	}
}

func (s *pushServiceServer) servePushJSON(ctx context.Context, resp http.ResponseWriter, req *http.Request) {
	var err error
	ctx = ctxsetters.WithMethodName(ctx, "Push")
	ctx, err = callRequestRouted(ctx, s.hooks)
	if err != nil {
		s.writeError(ctx, resp, err)
		return
	}

	d := json.NewDecoder(req.Body)
	rawReqBody := json.RawMessage{}
	if err := d.Decode(&rawReqBody); err != nil {
		s.handleRequestBodyError(ctx, resp, "the json request could not be decoded", err)
		return
	}
	reqContent := new(PushRequest)
	unmarshaler := protojson.UnmarshalOptions{DiscardUnknown: true}
	if err = unmarshaler.Unmarshal(rawReqBody, reqContent); err != nil {
		s.handleRequestBodyError(ctx, resp, "the json request could not be decoded", err)
		return
	}

	handler := s.PushService.Push
	if s.interceptor != nil {
		handler = func(ctx context.Context, req *PushRequest) (*PushResponse, error) {
			resp, err := s.interceptor(
				func(ctx context.Context, req interface{}) (interface{}, error) {
					typedReq, ok := req.(*PushRequest)
					if !ok {
						return nil, twirp.InternalError("failed type assertion req.(*PushRequest) when calling interceptor")
					}
					return s.PushService.Push(ctx, typedReq)
				},
			)(ctx, req)
			if resp != nil {
				typedResp, ok := resp.(*PushResponse)
				if !ok {
					return nil, twirp.InternalError("failed type assertion resp.(*PushResponse) when calling interceptor")
				}
				return typedResp, err
			}
			return nil, err
		}
	}

	// Call service method
	var respContent *PushResponse
	func() {
		defer ensurePanicResponses(ctx, resp, s.hooks)
		respContent, err = handler(ctx, reqContent)
	}()

	if err != nil {
		s.writeError(ctx, resp, err)
		return
	}
	if respContent == nil {
		s.writeError(ctx, resp, twirp.InternalError("received a nil *PushResponse and nil error while calling Push. nil responses are not supported"))
		return
	}

	ctx = callResponsePrepared(ctx, s.hooks)

	marshaler := &protojson.MarshalOptions{UseProtoNames: !s.jsonCamelCase, EmitUnpopulated: !s.jsonSkipDefaults}
	respBytes, err := marshaler.Marshal(respContent)
	if err != nil {
		s.writeError(ctx, resp, wrapInternal(err, "failed to marshal json response"))
		return
	}

	ctx = ctxsetters.WithStatusCode(ctx, http.StatusOK)
	resp.Header().Set("Content-Type", "application/json")
	resp.Header().Set("Content-Length", strconv.Itoa(len(respBytes)))
	resp.WriteHeader(http.StatusOK)

	if n, err := resp.Write(respBytes); err != nil {
		msg := fmt.Sprintf("failed to write response, %d of %d bytes written: %s", n, len(respBytes), err.Error())
		twerr := twirp.NewError(twirp.Unknown, msg)
		ctx = callError(ctx, s.hooks, twerr)
	}
	callResponseSent(ctx, s.hooks)
}

func (s *pushServiceServer) servePushProtobuf(ctx context.Context, resp http.ResponseWriter, req *http.Request) {
	var err error
	ctx = ctxsetters.WithMethodName(ctx, "Push")
	ctx, err = callRequestRouted(ctx, s.hooks)
	if err != nil {
		s.writeError(ctx, resp, err)
		return
	}

	buf, err := ioutil.ReadAll(req.Body)
	if err != nil {
		s.handleRequestBodyError(ctx, resp, "failed to read request body", err)
		return
	}
	reqContent := new(PushRequest)
	if err = proto.Unmarshal(buf, reqContent); err != nil {
		s.writeError(ctx, resp, malformedRequestError("the protobuf request could not be decoded"))
		return
	}

	handler := s.PushService.Push
	if s.interceptor != nil {
		handler = func(ctx context.Context, req *PushRequest) (*PushResponse, error) {
			resp, err := s.interceptor(
				func(ctx context.Context, req interface{}) (interface{}, error) {
					typedReq, ok := req.(*PushRequest)
					if !ok {
						return nil, twirp.InternalError("failed type assertion req.(*PushRequest) when calling interceptor")
					}
					return s.PushService.Push(ctx, typedReq)
				},
			)(ctx, req)
			if resp != nil {
				typedResp, ok := resp.(*PushResponse)
				if !ok {
					return nil, twirp.InternalError("failed type assertion resp.(*PushResponse) when calling interceptor")
				}
				return typedResp, err
			}
			return nil, err
		}
	}

	// Call service method
	var respContent *PushResponse
	func() {
		defer ensurePanicResponses(ctx, resp, s.hooks)
		respContent, err = handler(ctx, reqContent)
	}()

	if err != nil {
		s.writeError(ctx, resp, err)
		return
	}
	if respContent == nil {
		s.writeError(ctx, resp, twirp.InternalError("received a nil *PushResponse and nil error while calling Push. nil responses are not supported"))
		return
	}

	ctx = callResponsePrepared(ctx, s.hooks)

	respBytes, err := proto.Marshal(respContent)
	if err != nil {
		s.writeError(ctx, resp, wrapInternal(err, "failed to marshal proto response"))
		return
	}

	ctx = ctxsetters.WithStatusCode(ctx, http.StatusOK)
	resp.Header().Set("Content-Type", "application/protobuf")
	resp.Header().Set("Content-Length", strconv.Itoa(len(respBytes)))
	resp.WriteHeader(http.StatusOK)
	if n, err := resp.Write(respBytes); err != nil {
		msg := fmt.Sprintf("failed to write response, %d of %d bytes written: %s", n, len(respBytes), err.Error())
		twerr := twirp.NewError(twirp.Unknown, msg)
		ctx = callError(ctx, s.hooks, twerr)
	}
	callResponseSent(ctx, s.hooks)
}

func (s *pushServiceServer) ServiceDescriptor() ([]byte, int) {
	return twirpFileDescriptor5, 0
}

func (s *pushServiceServer) ProtocGenTwirpVersion() string {
	return "v8.1.0"
}

// PathPrefix returns the base service path, in the form: "/<prefix>/<package>.<Service>/"
// that is everything in a Twirp route except for the <Method>. This can be used for routing,
// for example to identify the requests that are targeted to this service in a mux.
func (s *pushServiceServer) PathPrefix() string {
	return baseServicePath(s.pathPrefix, "buf.alpha.registry.v1alpha1", "PushService")
}

var twirpFileDescriptor5 = []byte{
	// 347 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x52, 0xcd, 0x4a, 0x33, 0x31,
	0x14, 0x65, 0xbe, 0xfe, 0x40, 0xd3, 0x0f, 0x91, 0x20, 0x32, 0x54, 0xd0, 0xda, 0x85, 0x8c, 0x08,
	0x09, 0xad, 0x2b, 0x71, 0xd7, 0xb5, 0x42, 0x19, 0x71, 0x53, 0x17, 0x65, 0x32, 0x4d, 0x67, 0x22,
	0x69, 0x12, 0xf3, 0x53, 0xe9, 0x1b, 0xf8, 0x2c, 0x3e, 0xa5, 0x4c, 0xd2, 0xa1, 0xe3, 0x66, 0x70,
	0x97, 0x73, 0x72, 0xcf, 0x39, 0x37, 0xf7, 0x06, 0xdc, 0x10, 0xb7, 0xc1, 0x19, 0x57, 0x65, 0x86,
	0x35, 0x2d, 0x98, 0xb1, 0x7a, 0x8f, 0x77, 0x53, 0x4f, 0x4c, 0xb1, 0x72, 0xa6, 0x44, 0x4a, 0x4b,
	0x2b, 0xe1, 0x05, 0x71, 0x1b, 0xe4, 0x69, 0x54, 0xd7, 0xa1, 0xba, 0x6e, 0xd4, 0x30, 0xd9, 0xca,
	0xb5, 0xe3, 0xf4, 0x68, 0x11, 0x70, 0x30, 0x19, 0x25, 0x6d, 0x61, 0xcd, 0xca, 0xc9, 0x77, 0x04,
	0x86, 0x0b, 0x67, 0xca, 0x94, 0x7e, 0x38, 0x6a, 0x2c, 0x3c, 0x03, 0x3d, 0xf9, 0x29, 0xa8, 0x8e,
	0xa3, 0x71, 0x94, 0x0c, 0xd2, 0x00, 0xe0, 0x25, 0x00, 0x9a, 0x2a, 0x69, 0x98, 0x95, 0x7a, 0x1f,
	0xff, 0xf3, 0x57, 0x0d, 0x06, 0x9e, 0x83, 0x3e, 0xd1, 0x99, 0xc8, 0xcb, 0xb8, 0xe3, 0xef, 0x0e,
	0x08, 0x3e, 0x80, 0x7e, 0x48, 0x8b, 0xbb, 0xe3, 0x28, 0x19, 0xce, 0xae, 0xd1, 0xf1, 0x75, 0x87,
	0x36, 0xea, 0xb6, 0xd0, 0xb3, 0xc7, 0xe9, 0x41, 0x00, 0x21, 0xe8, 0xda, 0xac, 0x30, 0x71, 0x6f,
	0xdc, 0x49, 0x06, 0xa9, 0x3f, 0x4f, 0x28, 0xf8, 0x1f, 0x7a, 0x35, 0x4a, 0x0a, 0x43, 0xe1, 0x2b,
	0x38, 0xe5, 0x32, 0xcf, 0xf8, 0x2a, 0x68, 0x56, 0x8a, 0x89, 0xb8, 0xe7, 0x83, 0xee, 0x50, 0xcb,
	0x18, 0xd1, 0x53, 0x25, 0x0a, 0x79, 0x0b, 0x26, 0xd2, 0x13, 0xfe, 0x0b, 0xcf, 0xde, 0xc3, 0x48,
	0x5e, 0xa8, 0xde, 0xb1, 0x9c, 0xc2, 0x37, 0xd0, 0xad, 0x20, 0x4c, 0x5a, 0x3d, 0x1b, 0x43, 0x1c,
	0xdd, 0xfe, 0xa1, 0x32, 0x3c, 0x61, 0xfe, 0x15, 0x81, 0xab, 0x5c, 0x6e, 0xdb, 0x04, 0xf3, 0x41,
	0xa5, 0x58, 0x54, 0xeb, 0x5a, 0x2e, 0x0b, 0x66, 0x4b, 0x47, 0x50, 0x2e, 0xb7, 0x98, 0xb8, 0x0d,
	0x71, 0x8c, 0xaf, 0xab, 0x03, 0x66, 0xc2, 0x52, 0x2d, 0x32, 0x8e, 0x0b, 0x2a, 0xb0, 0x5f, 0x2d,
	0x2e, 0x24, 0x6e, 0xf9, 0x06, 0x8f, 0x35, 0x53, 0x13, 0xa4, 0xef, 0x65, 0xf7, 0x3f, 0x01, 0x00,
	0x00, 0xff, 0xff, 0xc0, 0x5b, 0x65, 0xb9, 0xaa, 0x02, 0x00, 0x00,
}
