package main

import (
	"context"
	"fmt"
	"time"

	"github.com/chromedp/chromedp"
)

func main() {
	ctx, cancel := chromedp.NewContext(
		context.Background(),
	)
	defer cancel()

	// create a timeout
	ctx, cancel = context.WithTimeout(ctx, 15*time.Second)
	defer cancel()

	var test []map[string]string
	// ByID取得資料
	err := chromedp.Run(ctx,
		chromedp.Navigate("https://www.taoyuan-airport.com/chinese/flight_arrival"),
		chromedp.WaitVisible(`#print`, chromedp.ByID),
		chromedp.AttributesAll(`#print`, &test, chromedp.ByID),
	)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(test)
		fmt.Println("YA")
	}

	/* old version 下path取得資料
	err := chromedp.Run(ctx,
		chromedp.Navigate("https://www.taoyuan-airport.com/chinese/flight_arrival"),
		chromedp.WaitVisible(`#print`, chromedp.ByID),

			chromedp.WaitVisible(`div.show >
			div.ng-scope >
			div.ng-scope >
			div.container >
			div.row >
			div.col-lg-12 >
			div.main_pan_col >
			div.main_pan_body >
			div >
			table`),
			chromedp.OuterHTML(`
			div.show >
			div.ng-scope >
			div.ng-scope >
			div.container >
			div.row >
			div.col-lg-12 >
			div.main_pan_col >
			div.main_pan_body >
			div >
			table`, &test, chromedp.ByQuery),
	)
	*/
}

/*
	client := resty.New()
	resp, err := client.SetRetryWaitTime(5).R().Get("https://www.taoyuan-airport.com/")
	if err != nil {
		fmt.Println("Error")
	}

	if resp.StatusCode() != 200 {
		fmt.Println("Status Error")
	}

	fmt.Println(string(resp.Body()))
*/
