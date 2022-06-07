package main

import (
	"bytes"
	"encoding/json"
	"io"
	"log"
	"net/http"
)

type createWalletRequest struct {
	Name  string `json:"Name"`
	Email string `json:"Email"`
	Phone string `json:"Phone"`
}

type topUpRequest struct {
	From   string `json:"From"`
	Amount int    `json:"Amount"`
}

func TestCreateWallet(name string, email string, phone string) error {
	reqData := createWalletRequest{name, email, phone}
	reqJson, _ := json.Marshal(reqData)
	reqBody := bytes.NewBuffer(reqJson)
	resp, err := http.Post("http://127.0.0.1:9090/create-wallet", "application/json", reqBody)
	if err != nil {
		log.Println(err)
		return err
	}
	sBody, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
		return err
	}
	log.Println(string(sBody))
	return nil
}

func TestTopUp(address string, amount int) error {
	reqData := topUpRequest{address, amount}
	reqJson, _ := json.Marshal(reqData)
	reqBody := bytes.NewBuffer(reqJson)
	resp, err := http.Post("http://127.0.0.1:9090/top-up", "application/json", reqBody)
	if err != nil {
		log.Println(err)
		return err
	}
	sBody, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
		return err
	}
	log.Println(string(sBody))
	return nil
}

func main() {
	log.Print("Start Main")
	//err := TestCreateWallet("DatDT", "datdt@microtecweb.com", "0123456789")
	//if err != nil {
	//	log.Println(err)
	//}

	err := TestTopUp("0x482B90D5AAD8340A216b7463860891356986AD5e", 100)
	if err != nil {
		log.Println(err)
	}
}
