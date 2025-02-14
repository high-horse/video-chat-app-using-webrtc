package server

import (
	"encoding/json"
	"io"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

var AllRooms RoomMap

func CreateRoomRequestHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*") 
	roomId := AllRooms.Create()
	
	type resp struct {
		RoomId string `json:"room_id"`
	} 
	
	json.NewEncoder(w).Encode(resp{RoomId: roomId})
}

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool{
		return true
	},
}

type broadcastMsg struct {
	Message map[string]interface{}
	RoomID string
	Client *websocket.Conn
}

var broadcast = make(chan broadcastMsg)

func broadcaster() {
	for {
		msg := <- broadcast
		log.Println("msg recieved,", msg.Message)
		for _, client := range AllRooms.Map[msg.RoomID] {
			if client.Conn != msg.Client {
				err :=client.Conn.WriteJSON(msg.Message)
				if err != nil {
					log.Println(err)
					client.Conn.Close() 
				}
			}
		}
	}
}


func JoinRoomRequestHandler(w http.ResponseWriter, r *http.Request) {
	roomID, ok := r.URL.Query()["roomID"]
	if !ok {
		log.Println("room id missing in url")
		return
	}
	
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Fatal("web socket upgrade error ", err)
	}
	
	AllRooms.InsertIntoRoom(roomID[0], false, ws)
	
	go broadcaster()
	for {
		var msg broadcastMsg
		err := ws.ReadJSON(&msg.Message)
		if err != nil {
			if err != io.EOF {
				log.Println(err)
			}
			return
		}
		msg.Client = ws
		msg.RoomID = roomID[0]
		
		broadcast <- msg
	}
	
}