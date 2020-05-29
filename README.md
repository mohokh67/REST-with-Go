### Endpoints
- [GET] `/`
- [GET, POST, HEAD, OPTIONS] `/organisation/accounts`
- [GET, DELETE, HEAD, OPTIONS] `/organisation/accounts/{accountId}`

Create a new account example:
- [POST] `/organisation/accounts/`
    - Headers:

        `"Content-Type": "application/vnd.api+json"`
    - Body
```json
{
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
}
```

#### Extras

- `HEAD` HTTP method has been added just as a good practice for the purpose of resource availability and to reduce the cost of expensive requests.

- `OPTIONS` HTTP methods has been added to give more information to the end user for available options.

 __Pagination__

Pagination is available through `page_number` and `page_size` parameters. e.g. `"/organisation/accounts?page_number=1&page_size=2"` Default `page_size` for the list of accounts is 100. 


__Third party libraries__: In this project I used the [Google UUID](github.com/google/uuid) for ID creations and also [storm](github.com/asdine/storm) for DB creation. By using them, there is no generated code or anything which made the process automated in general. (As you will see in the source code). 

__Database__: All data is saved to the `account.db` file as a database for simplicity. For the purpose of this task, a memory based solution also could work.

__Routing__: There are better libraries to compare `net/http`'s default [ServeMux](http://golang.org/pkg/net/http/#ServeMux), which is very limited and does not have especially good performance. In the real world projects maybe the [httprouter](https://github.com/julienschmidt/httprouter) would be a better solution.

I tried [Convey](https://github.com/smartystreets/goconvey) which seems to be good enough and easy. Also I tried [Ginkgo](https://onsi.github.io/ginkgo) with BDD pattern. I found these two almost good. However there should be more research on them on my side.

Apart from health checks which should be done on a regular basis on the API for monitoring down time, availability; a good set of end to end tests could also give us more confidence. Unit tests and functional tests are playing an important role here too.

The project could run with `go run ./scripts/main.go` . Tests result could be seen by running this command `go test -v  ./...`
