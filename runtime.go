package cloud_types

type RuntimeAudit struct {
	AccountID        string   `json:"accountID"`
	App              string   `json:"app"`
	AppId            string   `json:"appId"`
	AttackTechniques []string `json:"attackTechniques"`
	AttackType       []string `json:"attackType"`
	Collections      []string `json:"collections"`
	Cluster          string   `json:"Cluster"`
	RuleName         string   `json:"ruleName"`
	Count            int      `json:"count"`
	Command          string   `json:"Command"`
	ContainerName    string   `json:"containerName"`
	ContainerId      string   `json:"containerId"`
	ImageName        string   `json:"imageName"`
	ImageId          string   `json:"imageId"`
	Namespace        string   `json:"namespace"`
	Message          string   `json:"msg"`
}
