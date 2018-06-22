package main

import (
  "net/http"
  "net/http/httptest"
  "testing"
  "github.com/gorilla/websocket"
  "strings"
)

func TestGetHandler(t *testing.T) {
  req, err := http.NewRequest("GET", "/", nil)
  if err != nil {
      t.Fatal(err)
  }

  recorder := httptest.NewRecorder()
  handler := http.HandlerFunc(getHandler)

  handler.ServeHTTP(recorder, req)

  if status := recorder.Code; status != http.StatusOK {
      t.Errorf("handler returned wrong status code: got %v want %v",
          status, http.StatusOK)
  }
}

func TestWSHandler(t *testing.T) {
  s := httptest.NewServer(http.HandlerFunc(wsHandler))
  defer s.Close()

  // Convert http://127.0.0.1 to ws://127.0.0.
  u := "ws" + strings.TrimPrefix(s.URL, "http")

  // Connect to the server
  ws, _, err := websocket.DefaultDialer.Dial(u, nil)
  if err != nil {
      t.Fatalf("%v", err)
  }
  defer ws.Close()

  if err := ws.WriteMessage(websocket.TextMessage, []byte("https://google.com")); err != nil {
      t.Fatalf("%v", err)
  }
  item := Item{}
  if err = ws.ReadJSON(&item); err != nil {
      t.Fatalf("%v", err)
  }

  expectedItem := Item{"https://google.com", true}
  if item != expectedItem {
      t.Errorf("Different Item received from expected Item")
  }
}
