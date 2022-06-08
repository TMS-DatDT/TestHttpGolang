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
	Amount int64  `json:"Amount"`
}

type transferRequest struct {
	From   string `json:"From"`
	To     string `json:"To"`
	Amount int64  `json:"Amount"`
}

type createWalletResponse struct {
	WalletAddress string `json:"WalletAddress"`
	Error         string `json:"Error"`
}

func TestCreateWallet(name string, email string, phone string) (string, error) {
	address := ""
	reqData := createWalletRequest{name, email, phone}
	reqJson, _ := json.Marshal(reqData)
	reqBody := bytes.NewBuffer(reqJson)
	resp, err := http.Post("http://127.0.0.1:9090/create-wallet", "application/json", reqBody)
	if err != nil {
		log.Println(err)
		return address, err
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
		return address, err
	}
	sBody := string(body)
	log.Println(sBody)

	var respJson createWalletResponse
	json.Unmarshal(body, &respJson)
	address = respJson.WalletAddress
	return address, nil
}

func TestTopUp(address string, amount int64) error {
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

func TestTransfer(from string, to string, amount int64) error {
	reqData := transferRequest{from, to, amount}
	reqJson, _ := json.Marshal(reqData)
	reqBody := bytes.NewBuffer(reqJson)
	resp, err := http.Post("http://127.0.0.1:9090/transfer", "application/json", reqBody)
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
	//log.Print("Start Main")
	address, err := TestCreateWallet("Tran Van Huy", "huytv@microtecweb.com", "0123456789")
	if err != nil {
		log.Println(err)
	}

	//err = TestTopUp(address, 100)
	//if err != nil {
	//	log.Println(err)
	//}

	err = TestTransfer(address, "0x8Af8364b963092F765DCaAC76813ea4619be43fD", 100)
	if err != nil {
		log.Println(err)
	}
}
