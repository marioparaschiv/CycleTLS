package main

import (
	"encoding/json"
	"flag"
	"github.com/gorilla/websocket"
	// "io/ioutil"
	// "strings"
	"log"
	// "net/http"
	"strconv"
	"time"
	"net/url"
	"os"
	"fmt"

)

type myTLSRequest struct {
	RequestID string `json:"requestId"`
	Options   struct {
		URL     string            `json:"url"`
		Method  string            `json:"method"`
		Headers map[string]string `json:"headers"`
		Body    string            `json:"body"`
		Ja3     string            `json:"ja3"`
		UserAgent     string       `json:"userAgent"`
		ID     int            		`json:"id"`
		Proxy   string            `json:"proxy"`
	} `json:"options"`
}





type response struct {
	Status  int
	Body    string
}

type myTLSResponse struct {
	RequestID string
	Response  response
}

func getWebsocketAddr() string {
	port, exists := os.LookupEnv("WS_PORT")
	fmt.Println(port)
	var addr *string

	if exists {
		addr = flag.String("addr", "localhost:"+port, "http service address")
	} else {
		addr = flag.String("addr", "localhost:9112", "http service address")
	}

	u := url.URL{Scheme: "ws", Host: *addr, Path: "/"}

	return u.String()
}

// /////////////////////
func process(job []byte, i int, link chan<- []byte) {

	message := job
	mytlsrequest := new(myTLSRequest)
	e := json.Unmarshal(message, &mytlsrequest)
	if e != nil {
		log.Print(e)
	}


	s := strconv.Itoa(mytlsrequest.Options.ID)

	if  mytlsrequest.Options.ID == 2 {
		time.Sleep(4 *  time.Second)
		s = string("yaga")
	} else {
		time.Sleep(1 *  time.Second)
	}
	Response := response{200, s}

	reply := myTLSResponse{mytlsrequest.RequestID, Response}

	data, err := json.Marshal(reply)
	if err != nil {
		log.Print(mytlsrequest.RequestID + "Request_Id_On_The_Left" + err.Error())
		
	}
	
	m[s] = data
	
	
	

	
	
}
func consumer() {
	fmt.Println("called")
	for {
		for key, b := range m {
		
		
		
			err := greeting.WriteMessage(websocket.TextMessage, b)
			if err != nil {
				log.Print("err")		
			}
			fmt.Println(key)
			delete(m,key)
			
		}

	}
	

}

var greeting *websocket.Conn
func worker(jobChan <-chan  []byte, i int, link chan<- []byte) {
	for job := range jobChan {
		process(job,i, link)
	}
}
var m = map[string][]byte{}

func main() {
	flag.Parse()
	log.SetFlags(0)

	websocketAddress := getWebsocketAddr()

	c, _, err := websocket.DefaultDialer.Dial(websocketAddress, nil)
	if err != nil {
		log.Print(err)
		return
	}
	greeting = c

	workerCount := 2
	// make a channel with a capacity of 100.
	jobChan := make(chan []byte, 100) // Or jobChan := make(chan int)
	// done := make(chan bool)
	link := make(chan []byte)
	// start the worker
	for i:=0; i<workerCount; i++ {
		go worker(jobChan, i, link)
	}
	
	go consumer()

	for {
		_, message, err := c.ReadMessage()
		if err != nil {
			log.Print(err)
			continue
		}
		 
		if message != nil {
			jobChan <- message
		}
		
		
	
	}
}





// func main() {
// 	flag.Parse()
// 	log.SetFlags(0)

// 	websocketAddress := getWebsocketAddr()

// 	c, _, err := websocket.DefaultDialer.Dial(websocketAddress, nil)
// 	if err != nil {
// 		log.Print(err)
// 		return
// 	}

// 	for {
// 		_, message, err := c.ReadMessage()
// 		if err != nil {
// 			log.Print(err)
// 			continue
// 		}


		
// 		mytlsrequest := new(myTLSRequest)
// 		e := json.Unmarshal(message, &mytlsrequest)
// 		if e != nil {
// 			log.Print(e)
// 		}


// 		s := strconv.Itoa(mytlsrequest.Options.ID)

// 		if  mytlsrequest.Options.ID == 2 {
// 			time.Sleep(4 *  time.Second)
// 			s = string("yaga")
// 		}
// 		Response := response{200, s}

// 		reply := myTLSResponse{mytlsrequest.RequestID, Response}

// 		data, err := json.Marshal(reply)
// 		if err != nil {
// 			log.Print(mytlsrequest.RequestID + "Request_Id_On_The_Left" + err.Error())
			
// 		}
		
// 		fmt.Println(reply)
// 		err = c.WriteMessage(websocket.TextMessage, data)
// 		if err != nil {
// 			log.Print(mytlsrequest.RequestID + "Request_Id_On_The_Left" + err.Error())
			
// 		}


	
// 	}
// }


