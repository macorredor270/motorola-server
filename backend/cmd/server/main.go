package main


import(
	"fmt"
	"net/http"

	"motorola-server/internal/api"
	"motorola-server/internal/handlers"
)


func main(){

	http.HandleFunc("/api/status",handlers.Status)

	http.HandleFunc("/ws",api.DeviceSocket)


	fmt.Println("Motorola Server :8080")

	http.ListenAndServe(":8080",nil)

}
