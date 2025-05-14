package services

import (
	"github.com/strick-j/ark-sdk-golang/pkg/models/actions"
	sechubserviceinfo "github.com/strick-j/ark-sdk-golang/pkg/models/services/sechub/serviceinfo"
)

// SecHubActionToSchemaMap is a map that defines the mapping between SecHub action names and their corresponding schema types.
var SecHubActionToSchemaMap = map[string]interface{}{
	"service-info": &sechubserviceinfo.ArkSecHubGetServiceInfo{},
}

// SecHubActions is a struct that defines the SecHub action for the Ark service.
var SecHubActions = &actions.ArkServiceActionDefinition{
	ActionName: "sechub",
	Schemas:    SecHubActionToSchemaMap,
}
