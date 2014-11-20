package main
 
import (
    "net/http"
    "html/template"
   
)
 
var mainTemplate, _ = template.ParseFiles("topics.html")
 
func mainPage(w http.ResponseWriter, r *http.Request) {
    mainTemplate.Execute(w, nil)
}
 
func main() {
    http.HandleFunc("/", mainPage)
    http.ListenAndServe(":8080", nil)
}