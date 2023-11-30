package ws

import (
	"encoding/json"
	"pttcat-go/db"

	"github.com/gorilla/websocket"
)

type MessageJson struct {
}

func sync(conn *websocket.Conn) {
	state := db.GetState()
	var stateJson map[string]interface{}
	error := json.Unmarshal([]byte(state), &stateJson)
	if error != nil {
		println("Error Unmarshal:", error)
		return
	}
	error = conn.WriteJSON(map[string]interface{}{"name": "sync", "state": stateJson})
	if error != nil {
		println("Error WriteJSON:", error)
		return
	}
}

func AddWsClient(conn *websocket.Conn) {
	for {
		_, message, err := conn.ReadMessage()
		if err != nil {
			println("Error ReadMessage:", err)
			return
		}

		println("Message:", string(message))

		var messageJson map[string]interface{}

		err = json.Unmarshal(message, &messageJson)

		if err != nil {
			println("Error Unmarshal:", err)
			return
		}

		if messageJson["name"] == "uploadIdb" {
			databaseJson, error := json.Marshal(messageJson["database"])
			if error != nil {
				println("Error Marshal:", error)
				return
			}
			db.SaveState(string(databaseJson))
		}

		if messageJson["name"] == "sync" {
			sync(conn)
		}
	}
}
