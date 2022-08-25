package cloud_types

import "time"

type RegistryScan struct {
	Image           RegistryImageInfo `json:"image"`
	ScanTime        time.Time         `json:"scanTime"`
	OsDistro        string            `json:"osDistro"`
	OsDistroVersion string            `json:"osDistroVersion"`
	OsDistroRelease string            `json:"osDistroRelease"`
	Distro          string            `json:"distro"`
}

type RegistryImageInfo struct {
	Created time.Time `json:"created"`
}
