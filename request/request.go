package request

import (
	"fmt"
	"io"
	reader "lark/reader"
	"net/url"
	"strconv"
	"strings"
)

type Request struct {
	Method     string
	Proto      string
	RequestURI string
	Host       string

	URL *url.URL

	ProtoMajor int
	ProtoMinor int

	Header Header
}

func (request *Request) String() {
  fmt.Printf("\n==============================")
  fmt.Printf("\nMethod: %s", request.Method)
	fmt.Printf("\nProto: HTTP %d.%d", request.ProtoMajor, request.ProtoMinor)
	fmt.Printf("\nRequest URI: %s", request.RequestURI)
	fmt.Printf("\nHost: %s", request.Host)
  
  for header, value := range request.Header {
    fmt.Printf("\n%s: %s", header, value)
  }
  fmt.Printf("\n==============================")
}

func (request *Request) HydrateFromByteReader(reader *reader.ByteReader) error {
	line, err := reader.ReadLine()
	if err != nil {
		return err
	}

	var ok bool
	request.Method, request.Proto, request.RequestURI, ok = parseMethodProtocolAndURI(line)

	if !ok {
		return fmt.Errorf("Failed to parse either the method, protocol, or request URI")
	}

	var parsed bool
	request.ProtoMajor, request.ProtoMinor, parsed = parseHTTPMajorMinor(request.Proto)
	if !parsed {
		return fmt.Errorf("Failed to parse either major or minor HTTP version")
	}

	request.URL, err = url.ParseRequestURI(request.RequestURI)
	if err != nil {
		return err
	}

  err = request.hydrateHeaderData(reader)
  if err != nil {
    return err
  }

  request.Host = request.URL.Host
  if request.Host == "" {
    host, ok := request.Header["Host"]
    if ok {
      request.Host = host
    }
  }

  return nil
}

func parseHTTPMajorMinor(protocol string) (int, int, bool) {
	if !strings.HasPrefix(protocol, "HTTP/") {
		return 0, 0, false
	}

	if len(protocol) != len("HTTP/M.m") {
		return 0, 0, false
	}

	major, err := strconv.ParseUint(protocol[5:6], 10, 0)
	if err != nil {
		return 0, 0, false
	}

	minor, err := strconv.ParseUint(protocol[7:8], 10, 0)
	if err != nil {
		return 0, 0, false
	}

	return int(major), int(minor), true
}

func parseMethodProtocolAndURI(requestString string) (string, string, string, bool) {
	method, URLAndProto, ok1 := strings.Cut(requestString, " ")
	requestURI, protocol, ok2 := strings.Cut(URLAndProto, " ")

	if !ok1 || !ok2 {
		return "", "", "", false
	}

	return method, protocol, requestURI, true
}

func (request *Request) hydrateHeaderData(reader *reader.ByteReader) error {
	request.Header = make(Header)

	for {
		line, err := reader.ReadLine()
		if err != nil {
			if err == io.EOF {
				return nil
			}
			return err
		}

		key, value, ok := strings.Cut(line, ": ")
		if ok {
			request.Header[key] = value
		}
	}
}
