package middlewares

import (
	"net/http"
	"strconv"
	"sync"
	"sync/atomic"

	"api/app/responses/bad_request_responses"
	"api/helpers"
)

func GeneralApiRateLimitMiddlewareFactory() Middleware {
	rateLimit := helpers.GetEnvVariable("API_RATE_LIMIT")
	rateLimitInt, err := strconv.Atoi(rateLimit)
	if err != nil {
		return func(next http.HandlerFunc) http.HandlerFunc {
			return func(w http.ResponseWriter, r *http.Request) {
				next(w, r)
			}
		}
	}
	counter := int64(0)
	return func(next http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			atomic.AddInt64(&counter, 1)
			if counter > int64(rateLimitInt) {
				bad_request_responses.RateLimitResponse.Write(w)
				return
			}
			next(w, r)
			atomic.AddInt64(&counter, -1)
		}
	}
}

func EndpointGeneralRateLimitMiddlewareFactory(rateLimit int) Middleware {
	counter := int64(0)
	return func(next http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			atomic.AddInt64(&counter, 1)
			if counter > int64(rateLimit) {
				bad_request_responses.RateLimitResponse.Write(w)
				return
			}
			next(w, r)
			atomic.AddInt64(&counter, -1)
		}
	}
}

func EndpointIpRateLimitMiddlewareFactory(rateLimit int) Middleware {
	counters := sync.Map{}
	return func(next http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			ip := helpers.GetRequestClientIP(r)
			ipCounter, _ := counters.LoadOrStore(ip, int64(0))
			counter := ipCounter.(int64) + 1
			counters.Store(ip, counter)
			if counter > int64(rateLimit) {
				bad_request_responses.RateLimitResponse.Write(w)
				return
			}
			next(w, r)
			ipCounter, _ = counters.Load(ip)
			if ipCounter != nil {
				counter = ipCounter.(int64)
				if counter-1 <= 0 {
					counters.Delete(ip)
				} else {
					counters.Store(ip, counter-1)
				}
			}
		}
	}
}
