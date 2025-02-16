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
// source: buf/alpha/registry/v1alpha1/repository_tag.proto

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

// ==============================
// RepositoryTagService Interface
// ==============================

// RepositoryTagService is the Repository tag service.
type RepositoryTagService interface {
	// CreateRepositoryTag creates a new repository tag.
	CreateRepositoryTag(context.Context, *CreateRepositoryTagRequest) (*CreateRepositoryTagResponse, error)

	// ListRepositoryTags lists the repository tags associated with a Repository.
	ListRepositoryTags(context.Context, *ListRepositoryTagsRequest) (*ListRepositoryTagsResponse, error)
}

// ====================================
// RepositoryTagService Protobuf Client
// ====================================

type repositoryTagServiceProtobufClient struct {
	client      HTTPClient
	urls        [2]string
	interceptor twirp.Interceptor
	opts        twirp.ClientOptions
}

// NewRepositoryTagServiceProtobufClient creates a Protobuf client that implements the RepositoryTagService interface.
// It communicates using Protobuf and can be configured with a custom HTTPClient.
func NewRepositoryTagServiceProtobufClient(baseURL string, client HTTPClient, opts ...twirp.ClientOption) RepositoryTagService {
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
	serviceURL += baseServicePath(pathPrefix, "buf.alpha.registry.v1alpha1", "RepositoryTagService")
	urls := [2]string{
		serviceURL + "CreateRepositoryTag",
		serviceURL + "ListRepositoryTags",
	}

	return &repositoryTagServiceProtobufClient{
		client:      client,
		urls:        urls,
		interceptor: twirp.ChainInterceptors(clientOpts.Interceptors...),
		opts:        clientOpts,
	}
}

func (c *repositoryTagServiceProtobufClient) CreateRepositoryTag(ctx context.Context, in *CreateRepositoryTagRequest) (*CreateRepositoryTagResponse, error) {
	ctx = ctxsetters.WithPackageName(ctx, "buf.alpha.registry.v1alpha1")
	ctx = ctxsetters.WithServiceName(ctx, "RepositoryTagService")
	ctx = ctxsetters.WithMethodName(ctx, "CreateRepositoryTag")
	caller := c.callCreateRepositoryTag
	if c.interceptor != nil {
		caller = func(ctx context.Context, req *CreateRepositoryTagRequest) (*CreateRepositoryTagResponse, error) {
			resp, err := c.interceptor(
				func(ctx context.Context, req interface{}) (interface{}, error) {
					typedReq, ok := req.(*CreateRepositoryTagRequest)
					if !ok {
						return nil, twirp.InternalError("failed type assertion req.(*CreateRepositoryTagRequest) when calling interceptor")
					}
					return c.callCreateRepositoryTag(ctx, typedReq)
				},
			)(ctx, req)
			if resp != nil {
				typedResp, ok := resp.(*CreateRepositoryTagResponse)
				if !ok {
					return nil, twirp.InternalError("failed type assertion resp.(*CreateRepositoryTagResponse) when calling interceptor")
				}
				return typedResp, err
			}
			return nil, err
		}
	}
	return caller(ctx, in)
}

