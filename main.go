package main
import (
"flag"
"log"
"net/http"
"strconv"
)
var port = flag.Int("port", 8080, "web service port")
func main() {
flag.Parse()
r := NewLoggedRouter()
err := http.ListenAndServe(":"+strconv.Itoa(*port), r)
if err != nil {
log.Fatal(err)
}
log.Fatal("The server returned without error.")
}
