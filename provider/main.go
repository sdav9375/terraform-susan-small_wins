package main

import (
	"bytes"
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/hashicorp/terraform-plugin-framework/action"
	"github.com/hashicorp/terraform-plugin-framework/action/schema"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/provider"
	"github.com/hashicorp/terraform-plugin-framework/providerserver"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type httpProvider struct{}

func (p *httpProvider) Metadata(_ context.Context, _ provider.MetadataRequest, resp *provider.MetadataResponse) {
	resp.TypeName = "httpaction"
}

func (p *httpProvider) Schema(context.Context, provider.SchemaRequest, *provider.SchemaResponse) {}

func (p *httpProvider) Configure(context.Context, provider.ConfigureRequest, *provider.ConfigureResponse) {
}

func (p *httpProvider) Resources(context.Context) []func() resource.Resource {
	return nil
}

func (p *httpProvider) DataSources(context.Context) []func() datasource.DataSource {
	return nil
}

func (p *httpProvider) Actions(context.Context) []func() action.Action {
	return []func() action.Action{func() action.Action { return &httpRequest{} }}
}

type httpRequest struct{}

type httpRequestModel struct {
	URL     types.String `tfsdk:"url"`
	Method  types.String `tfsdk:"method"`
	Payload types.String `tfsdk:"payload"`
}

func (a *httpRequest) Metadata(_ context.Context, req action.MetadataRequest, resp *action.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_request"
}

func (a *httpRequest) Schema(_ context.Context, _ action.SchemaRequest, resp *action.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"url":     schema.StringAttribute{Required: true},
			"method":  schema.StringAttribute{Required: true},
			"payload": schema.StringAttribute{Required: true},
		},
	}
}

func (a *httpRequest) Invoke(ctx context.Context, req action.InvokeRequest, resp *action.InvokeResponse) {
	var config httpRequestModel
	resp.Diagnostics.Append(req.Config.Get(ctx, &config)...)
	if resp.Diagnostics.HasError() {
		return
	}

	request, err := http.NewRequestWithContext(
		ctx,
		config.Method.ValueString(),
		config.URL.ValueString(),
		bytes.NewBufferString(config.Payload.ValueString()),
	)
	if err != nil {
		resp.Diagnostics.AddError("HTTP request failed", err.Error())
		return
	}
	request.Header.Set("Content-Type", "application/json")

	client := http.Client{Timeout: 30 * time.Second}
	response, err := client.Do(request)
	if err != nil {
		resp.Diagnostics.AddError("HTTP request failed", err.Error())
		return
	}
	defer response.Body.Close()

	if response.StatusCode < http.StatusOK || response.StatusCode >= http.StatusMultipleChoices {
		resp.Diagnostics.AddError("HTTP request failed", fmt.Sprintf("received %s", response.Status))
		return
	}

	resp.SendProgress(action.InvokeProgressEvent{Message: "HTTP request completed with " + response.Status})
}

func main() {
	err := providerserver.Serve(
		context.Background(),
		func() provider.Provider { return &httpProvider{} },
		providerserver.ServeOpts{Address: "terraform.local/local/httpaction"},
	)
	if err != nil {
		log.Fatal(err.Error())
	}
}