func (c *repositoryTagServiceProtobufClient) callCreateRepositoryTag(ctx context.Context, in *CreateRepositoryTagRequest) (*CreateRepositoryTagResponse, error) {
	out := new(CreateRepositoryTagResponse)
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

func (c *repositoryTagServiceProtobufClient) ListRepositoryTags(ctx context.Context, in *ListRepositoryTagsRequest) (*ListRepositoryTagsResponse, error) {
	ctx = ctxsetters.WithPackageName(ctx, "buf.alpha.registry.v1alpha1")
	ctx = ctxsetters.WithServiceName(ctx, "RepositoryTagService")
	ctx = ctxsetters.WithMethodName(ctx, "ListRepositoryTags")
	caller := c.callListRepositoryTags
	if c.interceptor != nil {
		caller = func(ctx context.Context, req *ListRepositoryTagsRequest) (*ListRepositoryTagsResponse, error) {
			resp, err := c.interceptor(
				func(ctx context.Context, req interface{}) (interface{}, error) {
					typedReq, ok := req.(*ListRepositoryTagsRequest)
					if !ok {
						return nil, twirp.InternalError("failed type assertion req.(*ListRepositoryTagsRequest) when calling interceptor")
					}
					return c.callListRepositoryTags(ctx, typedReq)
				},
			)(ctx, req)
			if resp != nil {
				typedResp, ok := resp.(*ListRepositoryTagsResponse)
				if !ok {
					return nil, twirp.InternalError("failed type assertion resp.(*ListRepositoryTagsResponse) when calling interceptor")
				}
				return typedResp, err
			}
			return nil, err
		}
	}
	return caller(ctx, in)
}

func (c *repositoryTagServiceProtobufClient) callListRepositoryTags(ctx context.Context, in *ListRepositoryTagsRequest) (*ListRepositoryTagsResponse, error) {
	out := new(ListRepositoryTagsResponse)
	ctx, err := doProtobufRequest(ctx, c.client, c.opts.Hooks, c.urls[1], in, out)
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

// ================================
// RepositoryTagService JSON Client
// ================================

type repositoryTagServiceJSONClient struct {
	client      HTTPClient
	urls        [2]string
	interceptor twirp.Interceptor
	opts        twirp.ClientOptions
}

// NewRepositoryTagServiceJSONClient creates a JSON client that implements the RepositoryTagService interface.
// It communicates using JSON and can be configured with a custom HTTPClient.
func NewRepositoryTagServiceJSONClient(baseURL string, client HTTPClient, opts ...twirp.ClientOption) RepositoryTagService {
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
	serviceURL += baseServicePath(pathPrefix, "buf.alpha.registry.v1alpha1", "RepositoryTagService")
	urls := [2]string{
		serviceURL + "CreateRepositoryTag",
		serviceURL + "ListRepositoryTags",
	}

	return &repositoryTagServiceJSONClient{
		client:      client,
		urls:        urls,
		interceptor: twirp.ChainInterceptors(clientOpts.Interceptors...),
		opts:        clientOpts,
	}
}

func (c *repositoryTagServiceJSONClient) CreateRepositoryTag(ctx context.Context, in *CreateRepositoryTagRequest) (*CreateRepositoryTagResponse, error) {
	ctx = ctxsetters.WithPackageName(ctx, "buf.alpha.registry.v1alpha1")
	ctx = ctxsetters.WithServiceName(ctx, "RepositoryTagService")
	ctx = ctxsetters.WithMethodName(ctx, "CreateRepositoryTag")
	caller := c.callCreateRepositoryTag
	if c.interceptor != nil {
		caller = func(ctx context.Context, req *CreateRepositoryTagRequest) (*CreateRepositoryTagResponse, error) {
			resp, err := c.interceptor(
				func(ctx context.Context, req interface{}) (interface{}, error) {
					typedReq, ok := req.(*CreateRepositoryTagRequest)
					if !ok {
						return nil, twirp.InternalError("failed type assertion req.(*CreateRepositoryTagRequest) when calling interceptor")
					}
					return c.callCreateRepositoryTag(ctx, typedReq)
				},
			)(ctx, req)
			if resp != nil {
				typedResp, ok := resp.(*CreateRepositoryTagResponse)
				if !ok {
					return nil, twirp.InternalError("failed type assertion resp.(*CreateRepositoryTagResponse) when calling interceptor")
				}
				return typedResp, err
			}
			return nil, err
		}
	}
	return caller(ctx, in)
}

func (c *repositoryTagServiceJSONClient) callCreateRepositoryTag(ctx context.Context, in *CreateRepositoryTagRequest) (*CreateRepositoryTagResponse, error) {
	out := new(CreateRepositoryTagResponse)
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

func (c *repositoryTagServiceJSONClient) ListRepositoryTags(ctx context.Context, in *ListRepositoryTagsRequest) (*ListRepositoryTagsResponse, error) {
	ctx = ctxsetters.WithPackageName(ctx, "buf.alpha.registry.v1alpha1")
	ctx = ctxsetters.WithServiceName(ctx, "RepositoryTagService")
	ctx = ctxsetters.WithMethodName(ctx, "ListRepositoryTags")
	caller := c.callListRepositoryTags
	if c.interceptor != nil {
		caller = func(ctx context.Context, req *ListRepositoryTagsRequest) (*ListRepositoryTagsResponse, error) {
			resp, err := c.interceptor(
				func(ctx context.Context, req interface{}) (interface{}, error) {
					typedReq, ok := req.(*ListRepositoryTagsRequest)
					if !ok {
						return nil, twirp.InternalError("failed type assertion req.(*ListRepositoryTagsRequest) when calling interceptor")
					}
					return c.callListRepositoryTags(ctx, typedReq)
				},
			)(ctx, req)
			if resp != nil {
				typedResp, ok := resp.(*ListRepositoryTagsResponse)
				if !ok {
					return nil, twirp.InternalError("failed type assertion resp.(*ListRepositoryTagsResponse) when calling interceptor")
				}
				return typedResp, err
			}
			return nil, err
		}
	}
	return caller(ctx, in)
}

func (c *repositoryTagServiceJSONClient) callListRepositoryTags(ctx context.Context, in *ListRepositoryTagsRequest) (*ListRepositoryTagsResponse, error) {
	out := new(ListRepositoryTagsResponse)
	ctx, err := doJSONRequest(ctx, c.client, c.opts.Hooks, c.urls[1], in, out)
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

// ===================================
// RepositoryTagService Server Handler
// ===================================

type repositoryTagServiceServer struct {
	RepositoryTagService
	interceptor      twirp.Interceptor
	hooks            *twirp.ServerHooks
	pathPrefix       string // prefix for routing
	jsonSkipDefaults bool   // do not include unpopulated fields (default values) in the response
	jsonCamelCase    bool   // JSON fields are serialized as lowerCamelCase rather than keeping the original proto names
}

// NewRepositoryTagServiceServer builds a TwirpServer that can be used as an http.Handler to handle
// HTTP requests that are routed to the right method in the provided svc implementation.
// The opts are twirp.ServerOption modifiers, for example twirp.WithServerHooks(hooks).
func NewRepositoryTagServiceServer(svc RepositoryTagService, opts ...interface{}) TwirpServer {
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

	return &repositoryTagServiceServer{
		RepositoryTagService: svc,
		hooks:                serverOpts.Hooks,
		interceptor:          twirp.ChainInterceptors(serverOpts.Interceptors...),
		pathPrefix:           pathPrefix,
		jsonSkipDefaults:     jsonSkipDefaults,
		jsonCamelCase:        jsonCamelCase,
	}
}

// writeError writes an HTTP response with a valid Twirp error format, and triggers hooks.
// If err is not a twirp.Error, it will get wrapped with twirp.InternalErrorWith(err)
func (s *repositoryTagServiceServer) writeError(ctx context.Context, resp http.ResponseWriter, err error) {
	writeError(ctx, resp, err, s.hooks)
}

// handleRequestBodyError is used to handle error when the twirp server cannot read request
func (s *repositoryTagServiceServer) handleRequestBodyError(ctx context.Context, resp http.ResponseWriter, msg string, err error) {
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

// RepositoryTagServicePathPrefix is a convenience constant that may identify URL paths.
// Should be used with caution, it only matches routes generated by Twirp Go clients,
// with the default "/twirp" prefix and default CamelCase service and method names.
// More info: https://twitchtv.github.io/twirp/docs/routing.html
const RepositoryTagServicePathPrefix = "/twirp/buf.alpha.registry.v1alpha1.RepositoryTagService/"

func (s *repositoryTagServiceServer) ServeHTTP(resp http.ResponseWriter, req *http.Request) {
	ctx := req.Context()
	ctx = ctxsetters.WithPackageName(ctx, "buf.alpha.registry.v1alpha1")
	ctx = ctxsetters.WithServiceName(ctx, "RepositoryTagService")
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
	if pkgService != "buf.alpha.registry.v1alpha1.RepositoryTagService" {
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
	case "CreateRepositoryTag":
		s.serveCreateRepositoryTag(ctx, resp, req)
		return
	case "ListRepositoryTags":
		s.serveListRepositoryTags(ctx, resp, req)
		return
	default:
		msg := fmt.Sprintf("no handler for path %q", req.URL.Path)
		s.writeError(ctx, resp, badRouteError(msg, req.Method, req.URL.Path))
		return
	}
}

func (s *repositoryTagServiceServer) serveCreateRepositoryTag(ctx context.Context, resp http.ResponseWriter, req *http.Request) {
	header := req.Header.Get("Content-Type")
	i := strings.Index(header, ";")
	if i == -1 {
		i = len(header)
	}
	switch strings.TrimSpace(strings.ToLower(header[:i])) {
	case "application/json":
		s.serveCreateRepositoryTagJSON(ctx, resp, req)
	case "application/protobuf":
		s.serveCreateRepositoryTagProtobuf(ctx, resp, req)
	default:
		msg := fmt.Sprintf("unexpected Content-Type: %q", req.Header.Get("Content-Type"))
		twerr := badRouteError(msg, req.Method, req.URL.Path)
		s.writeError(ctx, resp, twerr)
	}
}

func (s *repositoryTagServiceServer) serveCreateRepositoryTagJSON(ctx context.Context, resp http.ResponseWriter, req *http.Request) {
	var err error
	ctx = ctxsetters.WithMethodName(ctx, "CreateRepositoryTag")
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
	reqContent := new(CreateRepositoryTagRequest)
	unmarshaler := protojson.UnmarshalOptions{DiscardUnknown: true}
	if err = unmarshaler.Unmarshal(rawReqBody, reqContent); err != nil {
		s.handleRequestBodyError(ctx, resp, "the json request could not be decoded", err)
		return
	}

	handler := s.RepositoryTagService.CreateRepositoryTag
	if s.interceptor != nil {
		handler = func(ctx context.Context, req *CreateRepositoryTagRequest) (*CreateRepositoryTagResponse, error) {
			resp, err := s.interceptor(
				func(ctx context.Context, req interface{}) (interface{}, error) {
					typedReq, ok := req.(*CreateRepositoryTagRequest)
					if !ok {
						return nil, twirp.InternalError("failed type assertion req.(*CreateRepositoryTagRequest) when calling interceptor")
					}
					return s.RepositoryTagService.CreateRepositoryTag(ctx, typedReq)
				},
			)(ctx, req)
			if resp != nil {
				typedResp, ok := resp.(*CreateRepositoryTagResponse)
				if !ok {
					return nil, twirp.InternalError("failed type assertion resp.(*CreateRepositoryTagResponse) when calling interceptor")
				}
				return typedResp, err
			}
			return nil, err
		}
	}

	// Call service method
	var respContent *CreateRepositoryTagResponse
	func() {
		defer ensurePanicResponses(ctx, resp, s.hooks)
		respContent, err = handler(ctx, reqContent)
	}()

	if err != nil {
		s.writeError(ctx, resp, err)
		return
	}
	if respContent == nil {
		s.writeError(ctx, resp, twirp.InternalError("received a nil *CreateRepositoryTagResponse and nil error while calling CreateRepositoryTag. nil responses are not supported"))
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

func (s *repositoryTagServiceServer) serveCreateRepositoryTagProtobuf(ctx context.Context, resp http.ResponseWriter, req *http.Request) {
	var err error
	ctx = ctxsetters.WithMethodName(ctx, "CreateRepositoryTag")
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
	reqContent := new(CreateRepositoryTagRequest)
	if err = proto.Unmarshal(buf, reqContent); err != nil {
		s.writeError(ctx, resp, malformedRequestError("the protobuf request could not be decoded"))
		return
	}

	handler := s.RepositoryTagService.CreateRepositoryTag
	if s.interceptor != nil {
		handler = func(ctx context.Context, req *CreateRepositoryTagRequest) (*CreateRepositoryTagResponse, error) {
			resp, err := s.interceptor(
				func(ctx context.Context, req interface{}) (interface{}, error) {
					typedReq, ok := req.(*CreateRepositoryTagRequest)
					if !ok {
						return nil, twirp.InternalError("failed type assertion req.(*CreateRepositoryTagRequest) when calling interceptor")
					}
					return s.RepositoryTagService.CreateRepositoryTag(ctx, typedReq)
				},
			)(ctx, req)
			if resp != nil {
				typedResp, ok := resp.(*CreateRepositoryTagResponse)
				if !ok {
					return nil, twirp.InternalError("failed type assertion resp.(*CreateRepositoryTagResponse) when calling interceptor")
				}
				return typedResp, err
			}
			return nil, err
		}
	}

	// Call service method
	var respContent *CreateRepositoryTagResponse
	func() {
		defer ensurePanicResponses(ctx, resp, s.hooks)
		respContent, err = handler(ctx, reqContent)
	}()

	if err != nil {
		s.writeError(ctx, resp, err)
		return
	}
	if respContent == nil {
		s.writeError(ctx, resp, twirp.InternalError("received a nil *CreateRepositoryTagResponse and nil error while calling CreateRepositoryTag. nil responses are not supported"))
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

func (s *repositoryTagServiceServer) serveListRepositoryTags(ctx context.Context, resp http.ResponseWriter, req *http.Request) {
	header := req.Header.Get("Content-Type")
	i := strings.Index(header, ";")
	if i == -1 {
		i = len(header)
	}
	switch strings.TrimSpace(strings.ToLower(header[:i])) {
	case "application/json":
		s.serveListRepositoryTagsJSON(ctx, resp, req)
	case "application/protobuf":
		s.serveListRepositoryTagsProtobuf(ctx, resp, req)
	default:
		msg := fmt.Sprintf("unexpected Content-Type: %q", req.Header.Get("Content-Type"))
		twerr := badRouteError(msg, req.Method, req.URL.Path)
		s.writeError(ctx, resp, twerr)
	}
}

func (s *repositoryTagServiceServer) serveListRepositoryTagsJSON(ctx context.Context, resp http.ResponseWriter, req *http.Request) {
	var err error
	ctx = ctxsetters.WithMethodName(ctx, "ListRepositoryTags")
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
	reqContent := new(ListRepositoryTagsRequest)
	unmarshaler := protojson.UnmarshalOptions{DiscardUnknown: true}
	if err = unmarshaler.Unmarshal(rawReqBody, reqContent); err != nil {
		s.handleRequestBodyError(ctx, resp, "the json request could not be decoded", err)
		return
	}

	handler := s.RepositoryTagService.ListRepositoryTags
	if s.interceptor != nil {
		handler = func(ctx context.Context, req *ListRepositoryTagsRequest) (*ListRepositoryTagsResponse, error) {
			resp, err := s.interceptor(
				func(ctx context.Context, req interface{}) (interface{}, error) {
					typedReq, ok := req.(*ListRepositoryTagsRequest)
					if !ok {
						return nil, twirp.InternalError("failed type assertion req.(*ListRepositoryTagsRequest) when calling interceptor")
					}
					return s.RepositoryTagService.ListRepositoryTags(ctx, typedReq)
				},
			)(ctx, req)
			if resp != nil {
				typedResp, ok := resp.(*ListRepositoryTagsResponse)
				if !ok {
					return nil, twirp.InternalError("failed type assertion resp.(*ListRepositoryTagsResponse) when calling interceptor")
				}
				return typedResp, err
			}
			return nil, err
		}
	}

	// Call service method
	var respContent *ListRepositoryTagsResponse
	func() {
		defer ensurePanicResponses(ctx, resp, s.hooks)
		respContent, err = handler(ctx, reqContent)
	}()

	if err != nil {
		s.writeError(ctx, resp, err)
		return
	}
	if respContent == nil {
		s.writeError(ctx, resp, twirp.InternalError("received a nil *ListRepositoryTagsResponse and nil error while calling ListRepositoryTags. nil responses are not supported"))
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

func (s *repositoryTagServiceServer) serveListRepositoryTagsProtobuf(ctx context.Context, resp http.ResponseWriter, req *http.Request) {
	var err error
	ctx = ctxsetters.WithMethodName(ctx, "ListRepositoryTags")
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
	reqContent := new(ListRepositoryTagsRequest)
	if err = proto.Unmarshal(buf, reqContent); err != nil {
		s.writeError(ctx, resp, malformedRequestError("the protobuf request could not be decoded"))
		return
	}

	handler := s.RepositoryTagService.ListRepositoryTags
	if s.interceptor != nil {
		handler = func(ctx context.Context, req *ListRepositoryTagsRequest) (*ListRepositoryTagsResponse, error) {
			resp, err := s.interceptor(
				func(ctx context.Context, req interface{}) (interface{}, error) {
					typedReq, ok := req.(*ListRepositoryTagsRequest)
					if !ok {
						return nil, twirp.InternalError("failed type assertion req.(*ListRepositoryTagsRequest) when calling interceptor")
					}
					return s.RepositoryTagService.ListRepositoryTags(ctx, typedReq)
				},
			)(ctx, req)
			if resp != nil {
				typedResp, ok := resp.(*ListRepositoryTagsResponse)
				if !ok {
					return nil, twirp.InternalError("failed type assertion resp.(*ListRepositoryTagsResponse) when calling interceptor")
				}
				return typedResp, err
			}
			return nil, err
		}
	}

	// Call service method
	var respContent *ListRepositoryTagsResponse
	func() {
		defer ensurePanicResponses(ctx, resp, s.hooks)
		respContent, err = handler(ctx, reqContent)
	}()

	if err != nil {
		s.writeError(ctx, resp, err)
		return
	}
	if respContent == nil {
		s.writeError(ctx, resp, twirp.InternalError("received a nil *ListRepositoryTagsResponse and nil error while calling ListRepositoryTags. nil responses are not supported"))
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

func (s *repositoryTagServiceServer) ServiceDescriptor() ([]byte, int) {
	return twirpFileDescriptor7, 0
}

func (s *repositoryTagServiceServer) ProtocGenTwirpVersion() string {
	return "v8.1.0"
}

// PathPrefix returns the base service path, in the form: "/<prefix>/<package>.<Service>/"
// that is everything in a Twirp route except for the <Method>. This can be used for routing,
// for example to identify the requests that are targeted to this service in a mux.
func (s *repositoryTagServiceServer) PathPrefix() string {
	return baseServicePath(s.pathPrefix, "buf.alpha.registry.v1alpha1", "RepositoryTagService")
}

var twirpFileDescriptor7 = []byte{
	// 510 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x54, 0x41, 0x8b, 0xd3, 0x40,
	0x14, 0x26, 0xdd, 0x75, 0xdd, 0x7d, 0xb5, 0x5d, 0x18, 0x45, 0x62, 0x8a, 0xb4, 0x54, 0x90, 0xc5,
	0xc3, 0x8c, 0x5b, 0xc1, 0x15, 0xf6, 0xb6, 0x9e, 0x04, 0x91, 0x35, 0xed, 0x69, 0x2f, 0x65, 0xd2,
	0xbe, 0x4e, 0x07, 0x9b, 0x4c, 0x9c, 0x99, 0x14, 0x77, 0xef, 0x82, 0x57, 0x2f, 0x82, 0x3f, 0x61,
	0xff, 0xa5, 0x64, 0xd2, 0xd4, 0x46, 0xdb, 0x40, 0x6f, 0xc9, 0xf7, 0xde, 0xf7, 0xe5, 0x7b, 0xdf,
	0x9b, 0x0c, 0xbc, 0x8e, 0xb2, 0x19, 0xe3, 0x8b, 0x74, 0xce, 0x99, 0x46, 0x21, 0x8d, 0xd5, 0xb7,
	0x6c, 0x79, 0xee, 0x80, 0x73, 0xa6, 0x31, 0x55, 0x46, 0x5a, 0xa5, 0x6f, 0xc7, 0x96, 0x0b, 0x9a,
	0x6a, 0x65, 0x15, 0xe9, 0x44, 0xd9, 0x8c, 0xba, 0x06, 0x5a, 0x32, 0x68, 0xc9, 0x08, 0xba, 0x42,
	0x29, 0xb1, 0x40, 0xe6, 0x5a, 0x73, 0x69, 0x2b, 0x63, 0x34, 0x96, 0xc7, 0x69, 0xc1, 0xee, 0xdf,
	0x7b, 0xd0, 0x0a, 0xd7, 0xb2, 0x23, 0x2e, 0x48, 0x1b, 0x1a, 0x72, 0xea, 0x7b, 0x3d, 0xef, 0xec,
	0x24, 0x6c, 0xc8, 0x29, 0xb9, 0x84, 0xe6, 0x44, 0x23, 0xb7, 0x38, 0xce, 0xb9, 0x7e, 0xa3, 0xe7,
	0x9d, 0x35, 0x07, 0x01, 0x2d, 0x84, 0x69, 0x29, 0x4c, 0x47, 0xa5, 0x70, 0x08, 0x45, 0x7b, 0x0e,
	0x10, 0x02, 0x87, 0x09, 0x8f, 0xd1, 0x3f, 0x74, 0x72, 0xee, 0x99, 0x74, 0xa1, 0x39, 0x51, 0x71,
	0x2c, 0xed, 0xd8, 0x95, 0x1e, 0xb8, 0x12, 0x14, 0xd0, 0xa7, 0xbc, 0xe1, 0x29, 0x1c, 0xf1, 0xcc,
	0xce, 0x95, 0xf6, 0x8f, 0x5c, 0x6d, 0xf5, 0xd6, 0x5f, 0x42, 0xf0, 0xde, 0x49, 0x57, 0x0c, 0x87,
	0xf8, 0x35, 0x43, 0x63, 0xc9, 0x0b, 0x68, 0x6d, 0xe4, 0xb3, 0x1e, 0xe1, 0xd1, 0x5f, 0xf0, 0xc3,
	0x74, 0xed, 0xa7, 0xb1, 0xdb, 0xcf, 0xc1, 0xbf, 0x7e, 0xfa, 0x29, 0x74, 0xb6, 0x7e, 0xd7, 0xa4,
	0x2a, 0x31, 0x48, 0x3e, 0x43, 0xbb, 0xba, 0x18, 0xf7, 0xe5, 0xe6, 0xe0, 0x15, 0xad, 0xd9, 0x0c,
	0xad, 0x6a, 0x6d, 0x58, 0x1f, 0x71, 0xd1, 0xff, 0xe5, 0xc1, 0xb3, 0x8f, 0xd2, 0xd8, 0x4a, 0x93,
	0xd9, 0x6b, 0xd2, 0x0e, 0x9c, 0xa4, 0x5c, 0xe0, 0xd8, 0xc8, 0xbb, 0x62, 0xdc, 0x56, 0x78, 0x9c,
	0x03, 0x43, 0x79, 0x87, 0xe4, 0x39, 0x80, 0x2b, 0x5a, 0xf5, 0x05, 0x93, 0xd5, 0xc4, 0xae, 0x7d,
	0x94, 0x03, 0xc4, 0x87, 0x87, 0x1a, 0x97, 0xa8, 0x4d, 0xb1, 0xb8, 0xe3, 0xb0, 0x7c, 0xed, 0xff,
	0xf6, 0x20, 0xd8, 0x66, 0x6c, 0x15, 0xc5, 0x10, 0x4e, 0xab, 0x51, 0x18, 0xdf, 0xeb, 0x1d, 0xec,
	0x99, 0x45, 0xbb, 0x92, 0x85, 0x21, 0x2f, 0xe1, 0x34, 0xc1, 0x6f, 0x76, 0xbc, 0xe1, 0xb8, 0x58,
	0x5f, 0x2b, 0x87, 0xaf, 0x4b, 0xd7, 0x83, 0xfb, 0x06, 0x3c, 0xa9, 0x28, 0x0d, 0x51, 0x2f, 0xe5,
	0x04, 0xc9, 0x0f, 0x0f, 0x1e, 0x6f, 0x59, 0x20, 0xb9, 0xa8, 0x35, 0xb5, 0xfb, 0xa8, 0x05, 0xef,
	0xf6, 0x27, 0xae, 0x02, 0xfa, 0xee, 0x01, 0xf9, 0x3f, 0x3f, 0xf2, 0xb6, 0x56, 0x70, 0xe7, 0x49,
	0x08, 0x2e, 0xf6, 0xe6, 0x15, 0x3e, 0xae, 0x7e, 0x7a, 0xd0, 0x9d, 0xa8, 0xb8, 0x8e, 0x7e, 0x45,
	0x2a, 0xdc, 0xeb, 0xfc, 0x47, 0xbf, 0xb9, 0x11, 0xd2, 0xce, 0xb3, 0x88, 0x4e, 0x54, 0xcc, 0xa2,
	0x6c, 0x16, 0x65, 0x72, 0x31, 0xcd, 0x1f, 0x98, 0x4c, 0x2c, 0xea, 0x84, 0x2f, 0x98, 0xc0, 0xa4,
	0xb8, 0x6d, 0x98, 0x50, 0xac, 0xe6, 0x32, 0xbb, 0x2c, 0x91, 0x12, 0x88, 0x8e, 0x1c, 0xed, 0xcd,
	0x9f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x18, 0xca, 0x42, 0xa5, 0x03, 0x05, 0x00, 0x00,
}
