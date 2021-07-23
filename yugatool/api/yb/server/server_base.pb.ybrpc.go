// Licensed to the Apache Software Foundation (ASF) under one
// or more contributor license agreements.  See the NOTICE file
// distributed with this work for additional information
// regarding copyright ownership.  The ASF licenses this file
// to you under the Apache License, Version 2.0 (the
// "License"); you may not use this file except in compliance
// with the License.  You may obtain a copy of the License at
//
//   http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing,
// software distributed under the License is distributed on an
// "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
// KIND, either express or implied.  See the License for the
// specific language governing permissions and limitations
// under the License.
//
// The following only applies to changes made to this file as part of YugaByte development.
//
// Portions Copyright (c) YugaByte, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License"); you may not use this file except
// in compliance with the License.  You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software distributed under the License
// is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express
// or implied.  See the License for the specific language governing permissions and limitations
// under the License.
//

// Code generated by protoc-gen-ybrpc. DO NOT EDIT.

package server

import (
	"github.com/go-logr/logr"
	"github.com/yugabyte/yb-tools/protoc-gen-ybrpc/pkg/message"
)

// service: yb.server.GenericService
// service: GenericService
type GenericService interface {
	SetFlag(request *SetFlagRequestPB) (*SetFlagResponsePB, error)
	GetFlag(request *GetFlagRequestPB) (*GetFlagResponsePB, error)
	RefreshFlags(request *RefreshFlagsRequestPB) (*RefreshFlagsResponsePB, error)
	FlushCoverage(request *FlushCoverageRequestPB) (*FlushCoverageResponsePB, error)
	ServerClock(request *ServerClockRequestPB) (*ServerClockResponsePB, error)
	GetStatus(request *GetStatusRequestPB) (*GetStatusResponsePB, error)
	Ping(request *PingRequestPB) (*PingResponsePB, error)
}

type GenericServiceImpl struct {
	Log       logr.Logger
	Messenger message.Messenger
}

func (s *GenericServiceImpl) SetFlag(request *SetFlagRequestPB) (*SetFlagResponsePB, error) {
	s.Log.V(1).Info("sending RPC request", "service", "yb.server.GenericService", "method", "SetFlag", "host", s.Messenger.GetHost(), "request", request)
	response := &SetFlagResponsePB{}

	err := s.Messenger.SendMessage("yb.server.GenericService", "SetFlag", request.ProtoReflect().Interface(), response.ProtoReflect().Interface())
	if err != nil {
		return nil, err
	}

	s.Log.V(1).Info("received RPC response", "service", "yb.server.GenericService", "method", "SetFlag", "host", s.Messenger.GetHost(), "response", response)

	return response, nil
}

func (s *GenericServiceImpl) GetFlag(request *GetFlagRequestPB) (*GetFlagResponsePB, error) {
	s.Log.V(1).Info("sending RPC request", "service", "yb.server.GenericService", "method", "GetFlag", "host", s.Messenger.GetHost(), "request", request)
	response := &GetFlagResponsePB{}

	err := s.Messenger.SendMessage("yb.server.GenericService", "GetFlag", request.ProtoReflect().Interface(), response.ProtoReflect().Interface())
	if err != nil {
		return nil, err
	}

	s.Log.V(1).Info("received RPC response", "service", "yb.server.GenericService", "method", "GetFlag", "host", s.Messenger.GetHost(), "response", response)

	return response, nil
}

func (s *GenericServiceImpl) RefreshFlags(request *RefreshFlagsRequestPB) (*RefreshFlagsResponsePB, error) {
	s.Log.V(1).Info("sending RPC request", "service", "yb.server.GenericService", "method", "RefreshFlags", "host", s.Messenger.GetHost(), "request", request)
	response := &RefreshFlagsResponsePB{}

	err := s.Messenger.SendMessage("yb.server.GenericService", "RefreshFlags", request.ProtoReflect().Interface(), response.ProtoReflect().Interface())
	if err != nil {
		return nil, err
	}

	s.Log.V(1).Info("received RPC response", "service", "yb.server.GenericService", "method", "RefreshFlags", "host", s.Messenger.GetHost(), "response", response)

	return response, nil
}

func (s *GenericServiceImpl) FlushCoverage(request *FlushCoverageRequestPB) (*FlushCoverageResponsePB, error) {
	s.Log.V(1).Info("sending RPC request", "service", "yb.server.GenericService", "method", "FlushCoverage", "host", s.Messenger.GetHost(), "request", request)
	response := &FlushCoverageResponsePB{}

	err := s.Messenger.SendMessage("yb.server.GenericService", "FlushCoverage", request.ProtoReflect().Interface(), response.ProtoReflect().Interface())
	if err != nil {
		return nil, err
	}

	s.Log.V(1).Info("received RPC response", "service", "yb.server.GenericService", "method", "FlushCoverage", "host", s.Messenger.GetHost(), "response", response)

	return response, nil
}

func (s *GenericServiceImpl) ServerClock(request *ServerClockRequestPB) (*ServerClockResponsePB, error) {
	s.Log.V(1).Info("sending RPC request", "service", "yb.server.GenericService", "method", "ServerClock", "host", s.Messenger.GetHost(), "request", request)
	response := &ServerClockResponsePB{}

	err := s.Messenger.SendMessage("yb.server.GenericService", "ServerClock", request.ProtoReflect().Interface(), response.ProtoReflect().Interface())
	if err != nil {
		return nil, err
	}

	s.Log.V(1).Info("received RPC response", "service", "yb.server.GenericService", "method", "ServerClock", "host", s.Messenger.GetHost(), "response", response)

	return response, nil
}

func (s *GenericServiceImpl) GetStatus(request *GetStatusRequestPB) (*GetStatusResponsePB, error) {
	s.Log.V(1).Info("sending RPC request", "service", "yb.server.GenericService", "method", "GetStatus", "host", s.Messenger.GetHost(), "request", request)
	response := &GetStatusResponsePB{}

	err := s.Messenger.SendMessage("yb.server.GenericService", "GetStatus", request.ProtoReflect().Interface(), response.ProtoReflect().Interface())
	if err != nil {
		return nil, err
	}

	s.Log.V(1).Info("received RPC response", "service", "yb.server.GenericService", "method", "GetStatus", "host", s.Messenger.GetHost(), "response", response)

	return response, nil
}

func (s *GenericServiceImpl) Ping(request *PingRequestPB) (*PingResponsePB, error) {
	s.Log.V(1).Info("sending RPC request", "service", "yb.server.GenericService", "method", "Ping", "host", s.Messenger.GetHost(), "request", request)
	response := &PingResponsePB{}

	err := s.Messenger.SendMessage("yb.server.GenericService", "Ping", request.ProtoReflect().Interface(), response.ProtoReflect().Interface())
	if err != nil {
		return nil, err
	}

	s.Log.V(1).Info("received RPC response", "service", "yb.server.GenericService", "method", "Ping", "host", s.Messenger.GetHost(), "response", response)

	return response, nil
}