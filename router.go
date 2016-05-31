package main

import (
    "net/http"
    "github.com/gorilla/mux"
    // "github.com/auth0/go-jwt-middleware"
)

func NewRouter() *mux.Router {

    router := mux.NewRouter().StrictSlash(true)
    for _, route := range routes {

        var handler http.Handler

        handler = route.HandlerFunc
        handler = Logger(handler, route.Name)
        if route.AuthRequired == true {
            handler = jwtMiddleware.Handler(handler)
        }
        
        router.
            Methods(route.Method).
            Path(route.Pattern).
            Name(route.Name).
            Handler(handler)
    }

    return router
}