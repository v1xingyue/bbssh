package ssh

import (
	option "bbssh/options/ssh"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type PrivateKeyInfo struct {
	Pkey string
}

func LoadSshPriKey(user string, host string) string {
	url := fmt.Sprintf("http://%s/pkey/%s", option.KeyHost, user)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatal("NewRequest: ", err)
		return ""
	}
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal("Do: ", err)
		return ""
	}
	defer resp.Body.Close()
	c, err := ioutil.ReadAll(resp.Body)
	var item PrivateKeyInfo
	err = json.Unmarshal(c, &item)
	if err != nil {
		log.Fatal("Do: ", err)
		return ""
	}

	return item.Pkey
}
