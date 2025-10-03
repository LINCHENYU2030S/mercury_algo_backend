package http

import (
	"context"
	"encoding/json"
	"mercury_algo_backend/application"
	"mercury_algo_backend/kitex_gen/api"
	"strconv"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
)

func ListBots(ctx context.Context, c *app.RequestContext) {
	// Parse query parameters based on IDL annotations
	// Get query params as a map
	queryParams := make(map[string]interface{})
	c.QueryArgs().VisitAll(func(key, value []byte) {
		// Convert values appropriately
		if string(key) == "trading_pair" {
			queryParams[string(key)] = string(value)
		} else {
			temp, _ := strconv.Atoi(string(value))
			queryParams[string(key)] = temp
		}
	})

	// Convert map to JSON and then to struct
	jsonData, _ := json.Marshal(queryParams)
	var rpcReq api.ListBotsRequest
	json.Unmarshal(jsonData, &rpcReq)

	// Directly call the handler method
	resp, _ := application.ListBots(ctx, &rpcReq)

	// Return the response
	c.JSON(consts.StatusOK, resp)
}
