package utils

import (
	"encoding/json"
	"github.com/valyala/fasthttp"
)

func HTTPResponse(ctx *fasthttp.RequestCtx, response interface{}, statusCode int) error {
	responseByte, err := json.Marshal(response)
	if err != nil {
		return err
	}

	// Create ctx response
	if _, err := ctx.Write(responseByte); err != nil {
		return err
	}

	ctx.SetStatusCode(statusCode)
	ctx.SetContentType("application/json")
	ctx.Response.Header.Set("Access-Control-Allow-Origin", "*")
	ctx.Response.Header.Set("Access-Control-Allow-Methods", "true")
	ctx.Response.Header.Set("Access-Control-Allow-Headers", "POST, GET, OPTIONS, PUT, DELETE")
	ctx.Response.Header.Set("Access-Control-Allow-Credentials", "Origin, Accept, Content-Type, Accept-Encoding, Authorization, Set-Cookie")
	return nil
}
