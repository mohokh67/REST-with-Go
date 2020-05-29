package handlers

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/google/uuid"

	. "github.com/smartystreets/goconvey/convey"
)

var sampleAccountID string

func TestAccountHandler(t *testing.T) {

	// Only pass t into top-level Convey calls
	Convey("Given the accounts route", t, func() {
		route := "/organisation/accounts"

		Convey("When creating an account", func() {
			Convey("Without body", func() {
				req, _ := http.NewRequest("POST", route, nil)
				req.Header.Set("Content-Type", "application/vnd.api+json")
				rr := httptest.NewRecorder()
				handler := http.HandlerFunc(AccountRouter)
				handler.ServeHTTP(rr, req)

				Convey("The return status should be status OK", func() {
					status := rr.Code
					So(status, ShouldEqual, http.StatusBadRequest)
				})

				Convey("The return value should be as expected", func() {
					body := rr.Body.String()
					So(body, ShouldEqual, "Bad Request\n")
				})
			})

			Convey("With body", func() {
				var jsonStr = []byte(`{
					"type": "accounts",
					"version": 0,
					"attributes": {
						"country": "GB",
						"base_currency": "GBP",
						"account_number": "41426819",
						"bank_id": "400300",
						"bank_id_code": "GBDSC",
						"bic": "NWBKGB22",
						"iban": "GB11NWBK40030041426819",
						"name": "MoHo Khaleqi",
						"title": "Mr",
						"account_classification": "Personal",
						"joint_account": false,
						"status": "confirmed"
					}
				}`)
				req, _ := http.NewRequest("POST", route, bytes.NewBuffer(jsonStr))
				req.Header.Set("Content-Type", "application/vnd.api+json")
				rr := httptest.NewRecorder()
				handler := http.HandlerFunc(AccountRouter)
				handler.ServeHTTP(rr, req)

				Convey("The return status should be status created", func() {
					status := rr.Code
					So(status, ShouldEqual, http.StatusCreated)
				})

				Convey("The return account ID in the header", func() {
					location := rr.HeaderMap.Get("Localtion")
					sampleAccountID = strings.Split(location, "/")[3]
					So(location, ShouldStartWith, "/organisation/accounts/")

				})

				Convey("The return value should be as expected", func() {
					body := rr.Body.String()
					So(body, ShouldEqual, "")
				})
			})
		})

		Convey("When fetching all accounts", func() {
			req, _ := http.NewRequest("GET", route, nil)
			rr := httptest.NewRecorder()
			handler := http.HandlerFunc(AccountRouter)
			handler.ServeHTTP(rr, req)

			Convey("The return status should be status OK", func() {
				status := rr.Code
				So(status, ShouldEqual, http.StatusOK)
			})

			Convey("The return value should not be empty", func() {
				body := rr.Body.String()
				So(body, ShouldStartWith, `{"data":[`)
				So(body, ShouldNotEqual, `{"data":[]}`)
			})

		})
	})

	Convey("Given the accounts route with account id", t, func() {
		Convey("When id does not exists", func() {
			accountID := uuid.New()
			route := "/organisation/accounts/" + accountID.String()

			Convey("And fetching an account", func() {
				req, _ := http.NewRequest("GET", route, nil)
				rr := httptest.NewRecorder()
				handler := http.HandlerFunc(AccountRouter)
				handler.ServeHTTP(rr, req)

				Convey("The return status should be status Not found", func() {
					status := rr.Code
					So(status, ShouldEqual, http.StatusNotFound)
				})

				Convey("The return value should not not found", func() {
					body := rr.Body.String()
					So(body, ShouldEqual, "Not Found\n")
				})
			})

			Convey("And trying to delete an accounts", func() {
				req, _ := http.NewRequest("DELETE", route, nil)
				rr := httptest.NewRecorder()
				handler := http.HandlerFunc(AccountRouter)
				handler.ServeHTTP(rr, req)

				Convey("The return status should be status Not found", func() {
					status := rr.Code
					So(status, ShouldEqual, http.StatusNotFound)
				})

				Convey("The return value should not not found", func() {
					body := rr.Body.String()
					So(body, ShouldEqual, "Not Found\n")
				})
			})
		})

		Convey("When the accound ID exists", func() {
			route := "/organisation/accounts/" + sampleAccountID

			Convey("And fetching an account", func() {
				req, _ := http.NewRequest("GET", route, nil)
				rr := httptest.NewRecorder()
				handler := http.HandlerFunc(AccountRouter)
				handler.ServeHTTP(rr, req)

				Convey("The return status should be status OK", func() {
					status := rr.Code
					So(status, ShouldEqual, http.StatusOK)
				})

				Convey("The return value should not not found", func() {
					body := rr.Body.String()
					expected := `"attributes":{"country":"GB","base_currency":"GBP","account_number":"41426819","bank_id":"400300","bank_id_code":"GBDSC","bic":"NWBKGB22","iban":"GB11NWBK40030041426819","title":"Mr","name":"MoHo Khaleqi","account_classification":"Personal","joint_account":false,"status":"confirmed"}}}`
					So(body, ShouldEndWith, expected)
				})
			})

			Convey("And deleting an account", func() {
				req, _ := http.NewRequest("DELETE", route, nil)
				rr := httptest.NewRecorder()
				handler := http.HandlerFunc(AccountRouter)
				handler.ServeHTTP(rr, req)

				Convey("The return status should be status no content", func() {
					status := rr.Code
					So(status, ShouldEqual, http.StatusNoContent)
				})

				Convey("The return value should not not found", func() {
					body := rr.Body.String()
					So(body, ShouldEndWith, "Not Found\n")
				})
			})
		})
	})
}
