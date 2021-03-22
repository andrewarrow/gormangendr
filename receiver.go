package main

import (
	"fmt"
	"io/ioutil"
)

func ReadPublicKey() []byte {
	dat, err := ioutil.ReadFile("receiver_public.key")
	if err != nil {
		fmt.Println("missing receiver_public.key")
		return []byte{}
	}
	return dat
}
