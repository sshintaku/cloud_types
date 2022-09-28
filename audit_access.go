package cloud_types

import (
	"time"
)

type AuditApiType struct {
	Api      string    `json:"api"`
	SourceIP string    `json:"sourceIp"`
	Time     time.Time `json:"time"`
	Username string    `json:"username"`
	Diff     string    `json:"diff"`
}
