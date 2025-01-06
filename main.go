package main

import (
	"Assignment2_TomirisTapen/handlers"
	"Assignment2_TomirisTapen/server"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	srv := server.NewServer()

	// Добавляем обработчик для корневого маршрута
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "<h1>Welcome to the Basic Web Server</h1><p>Available endpoints:</p><ul>"+
			"<li>POST /data - Add data</li>"+
			"<li>GET /data - Retrieve all data</li>"+
			"<li>GET /stats - Get server stats</li>"+
			"<li>DELETE /data/{key} - Remove a key</li>"+
			"</ul>")
	})

	// Другие маршруты
	http.HandleFunc("/data", handlers.DataHandler(srv))
	http.HandleFunc("/data/", handlers.DeleteDataHandler(srv))
	http.HandleFunc("/stats", handlers.StatsHandler(srv))

	http.HandleFunc("/view-data", func(w http.ResponseWriter, r *http.Request) {
		srv.Mu.Lock()
		defer srv.Mu.Unlock()
	
		fmt.Fprintf(w, "<h1>Stored Data</h1><ul>")
		for key, value := range srv.Data {
			fmt.Fprintf(w, "<li>%s: %s</li>", key, value)
		}
		fmt.Fprintf(w, "</ul>")
	})
	
	// Фоновый воркер
	go srv.StartBackgroundWorker()

	serverInstance := &http.Server{Addr: ":8080"}
	go func() {
		log.Println("Server started on http://localhost:8080")
		if err := serverInstance.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Server error: %v", err)
		}
	}()

	// Завершаем работу сервера
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutting down server...")

	srv.Shutdown()
	if err := serverInstance.Close(); err != nil {
		log.Fatalf("Server close error: %v", err)
	}
	log.Println("Server shut down gracefully.")
}



