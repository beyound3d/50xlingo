package main

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"errors"
)

type book struct{
	ID string         `json:"id"`
	Title string      `json:"title"`
	Author string     `json:"author"`
    Quantity int      `json:"quantity"`
}

var books = []book{
	{ID:"1", Title:"In search of Lost Time", Author: "Marcel Proust", Quantity:2},
	{ID:"2", Title:"the great gatby", Author:"F. Scott Fitzgenerald", Quantity:3},
	{ID:"3", Title:"War and Peace", Author:"Leo TolStoy", Quantity:6},
}

func getBooks(c *gin.Context){
	
}