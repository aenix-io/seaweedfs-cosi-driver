/*
Copyright 2023 SUSE, LLC.
Copyright 2024 s3gw maintainers.

Licensed under the Apache License, Version 2.0 (the "License");
You may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package driver

import (
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"k8s.io/klog/v2"

	cosispec "sigs.k8s.io/container-object-storage-interface-spec"
)

// identityServer implements cosi.IdentityServer interface.
type identityServer struct {
	provisioner string
}

// Interface guards.
var _ cosispec.IdentityServer = &identityServer{}

// NewIdentityServer returns IdentityServer with provisioner set to the "provisionerName" argument.
func NewIdentityServer(provisionerName string) (cosispec.IdentityServer, error) {
	return &identityServer{
		provisioner: provisionerName,
	}, nil
}

// DriverGetInfo call is meant to retrieve the unique provisioner Identity.
func (id *identityServer) DriverGetInfo(ctx context.Context,
	req *cosispec.DriverGetInfoRequest) (*cosispec.DriverGetInfoResponse, error) {

	if id.provisioner == "" {
		klog.ErrorS(ErrProvisionerNameEmpty, "invalid configuration")

		return nil, status.Error(codes.Internal, "Provisioner name is empty")
	}

	return &cosispec.DriverGetInfoResponse{
		Name: id.provisioner,
	}, nil
}
