package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"github.com/mrjones/oauth"
	"log"
)

type user struct {
	ID         int64  `json:"id"`
	Name       string `json:"name"`
	ScreenName string `json:"screen_name"`
}
type status struct {
	ID   int64  `json:"id"`
	Text string `json:"text"`
	User user   `json:"user"`
}

var consumerKey *string = flag.String(
	"consumerkey",
	"mEF22DxPk6cocNoc3lQQBoj55",
	"Consumer Key from Twitter. See: https://dev.twitter.com/apps/new")

var consumerSecret *string = flag.String(
	"consumersecret",
	"cGOq2NGmEqdwzVPPkQfMJuh6HEVFuVz5qFqBQJAteVuKC4ZQS9",
	"Consumer Secret from Twitter. See: https://dev.twitter.com/apps/new")

var accessToken *oauth.AccessToken
var selected_Command int

func main() {
	flag.Parse()
	loading()
	fmt.Println("please select next action")
	
	for {
	fmt.Println("(1)tweet (2)loadingtimeline (3)updatename ")
	selected_Command = 0
	fmt.Scanln(&selected_Command)
	switch selected_Command {
	case 1:
	 post_tweet()
	case 2:
	 get_timeline()
	case 3:
	 post_name()
		}}
	}

var c = oauth.NewConsumer(
	*consumerKey,
	*consumerSecret,
	oauth.ServiceProvider{
		RequestTokenUrl:   "https://api.twitter.com/oauth/request_token",
		AuthorizeTokenUrl: "https://api.twitter.com/oauth/authorize",
		AccessTokenUrl:    "https://api.twitter.com/oauth/access_token",
	})

func loading() {

	requestToken, url, err := c.GetRequestTokenAndUrl("oob")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("(1) Go to: " + url)
	fmt.Println("(3) Enter that verification code here: ")
	verificationCode := ""
	fmt.Scanln(&verificationCode)
	fmt.Println("loading successed.")
	accessToken, err = c.AuthorizeToken(requestToken, verificationCode)
	if err != nil {
		log.Fatal(err)
	}

}

func get_timeline() {

	response, err := c.Get(
		"https://api.twitter.com/1.1/statuses/home_timeline.json",
		map[string]string{},
		accessToken)
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()

	statuses := []status{}
	_ = json.NewDecoder(response.Body).Decode(&statuses)
	for _, s := range statuses {
		fmt.Printf("@%v: %v\n", s.User.ScreenName, s.Text)

	}
}

func post_name() {

	fmt.Println("write here your newname!")
	newname := ""
	fmt.Scanln(&newname)
	response, err := c.Post("https://api.twitter.com/1.1/account/update_profile.json",
		map[string]string{"name": newname}, accessToken)
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()

	fmt.Println("name posted!!")
	
}

func post_tweet() {

	fmt.Println("write here your newtweet!")
	newtweet := ""
	fmt.Scanln(&newtweet)
	response, err := c.Post("https://api.twitter.com/1.1/statuses/update.json",
		map[string]string{"status": newtweet}, accessToken)
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()

	fmt.Println("tweeted!!")

}
