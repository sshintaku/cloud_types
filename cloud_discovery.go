package cloud_types

type DiscoveryResult struct {
	AccountId    string          `json:"accountId"`
	Collections  []string        `json:"collections"`
	CredentialId string          `json:"credentialId"`
	Entities     []CloudEntities `json:"entities"`
	Defended     int             `json:"defended"`
	CloudError   string          `json:"err"`
	Project      string          `json:"project"`
	Provider     string          `json:"provider"`
	Region       string          `json:"region"`
	Registry     string          `json:"registry"`
	ServiceType  string          `json:"serviceType"`
	Name         string          `json:"name"`
}

type CloudTypeResult struct {
	Region string `json:"region"`
	Status string
}

type CloudEntities struct {
	ActiveServicesCount int64  `json:"activeServicesCount"`
	Arn                 string `json:"arn"`
	ContainerGroup      string `json:"containerGroup"`
	CreationDate        string `json:"createdAt"`
	ResourceGroup       string `json:"resourceGroup"`
	NodesCount          int    `json:"nodesCount"`
	Defended            bool   `json:"defended"`
	Status              string `json:"status"`
}
