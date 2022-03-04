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

type scanResult struct {
	Id              string           `json:"_id"`
	Collections     []string         `json:"collections"`
	AllCompliance   ComplianceIssues `json:"allCompliance"`
	Vulnerabilities []Vulnerable     `json:"vulnerabilities"`
}

type ComplianceIssues struct {
	Compliance []ComplianceAttributes `json:"compliance"`
}

type ComplianceAttributes struct {
	Id             int    `json:"id"`
	Severity       string `json:"severity"`
	Description    string `json:"description"`
	Title          string `json:"title"`
	PackageName    string `json:"packageNAme"`
	PackageVersion string `json:"packageVersion"`
}

type Vulnerable struct {
	CVE             string   `json:"cve"`
	Exploit         []string `json:"exploit"`
	ApplicableRules []string `json:"applicableRules"`
	Description     string   `json:"description"`
	Block           bool     `json:"block"`
	Cause           string   `json:"cause"`
}

type ComplianceObject struct {
	Type string `json:"type"`

	ComplianceIssues    []Compliance    `json:"complianceIssues"`
	VulnerabilityIssues []Vulnerability `json:"vulnerabilities"`
	Clusters            []string        `json:"clusters"`
	Labels              []string        `json:"labels"`
	PackageList         []PackageInfo   `json:"packages"`
	RepoTags            RepoTag         `json:"repoTag"`
}

type RepoTag struct {
	Registry string `json:"registry"`
	Repo     string `json:"repo"`
	Tag      string `json:"tag"`
}

type PackageInfo struct {
	PackageType string    `json:"pkgsType"`
	Packages    []Package `json:"pkgs"`
}

type Package struct {
	Version  string `json:"version"`
	CveCount int16  `json:"cveCount"`
	Name     string `json:"name"`
}

type Vulnerability struct {
	Severity       string  `json:"severity"`
	CVE            string  `json:"cve"`
	PackageName    string  `json:"packagename"`
	PackageVersion string  `json:"packageversion"`
	Description    string  `json:"description"`
	Status         string  `json:"status"`
	Type           string  `json:"type"`
	FixDate        int64   `json:"fixDate"`
	CVSS           float32 `json:"cvss"`
	Link           string  `json:"link"`
}

type Compliance struct {
	Severity       string `json:"severity"`
	CVE            string `json:"cve"`
	Cause          string `json:"cause"`
	Description    string `json:"description"`
	Title          string `json:"title"`
	PackageName    string `json:"packageName"`
	PackageVersion string `json:"packageVersion"`
	FixDate        int64  `json:fixDate`
}

type MaintainerSummary struct {
	ClusterName          []string
	Registry             string
	Repository           string
	Package              string
	PackageType          string
	FixDate              string
	CVE                  string
	CVSS                 float32
	Link                 string
	Description          string
	ComplianceSummary    AlarmCounter
	VulnerabilitySummary AlarmCounter
}

type AlarmCounter struct {
	Critical  int
	High      int
	Important int
	Medium    int
	Moderate  int
	Low       int
}
