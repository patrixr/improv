package main

import (
    "path"
    "net/http"
    "log"
    "math/rand"
	"time"

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

    r := mux.NewRouter()

    createRoutes("/", App.Routes, r)

    log.Println("Listenning on port " + App.Port)
    http.ListenAndServe(App.Port, r)

    defer App.Destroy()
}   