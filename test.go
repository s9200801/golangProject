package main

import (
    "fmt"
    "net/http"
    "log"
    "majun"
    "html/template"

)

func helloHandler(w http.ResponseWriter, r *http.Request){
    fmt.Fprintf(w,"hello world")
    username:=r.FormValue("username")
    userpass:=r.FormValue("password")
    fmt.Fprintf(w,"your username:%s",username)
    fmt.Fprintf(w,"your password:%s",userpass)
}

func homeHandler(w http.ResponseWriter, r *http.Request){
}

func majunHandler(w http.ResponseWriter, r *http.Request){
    t,err:=template.ParseFiles("try.html")
    if err!=nil{
        return
    }
    t.Execute(w,nil)
    go majun.Play(w)
}

func onclick(w http.ResponseWriter, r *http.Request){
    if r.Method=="POST"{
        //fmt.Printf("test")
        w.Header().Set("Access-Control-Allow-Origin","*")
        w.Header().Add("Access-Control-Allow-Headers","Content-Type")
        w.Header().Set("Content-Type","application/json")
        //w.Write([]byte("OK"))
    }
    majun.OnClick(w)
}


func main(){
    http.HandleFunc("/",homeHandler)
    http.HandleFunc("/loadin",helloHandler)
    http.HandleFunc("/majun",majunHandler)
    http.HandleFunc("/onclick",onclick)
    err:=http.ListenAndServe(":8080",nil)
    if err != nil {
        log.Fatal("ListenAndServer", err)
    }
    
}