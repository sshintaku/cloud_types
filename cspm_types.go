package cloud_types

type CSPMFilter struct {
	Detailed  bool            `json:"detailed"`
	Fields    []string        `json:"fields"`
	Filters   []FilterObjects `json:"filters"`
	GroupBy   []string        `json:"groupBy"`
	Limit     int             `json:"limit"`
	Offset    int             `json:"offset"`
	PageToken string          `json:"pageToken"`
	SortBy    string          `json:"sortBy"`
	TimeRange TimeObject      `json:"timeRange"`
}

type FilterObjects struct {
	Name     string `json:"name"`
	Operator string `json:"operator"`
	Value    string `json:"value"`
}

type TimeObject struct {
	Type  string    `json:"type"`
	Value TimeValue `json:"value"`
}

type TimeValue struct {
	StartTime int64 `json:"startTime"`
	EndTime   int64 `json:"endTime"`
}

type AlertQuery struct {
	TimeType    string  `json:"timeType"`
	TimeAmount  string  `json:"timeAmount"`
	TimeUnit    string  `json:"timeUnit"`
	Detailed    bool    `json:"detailed"`
	Offset      *int    `json:"offset"`
	Limit       *int    `json:"limit"`
	Severity    string  `json:"policy.severity"`
	PolicyName  string  `json:"policy.name"`
	AlertStatus string  `json:"alert.status"`
	ServiceName *string `json:"cloud.service"`
}

type AlertResponse struct {
	NextPageToken   string       `json:"nextPageToken"`
	AlertModelArray []AlertModel `json:"items"`
}

type AlertModel struct {
	Policy     PolicyModel   `json:"policy"`
	Resource   ResourceModel `json:"resource"`
	Reason     string        `json:"reason"`
	Id         string        `json:"id"`
	Status     string        `json:"status"`
	AlertCount int64         `json:"alertCount"`
}

type EC2Data struct {
	PublicDnsName    string          `json:"publicDnsName"`
	PrivateDnsName   string          `json:"privateDnsName"`
	VpcId            string          `json:"vpcId"`
	SecuirtyGroups   []SecurityGroup `json:"securityGroups"`
	PrivateIPAddress string          `json:"privateIpAddress"`
	PublicIPAddress  string          `json:"publicIpAddress"`
}

type SecurityGroup struct {
	GroupId   string `json:"groupId"`
	GroupName string `json:"groupName"`
}

type ResourceModel struct {
	Account               string                 `json:"account"`
	AccountId             string                 `json:"accountId"`
	CloudAccountAccestors []string               `json:"cloudAccountAncestors"`
	Data                  map[string]interface{} `json:"data"`
	CloudType             string                 `json:"cloudType"`
	CloudServiceName      string                 `json:"cloudServiceName"`
	CloudAccountGroups    []string               `json:"cloudAccountGroups"`
	CloudAccountOwners    []string               `json:"cloudAccountOwners"`
	ResourceType          string                 `json:"resourceType"`
	ResourceApiName       string                 `json:"resourceApiName"`
}

type PolicyModel struct {
	AlertAdditionalInfo interface{}               `json:"alertAdditionalInfo"`
	AlertAttribution    interface{}               `json:"alertAttribution"`
	ComplianceMetadata  []ComplianceMetadataModel `json:"complianceMetadata"`
	Labels              []string                  `json:"labels"`
	Name                string                    `json:"name"`
	PolicyType          string                    `json:"policyType"`
	PolicySubTypes      string                    `json:"policySubTypes"`
	Rule                RuleType                  `json:"rule"`
	Recommendation      string                    `json:"recommendation"`
	LastModifiedBy      string                    `json:"lastModifiedBy"`
	Severity            string                    `json:"severity"`
}

type RemediationModel struct {
	CliScriptTemplate string        `json:"cliScriptTemplate"`
	Description       string        `json:"description"`
	Actions           []ActionModel `json:"actions"`
}

type ActionModel struct {
	Operation string `json:"operation"`
	Payload   string `json:"payload"`
}

type RuleType struct {
	Criteria   string      `json:"criteria"`
	Name       string      `json:"name"`
	Parameters interface{} `json:"parameters"`
	Type       string      `json:"type"`
}

type PropertyTags struct {
	Property string `json:"property"`
}

type ComplianceMetadataModel struct {
	ComplianceId           string `json:"complianceId"`
	RequirementDescription string `json:"requirementDescription"`
	RequirementName        string `json:"requirementName"`
	SectionDescription     string `json:"sectionDescription"`
	StandardDescription    string `json:"standardDescription"`
	StandardName           string `json:"standardName"`
}
