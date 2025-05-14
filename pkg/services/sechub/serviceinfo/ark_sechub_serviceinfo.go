package serviceinfo

import (
	"context"
	"fmt"
	"io"
	"net/http"

	"github.com/cyberark/ark-sdk-golang/pkg/auth"
	"github.com/cyberark/ark-sdk-golang/pkg/common"
	"github.com/cyberark/ark-sdk-golang/pkg/common/isp"
	"github.com/cyberark/ark-sdk-golang/pkg/services"
	"github.com/mitchellh/mapstructure"
	serviceinfomodels "github.com/strick-j/ark-sdk-golang/pkg/models/services/sechub/serviceinfo"
)

const (
	sechubURL = "/api/info"
)

// SecHubServiceInfoServiceConfig is the configuration for the Secrets Hub Service Info service.
var SecHubServiceInfoServiceConfig = services.ArkServiceConfig{
	ServiceName:                "sechub-serviceinfo",
	RequiredAuthenticatorNames: []string{"isp"},
	OptionalAuthenticatorNames: []string{},
}

// SecHubServiceInfoService is the service for managing db secrets.
type ArkSecHubServiceInfoService struct {
	services.ArkService
	*services.ArkBaseService
	ispAuth *auth.ArkISPAuth
	client  *isp.ArkISPServiceClient
}

// NewSecHubServiceInfoService creates a new instance of SecHubServiceInfoService.
func NewSecHubServiceInfoService(authenticators ...auth.ArkAuth) (*ArkSecHubServiceInfoService, error) {
	serviceInfoService := &ArkSecHubServiceInfoService{}
	var serviceInfoServiceInterface services.ArkService = serviceInfoService
	baseService, err := services.NewArkBaseService(serviceInfoServiceInterface, authenticators...)
	if err != nil {
		return nil, err
	}
	ispBaseAuth, err := baseService.Authenticator("isp")
	if err != nil {
		return nil, err
	}
	ispAuth := ispBaseAuth.(*auth.ArkISPAuth)
	client, err := isp.FromISPAuth(ispAuth, "dpa", ".", "", serviceInfoService.refreshSecHubAuth)
	if err != nil {
		return nil, err
	}
	serviceInfoService.client = client
	serviceInfoService.ispAuth = ispAuth
	serviceInfoService.ArkBaseService = baseService
	return serviceInfoService, nil
}

func (s *ArkSecHubServiceInfoService) refreshSecHubAuth(client *common.ArkClient) error {
	err := isp.RefreshClient(client, s.ispAuth)
	if err != nil {
		return err
	}
	return nil
}

// ServiceInfo retrieves the service info from the Secrets Hub service.
func (s *ArkSecHubServiceInfoService) ServiceInfo(getSecret *serviceinfomodels.AArkSecHubGetServiceInfo) (*serviceinfomodels.ArkSecHubGetServiceInfo, error) {
	s.Logger.Info("Getting secret [%s]", getSecret.SecretID)
	response, err := s.client.Get(context.Background(), fmt.Sprintf(sechubURL, getSecret.SecretID), nil)
	if err != nil {
		return nil, err
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			common.GlobalLogger.Warning("Error closing response body")
		}
	}(response.Body)
	if response.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to get service info - [%d] - [%s]", response.StatusCode, common.SerializeResponseToJSON(response.Body))
	}
	serviceinfoJSON, err := common.DeserializeJSONSnake(response.Body)
	if err != nil {
		return nil, err
	}
	var serviceinfo serviceinfomodels.ArkSecHubGetServiceInfo
	err = mapstructure.Decode(serviceinfoJSON, &serviceinfo)
	if err != nil {
		return nil, err
	}
	return &serviceinfo, nil
}

// ServiceConfig returns the service configuration for the ArkSIASecretsVMService.
func (s *ArkSecHubServiceInfoService) ServiceConfig() services.ArkServiceConfig {
	return SecHubServiceInfoServiceConfig
}
