package groupServise

import (
	"bytes"
	"log"
	"strconv"
)

func (s *service) GetAll() *bytes.Buffer {
	var buf bytes.Buffer

	for key, val := range s.Users {
		if _, err := buf.WriteString("User ID: " + strconv.Itoa(key) + "\t" + val.toString() + "\n"); err != nil {
			log.Println(err)
		}
	}
	return &buf
}
