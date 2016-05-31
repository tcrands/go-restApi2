package main

import (
    "net/http"
    "time"
    "fmt"
    "github.com/dgrijalva/jwt-go"
    "github.com/auth0/go-jwt-middleware"
)

//set the secret
var signingKey = []byte("secret")

func getToken(w http.ResponseWriter, r *http.Request){
  
  userId := "Tom"
  token := generateToken(userId)

  w.Header().Set("jwt", token)
  w.WriteHeader(http.StatusOK)

  fmt.Fprintf(w, token)
}

func generateToken(userId string) string {
  //Create the token
  token := jwt.New(jwt.SigningMethodHS256)

  // Set token claims???????????
  token.Claims["admin"] = true
  token.Claims["name"] = userId
  token.Claims["exp"] = time.Now().Add(time.Hour * 7).Unix()

  //Sign the token with the secret
  tokenString, _ := token.SignedString(signingKey)

  return tokenString
}

var jwtMiddleware = jwtmiddleware.New(jwtmiddleware.Options{
  ValidationKeyGetter: func(token *jwt.Token) (interface{}, error) {
    return signingKey, nil
  },
  SigningMethod: jwt.SigningMethodHS256,
})

// decoded, err := base64.URLEncoding.DecodeString("wPCZGUt07zoiO7yhFAv2VCHyD2DoVRNQH_B29zm1zwZqWZ3A-qUJXCrR0B7cThrv")
//       if err != nil {
//         return nil, err
//       }
//       return decoded, nil