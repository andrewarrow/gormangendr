package main

import (
	"encoding/json"
	"fmt"
)

type Producer struct {
	Addr      string `json:"addr"`
	Port      int    `json:"port"`
	Continent string `json:"continent"`
	State     string `json:"state"`
}

type ProducerList struct {
	Producers []Producer `json:"Producers"`
}

func LoadRelays() {
	var pl ProducerList
	jsonString := DoGet("relays/topology.json")
	json.Unmarshal([]byte(jsonString), &pl)
	fmt.Println(pl.Producers[0])
}
