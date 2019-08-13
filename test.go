package main

import (
    "fmt"
    "net/http"
    "log"
    "majun"
)

func helloHandler(w http.ResponseWriter, r *http.Request){
    fmt.Fprintf(w,"hello world")
    username:=r.FormValue("username")
    userpass:=r.FormValue("password")
    fmt.Fprintf(w,"your username:%s",username)
    fmt.Fprintf(w,"your password:%s",userpass)
}

func homeHandler(w http.ResponseWriter, r *http.Request){
    fmt.Fprintf(w,"hello world")
}

func majunHandler(w http.ResponseWriter, r *http.Request){
    majun.Play(w)
}

func onclick(w http.ResponseWriter, r *http.Request){
    majun.OnClick()
}


func main(){
    http.HandleFunc("/",homeHandler)
    http.HandleFunc("/loadin",helloHandler)
    http.HandleFunc("/majun",majunHandler)
    http.HandleFunc("/onclick",onclick)
    err:=http.ListenAndServe(":7777",nil)
    if err != nil {
        log.Fatal("ListenAndServer", err)
    }
    
}