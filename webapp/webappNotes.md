# Go lang webapp example

> func (w http.ResponseWriter, r *http.Request)

The function receives `two parameters`:

An `http.ResponseWriter` which is where you `write your text/html response` to.

An `http.Request` which contains `all information about this HTTP request` including things like the `URL` or `header` fields.


```go
package main


```




```go

http.HandleFunc("/", func (w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Hello, you've requested: %s\n", r.URL.Path)
})

```


```go

package main

import (
    "fmt"
    "net/http"
)

func main() {
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintf(w, "Hello, you've requested: %s\n", r.URL.Path)
    })

    http.ListenAndServe(":80", nil)
}

```



```go
http.HandleFunc("/", func (w http.ResponseWriter, r *http.Request) {
    fmt.Fprint(w, "Welcome to my website!")
})
```

<br />
<br />

## HTTP Server

HTTP server has a few key jobs to take care

- Process dynamic requests
- Serve static assets
- Accept connections


<br />

### Process dynamic requests

##### net/http
- `net/http` package contains all utilities needed to `accept requests` and `handle` them dynamically

##### http.HandleFunc
- `http.HandleFunc`: register a new handler
    - It’s `first parameter` takes a `path` to match
    - `a function` to execute `as a second`

```go
http.HandleFunc("/", func (w http.ResponseWriter, r *http.Request) {
    fmt.Fprint(w, "Welcome to my website!")
})
```

##### http.Request

- `http.Request` contains `all information` about the `request` and `it’s parameters`

- `r.URL.Query().Get("token")`
    - To read `read GET parameters`

- `r.FormValue("email")`
    - POST parameters (fields from an HTML form)


<br />

### Serving static assets

##### http.FileServer

To serve `static assets` like JavaScript, CSS and images, we use the inbuilt `http.FileServer` point it to a url path.

For the file server to work properly it needs to know, where to serve files from

We can do this like so:

> fs := http.FileServer(http.Dir("static/"))

One thing to note: In order to serve files correctly, we need to strip away a part of the url path. Usually this is the name of the directory our files live in.

> http.Handle("/static/", http.StripPrefix("/static/", fs))


```go
package main

import (
    "fmt"
    "net/http"
)

func main() {
    http.HandleFunc("/", func (w http.ResponseWriter, r *http.Request) {
        fmt.Fprintf(w, "Welcome to my website!")
    })

    fs := http.FileServer(http.Dir("static/"))
    http.Handle("/static/", http.StripPrefix("/static/", fs))

    http.ListenAndServe(":80", nil)
}
```


> go get -u github.com/gorilla/mux


```go
package main

import (
    "fmt"
    "net/http"

    "github.com/gorilla/mux"
)

func main() {
    r := mux.NewRouter()

    r.HandleFunc("/books/{title}/page/{page}", func(w http.ResponseWriter, r *http.Request) {
        vars := mux.Vars(r)
        title := vars["title"]
        page := vars["page"]

        fmt.Fprintf(w, "You've requested the book: %s on page %s\n", title, page)
    })

    http.ListenAndServe(":80", r)
}
```


> r.HandleFunc("/books/{title}", BookHandler).Host("www.mybookstore.com")




