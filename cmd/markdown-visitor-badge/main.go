package main
 
import (
  "fmt"
  "net/http"
  "github.com/MustansirZia/markdown-visitor-badge/api"
)
 
func main() {
	http.HandleFunc("/", api.Handler)
	fmt.Println("Server is listening on port 8080...")
  	http.ListenAndServe(":8080", nil)
}