package serviceinfo

// ArkSecHubGetServiceInfo represents the response to get service info in Ark SIA.
type ArkSecHubGetServiceInfo struct {
	TenantRoleArn string `json:"tenant_role_arn" mapstructure:"tenant_role_arn" flag:"tenant_role_arn" desc:"Role ARN of the Secrets Hub Tenant" validate:"required"`
	TenantPamType string `json:"tenant_pam_type" mapstructure:"tenant_pam_type" flag:"tenant_pam_type" desc:"Tenant PAM Type" validate:"required"`
}
