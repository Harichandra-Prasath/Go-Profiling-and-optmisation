package main

import (
	"encoding/json"
	"os"
)

func checkEntry(cred *Creds) (bool, error) {

	var valid_creds []*Creds

	data, _ := os.ReadFile("dummy_creds.json")
	err := json.Unmarshal(data, &valid_creds)
	if err != nil {
		return true, err
	}

	for _, v_cred := range valid_creds {
		if v_cred.Username == cred.Username {
			if v_cred.Password == cred.Password {
				return true, nil
			}
		}
	}

	return false, nil

}
