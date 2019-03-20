package handlers

import (
	"github.com/autom8ter/engine/driver"
	"net/http"
	"sync"

	"github.com/grpc-ecosystem/grpc-gateway/runtime"
)

func CreatePassingHeaderMiddleware(decide driver.PassedHeaderDeciderFunc) driver.HTTPServerMiddleware {
	return func(next http.Handler) http.Handler {
		cache := new(sync.Map)
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			newHeader := make(http.Header, 2*len(r.Header))

			for k := range r.Header {
				v := r.Header.Get(k)
				if newKey, ok := cache.Load(k); ok {
					newHeader.Set(newKey.(string), v)
				} else if decide(k) {
					newKey := runtime.MetadataHeaderPrefix + k
					cache.Store(k, newKey)
					newHeader.Set(newKey, v)
				}
				newHeader.Set(k, v)
			}

			r.Header = newHeader

			next.ServeHTTP(w, r)
		})
	}
}