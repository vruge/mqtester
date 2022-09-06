package wrappedclient

import (
	"fmt"
)

func Case1() {
	connectionType := "tcp"
	port := 8883
	sslEnable := false
	fmt.Println("Start")
	fmt.Println("tcp connection, only publish")
	for i := 0; i < 10; i++ {
		clientData := CreateDefaultClientData(connectionType, port, i, sslEnable)
		switch i {
		case 1:
			go clientData.CreateClientDefaultPublish("company1/device1", "temp")
		case 2:
			go clientData.CreateClientDefaultPublish("company1/device2", "temp")
		case 3:
			go clientData.CreateClientDefaultPublish("company1/device2", "watt")
		case 4:
			go clientData.CreateClientDefaultPublish("company1/device3", "watt")
		case 5:
			go clientData.CreateClientDefaultPublish("company1/device3", "temp")
		case 6:
			go clientData.CreateClientDefaultPublish("company1/device3", "watt")
		case 7:
			go clientData.CreateClientDefaultPublish("company2/device1", "watt")
		case 8:
			go clientData.CreateClientDefaultPublish("company2/device2", "watt")
		case 9:
			go clientData.CreateClientDefaultPublish("company3/device1", "watt")
		}
	}
	fmt.Println("End")
}
