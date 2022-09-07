package wrappedclient

import (
	"fmt"
	"time"
)

type BaseCasesData struct {
	topic   string
	tagname string
}

func (base *BaseCasesData) GetDefaultPub(idx int) {
	switch idx {
	case 0:
		base.topic = "company1/device1"
		base.tagname = "temp"
	case 1:
		base.topic = "company1/device2"
		base.tagname = "watt"
	case 2:
		base.topic = "company1/device3"
		base.tagname = "watt"
	case 3:
		base.topic = "company1/device3"
		base.tagname = "watt"
	case 4:
		base.topic = "company2/device1"
		base.tagname = "temp"
	case 5:
		base.topic = "company2/device1"
		base.tagname = "temp"
	case 6:
		base.topic = "company2/device1"
		base.tagname = "watt"
	case 7:
		base.topic = "company2/device2"
		base.tagname = "watt"
	case 8:
		base.topic = "company3/device1"
		base.tagname = "watt"
	case 9:
		base.topic = "company3/device2"
		base.tagname = "watt"
	}
}

func Case1() {
	connectionType := "tcp"
	port := 1883
	sslEnable := false
	fmt.Println("Start")
	fmt.Println("tcp connection, only publish")
	for i := 0; i < 10; i++ {
		clientData := CreateDefaultClientData(connectionType, port, i, sslEnable)
		baseCaseData := BaseCasesData{topic: "company/device", tagname: "none"}
		baseCaseData.GetDefaultPub(i)
		go clientData.CreateClientDefaultPublish(baseCaseData.topic, baseCaseData.tagname)
		time.Sleep(time.Second)
	}
	fmt.Println("End")
	fmt.Println("Wait....")
	time.Sleep(time.Hour)
}
