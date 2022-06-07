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

func main() {
	//log.Printf("Start Main")
	//reqData := createWalletRequest{"DatDT", "datdt@microtecweb.com", "0123456789"}
	//reqJson, _ := json.Marshal(reqData)
	//log.Printf(string(reqJson))

	//resp, err := http.Get("https://jsonplaceholder.typicode.com/posts/1")
	//if err != nil {
	//	log.Fatalln(err)
	//}
	//body, err := io.ReadAll(resp.Body)
	//if err != nil {
	//	log.Fatal(err)
	//}
	//log.Printf(string(body))

	log.Print("Start Main")
}
