package main

import (
  "github.com/gorilla/websocket"
  "fmt"
  "html/template"
  "net/http"
  "time"
)

type Item struct {
  Url string
  Status bool
}

type CheckList struct {
  Items   []Item
}

func (item *Item) check() {
  timeout := time.Duration(800 * time.Millisecond)
  client := http.Client{
    Timeout: timeout,
  }
  resp, err := client.Get(item.Url)
  if err != nil {
    item.Status = false
    } else {
      fmt.Println(string(resp.StatusCode) + resp.Status)
      item.Status = true
    }
}

func (item *Item) scheduleCheck(conn *websocket.Conn) {
  for {
    time.Sleep(5 * time.Minute)
    item.check()
    var err error

    if err = conn.WriteJSON(item); err != nil {
        fmt.Println("Can't send")
        break
    }
  }
}

var upgrader = websocket.Upgrader{
    ReadBufferSize:  1024,
    WriteBufferSize: 1024,
}

func wsHandler(w http.ResponseWriter, r *http.Request) {
    checkList := CheckList{}
    conn, err := upgrader.Upgrade(w, r, nil)
    if err != nil {
        fmt.Println(err)
        return
    }

    for {
        _, message, err := conn.ReadMessage()
        if err != nil {
          fmt.Println(err)
            return
        }
        fmt.Println("Received from client: " + string(message))
        item := Item{string(message), false}
        checkList.Items = append(checkList.Items, item)
        item.check()
        go item.scheduleCheck(conn)

        if err = conn.WriteJSON(item); err != nil {
          fmt.Println(err)
          return
        }
    }
}

func getHandler(w http.ResponseWriter, r *http.Request) {
  tmpl := template.Must(template.ParseFiles("form.html"))
  tmpl.Execute(w, nil)
}

func main() {
  http.HandleFunc("/", getHandler)
  http.HandleFunc("/ws", wsHandler)
  if err := http.ListenAndServe(":8080", nil); err != nil {
    fmt.Println("ListenAndServe:", err)
  }
}
