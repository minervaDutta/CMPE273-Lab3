package main

import (
  "net/http"
  "strconv"
  "encoding/json"
	"Assignments/lab3/httprouter"
	"strings"
)
//maps and structs

var server_1 map[int] string
var server_2 map[int] string
var server_3 map[int] string

type Response struct{
	Key  int 		    `json:"key"`
	Value string    `json:"value"`
}
//*******************************
func main(){

	server_1 = make(map[int] string)
  server_2 = make(map[int] string)
  server_3 = make(map[int] string)

//for port 3000
	go func(){
		mux_3000 := httprouter.New()
	    mux_3000.PUT("/keys/:id/:value",put)
	    mux_3000.GET("/keys/:id",get)
	    mux_3000.GET("/keys",getall)
	    server1 := http.Server{
	            Addr:"0.0.0.0:3000",
	            Handler:mux_3000,
    	}
    	server1.ListenAndServe()
	}()

//for port 3001
    go func(){
	    mux_3001 := httprouter.New()
	    mux_3001.PUT("/keys/:id/:value",put)
	    mux_3001.GET("/keys/:id",get)
	    mux_3001.GET("/keys",getall)
	    server2 := http.Server{
	            Addr:"0.0.0.0:3001",
	            Handler:mux_3001,
	    }
	    server2.ListenAndServe()
    }()

//port 3002
  	mux_3002 := httprouter.New()
    mux_3002.PUT("/keys/:id/:value",put)
    mux_3002.GET("/keys/:id",get)
    mux_3002.GET("/keys",getall)
    server3 := http.Server{
            Addr:        "0.0.0.0:3002",
            Handler: mux_3002,
    }
    server3.ListenAndServe()


}


func put(rw http.ResponseWriter, req *http.Request, p httprouter.Params){

	key := p.ByName("id")
	value := p.ByName("value")
  var port []string
	key_int, _ := strconv.Atoi(key)

  port = strings.Split(req.Host,":")
  if(port[1]=="3000"){
      server_1[key_int] = value

  } else if (port[1]=="3001"){
      server_2[key_int] = value

  } else{
      server_3[key_int] = value

    }

}

func get(rw http.ResponseWriter, req *http.Request, p httprouter.Params){

	key := p.ByName("id")
	key_int, _ := strconv.Atoi(key)
  var port []string
	var response Response
    port = strings.Split(req.Host,":")
  if(port[1]=="3000"){
      response.Key = key_int
      response.Value = server_1[key_int]

  } else if (port[1]=="3001"){
      response.Key = key_int
      response.Value = server_2[key_int]

  } else{
      response.Key = key_int
      response.Value = server_3[key_int]

    }
  	resp, err := json.Marshal(response)
  	if err != nil {
    	 http.Error(rw,"Error in server" , http.StatusInternalServerError)
     	return
  	}
  	rw.Header().Set("Content-Type", "application/json")
  	rw.Write(resp)
}


func getall(rw http.ResponseWriter, req *http.Request, p httprouter.Params){
	var response []Response
	var key_pair Response
  var port_num []string
  port_num = strings.Split(req.Host,":")
  if(port_num[1]=="3000"){
      for key, value := range server_1 {
      key_pair.Key = key
      key_pair.Value = value
       response = append(response, key_pair)
      }

  } else if (port_num[1]=="3001"){
      for key, value := range server_2 {
      key_pair.Key = key
      key_pair.Value = value
       response = append(response, key_pair)
      }

  } else{
      for key, value := range server_3 {
      key_pair.Key = key
      key_pair.Value = value
       response = append(response, key_pair)
      }

    }


  	resp, err := json.Marshal(response)
  	if err != nil {
    	 http.Error(rw,"Error in server" , http.StatusInternalServerError)
     	return
  	}
  	rw.Header().Set("Content-Type", "application/json")
  	rw.Write(resp)
}
