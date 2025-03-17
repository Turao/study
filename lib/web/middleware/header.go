package middleware

import (
	"errors"
	"fmt"
	"net/http"
)

type middleware func(http.Handler) http.Handler

// HeaderValidatorFunc is a function that validates the headers of the request.
type HeaderValidatorFunc func(r *http.Request) error

// HeaderValidator validates the headers of the request. If any of the validators return an error, the request is aborted with a 400 status code.
func HeaderValidator(fs ...HeaderValidatorFunc) middleware {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			es := make([]error, 0, len(fs))
			for _, f := range fs {
				if err := f(r); err != nil {
					es = append(es, err)
				}
			}

			if len(es) > 0 {
				errs := errors.Join(es...)
				http.Error(w, errs.Error(), http.StatusBadRequest)
				return
			}

			next.ServeHTTP(w, r)
		})
	}
}

// HeaderExists validates if the header exists in the request.
func HeaderExists(header string) HeaderValidatorFunc {
	return func(r *http.Request) error {
		if r.Header.Get(header) == "" {
			return fmt.Errorf("header %s is required", header)
		}

		return nil
	}
}
