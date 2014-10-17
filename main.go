package trivette

import "github.com/google/go-github/github"

func main() {
	client := github.NewClient(nil)
	client.Activity.ListEventsPerformedByUser("user", false, nil)
	// credentials, err := os.Open(".credentials")
	// if nil != err {
	// 	panic(err)
	// }

}
