package main

import (
	"fmt"
	"time"

	"code.google.com/p/goauth2/oauth"

	"github.com/elkorn/trivette/app"
	"github.com/google/go-github/github"
)

func getDate(other time.Time) time.Time {
	return time.Date(other.Year(), other.Month(), other.Day(), 0, 0, 0, 0, time.UTC)
}

func printTodaysEvents(events []github.Event) {
	now := time.Now()
	now = getDate(now)
	for _, event := range events {
		fmt.Println(getDate(*(event.CreatedAt)).String())
		fmt.Println(getDate(now).String())
		if getDate(*(event.CreatedAt)).Equal(getDate(now)) {
			fmt.Printf("[%v] %v @%v\n", *(event.Type), *(event.Repo.Name))
		}
	}
}

func main() {
	credentials := trivette.ReadCredentials(".credentials")
	t := &oauth.Transport{
		Token: &oauth.Token{AccessToken: credentials.Token},
	}

	client := github.NewClient(t.Client())
	events, _, err := client.Activity.ListEventsPerformedByUser(credentials.User, true, nil)
	if nil != err {
		panic(err)
	}

	fmt.Println(events[0].CreatedAt.String())
	printTodaysEvents(events)
}
