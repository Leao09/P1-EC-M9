package main

import (
	"fmt"
	MQTT "github.com/eclipse/paho.mqtt.golang"
	"strconv"
	"strings"
)
const (
	AlertaBaixo = "**ALERTA**: Temperatura baixa"
	AlertaAlto = "**ALERTA**: Temperatura alta"
	limiteFB = -25
	limiteFA = -15
	limiteGA = 10
	limiteGB = 2
)

var messagePubHandler MQTT.MessageHandler = func(client MQTT.Client, msg MQTT.Message) {
	fmt.Printf("Recebido: %s do tópico: %s\n", msg.Payload(), msg.Topic())

	result := strings.Split(string(msg.Payload()), " - ")
	name := result[1]
	F1, _ := strconv.Atoi(result[2])
	F2, _ := strconv.Atoi(result[3])
	G1, _ := strconv.Atoi(result[4])
	G2, _ := strconv.Atoi(result[5])

	if int(F1) <= int(limiteFB){
		fmt.Printf("%s : Freezer1 %d : %s\n", name, F1, AlertaBaixo)
	}
	if int(F1) >= int(limiteFA){
		fmt.Printf("%s : Freezer 1 %d : %s\n", name, F1, AlertaAlto)
	}
	if int(F2) <= int(limiteFB){
		fmt.Printf("%s : Freezer 2 %d : %s\n", name, F2, AlertaBaixo)
	}
	if int(F2) >= int(limiteFA){
		fmt.Printf("%s : Freezer 2 %d : %s\n", name, F2, AlertaAlto)
	}
	if int(G1) <= int(limiteGB){
		fmt.Printf("%s : Geladeira 1 %d : %s\n", name, G1, AlertaBaixo)
	}
	if int(G1) >= int(limiteGA){
		fmt.Printf("%s : Geladeira 1 %d : %s\n", name, G1, AlertaAlto)
	}
	if int(G2) <= int(limiteGB){
		fmt.Printf("%s : Geladeira 2%d : %s\n", name, G2, AlertaBaixo)
	}
	if int(G2) >= int(limiteGA){
		fmt.Printf("%s : Geladeira 2 %d : %s\n", name, 2, AlertaAlto)
	}

}

func main() {
	opts := MQTT.NewClientOptions().AddBroker("tcp://localhost:4567")
	opts.SetClientID("go_subscriber")
	opts.SetDefaultPublishHandler(messagePubHandler)

	client := MQTT.NewClient(opts)
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}

	if token := client.Subscribe("/pond-2", 1, nil); token.Wait() && token.Error() != nil {
		fmt.Println(token.Error())
		return
	}

	fmt.Println("Subscriber está rodando. Pressione CTRL+C para sair.")
	select {} // Bloqueia indefinidamente
}