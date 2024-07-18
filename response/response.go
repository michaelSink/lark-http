package response

import "fmt"

type Status struct {
	StatusCode   uint16
	statusString string
}

var OK Status = Status{
	StatusCode:   200,
	statusString: "OK",
}

var BAD_REQUEST Status = Status{
	StatusCode:   400,
	statusString: "Bad Request",
}

var INTERNAL_ERROR Status = Status{
	StatusCode:   500,
	statusString: "Internal Server Error",
}

const line_break string = "\r\n"

type Response struct {
	method string
	proto  string

	StatusCode Status

	Body []byte
}

func BuildHttpResponse(status Status) []byte {
	return []byte(fmt.Sprintf("HTTP/1.1 %d %s", status.StatusCode, status.statusString))
}

func BuildHttpResponseWithBody(body []byte) []byte {
	return append(
		append(
			BuildHttpResponse(OK),
			[]byte(line_break+line_break)...,
		),
		body...,
	)
}
