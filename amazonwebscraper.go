/*package main

import(
   "github.com/gocolly/colly"
   "encoding/json"
   //"log"
   "os"
   "fmt"
    //"strings"
)

var name string
type Fact struct{
    Name string `json:"name"`
    ImageURL string `json:"imageURL"`
    Description string `json:"description"`
    Price int `json:"price"`
    TotalReviews int `json:"totalReviews"`
}

func main(){
    allFacts:=make([]Fact,0)
    var imageurl,desc string
    var price,reviews int
    
    collector:=colly.NewCollector()
    collector.OnHTML("#imgTagWrapperId",func(element *colly.HTMLElement){
        //name=strings.Trim(element.Text," ")
        fmt.Println("https"+element.ChildText("img"))
        
    })
    fmt.Println(name)
    imageurl="img"
    desc="desc"
    price=12
    reviews=10
    fact:=Fact{
        Name:name,
        ImageURL:imageurl,
        Description:desc,
        Price:price,
        TotalReviews:reviews,
    }
    allFacts=append(allFacts,fact)
    collector.OnRequest(func(request *colly.Request){
        fmt.Println("Visiting",request.URL.String())
    })
    collector.Visit("https://www.amazon.com/PlayStation-4-Pro-1TB-Console/dp/B01LOP8EZC/")
    
    enc:=json.NewEncoder(os.Stdout)
    enc.SetIndent(""," ")
    enc.Encode(allFacts)
}*/
package main

import (
    //"fmt"
    b "AmazonProject/writeJSON"
    "log"
    "net/http"
    //"os"
    "strings"
    "github.com/PuerkitoBio/goquery"
    "encoding/json"
    //"io/ioutil"
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

var name,imageURL,totalReviews string
var description [] string
var price [] string

func main() {
    
    url:="https://www.amazon.com/CHERYLON-Portable-Cleaner-Interior-Cleaning/dp/B085TNBFS9/ref=sr_1_2?crid=2UAELRGKOUG1L&dchild=1&keywords=car+accessories&qid=1608897502&sprefix=office+chair%2Caps%2C767&sr=8-2"
    response, err := http.Get(url)
    if err != nil {
        log.Fatal(err)
    }
    defer response.Body.Close()

    
    document, err := goquery.NewDocumentFromReader(response.Body)
    if err != nil {
        log.Fatal("Error loading HTTP response body. ", err)
    }

    document.Find("#imgTagWrapperId").Each(func(index int, parent *goquery.Selection){
    parent.Find("img").Each(func(index int, element *goquery.Selection) {
        imgALT, existsALT := element.Attr("alt")
        imgURL, existsURL := element.Attr("data-a-dynamic-image")
        split:=strings.Split(imgURL,",")
        indexOfJPG:=strings.Index(split[0],".jpg")
        if existsALT && existsURL  {
            name=imgALT
            imageURL=split[0][2:(indexOfJPG+4)]
        }
    })})
    
    var reviewsFlag=1
    document.Find("#acrCustomerReviewText").Each(func(index int, element *goquery.Selection) {
        reviews:=element.Text()
        if reviewsFlag==1{
            totalReviews=reviews
            reviewsFlag=0
        }
        
    })
    
    priceFlag:=0
    
    document.Find("#priceblock_ourprice").Each(func(index int, element *goquery.Selection) {
        priceFlag=1
        //price:=element.Text()
        price=append(price,element.Text())
    })
    
    if priceFlag==0{
      document.Find("#edition_0_price").Each(func(index int, element *goquery.Selection) {

          str:=strings.ReplaceAll(element.Text(),"\n","")
          price=append(price,strings.Trim(str," "))
        
    })
        document.Find("#edition_1_price").Each(func(index int, element *goquery.Selection) {
        
       str:=strings.ReplaceAll(element.Text(),"\n","")        
            price=append(price,strings.Trim(str," "))
    }) 
        //fmt.Println(prices)
    }
    
      document.Find("#feature-bullets").Each(func(index int, feature *goquery.Selection) {
      
          feature.Find("li").Each(func(index int, list *goquery.Selection) {
              str:=strings.ReplaceAll(list.Text(),"\n","")
              description=append(description,strings.Trim(str," "))
    })    
    })
    
    
    fact:=Fact{
        Name:name,
        ImageURL:imageURL,
        Description:description,
        Price:price,
        TotalReviews:totalReviews,
    }
    
    
    allFacts:=make([]Fact,0)
    allFacts=append(allFacts,fact)
    
    
    
    product:=Product{
        Url:url,
        ProductDetails:fact,
    }
    productDetails:=make([]Product,0)
    productDetails=append(productDetails,product)
    //enc:=json.NewEncoder(os.Stdout)
    //enc.SetIndent(""," ")
    //enc.Encode(allFacts)
    s,_:=json.MarshalIndent(product,""," ")
    //log.Println(string(b))
    b.WriteJSON(string(s))
    
}

/*func writeJSON(data []Product){
    file,_:=json.MarshalIndent(data,""," ")
    //log.Println("Hello")
    f, err := os.OpenFile("AmazonProductDetails.json", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
    if err != nil {
        log.Println(err)
    }
    defer f.Close()
    //log.Println(file)
    //_=ioutil.WriteFile("AmazonProductDetails.json",file,0777)
    if _, err := f.Write(file); err != nil {
        log.Println(err)
    }
}*/
    

