package main

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"log"
)

// WSMessage represents websocket message format
type WSMessage struct {
	Message string `json:"message"`
	Action  string `json:"action"`
}

var wsupgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func (s *ServerSettings) sendWSMessage(action string, message string) {
	response := WSMessage{
		Action:  action,
		Message: message,
	}
	responseJSON, err := json.Marshal(response)
	if err != nil {
		fmt.Println("Can't serialize", response)
	}
	if s.statusWebSocket != nil {
		s.statusWebSocket.WriteMessage(websocket.TextMessage, responseJSON)
	}
}

func (s *ServerSettings) handleStatusViaWS(c *gin.Context) {
	log.Println("handleJoinViaWS")
	conn, err := wsupgrader.Upgrade(c.Writer, c.Request, nil)

	if err != nil {
		log.Printf("Failed to upgrade ws: %+v", err)
		return
	}

	s.statusWebSocket = conn

	for {
		t, msg, err := conn.ReadMessage()
		log.Printf("Got ws message: %s", msg)
		if err != nil {
			log.Printf("Error reading message: %+v", err)
			break
		}
		if t != websocket.TextMessage {
			log.Printf("Not a text message: %d", t)
			continue
		}
		var m WSMessage
		err = json.Unmarshal(msg, &m)
		if err != nil {
			log.Printf("Failed to unmarshal message '%+v': %+v", string(msg), err)
			continue
		}
		log.Printf("WS message: %+v", m)
		if m.Action == "new" {
			s.createNewPrometheus(m.Message)
		}
	}
}

func (s *ServerSettings) createNewPrometheus(url string) {
	log.Printf("Creating new prom instance for %s", url)
	metricsTar := getMetricsTar(url)
	s.sendWSMessage("status", fmt.Sprintf("metrics tar: %s", metricsTar))

	// Create namespace
	appLabel := generateAppLabel()
	s.sendWSMessage("status", fmt.Sprintf("Generating app %s", appLabel))

	if output, err := applyKustomize(appLabel, metricsTar); err != nil {
		s.sendWSMessage("failure", fmt.Sprintf("%s\n%s", output, err.Error()))
		return
	} else {
		s.sendWSMessage("status", output)
	}

	if output, err := exposeService(appLabel); err != nil {
		s.sendWSMessage("failure", fmt.Sprintf("%s\n%s", output, err.Error()))
		return
	} else {
		s.sendWSMessage("status", output)
	}

	// Return route name
	promRoute, err := getRouteHost(appLabel)
	if err != nil {
		s.sendWSMessage("failure", err.Error())
		return
	}
	s.sendWSMessage("done", promRoute)
}
