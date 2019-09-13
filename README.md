## News API (https://newsapi.org) SDK for Go

### Sample Usage

```
    package main

    import (
        "github.com/naganandakk/news-api-sdk-go/news"
        "fmt"
    )

    func main() {
        newsClient := news.NewClient("f251266458e947ba94e465c731e10a2a")

        // Sources
        sources, err := newsClient.Sources(map[string]string{})
        if err != nil {
            fmt.Printf("%v\n", err)
        } else {
            // First 10 sources
            fmt.Printf("\nFirst 10 sources\n")
            for index, source := range(sources[:10]) {
                fmt.Printf("\t%2d : %v\n", index + 1, source.Name)
            }
        }

        // Top Headlines
        topHeadlines, err := newsClient.TopHeadlines(map[string]string{
            "country": "in",
        })
        if err != nil {
            fmt.Printf("%v\n", err)
        } else {
            // Total headline count
            fmt.Printf("\n\nTotal headlines count %2d\n", topHeadlines.TotalResults)

            // First 10 headlines
            fmt.Printf("First 10 headlines\n")
            for index, article := range(topHeadlines.Articles[:10]) {
                fmt.Printf("\t%2d : %v\n", index + 1, article.Description)
            }
        }

        // Everything
        everything, err := newsClient.Everything(map[string]string{
            "domains": "google.com",
        })
        if err != nil {
            fmt.Printf("%v\n", err)
        } else {
            // Total articles count
            fmt.Printf("\n\nTotal articles count %2d\n", everything.TotalResults)

            // First 10 articles
            fmt.Printf("First 10 articles\n")
            for index, article := range(everything.Articles[:10]) {
                fmt.Printf("\t%2d : %v\n", index + 1, article.Description)
            }
        }
    }
```
