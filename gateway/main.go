package main

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"github.com/sirupsen/logrus"
	"log"
	"net/http"
	"toll-calculator/aggregator/client"
)

type apiFunc func(w http.ResponseWriter, r *http.Request) error

func main() {
	listenAddr := flag.String("listenAddr", ":6000", "port for gateway server")
	flag.Parse()
	var (
		client     = client.NewHTTPClient("http://127.0.0.1:3000")
		invHandler = NewInvoiceHandler(client)
	)

	http.HandleFunc("/invoice", makeAPIFunc(invHandler.handleGetInvoice))
	logrus.Infof("Gateway HTTP server running on port: %s", *listenAddr)
	log.Fatal(http.ListenAndServe(*listenAddr, nil))
}

type InvoiceHandler struct {
	client client.Client
}

func NewInvoiceHandler(c client.Client) *InvoiceHandler {
	return &InvoiceHandler{
		client: c,
	}
}

func (h *InvoiceHandler) handleGetInvoice(w http.ResponseWriter, r *http.Request) error {
	// we need to access aggregator client
	invoice, err := h.client.GetInvoice(context.Background(), 1234)
	if err != nil {
		return err
	}
	fmt.Println(invoice)
	return writeJSON(w, http.StatusOK, invoice)
}

func writeJSON(w http.ResponseWriter, code int, v any) error {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(code)
	return json.NewEncoder(w).Encode(&v)
}

func makeAPIFunc(fn apiFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := fn(w, r); err != nil {
			writeJSON(w, http.StatusInternalServerError, map[string]string{"error": err.Error()})
		}
	}
}
