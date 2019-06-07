package main
import (
  "fmt"
  "net/http"
  "os"
  "context"
  "time"
  "log"
)
func homePage(w http.ResponseWriter, r *http.Request) {
  //fmt.Fprint(w, "Welcome to the HomePage!")
  exHost := os.Getenv("HOST")
  exPort := os.Getenv("PORT")
  exUri := os.Getenv("URI")
  uri := exHost + ":" + exPort + exUri
  req, err := http.NewRequest("GET", uri, nil)
//  fmt.Fprint(w, uri)
  fmt.Println(w,uri)
        if err != nil {
                log.Fatalf("http.NewRequest() failed with '%s'\n", err)
        }

        // Timeout = 100 ms
        ctx, _ := context.WithTimeout(context.Background(), time.Millisecond*100990000)
        req = req.WithContext(ctx)

        resp, err := http.DefaultClient.Do(req)
        if err != nil {
                log.Fatalf("http.DefaultClient.Do() failed with:\n'%s'\n", err)
        }
        fmt.Println(resp)
        fmt.Fprint(w, resp.StatusCode)
        defer resp.Body.Close()


}
func handleRequest() {
  http.HandleFunc("/", homePage)
  http.ListenAndServe(":1323", nil)
}
func main() {
  handleRequest() // ต้องนำ handleRequest มาใส่ใน main ด้วยนะครับ
}
