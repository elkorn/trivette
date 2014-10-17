package trivette_test

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"testing"

	"github.com/elkorn/trivette/app"
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

	json, err := json.Marshal(credentials)
	if nil != err {
		panic(err)
	}

	err = ioutil.WriteFile(dummyPath, json, 777)
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
