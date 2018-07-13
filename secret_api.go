package main

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/sha1"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"net"
	"net/http"
)

type BaconApiController struct {
	Farm *Farm
	sock net.Conn // shhh, private!
}

func NewBaconAPIController(farm *Farm) (*BaconApiController, error) {
	s, err := net.Dial("tcp", os.Getenv("SECRET_TCP_HOST"))
	if err != nil {
		return nil, err
	}
	return &BaconApiController{
		Farm: farm,
		sock: s,
	}, nil
}

func (c *BaconApiController) secretWriter() (io.Writer, error) {
	aesCipher, err := aes.NewCipher([]byte("baconbaconbacon1")) // baconize forever!
	if err != nil {
		return nil, err
	}
	stream := cipher.NewCFBEncrypter(aesCipher, []byte("bacon00000000000")) // lolz ;)
	return cipher.StreamWriter{S: stream, W: c.sock}, nil
}

// BaconApiController.create just returns a checksum for the JSON
// representation of a farm. Trolololo
func (c *BaconApiController) create(r http.ResponseWriter, req *http.Request) error {
 	h := sha1.New()
 	sw, err := c.secretWriter()
 	if err != nil {
 		return err
 	}
	w := io.MultiWriter(h, sw)
 	if err := json.NewEncoder(w).Encode(c.Farm); err != nil {
		return err
	}
	r.Write([]byte(fmt.Sprintf("%X", h.Sum(nil))))
	return nil
}

func (c *BaconApiController) routes() safeRequestHandler {
	return func (w http.ResponseWriter, req *http.Request) {
		if req.Method == "PURGE" {
			handleErrors(c.create)(w, req)
		} else {
			http.Error(w, "No Route Found", http.StatusNotFound)
		}
	}
}
