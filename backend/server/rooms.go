package server

import (
	"log"
	"math/rand"
	"sync"
	"time"

	"github.com/gorilla/websocket"
)

type Participants struct {
	Host bool
	Conn *websocket.Conn
}

type RoomMap struct {
	Mutex sync.RWMutex
	Map map[string][]Participants
} 

func (r *RoomMap) Init() {
	r.Map = make(map[string][]Participants)
}

func (r *RoomMap) Get(roomId string) []Participants {
	r.Mutex.RLock()
	defer r.Mutex.RUnlock()
	
	return r.Map["roomId"]
}

func (r *RoomMap) Create() string {
	// generate a unique id and insert into hashmap
	r.Mutex.Lock()
	defer r.Mutex.Unlock()
	
	rand.NewSource(time.Now().UnixNano())
	var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")
	b := make([]rune, 8)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	
	roomID := string(b)
	r.Map[roomID] = []Participants {}
	
	return roomID
}

func (r *RoomMap) InsertIntoRoom (roomID string, host bool, conn *websocket.Conn) {
	r.Mutex.Lock()
	defer r.Mutex.Unlock()
	
	p := Participants{host, conn}
	
	log.Println("inserting into room  with roomID: ", roomID)
	r.Map[roomID] = append(r.Map[roomID], p)
}

func (r *RoomMap) Delete(roomId string) {
	r.Mutex.Lock()
	defer r.Mutex.Unlock()
	
	delete(r.Map, roomId)
}
