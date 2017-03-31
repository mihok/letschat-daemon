package main

import (
  "log"
  "fmt"
  "flag"
  "net/http"

  "github.com/julienschmidt/httprouter" // Http router
  // "github.com/googollee/go-socket.io" // Socket

  // "github.com/wayn3h0/go-uuid" // UUID (RFC 4122)

  "github.com/minimalchat/mnml-daemon/rest"
  "github.com/minimalchat/mnml-daemon/socket"
  "github.com/minimalchat/mnml-daemon/store"
  // "github.com/minimalchat/mnml-daemon/operator"
  // "github.com/minimalchat/mnml-daemon/client"
  // "github.com/minimalchat/mnml-daemon/chat"
 )

 // Log levels
 const (
   DEBUG string = "DEBUG"
   INFO string = "INFO"
   WARNING string = "WARN"
   ERROR string = "ERROR"
   FATAL string = "FATAL"
 )


// Configuration object
type configuration struct {
  Protocol string
  IP string
  Port int
  Host string
}

var config configuration
var needHelp bool

func help() {
  fmt.Println("mnml-daemon live chat API daemon\n")
  fmt.Println("Find more information at https://github.com/minimalchat/mnml-daemon\n")

  fmt.Println("Flags:")
  flag.PrintDefaults()
}

func init() {
    // Configuration
    flag.IntVar(&config.Port, "port", 8000, "Port used to serve http and websocket traffic on")
    flag.StringVar(&config.IP, "host", "localhost", "IP to serve http and websocket traffic on")
    flag.BoolVar(&needHelp, "h", false, "Get help")
}

