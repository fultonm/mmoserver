package packet

import (
	"encoding/binary"
	"fmt"

	"github.com/superp00t/etc"
	"github.com/superp00t/gophercraft/vsn"
)

var ClientAddonData = []byte{
	0x9E, 0x2, 0x0, 0x0, 0x78, 0x9C, 0x75, 0xD2, 0x4D, 0x4E, 0xC4, 0x30, 0xC, 0x5, 0xE0, 0x70, 0xA, 0x36, 0xDC, 0x84, 0x15, 0x9D, 0x91, 0xAA, 0x8A, 0xC9, 0x86, 0x66, 0xD6, 0xC8, 0x4D, 0x1E, 0xAD, 0xD5, 0xC4, 0xA9, 0xD2, 0x74, 0xFE, 0xEE, 0xC1, 0x7D, 0x11, 0x3B, 0x90, 0xDC, 0xF5, 0x67, 0x3D, 0x5B, 0x4F, 0x7E, 0x36, 0xC6, 0x34, 0x91, 0x1F, 0xF, 0x2A, 0xE1, 0xF3, 0xCD, 0x4F, 0x8C, 0xB, 0x12, 0xA4, 0x9E, 0x3B, 0xF3, 0x94, 0xAE, 0x2F, 0x27, 0xF3, 0xCF, 0xB, 0x84, 0x74, 0xD9, 0x7C, 0xE5, 0x2C, 0xAA, 0x35, 0x54, 0x6, 0x94, 0x75, 0xCA, 0xCB, 0xE, 0xD7, 0x1A, 0xF1, 0xC5, 0x88, 0xC1, 0xB2, 0x70, 0xA2, 0x45, 0x1B, 0x62, 0x9, 0x2C, 0xA3, 0x1A, 0x70, 0xA0, 0x8, 0x9, 0x54, 0x34, 0xCA, 0x69, 0xA0, 0x7A, 0xCA, 0xE3, 0xAE, 0x39, 0xDC, 0xAA, 0x82, 0x47, 0xC, 0xDB, 0xE8, 0x72, 0x8E, 0xAB, 0x82, 0x6D, 0xBC, 0x2F, 0x93, 0x7A, 0x4A, 0x6B, 0xF, 0x13, 0xE9, 0xDD, 0xB5, 0xB6, 0xDF, 0xCA, 0x5, 0x77, 0x1D, 0x37, 0x8E, 0xA1, 0x21, 0x99, 0x55, 0xED, 0x64, 0x5D, 0xE0, 0xF5, 0xD8, 0xAE, 0x22, 0xF5, 0xD9, 0xCF, 0xA8, 0x7B, 0xF5, 0x58, 0xF2, 0x25, 0xAB, 0xF2, 0x41, 0x1C, 0x54, 0x70, 0xBF, 0x8D, 0xEA, 0xFB, 0x1C, 0x27, 0x58, 0x12, 0x1A, 0xA1, 0xF5, 0xED, 0xF2, 0xC, 0xFD, 0x9, 0x5C, 0xA1, 0x80, 0x7E, 0xE6, 0x18, 0xF7, 0x98, 0x5, 0xE5, 0xAF, 0xBD, 0x7E, 0x1F, 0xDF, 0x7F, 0x0, 0xE1, 0x27, 0xC8, 0x8D,
}

