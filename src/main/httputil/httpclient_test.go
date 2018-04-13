package httputil

import "testing"


func TestClient_SendGetRequest(t *testing.T) {
	client:=Client{"http://localhost:8080/info/test","POST"}
	client.SendRequest(client.url,client.method)
}