func main() {
  // Configuration
  flag.Parse()

  if (needHelp) {
    help();
    return;
  }

  config.Host = fmt.Sprintf("%s:%d", config.IP, config.Port)

  db := new(store.InMemory)

  // Socket.io
  socket := socket.Listen(db)

  // if err != nil {
  //     log.Fatal(err)
  // }

  // Socket.io - Connection event
  // socket.On("connection", func (sock socketio.Socket) {
  //   log.Println(DEBUG, "socket:", fmt.Sprintf("Incoming connection %s %s", sock.Id(), sock.Request().URL.Query().Get("type")))
  //
  //   var cl *client.Client
  //   var op *operator.Operator
  //   t := sock.Request().URL.Query().Get("type")
  //
  //   if (t === "operator") {
  //     log.Println(DEBUG, "socket: OPERATOR!!")
  //   } else if (t === "client") {
  //
  //     hasFingerprint := false
  //     hasCookie := false
  //     hasIP := false
  //     // Does this user match a previous fingerprint ?
  //     //  Does user have cookie?
  //     //  Does user have known IP?
  //
  //     // If yes, lets get/update the user
  //     if (hasFingerprint && hasCookie && hasIP) {
  //
  //     } else { // If no, lets create new user
  //       cl = client.Create(client.Client{
  //           Name: "Site Visitor",
  //         }, sock)
  //
  //       db.Put(cl)
  //     }
  //
  //     // Create new chat, assign user
  //     ch := chat.Chat{
  //       ID: sock.Id(),
  //       Client: cl,
  //       Operator: op,
  //       Open: true,
  //       CreationTime: time.Now(),
  //       UpdatedTime: time.Now(),
  //     }
  //
  //     db.Put(ch)
  //   } else {
  //     log.Println(ERROR, "unable to find connection type.")
  //   }
  //
  //
  //   // Message event
  //   sock.On("client:message", func (msg string) {
  //     log.Println(DEBUG, "client", fmt.Sprintf("%s: %s", sock.Id(), msg, sock.Request().Header))
  //
  //     // Create and Save message
  //     m := chat.Message{
  //       Timestamp: time.Now(),
  //       Content: msg,
  //       Author: ch.Client.StoreKey(),
  //       Chat: ch.ID,
  //     }
  //     db.Put(m)
  //
  //     // Update and Save chat
  //     ch.UpdatedTime = time.Now()
  //     db.Put(ch)
  //
  //   })
  //
  //   // Disconnection event
  //   sock.On("disconnection", func () {
  //     log.Println(DEBUG, "socket:", fmt.Sprintf("%s disconnected", sock.Id()))
  //
  //     // Save chat
  //   })
  // })
  //
  // socket.On("error", func (so socketio.Socket, err error) {
  //     log.Println(ERROR, "socket:", err)
  // })

  router := httprouter.New()

  // 404
  router.NotFound = http.HandlerFunc(func (resp http.ResponseWriter, req *http.Request) {
    resp.Header().Set("Content-Type", "text/plain; charset=UTF-8")
    resp.WriteHeader(http.StatusNotFound)

    fmt.Fprintf(resp, "Not Found")
  })

  // 405
  router.HandleMethodNotAllowed = true
  router.MethodNotAllowed = http.HandlerFunc(func (resp http.ResponseWriter, req *http.Request) {
    resp.Header().Set("Content-Type", "text/plain; charset=UTF-8")
    resp.WriteHeader(http.StatusMethodNotAllowed)

    fmt.Fprintf(resp, "Method Not Allowed")
  })


  // Socket.io handler
  router.HandlerFunc("GET", "/socket.io/", func (resp http.ResponseWriter, req *http.Request) {
    resp.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
    resp.Header().Set("Access-Control-Allow-Credentials", "true")
    // resp.Header().Set("Access-Control-Allow-Headers", "X-Socket-Type")

    socket.ServeHTTP(resp, req)
  })


  router.GET("/api", rest.CatchAll)
  router.GET("/api/", rest.CatchAll)

  // Operators
  router.GET("/api/operators", rest.ReadOperators(db)) // Check
  router.GET("/api/operator/:id", rest.ReadOperator(db)) // Check
  router.POST("/api/operator", rest.CreateOrUpdateOperator(db)) // Check
  router.POST("/api/operator/", rest.CreateOrUpdateOperator(db)) // Check
  router.PUT("/api/operator/:id", rest.CreateOrUpdateOperator(db)) // Check
  router.PATCH("/api/operator/:id", rest.CreateOrUpdateOperator(db)) // Check
  router.DELETE("/api/operator/:id", rest.DeleteOperator(db)) // Check

  // Clients
  router.GET("/api/clients", rest.ReadClients(db)) // Check
  router.GET("/api/client/:id", rest.ReadClient(db)) // Check
  router.POST("/api/client", rest.CreateOrUpdateClient(db)) // Not Implement
  router.POST("/api/client/", rest.CreateOrUpdateClient(db)) // Not Implement
  router.PUT("/api/client/:id", rest.CreateOrUpdateClient(db)) // Not Implement
  router.PATCH("/api/client/:id", rest.CreateOrUpdateClient(db)) // Not Implement
  router.DELETE("/api/client/:id", rest.DeleteClient(db)) // Not Implement

  // Chats
  router.GET("/api/chats", rest.ReadChats(db)) // Check
  router.GET("/api/chat/:id", rest.ReadChat(db)) // Check
  router.POST("/api/chat", rest.CreateOrUpdateChat(db)) // Not Implement
  router.POST("/api/chat/", rest.CreateOrUpdateChat(db)) // Not Implement
  router.PUT("/api/chat/:id", rest.CreateOrUpdateChat(db)) // Not Implement
  router.PATCH("/api/chat/:id", rest.CreateOrUpdateChat(db)) // Not Implement
  router.DELETE("/api/chat/:id", rest.DeleteChat(db)) // Not Implement

  // Chat Messages
  router.GET("/api/chat/:id/messages", rest.ReadMessages(db)) // Check
  router.GET("/api/chat/:id/message/:mid", rest.ReadMessage(db)) // Not Implement
  router.POST("/api/chat/:id/message", rest.CreateMessage(db)) // Check
  router.POST("/api/chat/:id/message/", rest.CreateMessage(db)) // Check
  router.PUT("/api/chat/:id/message/:mid", rest.UpdateMessage(db)) // Not Implement
  router.PATCH("/api/chat/:id/message/:mid", rest.UpdateMessage(db)) // Not Implement
  router.DELETE("/api/chat/:id/message/:mid", rest.DeleteMessage(db)) // Not Implement


  // Server
  log.Println(INFO, "server:", fmt.Sprintf("Listening for HTTP Requests on %s ...", config.Host))
  log.Fatal(http.ListenAndServe(config.Host, router))
}
