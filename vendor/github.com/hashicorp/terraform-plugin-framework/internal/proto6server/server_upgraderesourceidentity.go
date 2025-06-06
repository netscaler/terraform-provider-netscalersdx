// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package proto6server

import (
	"context"
	"github.com/hashicorp/terraform-plugin-framework/internal/fromproto6"
	"github.com/hashicorp/terraform-plugin-framework/internal/fwserver"
	"github.com/hashicorp/terraform-plugin-framework/internal/logging"
	"github.com/hashicorp/terraform-plugin-framework/internal/toproto6"

	"github.com/hashicorp/terraform-plugin-go/tfprotov6"
)

// UpgradeResourceIdentity satisfies the tfprotov6.ProviderServer interface.
func (s *Server) UpgradeResourceIdentity(ctx context.Context, proto6Req *tfprotov6.UpgradeResourceIdentityRequest) (*tfprotov6.UpgradeResourceIdentityResponse, error) {
	ctx = s.registerContext(ctx)
	ctx = logging.InitContext(ctx)

	fwResp := &fwserver.UpgradeResourceIdentityResponse{}

	if proto6Req == nil {
		return toproto6.UpgradeResourceIdentityResponse(ctx, fwResp), nil
	}

	resource, diags := s.FrameworkServer.Resource(ctx, proto6Req.TypeName)

	fwResp.Diagnostics.Append(diags...)

	if fwResp.Diagnostics.HasError() {
		return toproto6.UpgradeResourceIdentityResponse(ctx, fwResp), nil
	}

	identitySchema, diags := s.FrameworkServer.ResourceIdentitySchema(ctx, proto6Req.TypeName)

	fwResp.Diagnostics.Append(diags...)

	if fwResp.Diagnostics.HasError() {
		return toproto6.UpgradeResourceIdentityResponse(ctx, fwResp), nil
	}

	fwReq, diags := fromproto6.UpgradeResourceIdentityRequest(ctx, proto6Req, resource, identitySchema)

	fwResp.Diagnostics.Append(diags...)

	if fwResp.Diagnostics.HasError() {
		return toproto6.UpgradeResourceIdentityResponse(ctx, fwResp), nil
	}

	s.FrameworkServer.UpgradeResourceIdentity(ctx, fwReq, fwResp)

	return toproto6.UpgradeResourceIdentityResponse(ctx, fwResp), nil
}
