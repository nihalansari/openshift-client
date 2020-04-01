package main
 
import (
    "fmt"
    "net/http"
)
 
func handler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Test Go on Openshift. Endpoint:  %s", r.URL.Path[1:])
}
 
func main() {
    
    fmt.Println("http Web Server is listening at port 9000")
    http.HandleFunc("/", handler)
    http.ListenAndServe(":9000", nil)
 
    
}