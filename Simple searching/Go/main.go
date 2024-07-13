package main

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

func main() {
	var err error
	db, err = sql.Open("mysql", "myuser:mypassword@tcp(mysql:3306)/mydatabase")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	http.HandleFunc("/api/query", queryHandler)
	http.HandleFunc("/api/add", addHandler)
	http.HandleFunc("/api/delete", deleteHandler)

	http.Handle("/", http.FileServer(http.Dir("/static")))

	log.Println("Starting server on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func queryHandler(w http.ResponseWriter, r *http.Request) {
	rows, err := db.Query("SELECT * FROM mytable")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var results []map[string]interface{}
	cols, _ := rows.Columns()
	for rows.Next() {
		columns := make([]interface{}, len(cols))
		columnPointers := make([]interface{}, len(cols))
		for i := range columns {
			columnPointers[i] = &columns[i]
		}
		if err := rows.Scan(columnPointers...); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		result := make(map[string]interface{})
		for i, colName := range cols {
			val := columnPointers[i].(*interface{})
			result[colName] = *val
		}
		results = append(results, result)
	}

	json.NewEncoder(w).Encode(results)
}

func addHandler(w http.ResponseWriter, r *http.Request) {
	var requestData struct {
		Data string `json:"data"`
	}
	if err := json.NewDecoder(r.Body).Decode(&requestData); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	_, err := db.Exec("INSERT INTO mytable (data) VALUES (?)", requestData.Data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(map[string]string{"message": "Data added successfully"})
}

func deleteHandler(w http.ResponseWriter, r *http.Request) {
	var requestData struct {
		ID int `json:"id"`
	}
	if err := json.NewDecoder(r.Body).Decode(&requestData); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	_, err := db.Exec("DELETE FROM mytable WHERE id = ?", requestData.ID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(map[string]string{"message": "Data deleted successfully"})
}
