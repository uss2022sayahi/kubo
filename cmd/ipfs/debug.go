package main

import (
	"net/http"

	"github.com/uss2022sayahi/kubo/profile"
)

func init() {
	http.HandleFunc("/debug/stack",
		func(w http.ResponseWriter, _ *http.Request) {
			_ = profile.WriteAllGoroutineStacks(w)
		},
	)
}
