package hikvision

import (
	"encoding/xml"
	"log"
)

const (
	path = "/ISAPI/System/deviceInfo"
)

// DeviceInfo represents to XML_DeviceInfo resource.
type DeviceInfo struct {
	XMLName              xml.Name              `xml:"DeviceInfo,omitempty"`
	DeviceName           string                `xml:"deviceName,omitempty"`
	DeviceID             string                `xml:"deviceID,omitempty"`
	DeviceDescription    string                `xml:"deviceDescription,omitempty"`
	DeviceLocation       string                `xml:"deviceLocation,omitempty"`
	DeviceStatus         string                `xml:"deviceStatus,omitempty"`
	DetailAbnormalStatus *DetailAbnormalStatus `xml:"DetailAbnormalStatus,omitempty"`
	SystemContact        string                `xml:"systemContact,omitempty"`
	Model                string                `xml:"model,omitempty"`
	SerialNumber         string                `xml:"serialNumber,omitempty"`
	MacAddress           string                `xml:"macAddress,omitempty"`
	FirmwareVersion      string                `xml:"firmwareVersion,omitempty"`
	FirmwareReleasedDate string                `xml:"firmwareReleasedDate,omitempty"`
	BootVersion          string                `xml:"bootVersion,omitempty"`
	BootReleasedDate     string                `xml:"bootReleasedDate,omitempty"`
	HardwareVersion      string                `xml:"hardwareVersion,omitempty"`
	EncoderVersion       string                `xml:"encoderVersion,omitempty"`
	EncoderReleasedDate  string                `xml:"encoderReleasedDate,omitempty"`
	DecoderVersion       string                `xml:"decoderVersion,omitempty"`
	DecoderReleasedDate  string                `xml:"decoderReleasedDate,omitempty"`
	SoftwareVersion      string                `xml:"softwareVersion,omitempty"`
	Capacity             int                   `xml:"capacity,omitempty"`
	UsedCapacity         int                   `xml:"usedCapacity,omitempty"`
	DeviceType           string                `xml:"deviceType,omitempty"`
	TelecontrolID        int                   `xml:"telecontrolID,omitempty"`
	SupportBeep          bool                  `xml:"supportBeep,omitempty"`
	ActualFloorNum       int                   `xml:"actualFloorNum,omitempty"`
	SubChannelEnabled    bool                  `xml:"subChannelEnabled,omitempty"`
	ThrChannelEnabled    bool                  `xml:"thrChannelEnabled,omitempty"`
	RadarVersion         string                `xml:"radarVersion,omitempty"`
	LocalZoneNum         int                   `xml:"localZoneNum,omitempty"`
	AlarmOutNum          int                   `xml:"alarmOutNum,omitempty"`
	DistanceResolution   float32               `xml:"distanceResolution,omitempty"`
	AngleResolution      float32               `xml:"angleResolution,omitempty"`
	SpeedResolution      float32               `xml:"speedResolution,omitempty"`
	DetectDistance       float32               `xml:"detectDistance,omitempty"`
	LanguageType         string                `xml:"languageType,omitempty"`
	RelayNum             int                   `xml:"relayNum,omitempty"`
	ElectroLockNum       int                   `xml:"electroLockNum,omitempty"`
	RS485Num             int                   `xml:"RS485Num,omitempty"`
	PowerOnMode          string                `xml:"powerOnMode,omitempty"`
}

// DetailAbnormalStatus represents the device abnormal status,
// which is used in the XML_DeviceInfo resource.
type DetailAbnormalStatus struct {
	HardDiskFull         bool `xml:"hardDiskFull,omitempty"`
	HardDiskError        bool `xml:"hardDiskError,omitempty"`
	EthernetBroken       bool `xml:"ethernetBroken,omitempty"`
	IPAddrConflict       bool `xml:"ipaddrConflict,omitempty"`
	IllegalAccess        bool `xml:"illegalAccess,omitempty"`
	RecordError          bool `xml:"recordError,omitempty"`
	RaidLogicDiskError   bool `xml:"raidLogicDiskError,omitempty"`
	SpareWorkDeviceError bool `xml:"spareWorkDeviceError,omitempty"`
}

// GetDeviceInfo executes a HTTP GET request to the
// /ISAPI/System/deviceInfo endpoint.
// TODO: Returns XML_Response when error
func (c *Client) GetDeviceInfo() (*DeviceInfo, error) {
	method := "GET"
	log.Println(method, path)
	body, err := c.Do(method, path)
	if err != nil {
		return nil, err
	}
	data := DeviceInfo{}
	err = xml.Unmarshal(body, &data)
	if err != nil {
		return nil, err
	}
	return &data, nil
}
