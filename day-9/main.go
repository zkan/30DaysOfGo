package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"fmt"
)

// Response:
// {
//   "qotd_date": "2025-11-25T00:00:00.000+00:00",
//   "quote": {
//     "id": 62451,
//     "dialogue": false,
//     "source": "Becoming Warren Buffett",
//     "private": false,
//     "tags": [],
//     "url": "https://favqs.com/quotes/warren-buffett/62451-if-you-re-emo-",
//     "favorites_count": 2,
//     "upvotes_count": 0,
//     "downvotes_count": 1,
//     "author": "Warren Buffett",
//     "author_permalink": "warren-buffett",
//     "body": "If you're emotional about investment, you're not going to do well. You may have all these feelings about the stock. The stock has no feelings about you."
//   }
// }

type Response struct {
	QotdDate string `json:"qotd_date"`
	Quote Quote `json:"quote"`
}

type Quote struct {
	Author string `json:"author"`
	Body string `json:"body"`
}

func main(){
	response, err := http.Get("https://favqs.com/api/qotd")
	if err != nil {
		fmt.Println("We can't retrieve your quote. You have internet connection?")
	} else {
		rawQuotes, _ := ioutil.ReadAll(response.Body)

		var quote Response
		json.Unmarshal(rawQuotes, &quote)

		fmt.Println(quote.QotdDate)
		fmt.Println(quote.Quote.Body, "by", quote.Quote.Author)
	}
}
