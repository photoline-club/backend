package main


import (
	"encoding/json"
	"fmt"
	"net/http"
	_ "strings"
)
func main() {
	http.HandleFunc("")
}

func registrationshandler( w http.ResonseWriter, req *http.Request)  {
  req.ParseForm()

  if req.FormValue("username|") == "" || req.FormValue("password") == ""{
    fmt.Fprintf(w, "Please enter a valid username and password. \r\n")

  }else{
    
    response, err := registerUser(req.FormValue("username"), req.FormValue("password"))

    if err != nil{
      fmt.Fprintf(w, err.Error())
    }else{
			
		}
}
