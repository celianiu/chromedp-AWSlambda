package main

import (
	"fmt"
	"context"
	"log"
	"time"
	// //"io/ioutil"
	// "os"
	// "strings"
	// "path/filepath"
	// "bytes"
	// "sync"

	"github.com/chromedp/chromedp"
	// "github.com/chromedp/cdproto/browser"
	// "github.com/chromedp/cdproto/cdp"
	// "github.com/chromedp/cdproto/network"
	// "github.com/aws/aws-sdk-go-v2/aws"
    // "github.com/aws/aws-sdk-go-v2/service/s3"
    // "github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-lambda-go/lambda"

)

type MyEvent struct {
	Name string `json:"name"`
}

// when login button available, enter the username and password and login 
func Login(usrnamebox,pwdbox,usrname,pwd,loginBtn string) chromedp.Tasks{
	return chromedp.Tasks{
		chromedp.SetValue(usrnamebox,usrname,chromedp.ByQuery),
		chromedp.SetValue(pwdbox,pwd,chromedp.ByQuery),
		chromedp.Sleep(5*time.Second),
		chromedp.Submit(loginBtn,chromedp.ByQuery,chromedp.NodeEnabled),
		chromedp.Sleep(10*time.Second),


	}
}


func Handler(ctx context.Context, name MyEvent) (string, error) {
	opts := []chromedp.ExecAllocatorOption{
		chromedp.Headless,
		chromedp.NoSandbox,
		chromedp.WindowSize(1200, 1000),
		// chromedp.Flag("disable-setuid-sandbox", true),
		chromedp.Flag("disable-dev-shm-usage", true),
		// chromedp.Flag("single-process", true),
		//chromedp.Flag("no-zygote", true),
		chromedp.Flag("use-gl", "angle"),
		chromedp.Flag("use-angle", "swiftshader"),
	}


	allocCtx, cancel := chromedp.NewExecAllocator(context.Background(), opts...)
	defer cancel()

	// also set up a custom logger
	taskCtx, cancel := chromedp.NewContext(allocCtx,
		chromedp.WithLogf(log.Printf),
		chromedp.WithDebugf(log.Printf),
		chromedp.WithErrorf(log.Printf),
	)
	defer cancel()

	const (
		url ="https://app.drivecentric.com/"
        	username = "xxx"
		password = "xxx"
    	)

	const (
		usernamebox = `#signInFormUsername`
		pwdbox = `#signInFormPassword`
		loginBtn = `form[name=cognitoSignInForm]`
	)

	//User login
	err := chromedp.Run(taskCtx,
		chromedp.Navigate(url),
	)
	fmt.Println("reach here")
	err = chromedp.Run(taskCtx,
		Login(usernamebox,pwdbox,username,password,loginBtn),
	)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Successful logged in!")
	
	return fmt.Sprintf("Event name is: %s\n",name.Name),nil


}
func main() {
	lambda.Start(Handler)
}



