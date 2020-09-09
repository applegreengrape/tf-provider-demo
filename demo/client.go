package demo

import (
	"io/ioutil"
	"net/http"
	"bytes"
	"encoding/json"
	"log"
)

type payload struct {
	User string `json:"user"`
}

type update struct {
	Old string `json:"old"`
	New string `json:"new"`
}

type response struct {
	ID    string `json:"id"`
	Stats string `json:"stats"`
	User  string `json:"user"`
}
// ClientCreate ..
func ClientCreate(userName string) error{
	data := payload{
		User: userName, 
	}
	payloadBytes, err := json.Marshal(data)
	if err != nil {
		log.Fatal(err)
	}

	req, err := http.NewRequest("POST", "http://127.0.0.1:5000/create", bytes.NewReader(payloadBytes))
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	return nil
}

// ClientUpdate ..
func ClientUpdate(old, new string) error{
	data := update{
		Old: old, 
		New: new,
	}
	payloadBytes, err := json.Marshal(data)
	if err != nil {
		log.Fatal(err)
	}

	req, err := http.NewRequest("POST", "http://127.0.0.1:5000/update", bytes.NewReader(payloadBytes))
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	return nil 
}

// ClientDelete ..
func ClientDelete(userName string) error{
	data := payload{
		User: userName, 
	}
	payloadBytes, err := json.Marshal(data)
	if err != nil {
		log.Fatal(err)
	}

	req, err := http.NewRequest("POST", "http://127.0.0.1:5000/delete", bytes.NewReader(payloadBytes))
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	return nil
}

// ClientGet ..
func ClientGet(userName string) (id, stats string){
	data := payload{
		User: userName, 
	}
	payloadBytes, err := json.Marshal(data)
	if err != nil {
		log.Fatal(err)
	}

	req, err := http.NewRequest("POST", "http://127.0.0.1:5000/user", bytes.NewReader(payloadBytes))
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	
	var record response
	if err := json.Unmarshal(body, &record); err != nil {
		log.Println(err)
	}

	return record.ID, record.Stats
}