package ticketmatic

import (
	"bufio"
	"encoding/json"
	"io"
	"net/http"
)

type Stream struct {
	resp   *http.Response
	reader *bufio.Reader
}

func NewStream(resp *http.Response) *Stream {
	return &Stream{
		resp:   resp,
		reader: bufio.NewReader(resp.Body),
	}
}

func (s *Stream) Next(obj interface{}) error {
	line, err := s.reader.ReadBytes('\n')
	if err != nil {
		return err
	}

	if len(line) == 0 {
		return io.EOF
	}

	return json.Unmarshal(line, obj)
}

func (s *Stream) Close() {
	s.resp.Body.Close()
}
