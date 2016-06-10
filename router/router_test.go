package router_test

import (
	"net/http"
	"testing"

	"github.com/bcokert/terragen/controller"
	"github.com/bcokert/terragen/router"
)

func TestCreateDefaultRouter(t *testing.T) {
	var r http.Handler = router.CreateDefaultRouter(controller.CreateDefaultServer())

	if r == nil {
		t.Errorf("Did not create a router")
	}
}
