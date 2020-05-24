package server

import (
	"net/http"
	"testing"

	"github.com/steinfletcher/apitest"
)

func TestPing(t *testing.T) {
	apitest.New().
		Handler(setupRouter()).
		Get("/ping").
		Expect(t).
		Body(`{"message":"pong"}`).
		Status(http.StatusOK).
		End()
}
