package wrappedclient

import (
	"crypto/tls"
	"crypto/x509"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"math"
	"os"
	"time"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

type ClientData struct {
	ClientCert     string
	ServerCert     string
	ClientKey      string
	SslEnable      bool
	Url            string
	Port           int
	UserName       string
	PassWord       string
	ConnectionType string
	ClientId       string
}

func (data *ClientData) get_address() string {
	return fmt.Sprintf("%s://%s:%d", data.ConnectionType, data.Url, data.Port)
}

func messagePubHandlerCreate(clientId string) mqtt.MessageHandler {
	var messagePubHandler mqtt.MessageHandler = func(client mqtt.Client, msg mqtt.Message) {

		data := parse_messsage(string(msg.Payload()))
		fmt.Printf("%s -- %+v\n", clientId, data)
	}
	return messagePubHandler
}

var connectHandler mqtt.OnConnectHandler = func(client mqtt.Client) {
	fmt.Println("Connected")
}

var connectLostHandler mqtt.ConnectionLostHandler = func(client mqtt.Client, err error) {
	fmt.Printf("Connect lost: %v", err)
}

func (certPaths *ClientData) NewTlsConfig() *tls.Config {
	certpool := x509.NewCertPool()
	ca, err := ioutil.ReadFile(certPaths.ServerCert)
	if err != nil {
		log.Fatalln(err.Error())
	}
	certpool.AppendCertsFromPEM(ca)
	// Import client certificate/key pair
	clientKeyPair, err := tls.LoadX509KeyPair(certPaths.ClientCert, certPaths.ClientKey)
	if err != nil {
		panic(err)
	}
	return &tls.Config{
		RootCAs:            certpool,
		ClientAuth:         tls.NoClientCert,
		ClientCAs:          nil,
		InsecureSkipVerify: true,
		Certificates:       []tls.Certificate{clientKeyPair},
	}
}

func (clientData *ClientData) createClient() *mqtt.Client {

	opts := mqtt.NewClientOptions()
	opts.AddBroker(clientData.get_address())
	opts.SetClientID(clientData.ClientId)
	if clientData.SslEnable {
		tlsconfig := clientData.NewTlsConfig()
		opts.SetTLSConfig(tlsconfig)
	}

	//opts.SetUsername(userName)
	//opts.SetPassword(passWord)
	opts.SetDefaultPublishHandler(messagePubHandlerCreate(clientData.ClientId))
	opts.OnConnect = connectHandler
	opts.OnConnectionLost = connectLostHandler
	client := mqtt.NewClient(opts)
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		log.Fatalln(token.Error())
		panic(token.Error())
	}
	return &client
}

func (clientData ClientData) CreateClientDefaultPublish(topic, tag_name string) {
	fmt.Println("CreateClientDefaultPublish")
	client := clientData.createClient()
	publish(client, topic, tag_name)
}

func (clientData ClientData) createClientDefaultSub(topic string) {
	client := clientData.createClient()
	sub(client, topic)
}

func makeTimestamp() float64 {
	return float64(time.Now().UnixNano()) / float64(time.Millisecond) * 1e-8
}

type MeasureDataRaw struct {
	Value     float32 `json:"value"`
	TimeStamp float64 `json:"timestamp"`
	Name      string  `json:"name"`
}

func parse_messsage(msg string) MeasureDataRaw {
	data := MeasureDataRaw{}
	err := json.Unmarshal([]byte(msg), &data)
	if err != nil {
		fmt.Println(err.Error())
	}
	return data
}

func sub(client *mqtt.Client, topic string) {
	token := (*client).Subscribe(topic, 1, nil)
	token.Wait()
}

func publish(client *mqtt.Client, topic string, tag_name string) {
	a := 0.0
	b := 1.0
	c := 0.0
	d := 1.0
	switch tag_name {
	case "temp":
		a = 1.0
		b = -1.4
		d = 3.14
	case "watt":
		a = 7.0
		b = 1.3
		d = 0.1
	}

	for range time.Tick(time.Duration(1) * time.Second) {
		text := fmt.Sprintf("{\"value\": %f, \"timestamp\": %f, \"name\": \"%s\"}", a+b*math.Sin(c+d*makeTimestamp()), makeTimestamp(), tag_name)
		token := (*client).Publish(topic, 0, false, text)
		token.Wait()
	}
}

func fileExist(filename string) bool {
	ret := true
	if _, err := os.Stat(filename); errors.Is(err, os.ErrNotExist) {
		ret = false
	}
	return ret
}

func fileNotExistWarning(filename string) {
	if !fileExist(filename) {
		fmt.Printf("Not found file: %s\n", filename)
	}

}

func CreateDefaultClientData(connectionType string, port int, idx int, sslEnable bool) *ClientData {
	brokerUrl := "localhost"
	clientId := fmt.Sprintf("client-id-%d", idx)
	clientCert := fmt.Sprintf("certs/client%d/client.pem", idx)
	clientKey := fmt.Sprintf("certs/client%d/client.key", idx)
	serverCert := "certs/server/hivemq-server-cert.pem"

	fileNotExistWarning(clientCert)
	fileNotExistWarning(clientKey)
	fileNotExistWarning(serverCert)

	return &ClientData{
		ClientCert:     clientCert,
		ServerCert:     serverCert,
		ClientKey:      clientKey,
		SslEnable:      sslEnable,
		Url:            brokerUrl,
		Port:           port,
		ConnectionType: connectionType,
		ClientId:       clientId,
	}
}
