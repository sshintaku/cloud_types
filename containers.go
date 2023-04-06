package cloud_types

type ContainerType struct {
	Info ContainerInfo `json:"info"`
}

type ContainerInfo struct {
	Image            string              `json:"image"`
	ImageName        string              `json:"imageName"`
	Cluster          string              `json:"cluster"`
	Namespace        string              `json:"namespace"`
	App              string              `json:"app"`
	ImageId          string              `json:"imageId"`
	Name             string              `json:"name"`
	Labels           []string            `json:"labels"`
	ExternalLabels   []ExternalLabelType `json:"externalLabels"`
	ComplianceIssues []Compliance        `json:"complianceIssues"`
}

type ExternalLabelType struct {
	Key        string `json:"key"`
	SourceName string `json:"sourceName"`
	Value      string `json:"value"`
}

type ContainerResult struct {
	Hostname         string        `json:"hostname"`
	Collections      []string      `json:"collections"`
	ContainerDetails ContainerInfo `json:"info"`
}
