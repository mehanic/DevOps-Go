package main

import (
    "database/sql"
    "net/http"
    "strings"
    "fmt"

    vault "github.com/mittwald/vaultgo"
    "github.com/gorilla/mux"
    _ "github.com/lib/pq"
)

type VaultCreds struct {
    Data struct {
        User        string `json:"username"`
        Password    string  `json:"password"`
    } `json:"data"`
}

type DbConnection struct {
    Dbname string
    Host string
    Port int
    User string 
    Password string
}

// Read (generate) credentials from our Vault server.
// Don't forget to update your Vault address and token.
func getDBConnectionConfig() DbConnection {
    c, err := vault.NewClient("VAULT SERVER ADDRESS", vault.WithCaPath(""), vault.WithAuthToken("VAULT AUTH TOKEN"))
    if err != nil {
        panic(err)
    }

    key := []string{"v1", "database", "creds", "dbuser"}
    options := &vault.RequestOptions{}
    response := &VaultCreds{}

    err = c.Read(key, response, options)
    if err != nil {
        panic(err)
    }

    return DbConnection{
        Dbname: "app",
        Host: "192.168.56.101",
        Port: 5432,
        User: response.Data.User,
        Password: response.Data.Password,
    }
}

// This function opens up a new Postgres connection to our server and returns it.
func openConnection() *sql.DB {
    config := getDBConnectionConfig()
    connStr := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", config.Host, config.Port, config.User, config.Password, config.Dbname)

    db, err := sql.Open("postgres", connStr)
    if err != nil {
        panic(err)
    }

    err = db.Ping()
    if err != nil {
        panic(err)
    }

    fmt.Printf("Connected to PostgreSQL db using user <%s> and password <%s>\n", config.User, config.Password);
    return db
}

// Factory to create the function that handles the index request "/".
// It queries the database and return a join of all names in the users table. Pretty simple.
func newIndexHandler(db *sql.DB) func(http.ResponseWriter, *http.Request) {
    return func(w http.ResponseWriter, r *http.Request) {
        rows, err := db.Query(`SELECT name FROM users`)
        if err != nil {
            panic(err)
        }
        defer rows.Close()


        names := []string{}
        for rows.Next() {
            var name string
            err = rows.Scan(&name)
            if err != nil {
                panic(err)
            }

            names = append(names, name)
        }

        w.WriteHeader(http.StatusOK)
        fmt.Fprint(w, strings.Join(names, ", "))
    }
}

func main() {
    // Setup DB connection
    db := openConnection()
    defer db.Close();

    // Setup router 
    r := mux.NewRouter()
    r.HandleFunc("/", newIndexHandler(db))
    http.Handle("/", r)

    // Start listening
    fmt.Println("Listening on 127.0.0.1:8002")
    http.ListenAndServe("127.0.0.1:8002", r)
}


//go build
//./go-vault-example

#
# Output is something like
# Connected to PostgreSQL db using user <v-root-truly-OQZIsxFo71ytoXb3aIKo-1660475644> and password <4u-A74pocYWqzUaSqW5L>                                                                     │
# Listening on 127.0.0.1:8002
#
