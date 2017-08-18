package main

import (
    "net/http"
)

type Route struct {
    Name        string
    Method      string
    Pattern     string
    HandlerFunc http.HandlerFunc
    AuthRequired bool 
}

type Routes []Route

var routes = Routes{
    Route{
        "HomePage",
        "GET",
        "/",
        homePage,
        false,
    },
    Route{
        "GetSomething",
        "GET",
        "/get",
        getSomething,
        true,
    },
    Route{
        "GetTokens",
        "GET",
        "/getToken",
        getToken,
        false,
    },
}