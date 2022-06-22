package main

import (
	"bytes"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"time"
)

type createWalletRequest struct {
	Name  string `json:"Name"`
	Email string `json:"Email"`
	Phone string `json:"Phone"`
}

type topUpRequest struct {
	From   string  `json:"From"`
	Amount float64 `json:"Amount"`
}

type transferRequest struct {
	From   string  `json:"From"`
	To     string  `json:"To"`
	Amount float64 `json:"Amount"`
}

type withdrawRequest struct {
	From   string  `json:"From"`
	Amount float64 `json:"Amount"`
}

type createWalletResponse struct {
	WalletAddress string `json:"WalletAddress"`
	Error         string `json:"Error"`
}

type getBalanceRequest struct {
	WalletAddress string `json:"WalletAddress"`
}

type getTransactionByAccountRequest struct {
	From   string `json:"From"`
	To     string `json:"To"`
	Type   int32  `json:"Type"`   // -1: all, 0: topup, 1: transfer, 2: withdraw
	Status int32  `json:"Status"` // -1: all, 0: success, 1: failed
}

type getTransactionByHashRequest struct {
	Hash string `json:"Hash"`
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

func TestTopUp(address string, amount float64) error {
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

func TestTopUpContract(address string, amount float64) error {
	reqData := topUpRequest{address, amount}
	reqJson, _ := json.Marshal(reqData)
	reqBody := bytes.NewBuffer(reqJson)
	resp, err := http.Post("http://127.0.0.1:9090/top-up-contract", "application/json", reqBody)
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

func TestTransfer(from string, to string, amount float64) error {
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

func TestTransferContract(from string, to string, amount float64) error {
	reqData := transferRequest{from, to, amount}
	reqJson, _ := json.Marshal(reqData)
	reqBody := bytes.NewBuffer(reqJson)
	resp, err := http.Post("http://127.0.0.1:9090/transfer-contract", "application/json", reqBody)
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

func TestGetBalance(walletAddress string) error {
	reqData := getBalanceRequest{walletAddress}
	reqJson, _ := json.Marshal(reqData)
	reqBody := bytes.NewBuffer(reqJson)
	resp, err := http.Post("http://127.0.0.1:9090/balance", "application/json", reqBody)
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

func TestWithdrawContract(walletAddress string, Amount float64) error {
	reqData := withdrawRequest{walletAddress, Amount}
	reqJson, _ := json.Marshal(reqData)
	reqBody := bytes.NewBuffer(reqJson)
	resp, err := http.Post("http://127.0.0.1:9090/withdraw-contract", "application/json", reqBody)
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

func TestTransactionByAccount(from string, to string, trxnType int32, trxnStatus int32) error {
	reqData := getTransactionByAccountRequest{from, to, trxnType, trxnStatus}
	reqJson, _ := json.Marshal(reqData)
	reqBody := bytes.NewBuffer(reqJson)
	resp, err := http.Post("http://127.0.0.1:9090/transaction-by-account", "application/json", reqBody)
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

func TestTransactionByHash(hash string) error {
	reqData := getTransactionByHashRequest{hash}
	reqJson, _ := json.Marshal(reqData)
	reqBody := bytes.NewBuffer(reqJson)
	resp, err := http.Post("http://127.0.0.1:9090/transaction-by-hash", "application/json", reqBody)
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
	//address1, err := TestCreateWallet("Dinh Van Vuong", "VuongDV@microtecweb.com", "0123456789")
	//if err != nil {
	//	log.Println(err)
	//}
	//
	//address2, err := TestCreateWallet("Tran Van Huy", "HuyTV@microtecweb.com", "0123456789")
	//if err != nil {
	//	log.Println(err)
	//}
	//
	////err = TestTopUpContract(address, 1)
	////if err != nil {
	////	log.Println(err)
	////}
	//
	//err = TestTransferContract(address1, address2, 2)
	//if err != nil {
	//	log.Println(err)
	//}

	//log.Print("Start Main")
	//address, err := TestCreateWallet("Dinh Van Vuong", "VuongDV@microtecweb.com", "0123456789")
	//if err != nil {
	//	log.Println(err)
	//}
	//err = TestTopUpContract(address2, 100)
	//if err != nil {
	//	log.Println(err)
	//}
	//err = TestTransferContract(address1, address2, 10)
	//if err != nil {
	//	log.Println(err)
	//}
	//time.Sleep(3000 * time.Millisecond)
	//err = TestGetBalance(address1)
	//if err != nil {
	//	log.Println(err)
	//}
	//err = TestGetBalance(address2)
	//if err != nil {
	//	log.Println(err)
	//}
	//err = TestTransactionByAccount(address1, "", -1, -1)
	//if err != nil {
	//	log.Println(err)
	//}
	//err = TestTransactionByAccount(address2, "", -1, -1)
	//if err != nil {
	//	log.Println(err)
	//}
	err := TestWithdrawContract("0x68917b6956201309DF637428302d7e5EBb1EDffd", 30)
	if err != nil {
		log.Print(err)
	}
	time.Sleep(20000 * time.Millisecond)
	err = TestGetBalance("0x8Af8364b963092F765DCaAC76813ea4619be43fD")
	if err != nil {
		log.Print(err)
	}
}
