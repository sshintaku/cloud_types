package cloud_types

type AuditType struct {
	ContainerId   string  `json:"containerID"`
	ContainerName string  `json:"containerName"`
	Audits        []Audit `json:"audits"`
}

type Audit struct {
	AppId            string   `json:"appID"`
	AppName          string   `json:"appName"`
	AttackTechniques []string `json:"attackTechniques"`
	AttackType       string   `json:"attackType"`
	ContainerId      string   `json:"containerID"`
	ContainerName    string   `json:"containerName"`
	Command          string   `json:"command"`
	ProcessPath      string   `json:"processPath"`
	RawEvent         string   `json:"rawEvent"`
	Message          string   `json:"msg"`
	FilePath         string   `json:"filepath"`
}
