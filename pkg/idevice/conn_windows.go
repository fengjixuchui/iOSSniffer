package idevice

func usbmuxdDial() (net.Conn, error) {
	return net.Dial("tcp", "localhost:27015")
}
