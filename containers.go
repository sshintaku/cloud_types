package cloud_types

type ContainerType struct {
	Info ContainerInfo `json:"info"`
}

type ContainerInfo struct {
	Image          string            `json:"image"`
	ImageName      string            `json:"imageName"`
	Labels         []string          `json:"labels"`
	ExternalLabels ExternalLabelType `json:"externalLabels"`
}

type ExternalLabelType struct {
	Key        string `json:"key"`
	SourceName string `json:"sourceName"`
	Value      string `json:"value"`
}
