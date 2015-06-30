package wsapi

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"testing"

	"github.com/hoisie/web"
)

func TestHandleBlockHeight(t *testing.T) {
	var server = web.NewServer()
	server.Get(`/v1/blockheight/?`, handleBlockHeight)
	go server.Run("localhost:8088")
	
	resp, err := http.Get("http://localhost:8088/v1/blockheight/")
	if err != nil {
		t.Errorf(err.Error())
	}
	defer resp.Body.Close()
	
	p, err := ioutil.ReadAll(resp.Body)
	fmt.Println(string(p))
}
