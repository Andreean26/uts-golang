package main

import (
    "log"
    "net/http"
    "citizen"
)

func main() {
    http.HandleFunc("/find-person", citizen.FindPerson)
    http.HandleFunc("/get-people", citizen.GetPeople)
    log.Fatal(http.ListenAndServe(":8080", nil))
}
