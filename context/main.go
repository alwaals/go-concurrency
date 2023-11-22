package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

func main() {
	st := time.Now()
	cont := context.Background()

	ctx, cancelCtx := context.WithTimeout(cont, time.Millisecond*300)
	defer cancelCtx()

	ctx = context.WithValue(ctx, "User", "Swetha")
	httpReq, err := http.NewRequestWithContext(ctx, http.MethodGet, "https://www.google.com/maps/place/Highland+Luxury+Living/", nil)
	if err != nil {
		log.Fatal("Got error at http.NewRequest:", err.Error())
	}
	client := &http.Client{
		Transport: &http.Transport{
			MaxIdleConns:    10,
			MaxConnsPerHost: 20,
		},
	}

	httpRes, err := client.Do(httpReq)
	if err != nil {
		log.Fatal("Got error at client.Do:", err.Error())
	}
	defer httpRes.Body.Close()

	res, err := ioutil.ReadAll(httpRes.Body)
	if err != nil {
		log.Fatal("Got error at ioutil.ReadAll:", err.Error())
	}
	fmt.Println(string(res), "\n", time.Since(st),ctx.Value("User"))
}
