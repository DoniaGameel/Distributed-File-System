// Copyright 2022 Google LLC
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

// Code generated by aliasgen. DO NOT EDIT.

// Package shell aliases all exported identifiers in package
// "cloud.google.com/go/shell/apiv1/shellpb".
//
// Deprecated: Please use types in: cloud.google.com/go/shell/apiv1/shellpb.
// Please read https://github.com/googleapis/google-cloud-go/blob/main/migration.md
// for more details.
package shell

import (
	src "cloud.google.com/go/shell/apiv1/shellpb"
	grpc "google.golang.org/grpc"
)

// Deprecated: Please use consts in: cloud.google.com/go/shell/apiv1/shellpb
const (
	CloudShellErrorDetails_CLOUD_SHELL_DISABLED               = src.CloudShellErrorDetails_CLOUD_SHELL_DISABLED
	CloudShellErrorDetails_CLOUD_SHELL_ERROR_CODE_UNSPECIFIED = src.CloudShellErrorDetails_CLOUD_SHELL_ERROR_CODE_UNSPECIFIED
	CloudShellErrorDetails_IMAGE_UNAVAILABLE                  = src.CloudShellErrorDetails_IMAGE_UNAVAILABLE
	CloudShellErrorDetails_QUOTA_EXCEEDED                     = src.CloudShellErrorDetails_QUOTA_EXCEEDED
	CloudShellErrorDetails_TOS_VIOLATION                      = src.CloudShellErrorDetails_TOS_VIOLATION
	Environment_DELETING                                      = src.Environment_DELETING
	Environment_PENDING                                       = src.Environment_PENDING
	Environment_RUNNING                                       = src.Environment_RUNNING
	Environment_STATE_UNSPECIFIED                             = src.Environment_STATE_UNSPECIFIED
	Environment_SUSPENDED                                     = src.Environment_SUSPENDED
	StartEnvironmentMetadata_AWAITING_COMPUTE_RESOURCES       = src.StartEnvironmentMetadata_AWAITING_COMPUTE_RESOURCES
	StartEnvironmentMetadata_FINISHED                         = src.StartEnvironmentMetadata_FINISHED
	StartEnvironmentMetadata_STARTING                         = src.StartEnvironmentMetadata_STARTING
	StartEnvironmentMetadata_STATE_UNSPECIFIED                = src.StartEnvironmentMetadata_STATE_UNSPECIFIED
	StartEnvironmentMetadata_UNARCHIVING_DISK                 = src.StartEnvironmentMetadata_UNARCHIVING_DISK
)

// Deprecated: Please use vars in: cloud.google.com/go/shell/apiv1/shellpb
var (
	CloudShellErrorDetails_CloudShellErrorCode_name  = src.CloudShellErrorDetails_CloudShellErrorCode_name
	CloudShellErrorDetails_CloudShellErrorCode_value = src.CloudShellErrorDetails_CloudShellErrorCode_value
	Environment_State_name                           = src.Environment_State_name
	Environment_State_value                          = src.Environment_State_value
	File_google_cloud_shell_v1_cloudshell_proto      = src.File_google_cloud_shell_v1_cloudshell_proto
	StartEnvironmentMetadata_State_name              = src.StartEnvironmentMetadata_State_name
	StartEnvironmentMetadata_State_value             = src.StartEnvironmentMetadata_State_value
)

// Message included in the metadata field of operations returned from
// [AddPublicKey][google.cloud.shell.v1.CloudShellService.AddPublicKey].
//
// Deprecated: Please use types in: cloud.google.com/go/shell/apiv1/shellpb
type AddPublicKeyMetadata = src.AddPublicKeyMetadata

// Request message for
// [AddPublicKey][google.cloud.shell.v1.CloudShellService.AddPublicKey].
//
// Deprecated: Please use types in: cloud.google.com/go/shell/apiv1/shellpb
type AddPublicKeyRequest = src.AddPublicKeyRequest

// Response message for
// [AddPublicKey][google.cloud.shell.v1.CloudShellService.AddPublicKey].
//
// Deprecated: Please use types in: cloud.google.com/go/shell/apiv1/shellpb
type AddPublicKeyResponse = src.AddPublicKeyResponse

// Message included in the metadata field of operations returned from
// [AuthorizeEnvironment][google.cloud.shell.v1.CloudShellService.AuthorizeEnvironment].
//
// Deprecated: Please use types in: cloud.google.com/go/shell/apiv1/shellpb
type AuthorizeEnvironmentMetadata = src.AuthorizeEnvironmentMetadata

// Request message for
// [AuthorizeEnvironment][google.cloud.shell.v1.CloudShellService.AuthorizeEnvironment].
//
// Deprecated: Please use types in: cloud.google.com/go/shell/apiv1/shellpb
type AuthorizeEnvironmentRequest = src.AuthorizeEnvironmentRequest

// Response message for
// [AuthorizeEnvironment][google.cloud.shell.v1.CloudShellService.AuthorizeEnvironment].
//
// Deprecated: Please use types in: cloud.google.com/go/shell/apiv1/shellpb
type AuthorizeEnvironmentResponse = src.AuthorizeEnvironmentResponse

