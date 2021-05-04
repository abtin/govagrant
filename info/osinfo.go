package info

import (
	"crypto/md5"
	"encoding/base64"
	"fmt"

	"github.com/shirou/gopsutil/v3/host"
)

type OSInfo struct {
	Platform string
	Family   string
	Version  string
}

func NewOsInfo() (OSInfo, error) {
	platform, family, version, err := host.PlatformInformation()
	if err != nil {
		return OSInfo{}, err
	}
	return OSInfo{
		Platform: platform,
		Family:   family,
		Version:  version,
	}, nil
}

func (o OSInfo) String() string {
	return fmt.Sprintf("Platform: %s, Family: %s, Version: %s", o.Platform, o.Family, o.Version)
}

func (o OSInfo) Md5sum() string {
	sum := md5.Sum([]byte(o.String()))
	return base64.StdEncoding.EncodeToString(sum[:])
}
