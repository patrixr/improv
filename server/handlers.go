package main

import (
    "net/http"
    "encoding/json"
    "log"
    "strconv"
)

func HandleVersion(w http.ResponseWriter, r *http.Request) {
    w.Write([]byte( App.Version ))
}

func HandleData(w http.ResponseWriter, r *http.Request) {

    count := 10
    from  := 0

    tmp, err := strconv.Atoi( r.URL.Query().Get("count") ) 
    if err == nil {
        count = tmp
    }

    tmp, err = strconv.Atoi( r.URL.Query().Get("from") ) 
    if err == nil {
        from = tmp
    }

    chunk := App.Storage.Read(from, count)
    str, err := json.Marshal(chunk)
    if (err != nil) {
        log.Panic("JSON Marshal error")
    }
    
    w.Write([]byte( str ))
}