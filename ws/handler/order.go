package handler

import (
	"log"
	"net/http"

	"context"
	"encoding/json"

	"github.com/gorilla/websocket"
	"github.com/hopehook/micro-project/common/go_client"
	"github.com/hopehook/micro-project/proto/go/service_order"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool { return true },
}

func Stream(w http.ResponseWriter, r *http.Request) {
	// Upgrade request to websocket
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Fatal("Upgrade: ", err)
		return
	}
	defer conn.Close()

	// Handle websocket request
	if err := stream(conn); err != nil {
		log.Fatal("Echo: ", err)
		return
	}
	log.Println("Stream complete")
}

func stream(ws *websocket.Conn) error {
	// Even if we aren't expecting further requests from the websocket, we still need to read from it to ensure we
	// get close signals
	//go func() {
	//	for {
	//		if _, _, err := ws.NextReader(); err != nil {
	//			break
	//		}
	//	}
	//}()

	// Read from the stream server and pass responses on to websocket
	for {
		// Read initial request from websocket
		var req service_order.GetOrdersRequest
		err := ws.ReadJSON(&req)
		if err != nil {
			return err
		}

		log.Printf("Received Request: %v", req)

		// Send request to stream server
		rsp, err := go_client.OrderServiceClient.GetOrders(context.TODO(), &service_order.GetOrdersRequest{
			PageIndex: 0,
			PageCount: 10,
		})
		if err != nil {
			return err
		}

		// Write server response to the websocket
		result, _ := json.Marshal(rsp.Orders)
		err = ws.WriteJSON(result)
		if err != nil {
			// End request if socket is closed
			if isExpectedClose(err) {
				log.Println("Expected Close on socket", err)
				break
			} else {
				return err
			}
		}
	}

	return nil
}

func isExpectedClose(err error) bool {
	if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway) {
		log.Println("Unexpected websocket close: ", err)
		return false
	}

	return true
}
