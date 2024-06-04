package main

import (
	"fmt"

	"github.com/go-rod/rod"
	"github.com/go-rod/rod/lib/launcher"
)

func main() {
	u := launcher.New().
		UserDataDir("./cache").
		Headless(false).
		MustLaunch()
	browser := rod.New().ControlURL(u).MustConnect()
	page := browser.MustPage("https://x.com/i/flow/login")

	page.MustWaitLoad()

	loginTwitter(page)

	waitExit()

}

func loginTwitter(page *rod.Page) {

	email := "xcoder69.anonymous@gmail.com"
	password := "heaven123HELL"
	uname := "xcoder69"

	page.MustElement("input[name='text']").Input(email)
	page.MustElement("button[class='css-175oi2r r-sdzlij r-1phboty r-rs99b7 r-lrvibr r-ywje51 r-184id4b r-13qz1uu r-2yi16 r-1qi8awa r-3pj75a r-1loqt21 r-o7ynqc r-6416eg r-1ny4l3l']").MustClick()
	page.WaitLoad()

	page.MustElement("input[name='text']").Input(uname)
	page.MustElement("button[data-testid='ocfEnterTextNextButton']").MustClick()
	page.WaitLoad()

	page.MustElement("input[name='password']").Input(password)
	page.MustElement("button[data-testid='LoginForm_Login_Button']").MustClick()
	page.WaitLoad()

	scrapeTweets(page)
	fmt.Println("Logged in!")
}

func scrapeTweets(page *rod.Page) {
	page.MustWaitLoad()
}

func waitExit() {
	fmt.Println("Press Enter to exit...")
	fmt.Scanln()
}
