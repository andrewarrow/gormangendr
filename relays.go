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

func LoadRelays() []string {
	var pl ProducerList
	jsonString := DoGet("relays/topology.json")
	json.Unmarshal([]byte(jsonString), &pl)
	m := map[string][]string{}
	for _, p := range pl.Producers {
		s := fmt.Sprintf("%s:%d", p.Addr, p.Port)
		key := p.Continent + "|" + p.State
		m[key] = append(m[key], s)
	}
	na := m["North America|California"]
	fmt.Println(len(na))
	return na

}
