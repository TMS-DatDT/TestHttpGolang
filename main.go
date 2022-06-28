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

type transactionHistoryRequest struct {
	From     string `json:"From"`
	To       string `json:"To"`
	Type     int32  `json:"Type"`   // -1: all, 0: topup, 1: transfer, 2: withdraw
	Status   int32  `json:"Status"` // -1: all, 0: success, 1: failed
	FromDate string `json:"FromDate"`
	ToDate   string `json:"ToDate"`
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
	link := "http://127.0.0.1:9090/balance?WalletAddress=" + walletAddress
	resp, err := http.Get(link)
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

func TestTransactionHistory(from string, to string, trxnType int32, trxnStatus int32, fromDate string, toDate string) error {
	reqData := transactionHistoryRequest{from, to, trxnType, trxnStatus, fromDate, toDate}
	reqJson, _ := json.Marshal(reqData)
	reqBody := bytes.NewBuffer(reqJson)
	resp, err := http.Post("http://127.0.0.1:9090/transaction-history", "application/json", reqBody)
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
	link := "http://127.0.0.1:9090/transaction-by-hash?Hash=" + hash
	resp, err := http.Get(link)
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
	// Test create wallet
	address, err := TestCreateWallet("Doan Trong Dat", "datdt@microtecweb.com", "0398881726")
	if err != nil {
		log.Println(err)
	}
	address1, err := TestCreateWallet("Dinh Van Vuong", "vuongdv@microtecweb.com", "123456789")
	if err != nil {
		log.Println(err)
	}
	log.Println("Address: ", address)
	//Test topup
	//err = TestTopUpContract(address, 100)
	//if err != nil {
	//	log.Println(err)
	//}
	//time.Sleep(1000 * time.Millisecond)
	// Test transfer
	//err = TestTransferContract(address, address1, 190)
	//if err != nil {
	//	log.Println(err)
	//}
	// Test withdraw
	//err = TestWithdrawContract(address1, 10)
	//if err != nil {
	//	log.Println(err)
	//}
	// Test getbalance
	//time.Sleep(5000 * time.Millisecond)
	//err = TestGetBalance(address)
	//if err != nil {
	//	log.Println(err)
	//}
	//Test transaction history
	_ = address1
	err = TestTransactionHistory(address, "", -1, -1, "2022-06-28 11:45:35", "2022-06-28 11:50:39")
	if err != nil {
		log.Println(err)
	}
}
