package handlers

import (
	"net/http"
	"net/http/httptest"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestRootHandler(t *testing.T) {

	// Only pass t into top-level Convey calls
	Convey("Given the root route", t, func() {
		req, err := http.NewRequest("GET", "/", nil)
		if err != nil {
			t.Fatal(err)
		}
		Convey("When serving the request", func() {
			rr := httptest.NewRecorder()
			handler := http.HandlerFunc(RootHandler)
			handler.ServeHTTP(rr, req)

			Convey("The return status should be status OK", func() {
				status := rr.Code
				So(status, ShouldEqual, http.StatusOK)
			})

			Convey("The return value should be as expected", func() {
				body := rr.Body.String()
				So(body, ShouldEqual, "Running API v1\n")
			})

		})
	})

	Convey("Given the route which does not exists", t, func() {
		req, err := http.NewRequest("GET", "/something", nil)
		if err != nil {
			t.Fatal(err)
		}
		Convey("When serving the request", func() {
			rr := httptest.NewRecorder()
			handler := http.HandlerFunc(RootHandler)
			handler.ServeHTTP(rr, req)

			Convey("The return status should be status not found", func() {
				status := rr.Code
				So(status, ShouldEqual, http.StatusNotFound)
			})

			Convey("The return value should be as expected", func() {
				body := rr.Body.String()
				So(body, ShouldEqual, "Not found\n")
			})

		})
	})
}
