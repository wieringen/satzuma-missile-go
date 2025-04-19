package missile

import (
	"fmt"

	"github.com/google/gousb"
)

type MissileDevice struct {
	ctx    *gousb.Context
	device *gousb.Device
}

var vendorProductIDs = []struct {
	vendorID  gousb.ID
	productID gousb.ID
}{
	{0x1130, 0x0202},
	{0x0416, 0x9391},
}

func NewMissileDevice() *MissileDevice {
	return &MissileDevice{
		ctx: gousb.NewContext(),
	}
}

func (m *MissileDevice) Open() error {
	for _, vp := range vendorProductIDs {
		dev, err := m.ctx.OpenDeviceWithVIDPID(vp.vendorID, vp.productID)
		if err != nil {
			return fmt.Errorf("failed to open USB device: %w", err)
		}

		if dev != nil {
			m.device = dev
			if err := m.device.SetAutoDetach(true); err != nil {
				return fmt.Errorf("failed to set auto-detach: %w", err)
			}
			return nil
		}
	}
	return fmt.Errorf("no compatible missile devices found")
}

func (m *MissileDevice) Move(direction byte) error {
	data := []byte{0x5f, direction, 0xe0, 0xff, 0xfe}
	_, err := m.device.Control(0x21, 0x09, 0x0300, 0x00, data)
	if err != nil {
		return fmt.Errorf("failed to send control message: %w", err)
	}
	return nil
}

func (m *MissileDevice) Close() {
	if m.device != nil {
		m.device.Close()
	}
	if m.ctx != nil {
		m.ctx.Close()
	}
}
