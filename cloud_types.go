package cloud_types

type JwtToken struct {
	Token string `json:"token"`
}

type Authentication struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type InputObject struct {
	Auth       Authentication
	Token      JwtToken
	Host       string
	AzureInput AzureInfo
}

type AzureInfo struct {
	AccountInfo        CloudAccount `json:"cloudAccount"`
	ClientId           string       `json:"clientId"`
	Key                string       `json:"key"`
	MonitorFlowLogs    bool         `json:"monitorFlowLogs"`
	TenantId           string       `json:"tenantId"`
	ServicePrincipalId string       `json:"servicePrincipalId"`
}

type PrismaClient struct {
	Token    string
	HostName string
}

type Registry struct {
	Tag RegistryTag `json:tag`
}

type RegistryTag struct {
	Registry string `json:"registry"`
	Repo     string `json:"repo"`
	Tag      string `json:"tag"`
	Digest   string `json:"digest"`
}

type CloudAccount struct {
	AccountId     string   `json:"accountId"`
	Enabled       bool     `json:"enabled"`
	GroupIds      []string `json:"groupIds"`
	Name          string   `json:"name"`
	AccountType   string   `json:"accountType"`
	ProtectedMode string   `json:"protectedMode"`
}

type PrismaCredentials struct {
	Secret         Credentials `json:"secret"`
	ServiceAccount interface{} `json:"serviceAccount"`
	Type           string      `json:"type"`
	Description    string      `json:"description"`
	Name           string      `json:"_id"`
}

type Credentials struct {
	Encrypted *string     `json:"encrypted"`
	Plain     interface{} `json:"plain"`
}

type Input struct {
	SubscriptionId string
	Name           string
}

type InputArray struct {
	InputCollection []Input
}

type Twistlock struct {
	BaseURI string `json:"twistlockUrl"`
}
