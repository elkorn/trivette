package trivette_test

import (
	"encoding/json"
	"fmt"
	"os"
	"testing"

	"github.com/elkorn/trivette"
	"github.com/stretchr/testify/assert"
)

const dummyPath = ".credentials_dummy"
const dummyUser = "dummyUser"
const dummyToken = "dummyToken"

func writeDummyCredentials() {
	credentials := &trivette.Credentials{
		User:  "dummyUser",
		Token: "dummyToken",
	}

	file, err := os.OpenFile(dummyPath, os.O_CREATE|os.O_RDWR, 555)
	if nil != err {
		panic(err)
	}

	json, err := json.Marshal(credentials)
	if nil != err {
		panic(err)
	}

	fmt.Println(json)
	_, err = file.Write(json)
	if nil != err {
		panic(err)
	}
}

func removeDummyCredentials() {
	os.Remove(dummyPath)
}

func TestConfigLoader(t *testing.T) {
	writeDummyCredentials()
	credentials := trivette.ReadCredentials(dummyPath)
	assert.Equal(t, dummyUser, credentials.User)
	assert.Equal(t, dummyToken, credentials.Token)
	removeDummyCredentials()
}
