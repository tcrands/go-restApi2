package main

type DataResult struct {
    Id string `json:"id"`
    Name string `json:"name"`
    Content string `json:"content"`
}

type DataResults []DataResult 