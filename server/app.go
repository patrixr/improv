package main

import (
    "net/http"
)

type RouteMap map[string]Route
type Route struct {
    GET         http.HandlerFunc
    POST        http.HandlerFunc
    Subroutes   RouteMap
}

//
// Main application structure
//
type ImprovWebApplication struct{
    Title           string
    Version         string
    Port            string
    Routes          RouteMap
    Storage         ImprovStorage
    Initialized     bool
}

func (this *ImprovWebApplication) Init() {
    if this.Initialized {
        return
    }
    this.Routes = RouteMap{
        "/version": Route{
            GET: HandleVersion,
        },
        "/api": Route{
            Subroutes: RouteMap{
                "/data": Route{
                    GET: HandleData,
                },
            },
        },
    }
    this.Storage.Open("improv_content")
    this.Initialized = true
}

func (this *ImprovWebApplication) Destroy() {
    if !this.Initialized {
        return;
    }
    this.Storage.Close()
    this.Initialized = false
}

var App = ImprovWebApplication{
    Title:          "Improv",
    Version:        "0.0.1",
    Port:           ":8000",
    Initialized:    false,
}