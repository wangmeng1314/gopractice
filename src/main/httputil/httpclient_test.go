package httputil

import (
	"testing"
	"fmt"
	"sync"
	//"time"
)

func TestClient_SendGetRequest(t *testing.T) {
	gr := sync.WaitGroup{}
	gr.Add(1000)
	client := &Client{"http://localhost:8080", "GET"}
	for i := 0; i < 1000; i++ {
		go func() {
			for a := 0; a < 10; a++ {
				client.SendRequest(client.url, client.method)
			}
			fmt.Println(i)
			gr.Done()
		}()
	}
	gr.Wait()
}

func TestClient_SendGetRequest2(t *testing.T) {
	client := &Client{"http://localhost:8080", "GET"}
	for i := 0; i < 100; i++ {
		client.SendRequest(client.url,client.method)
		fmt.Println(i)
	}
}
