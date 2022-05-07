package syslog

import (
	"io"

	"github.com/gofmt/iOSSniffer/pkg/idevice/lockdownd"
)

func Syslog(udid string) (io.ReadCloser, error) {
	cli, err := lockdownd.NewClientForService("com.apple.syslog_relay", udid, false)
	if err != nil {
		return nil, err
	}

	if err := cli.Send("watch"); err != nil {
		return nil, err
	}

	return cli.Conn(), nil
}
