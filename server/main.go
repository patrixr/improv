package main

import (
    "path"
    "net/http"
    "log"
    "math/rand"
	"time"
    "os"

    "github.com/gorilla/mux"
)

func createRoutes( root string, routes RouteMap, r *mux.Router ) {
    for uri, route := range routes {

        fullPath := path.Join(root, uri)

        createRoutes(fullPath, route.Subroutes, r)

        if route.GET != nil {
            log.Println("Creating route: GET " + fullPath)
            r.HandleFunc(fullPath, route.GET).Methods("GET")
        }
        if route.POST != nil {
            log.Println("Creating route: POST " + fullPath)
            r.HandleFunc(fullPath, route.POST).Methods("POST")
        }
    }
}

func main() {
    App.Init()

    rand.Seed(time.Now().UTC().UnixNano())

    cwd, _ := os.Getwd()
    static := path.Join(cwd, "documentation/pdf")

    r := mux.NewRouter()
    
    r.PathPrefix("/doc/pdf").Handler(http.StripPrefix("/doc/pdf", http.FileServer(http.Dir(static))))

    createRoutes("/", App.Routes, r)

    log.Println("Listenning on port " + App.Port)
    http.ListenAndServe(App.Port, r)

    defer App.Destroy()
}   