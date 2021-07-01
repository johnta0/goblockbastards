package main

import (
	"os"
	"fmt"
	"log"
	"github.com/ChimeraCoder/anaconda"
)

func main() {
	CK := os.Getenv("TW_CK")
	CS := os.Getenv("TW_CS")
	AT := os.Getenv("TW_AT")
	AS := os.Getenv("TW_AS")
	if CK == "" || CS == "" || AT == "" || AS == "" {
		log.Fatal("Consumer key/secret and Access token/secret required")
	}

	anaconda.SetConsumerKey(CK)
	anaconda.SetConsumerSecret(CS)
}

func getBlockIds(api *anaconda.TwitterApi) {
	
}

