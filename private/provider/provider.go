package provider

import (
	"context"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/provider"
	"github.com/hashicorp/terraform-plugin-framework/resource"
)

type DexProvider struct{}

func (p DexProvider) Metadata(_ context.Context, _ provider.MetadataRequest, resp *provider.MetadataResponse) {
	resp.TypeName = "dex"
}

func (p DexProvider) Schema(_ context.Context, _ provider.SchemaRequest, _ *provider.SchemaResponse) {
}

func (p DexProvider) Configure(_ context.Context, _ provider.ConfigureRequest, _ *provider.ConfigureResponse) {
}

func (p DexProvider) DataSources(_ context.Context) []func() datasource.DataSource {
	return []func() datasource.DataSource{
		NewDexSubDataSource,
	}
}

func (p DexProvider) Resources(_ context.Context) []func() resource.Resource {
	return nil
}

func NewDexProvider() provider.Provider {
	return &DexProvider{}
}