var addonPublicKey = []byte{
	0xC3, 0x5B, 0x50, 0x84, 0xB9, 0x3E, 0x32, 0x42, 0x8C, 0xD0, 0xC7, 0x48, 0xFA, 0x0E, 0x5D, 0x54,
	0x5A, 0xA3, 0x0E, 0x14, 0xBA, 0x9E, 0x0D, 0xB9, 0x5D, 0x8B, 0xEE, 0xB6, 0x84, 0x93, 0x45, 0x75,
	0xFF, 0x31, 0xFE, 0x2F, 0x64, 0x3F, 0x3D, 0x6D, 0x07, 0xD9, 0x44, 0x9B, 0x40, 0x85, 0x59, 0x34,
	0x4E, 0x10, 0xE1, 0xE7, 0x43, 0x69, 0xEF, 0x7C, 0x16, 0xFC, 0xB4, 0xED, 0x1B, 0x95, 0x28, 0xA8,
	0x23, 0x76, 0x51, 0x31, 0x57, 0x30, 0x2B, 0x79, 0x08, 0x50, 0x10, 0x1C, 0x4A, 0x1A, 0x2C, 0xC8,
	0x8B, 0x8F, 0x05, 0x2D, 0x22, 0x3D, 0xDB, 0x5A, 0x24, 0x7A, 0x0F, 0x13, 0x50, 0x37, 0x8F, 0x5A,
	0xCC, 0x9E, 0x04, 0x44, 0x0E, 0x87, 0x01, 0xD4, 0xA3, 0x15, 0x94, 0x16, 0x34, 0xC6, 0xC2, 0xC3,
	0xFB, 0x49, 0xFE, 0xE1, 0xF9, 0xDA, 0x8C, 0x50, 0x3C, 0xBE, 0x2C, 0xBB, 0x57, 0xED, 0x46, 0xB9,
	0xAD, 0x8B, 0xC6, 0xDF, 0x0E, 0xD6, 0x0F, 0xBE, 0x80, 0xB3, 0x8B, 0x1E, 0x77, 0xCF, 0xAD, 0x22,
	0xCF, 0xB7, 0x4B, 0xCF, 0xFB, 0xF0, 0x6B, 0x11, 0x45, 0x2D, 0x7A, 0x81, 0x18, 0xF2, 0x92, 0x7E,
	0x98, 0x56, 0x5D, 0x5E, 0x69, 0x72, 0x0A, 0x0D, 0x03, 0x0A, 0x85, 0xA2, 0x85, 0x9C, 0xCB, 0xFB,
	0x56, 0x6E, 0x8F, 0x44, 0xBB, 0x8F, 0x02, 0x22, 0x68, 0x63, 0x97, 0xBC, 0x85, 0xBA, 0xA8, 0xF7,
	0xB5, 0x40, 0x68, 0x3C, 0x77, 0x86, 0x6F, 0x4B, 0xD7, 0x88, 0xCA, 0x8A, 0xD7, 0xCE, 0x36, 0xF0,
	0x45, 0x6E, 0xD5, 0x64, 0x79, 0x0F, 0x17, 0xFC, 0x64, 0xDD, 0x10, 0x6F, 0xF3, 0xF5, 0xE0, 0xA6,
	0xC3, 0xFB, 0x1B, 0x8C, 0x29, 0xEF, 0x8E, 0xE5, 0x34, 0xCB, 0xD1, 0x2A, 0xCE, 0x79, 0xC3, 0x9A,
	0x0D, 0x36, 0xEA, 0x01, 0xE0, 0xAA, 0x91, 0x20, 0x54, 0xF0, 0x72, 0xD8, 0x1E, 0xC7, 0x89, 0xD2,
}

type AddonList struct {
	Addons      []AddonInfo
	CurrentTime uint32
}

type AddonInfo struct {
	Name    string
	Enabled bool
	CRC     uint32
	Unk1    uint32
}

func ParseAddonList(build vsn.Build, addonData []byte) (*AddonList, error) {
	if len(addonData) < 4 {
		return nil, fmt.Errorf("no addon info")
	}

	aiList := &AddonList{}

	size := binary.LittleEndian.Uint32(addonData[0:4])
	if size == 0 {
		return nil, fmt.Errorf("AddonInfo too small")
	}

	if size > 0xFFFFF {
		return nil, fmt.Errorf("AddonInfo too large")
	}

	addonInfo := etc.FromBytes(Uncompress(addonData[4:]))

	if build.AddedIn(vsn.V3_3_5a) {
		// Size of buffer
		_ = addonInfo.ReadUint32()
	}

	for addonInfo.Available() > 0 {
		ai := AddonInfo{}
		ai.Name = addonInfo.ReadCString()
		ai.Enabled = addonInfo.ReadBool()
		ai.CRC = addonInfo.ReadUint32()
		ai.Unk1 = addonInfo.ReadUint32()
		aiList.Addons = append(aiList.Addons, ai)
	}

	return aiList, nil
}

// It is unknown why this is necessary
func BuildServerAddonResponse(build vsn.Build, info *AddonList) *WorldPacket {
	// var crcCheck uint32 = 0x1c776d01
	// if build == vsn.V3_3_5a {
	// }

	var crcCheck uint32 = 0x4c1c776d

	pkt := NewWorldPacket(SMSG_ADDON_INFO)

	for _, v := range info.Addons {
		state := 1
		if v.Enabled {
			state = 2
		}

		pkt.WriteByte(uint8(state))
		pkt.WriteBool(v.Enabled)

		if v.Enabled {
			usePublicKey := v.CRC != crcCheck
			pkt.WriteBool(usePublicKey)
			if usePublicKey {
				pkt.Write(addonPublicKey)
			}

			pkt.WriteUint32(0)
		}

		pkt.WriteBool(!v.Enabled)
		if !v.Enabled {
			// String? I wonder if this has to do with addons being banned
			pkt.WriteByte(0)
		}
	}

	// zero banned addons
	pkt.WriteUint32(0)
	return pkt
}