// Cloud-shell specific information that will be included as details in
// failure responses.
//
// Deprecated: Please use types in: cloud.google.com/go/shell/apiv1/shellpb
type CloudShellErrorDetails = src.CloudShellErrorDetails

// Set of possible errors returned from API calls.
//
// Deprecated: Please use types in: cloud.google.com/go/shell/apiv1/shellpb
type CloudShellErrorDetails_CloudShellErrorCode = src.CloudShellErrorDetails_CloudShellErrorCode

// CloudShellServiceClient is the client API for CloudShellService service.
// For semantics around ctx use and closing/ending streaming RPCs, please refer
// to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
//
// Deprecated: Please use types in: cloud.google.com/go/shell/apiv1/shellpb
type CloudShellServiceClient = src.CloudShellServiceClient

// CloudShellServiceServer is the server API for CloudShellService service.
//
// Deprecated: Please use types in: cloud.google.com/go/shell/apiv1/shellpb
type CloudShellServiceServer = src.CloudShellServiceServer

// Message included in the metadata field of operations returned from
// [CreateEnvironment][google.cloud.shell.v1.CloudShellService.CreateEnvironment].
//
// Deprecated: Please use types in: cloud.google.com/go/shell/apiv1/shellpb
type CreateEnvironmentMetadata = src.CreateEnvironmentMetadata

// Message included in the metadata field of operations returned from
// [DeleteEnvironment][google.cloud.shell.v1.CloudShellService.DeleteEnvironment].
//
// Deprecated: Please use types in: cloud.google.com/go/shell/apiv1/shellpb
type DeleteEnvironmentMetadata = src.DeleteEnvironmentMetadata

// A Cloud Shell environment, which is defined as the combination of a Docker
// image specifying what is installed on the environment and a home directory
// containing the user's data that will remain across sessions. Each user has
// at least an environment with the ID "default".
//
// Deprecated: Please use types in: cloud.google.com/go/shell/apiv1/shellpb
type Environment = src.Environment

// Possible execution states for an environment.
//
// Deprecated: Please use types in: cloud.google.com/go/shell/apiv1/shellpb
type Environment_State = src.Environment_State

// Request message for
// [GetEnvironment][google.cloud.shell.v1.CloudShellService.GetEnvironment].
//
// Deprecated: Please use types in: cloud.google.com/go/shell/apiv1/shellpb
type GetEnvironmentRequest = src.GetEnvironmentRequest

// Message included in the metadata field of operations returned from
// [RemovePublicKey][google.cloud.shell.v1.CloudShellService.RemovePublicKey].
//
// Deprecated: Please use types in: cloud.google.com/go/shell/apiv1/shellpb
type RemovePublicKeyMetadata = src.RemovePublicKeyMetadata

// Request message for
// [RemovePublicKey][google.cloud.shell.v1.CloudShellService.RemovePublicKey].
//
// Deprecated: Please use types in: cloud.google.com/go/shell/apiv1/shellpb
type RemovePublicKeyRequest = src.RemovePublicKeyRequest

// Response message for
// [RemovePublicKey][google.cloud.shell.v1.CloudShellService.RemovePublicKey].
//
// Deprecated: Please use types in: cloud.google.com/go/shell/apiv1/shellpb
type RemovePublicKeyResponse = src.RemovePublicKeyResponse

// Message included in the metadata field of operations returned from
// [StartEnvironment][google.cloud.shell.v1.CloudShellService.StartEnvironment].
//
// Deprecated: Please use types in: cloud.google.com/go/shell/apiv1/shellpb
type StartEnvironmentMetadata = src.StartEnvironmentMetadata

// Possible states an environment might transition between during startup.
// These states are not normally actionable by clients, but may be used to show
// a progress message to the user. An environment won't necessarily go through
// all of these states when starting. More states are likely to be added in the
// future.
//
// Deprecated: Please use types in: cloud.google.com/go/shell/apiv1/shellpb
type StartEnvironmentMetadata_State = src.StartEnvironmentMetadata_State

// Request message for
// [StartEnvironment][google.cloud.shell.v1.CloudShellService.StartEnvironment].
//
// Deprecated: Please use types in: cloud.google.com/go/shell/apiv1/shellpb
type StartEnvironmentRequest = src.StartEnvironmentRequest

// Message included in the response field of operations returned from
// [StartEnvironment][google.cloud.shell.v1.CloudShellService.StartEnvironment]
// once the operation is complete.
//
// Deprecated: Please use types in: cloud.google.com/go/shell/apiv1/shellpb
type StartEnvironmentResponse = src.StartEnvironmentResponse

// UnimplementedCloudShellServiceServer can be embedded to have forward
// compatible implementations.
//
// Deprecated: Please use types in: cloud.google.com/go/shell/apiv1/shellpb
type UnimplementedCloudShellServiceServer = src.UnimplementedCloudShellServiceServer

// Deprecated: Please use funcs in: cloud.google.com/go/shell/apiv1/shellpb
func NewCloudShellServiceClient(cc grpc.ClientConnInterface) CloudShellServiceClient {
	return src.NewCloudShellServiceClient(cc)
}

// Deprecated: Please use funcs in: cloud.google.com/go/shell/apiv1/shellpb
func RegisterCloudShellServiceServer(s *grpc.Server, srv CloudShellServiceServer) {
	src.RegisterCloudShellServiceServer(s, srv)
}
