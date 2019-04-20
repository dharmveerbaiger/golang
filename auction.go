package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func Max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

type AdPrice struct {
	Price int `json:"price"`
}

func main() {

	var client1 http.Client
	resp1, err := client1.Get("http://localhost:8080/bid")
	if err != nil {
		log.Fatal(err)
	}
	defer resp1.Body.Close()

	bodyByte1, err := ioutil.ReadAll(resp1.Body)
	if err != nil {
		log.Fatal(err)
	}
	adPrice1 := &AdPrice{}
	if err := json.Unmarshal(bodyByte1, adPrice1); err != nil {
		log.Fatal(err)
	}
	fmt.Println("Price of Ad1 is:", adPrice1.Price)

	var client2 http.Client
	resp2, err := client2.Get("http://localhost:8081/bid")
	if err != nil {
		log.Fatal(err)
	}
	defer resp2.Body.Close()

	bodyByte2, err := ioutil.ReadAll(resp2.Body)
	if err != nil {
		log.Fatal(err)
	}
	adPrice2 := &AdPrice{}
	if err := json.Unmarshal(bodyByte2, adPrice2); err != nil {
		log.Fatal(err)
	}
	fmt.Println("Price of Ad2 is:", adPrice2.Price)

	ans := Max(adPrice1.Price, adPrice2.Price)

	if ans == adPrice1.Price && ans == adPrice2.Price {
		fmt.Println("Both are same  :", ans)
	} else if ans == adPrice1.Price {
		fmt.Println("max bid = first bid :", ans)
	} else {
		fmt.Println("max bid = Second bid:", ans)
	}

}
