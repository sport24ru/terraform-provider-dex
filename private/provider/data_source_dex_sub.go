package provider

import (
	"context"
	"encoding/base64"
	"github.com/golang/protobuf/proto"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var _ datasource.DataSource = &DexSubDataSource{}

type DexSubDataSource struct{}

type DexSubDataSourceModel struct {
	Id     types.String `tfsdk:"id"`
	UserId types.String `tfsdk:"user_id"`
	ConnId types.String `tfsdk:"conn_id"`
	Sub    types.String `tfsdk:"sub"`
}

func (d DexSubDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_sub"
}

func (d DexSubDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"user_id": schema.StringAttribute{
				Required:            true,
				MarkdownDescription: "User ID of Dex ID token.",
			},
			"conn_id": schema.StringAttribute{
				Required:            true,
				MarkdownDescription: "Conn ID of Dex ID token.",
			},
			"sub": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "Subject string of Dex ID token.",
			},
		},
	}
}

type IDTokenSubject struct {
	UserId string `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	ConnId string `protobuf:"bytes,2,opt,name=conn_id,json=connId,proto3" json:"conn_id,omitempty"`
}

func (m *IDTokenSubject) Reset()         { *m = IDTokenSubject{} }
func (m *IDTokenSubject) String() string { return proto.CompactTextString(m) }
func (*IDTokenSubject) ProtoMessage()    {}

func Marshal(message proto.Message) (string, error) {
	data, err := proto.Marshal(message)
	if err != nil {
		return "", err
	}
	return base64.RawURLEncoding.EncodeToString(data), nil
}

func (d DexSubDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var state DexSubDataSourceModel

	resp.Diagnostics.Append(req.Config.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}

	var idTokenSubject IDTokenSubject

	idTokenSubject.ConnId = state.ConnId.ValueString()
	idTokenSubject.UserId = state.UserId.ValueString()

	sub, err := Marshal(&idTokenSubject)
	if err != nil {
		panic(err)
	}

	state.Sub = types.StringValue(sub)
	state.Id = types.StringValue(sub)

	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}
}

func NewDexSubDataSource() datasource.DataSource {
	return &DexSubDataSource{}
}
