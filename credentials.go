package trivette

import (
	"encoding/json"
	"io/ioutil"
)

type Credentials struct {
	User, Token string
}

func e(err error) {
	if nil != err {
		panic(err)
	}
}

func ReadCredentials(path string) *Credentials {
	raw, err := ioutil.ReadFile(path)
	var credentials Credentials
	err = json.Unmarshal(raw, &credentials)
	e(err)
	return &credentials
}
