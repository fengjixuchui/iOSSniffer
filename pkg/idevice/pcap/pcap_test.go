package pcap

import (
	"context"
	"encoding/hex"
	"fmt"
	"os"
	"testing"

	"github.com/gofmt/iOSSniffer/pkg/idevice"
)

func TestClient_ReadOPacket(t *testing.T) {
	conn, err := idevice.NewConn()
	if err != nil {
		t.Fatal(err)
	}
	defer func(conn *idevice.Conn) {
		_ = conn.Close()
	}(conn)

	devices, err := conn.ListDevices()
	if err != nil {
		t.Fatal(err)
	}

	for _, device := range devices {
		cli, err := NewClient(device.UDID)
		if err != nil {
			t.Fatal(err)
		}

		err = cli.ReadPacket(context.Background(), "", os.Stdout, func(data []byte) {
			fmt.Println("\n----")
			fmt.Println(hex.Dump(data))
		})
		if err != nil {
			t.Fatal(err)
		}
		_ = cli.Close()
	}
}
