package whatismyip

import (
	"fmt"
	"net/http"
)

// WhatIsMyIp is a function that prints the client ip-address
func WhatIsMyIp(w http.ResponseWriter, req *http.Request) {
	ip := req.Header.Get("x-forwarded-for")
	w.Header().Set("Content-Type", "text/plain")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintln(w, ip)
}
