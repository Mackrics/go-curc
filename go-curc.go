package main

import (
    "fmt"
    "flag"
    "net/http"
    "os"
    "encoding/json"
    "strings"
    "regexp"
)

func main() {
  fromCur := flag.String("to", "dkk", "The currency converting from")
  toCur := flag.String("from", "sek", "The currency converting to")
  amountCur := flag.String("amount", "1", "The currency converting to")

  flag.Parse()

  r, _ := regexp.Compile("[a-zA-Z]")
  if r.MatchString(*amountCur) {
    fmt.Println("Error: Amount must be a digit.")
    os.Exit(1)
  }

  if len(*toCur) != 3 || len(*fromCur) != 3 {
    fmt.Println("To and from currency must be 3 characters")
    os.Exit(1)
  }

  if *toCur == *fromCur {
    fmt.Print("1")
  } else {
    url := "https://api.frankfurter.app/latest?amount=" + *amountCur + "&from=" + *fromCur + "&to=" + *toCur 

    resp, err := http.Get(url)
    if err != nil {
      fmt.Println("Invalid from or to currency")
      os.Exit(1)
    }
    var conversion map[string]interface{}
    json.NewDecoder(resp.Body).Decode(&conversion)
    fmt.Print(conversion["rates"].(map[string]interface{})[strings.ToUpper(*toCur)])
  }
}
