package gowebdav

import (
	"bytes"
	"encoding/xml"
	"io"
	"strconv"
	"strings"
	"time"
)

func Join(path0 string, path1 string) string {
	return strings.TrimSuffix(path0, "/") + "/" + strings.TrimPrefix(path1, "/")
}

func String(r io.Reader) string {
	buf := new(bytes.Buffer)
	buf.ReadFrom(r)
	return buf.String()
}

func parseUint(s *string) uint {
	if n, e := strconv.ParseUint(*s, 10, 32); e == nil {
		return uint(n)
	}
	return 0
}

func parseModified(s *string) time.Time {
	if t, e := time.Parse(time.RFC1123, *s); e == nil {
		return t
	}
	return time.Unix(0, 0)
}

func parseXML(data io.Reader, resp interface{}, parse func(resp interface{})) {
	decoder := xml.NewDecoder(data)
	for t, _ := decoder.Token(); t != nil; t, _ = decoder.Token() {
		switch se := t.(type) {
		case xml.StartElement:
			if se.Name.Local == "response" {
				if e := decoder.DecodeElement(resp, &se); e == nil {
					parse(resp)
				}
			}
		}
	}
}
