package api_handlers

import (
	"fmt"
	"net/http"

	"github.com/numan-khan/g-store/internal/pkg/store"
)

type StoreHandler struct {
	store *store.Store
}

func NewStoreHandler(s *store.Store) *StoreHandler {
	return &StoreHandler{store: s}
}

func (sh *StoreHandler) GetKey(rw http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	key := r.Form.Get("key")
	value, err := sh.store.GetKey(key)

	fmt.Fprintf(rw, "Value = %q, Error = %v", value, err)

}

func (sh *StoreHandler) SetKey(rw http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	key := r.Form.Get("key")
	value := r.Form.Get("value")

	err := sh.store.SetKey(key, []byte(value))

	fmt.Fprintf(rw, "Error = %v", err)
}
