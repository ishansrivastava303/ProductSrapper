package writeJSON

import (
    //"fmt"
    "log"
    //"net/http"
    //"os"
    //"strings"
    //"github.com/PuerkitoBio/goquery"
    //"encoding/json"
    "io/ioutil"
    //"github.com/gocolly/colly"
)

type Fact struct{
    Name string `json:"name"`
    ImageURL string `json:"imageURL"`
    Description []string `json:"description"`
    Price []string `json:"price"`
    TotalReviews string `json:"totalReviews"`
}


type Product struct{
    Url string `json:"Url"`
    ProductDetails Fact `json:"ProductDetails"` 
}

func WriteJSON(data string){
    /*file,err:=json.MarshalIndent(data,""," ")
    
    if err!=nil{
        log.Println("Unable to create JSON file")
        return
    }*/
    _=ioutil.WriteFile("/home/ishu/Desktop/AmazonProject/AmazonProductDetails.json",[]byte(data),0644)
    log.Println(data)
}
