package main

import (
    "database/sql"
    "encoding/json"
    "log"
    "net/http"

    _ "github.com/go-sql-driver/mysql"
)

type Person struct {
    ID    int    `json:"id"`
    Name  string `json:"name"`
    Email string `json:"email"`
}

var db *sql.DB

func initDB() {
    var err error
    db, err = sql.Open("mysql", "may:12345@tcp(mysql:3306)/db")
    if err != nil {
        log.Fatal(err)
    }
}

func main() {
    initDB()
    http.HandleFunc("/api/person", personHandler)
    log.Fatal(http.ListenAndServe(":8080", nil))
}

func personHandler(w http.ResponseWriter, r *http.Request) {
    switch r.Method {
    case http.MethodGet:
        getPersons(w, r)
    case http.MethodPost:
        addPerson(w, r)
    case http.MethodDelete:
        deletePerson(w, r)
    default:
        http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
    }
}

func getPersons(w http.ResponseWriter, r *http.Request) {
    rows, err := db.Query("SELECT id, name, email FROM persons")
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    defer rows.Close()

    persons := []Person{}
    for rows.Next() {
        var p Person
        if err := rows.Scan(&p.ID, &p.Name, &p.Email); err != nil {
            http.Error(w, err.Error(), http.StatusInternalServerError)
            return
        }
        persons = append(persons, p)
    }

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(persons)
}

func addPerson(w http.ResponseWriter, r *http.Request) {
    var p Person
    if err := json.NewDecoder(r.Body).Decode(&p); err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    result, err := db.Exec("INSERT INTO persons (name, email) VALUES (?, ?)", p.Name, p.Email)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    id, err := result.LastInsertId()
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    p.ID = int(id)
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(p)
}

func deletePerson(w http.ResponseWriter, r *http.Request) {
    var p Person
    if err := json.NewDecoder(r.Body).Decode(&p); err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    _, err := db.Exec("DELETE FROM persons WHERE id = ?", p.ID)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    w.WriteHeader(http.StatusNoContent)
}

