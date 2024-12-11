package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

// Upgrader untuk meng-upgrade koneksi HTTP ke WebSocket
var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		// Izinkan semua origin (ubah sesuai kebutuhan)
		return true
	},
}

func handleWebSocket(c *gin.Context) {
	// Upgrade koneksi HTTP ke WebSocket
	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		fmt.Println("Failed to upgrade connection:", err)
		return
	}
	defer conn.Close()

	fmt.Println("Client connected")

	// Kirim data real-time ke klien
	for {
		// Simulasi data acak
		data := fmt.Sprintf(`{"timestamp": "%s", "value": %d}`, time.Now().Format(time.RFC3339), time.Now().Second())

		// Kirim data ke klien
		err := conn.WriteMessage(websocket.TextMessage, []byte(data))
		if err != nil {
			fmt.Println("Error writing message:", err)
			break
		}

		// Delay sebelum mengirim data berikutnya
		time.Sleep(1 * time.Second)
	}
}

func main() {
	r := gin.Default()

	// Route untuk WebSocket
	r.GET("/ws", handleWebSocket)

	// Jalankan server
	r.Run(":8080")
}
