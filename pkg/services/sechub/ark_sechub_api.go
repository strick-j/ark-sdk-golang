package sechub

import (
	"github.com/cyberark/ark-sdk-golang/pkg/auth"
	"github.com/strick-j/ark-sdk-golang/pkg/services/sechub/serviceinfo"
)

// ArkSecHubAPI is a struct that provides access to the Ark SecHub API as a wrapped set of services.
type ArkSecHubAPI struct {
	serviceinfoService *serviceinfo.ArkSecHubServiceInfoService
}

// NewArkSIAAPI creates a new instance of ArkSIAAPI with the provided ArkISPAuth.
func NewArkSIAAPI(ispAuth *auth.ArkISPAuth) (*ArkSecHubAPI, error) {
	var baseIspAuth auth.ArkAuth = ispAuth
	serviceinfoService, err := serviceinfo.NewArkSecHubServiceInfoService(baseIspAuth)
	if err != nil {
		return nil, err
	}
	return &ArkSecHubAPI{
		serviceinfoService: serviceinfoService,
	}, nil
}
