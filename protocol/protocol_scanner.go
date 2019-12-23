package protocol

import (
	"github.com/mhatta/onionscan/config"
	"github.com/mhatta/onionscan/report"
)

type Scanner interface {
	ScanProtocol(hiddenService string, onionscanConfig *config.OnionScanConfig, report *report.OnionScanReport)
}
