package def

import (
	"fmt"
)

const (
	AppName        = "msg"
	AppUsage       = "Message Center"
	AppDescription = "msg is a message manage service"
)

var (
	Branch    string
	Version   string
	Commit    string
	BuildTime string
)

func AppVersion() string {
	return fmt.Sprintf("%s-%s_%s (%s)", Version, Branch, Commit, BuildTime)
}
