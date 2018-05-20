package handler

import (
	"context"
	"net/http"
	"time"

	"github.com/hopehook/micro-project/common/go_client"
	"github.com/hopehook/micro-project/proto/go/service_order"
	"github.com/micro/go-micro/metadata"
)

// GetOrderList handler
func GetOrderList(w http.ResponseWriter, r *http.Request) {
	//time.Sleep(time.Second * 10)
	// Set arbitrary headers in context
	ctx := metadata.NewContext(context.Background(), map[string]string{
		"X-User-Id": "john",
		"X-From-Id": "script",
	})

	rsp, err := go_client.OrderServiceClient.GetOrders(ctx, &service_order.GetOrdersRequest{
		PageIndex: 0,
		PageCount: 10,
	})
	if err != nil {
		CommonWrite(w, r, -1, err.Error(), nil)
		return
	}
	CommonWriteSuccess(w, r, rsp.Orders)
}

// UpdateOrder async handler
func UpdateOrder(w http.ResponseWriter, r *http.Request) {
	for {
		time.Sleep(time.Second * 1)
		go_client.Pub.Publish(context.TODO(), "/test", "this is a json data")
	}
	CommonWriteSuccess(w, r, nil)
}
