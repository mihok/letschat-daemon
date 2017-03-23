package rest

import (
	"encoding/json"
	"fmt"
	"github.com/julienschmidt/httprouter"
	"github.com/minimalchat/mnml-daemon/chat"
	"github.com/minimalchat/mnml-daemon/store"
	"github.com/minimalchat/mnml-daemon/utils"
	"log"
	"net/http"
)

// Chats

// ReadChats performs a GET /api/chats
func ReadChats(db *store.InMemory) func(resp http.ResponseWriter, req *http.Request, params httprouter.Params) {
	return func(resp http.ResponseWriter, req *http.Request, params httprouter.Params) {
		chats, _ := db.Search("chat.")
		result := make(map[string]interface{})

		result["chats"] = chats

		log.Println(DEBUG, "chat:", "Reading chats", fmt.Sprintf("(%d records)", len(chats)))

		resp.Header().Set("Content-Type", "application/json; charset=UTF-8")
		resp.WriteHeader(http.StatusOK)
		if err := json.NewEncoder(resp).Encode(result); err != nil {
			panic(err)
		}
	}
}

// Chat

// ReadChat performs GET at /api/chat/:id
func ReadChat(db *store.InMemory) func(resp http.ResponseWriter, req *http.Request, params httprouter.Params) {
	return func(resp http.ResponseWriter, req *http.Request, params httprouter.Params) {
		ch, _ := db.Get(fmt.Sprintf("chat.%s", params.ByName("id")))

		log.Println(DEBUG, "chat:", "Reading chat", params.ByName("id"))

		resp.Header().Set("Content-Type", "application/json; charset=UTF-8")
		if ch != nil {
			resp.WriteHeader(http.StatusOK)
			if err := json.NewEncoder(resp).Encode(ch); err != nil {
				panic(err)
			}
		} else {
			resp.WriteHeader(http.StatusNotFound)

			fmt.Fprintf(resp, "Not Found")
		}
	}
}

// CreateOrUpdateChat POST / PUT / PATCH /api/chat/:id?
func CreateOrUpdateChat(db *store.InMemory) func(resp http.ResponseWriter, req *http.Request, params httprouter.Params) {
	return NotImplemented
}

// DeleteChat DELETE /api/chat/:id?
func DeleteChat(db *store.InMemory) func(resp http.ResponseWriter, req *http.Request, params httprouter.Params) {
	return NotImplemented
}

// Chat Messages

// ReadMessages GET /api/chat/:id/messages
func ReadMessages(db *store.InMemory) func(resp http.ResponseWriter, req *http.Request, params httprouter.Params) {
	return func(resp http.ResponseWriter, req *http.Request, params httprouter.Params) {
		messages, _ := db.Search(fmt.Sprintf("message.%s-", params.ByName("id")))
		result := make(map[string]interface{})

		// result["messages"] = messages
		result["messages"] = utils.MakeDummy(20)

		log.Println(DEBUG, "message:", "Reading messages", fmt.Sprintf("(%d records)", len(messages)))

		resp.Header().Set("Content-Type", "application/json; charset=UTF-8")
		resp.WriteHeader(http.StatusOK)
		if err := json.NewEncoder(resp).Encode(result); err != nil {
			panic(err)
		}
	}
}

// Chat Message

// ReadMessage GET /api/chat/:id/message/:mid
func ReadMessage(db *store.InMemory) func(resp http.ResponseWriter, req *http.Request, params httprouter.Params) {
	return NotImplemented
}

// CreateMessage POST / PUT /api/chat/:id/message
func CreateMessage(db *store.InMemory) func(resp http.ResponseWriter, req *http.Request, params httprouter.Params) {
	return func(resp http.ResponseWriter, req *http.Request, params httprouter.Params) {
		var msg *chat.Message

		id := params.ByName("id")
		decoder := json.NewDecoder(req.Body)

		resp.Header().Set("Content-Type", "application/json; charset=UTF-8")

		if err := decoder.Decode(&msg); err != nil {
			log.Println(DEBUG, "message:", "Bad Request", err)
			resp.WriteHeader(http.StatusBadRequest)

			fmt.Fprintf(resp, "Bad Request")
			return
		}

		if id == "" {
			log.Println(DEBUG, "message:", "Bad Request ID", id)
			resp.WriteHeader(http.StatusBadRequest)

			fmt.Fprintf(resp, "Bad Request")
			return
		}

		result, _ := db.Get(fmt.Sprintf("chat.%s", id))

		if result == nil {
			log.Println(DEBUG, "message:", "Unknown Chat ID", id, result)
			resp.WriteHeader(http.StatusNotFound)

			fmt.Fprintf(resp, "Not Found")
			return
		}

		if ch, ok := result.(chat.Chat); ok {
			log.Println(DEBUG, "operator:", msg.Content)

			// Fix if missing in Message object
			if msg.Chat == "" {
				msg.Chat = id
			}

			db.Put(msg)

			ch.Client.Socket.Emit("operator:message", msg.Content, nil)
		}

		resp.WriteHeader(http.StatusOK)
	}
}

// UpdateMessage PATCH /api/chat/:id/message/:mid
func UpdateMessage(db *store.InMemory) func(resp http.ResponseWriter, req *http.Request, params httprouter.Params) {
	return NotImplemented
}

// DeleteMessage DELETE /api/chat/:id/message/:mid
func DeleteMessage(db *store.InMemory) func(resp http.ResponseWriter, req *http.Request, params httprouter.Params) {
	return NotImplemented
}
