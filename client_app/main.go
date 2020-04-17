package main


import (
	"fmt"
	"github.com/dmskdlghs213/golang_api_client/client_app/client"
)

func main() {
	result, err := client.TestClient()

	if err != nil {
		fmt.Println(err)
		fmt.Println("conection error")
	}

	fmt.Println(result)
}