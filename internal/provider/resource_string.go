package provider

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
)

type resourceStringType struct{}

func (r resourceStringType) GetSchema(context.Context) (tfsdk.Schema, diag.Diagnostics) {
	description := "The resource `random_string` generates a random permutation of alphanumeric " +
		"characters and optionally special characters.\n" +
		"\n" +
		"This resource *does* use a cryptographic random number generator.\n" +
		"\n" +
		"Historically this resource's intended usage has been ambiguous as the original example used " +
		"it in a password. For backwards compatibility it will continue to exist. For unique ids please " +
		"use [random_id](id.html), for sensitive random values please use [random_password](password.html)."

	schema := getStringSchemaV1(false, description)
	schema.Version = 1

	return schema, nil
}

func (r resourceStringType) NewResource(_ context.Context, p tfsdk.Provider) (tfsdk.Resource, diag.Diagnostics) {
	return resourceString{
		p: *(p.(*provider)),
	}, nil
}

type resourceString struct {
	p provider
}

func (r resourceString) Create(ctx context.Context, req tfsdk.CreateResourceRequest, resp *tfsdk.CreateResourceResponse) {
	createString(ctx, req, resp, false)
}

// Read does not need to perform any operations as the state in ReadResourceResponse is already populated.
func (r resourceString) Read(ctx context.Context, req tfsdk.ReadResourceRequest, resp *tfsdk.ReadResourceResponse) {
}

// Update is intentionally left blank as all required and optional attributes force replacement of the resource
// through the RequiresReplace AttributePlanModifier.
func (r resourceString) Update(ctx context.Context, req tfsdk.UpdateResourceRequest, resp *tfsdk.UpdateResourceResponse) {
}

// Delete does not need to explicitly call resp.State.RemoveResource() as this is automatically handled by the
// [framework](https://github.com/hashicorp/terraform-plugin-framework/pull/301).
func (r resourceString) Delete(ctx context.Context, req tfsdk.DeleteResourceRequest, resp *tfsdk.DeleteResourceResponse) {
}

func (r resourceString) ImportState(ctx context.Context, req tfsdk.ImportResourceStateRequest, resp *tfsdk.ImportResourceStateResponse) {
	importString(ctx, req, resp, false)
}

func (r resourceString) ValidateConfig(ctx context.Context, req tfsdk.ValidateResourceConfigRequest, resp *tfsdk.ValidateResourceConfigResponse) {
	validateLength(ctx, req, resp)
}
