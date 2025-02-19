package planmodifiers

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/terraform-providers/terraform-provider-datadog/datadog/internal/utils"
)

func NormalizeIP() planmodifier.String {
	return normalizeIpModifier{}
}

type normalizeIpModifier struct{}

func (m normalizeIpModifier) Description(_ context.Context) string {
	return "Normalize IP value."
}

func (m normalizeIpModifier) MarkdownDescription(ctx context.Context) string {
	return m.Description(ctx)
}

func (m normalizeIpModifier) PlanModifyString(ctx context.Context, req planmodifier.StringRequest, resp *planmodifier.StringResponse) {
	if req.ConfigValue.IsUnknown() || req.ConfigValue.IsNull() {
		return
	}

	val := req.ConfigValue.ValueString()
	resp.PlanValue = types.StringValue(utils.NormalizeIPAddress(val))
}
