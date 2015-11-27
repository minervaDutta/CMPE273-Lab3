package main

import (
    "net/http"
    "fmt"
    "os"
    "strings"
    "log"
    "io/ioutil"
    "strconv"
  )


 type Response struct{
 	Key  int 		    `json:"key"`
 	Value string    `json:"value"`
 }

var hash_val int
var port_no string


func main(){

//###################cases for different commandline arguments

	if(len(os.Args)==1){
		os.Exit(1)
	}
  //for /keys/id or /id

	if(len(os.Args)==2){

		url := fmt.Sprintf("http://localhost:3000/keys")
		res, err := http.Get(url)
		if err != nil {
			log.Fatal(err)
		}
		kval, err := ioutil.ReadAll(res.Body)
		res.Body.Close()
		fmt.Println("Key Value pairs are  : ", string(kval))

// gor get URL

	}else if os.Args[1]== "GET" {
		request_string := os.Args[2]
		key := strings.Split(request_string,"/")
		key_int,_ := strconv.Atoi(key[2])

		hash_val = key_int % 3
		if(hash_val == 0){
			port_no = "3000"
		}else if(hash_val == 1){
			port_no = "3001"
		}else {
			port_no = "3002"
		}
		url := fmt.Sprintf("http://localhost:%s/keys/%s",port_no,key[2])
		get_key, err := http.Get(url)
		if err != nil {
			log.Fatal(err)
		}
		data, err := ioutil.ReadAll(get_key.Body)
		get_key.Body.Close()
		fmt.Println("Key Value = ", string(data))
	}else

  {
  //in all other cases
		req_string := os.Args[2]
		key_put := strings.Split(req_string,"/")
		key_int,_ := strconv.Atoi(key_put[2])
		hash_val = key_int % 3
		if(hash_val == 0){
			port_no = "3000"
		}else if(hash_val == 1){
			port_no = "3001"
		}else {
			port_no = "3002"
		}
		put_url:= fmt.Sprintf("http://localhost:%s/keys/%s/%s",port_no,key_put[2],key_put[3])

		client := &http.Client{}
		req, _ := http.NewRequest("PUT", put_url, nil)
		resp, _ := client.Do(req)
		resp.Body.Close()
	}

}
