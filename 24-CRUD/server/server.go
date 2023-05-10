package server

import (
	"crud/db"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type user struct {
	ID    uint32 `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	//le o body com a biblioteca ioutil
	bodyReq, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.Write([]byte("Falha ao ler o corpo da requisição!"))
		return
	}

	var user user
	//converte o body para a struct
	if err = json.Unmarshal(bodyReq, &user); err != nil {
		w.Write([]byte("Erro ao converter usuário para struct!"))
		return
	}

	db, err := db.Connect()
	if err != nil {
		w.Write([]byte("Erro ao conectar!"))
		return
	}
	defer db.Close()

	//PREPARE STATEMENT - usado para inserts, updates e deletes
	//sempre colocar o ? para evitar o ataque sql injection
	statement, err := db.Prepare("insert into users (name, email) values (?,?)")
	if err != nil {
		w.Write([]byte("Erro ao criar statement!"))
		return
	}
	defer statement.Close()

	insertion, err := statement.Exec(user.Name, user.Email)
	if err != nil {
		w.Write([]byte("Erro ao executar statement!"))
		return
	}

	idInserted, err := insertion.LastInsertId()
	if err != nil {
		w.Write([]byte("Erro ao obter id inserido!"))
		return
	}

	//STATUS CODE 201 é created
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(fmt.Sprintf("Usuário inserido com sucesso! Id: %d", idInserted)))
}

func FindUsers(w http.ResponseWriter, r *http.Request) {
	db, err := db.Connect()
	if err != nil {
		w.Write([]byte("Erro ao conectar!"))
		return
	}
	defer db.Close()

	//QUERY - usado para consultas
	lines, err := db.Query("select * from users")
	if err != nil {
		w.Write([]byte("Erro ao buscar usuarios!"))
		return
	}
	defer lines.Close()

	var users []user

	for lines.Next() {
		var user user

		if err := lines.Scan(&user.ID, &user.Name, &user.Email); err != nil {
			w.Write([]byte("Erro ao escanear usuarios!"))
			return
		}

		users = append(users, user)
	}

	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(users); err != nil {
		w.Write([]byte("Erro ao converter os usuarios para json!"))
		return
	}
}

func FindUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	ID, err := strconv.ParseUint(params["id"], 10, 32)
	if err != nil {
		w.Write([]byte("Erro ao converter o parametro!"))
		return
	}

	db, err := db.Connect()
	if err != nil {
		w.Write([]byte("Erro ao conectar!"))
		return
	}
	defer db.Close()

	line, err := db.Query("select * from users where id = ?", ID)
	if err != nil {
		w.Write([]byte("Erro ao buscar o usuario!"))
		return
	}
	defer line.Close()

	var user user
	if line.Next() {
		if err := line.Scan(&user.ID, &user.Name, &user.Email); err != nil {
			w.Write([]byte("Erro ao escanear usuarios!"))
			return
		}
	}

	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(user); err != nil {
		w.Write([]byte("Erro ao converter o usuario para json!"))
		return
	}
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	ID, err := strconv.ParseUint(params["id"], 10, 32)
	if err != nil {
		w.Write([]byte("Erro ao converter o parametro!"))
		return
	}

	bodyReq, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.Write([]byte("Falha ao ler o corpo da requisição!"))
		return
	}

	var user user
	if err = json.Unmarshal(bodyReq, &user); err != nil {
		w.Write([]byte("Erro ao converter usuário para struct!"))
		return
	}

	db, err := db.Connect()
	if err != nil {
		w.Write([]byte("Erro ao conectar!"))
		return
	}
	defer db.Close()

	statement, err := db.Prepare("update users set name = ?, email = ? where id = ?")
	if err != nil {
		w.Write([]byte("Erro ao criar statement!"))
		return
	}
	defer statement.Close()

	if _, err := statement.Exec(user.Name, user.Email, ID); err != nil {
		w.Write([]byte("Erro ao executar statement!"))
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	ID, err := strconv.ParseUint(params["id"], 10, 32)
	if err != nil {
		w.Write([]byte("Erro ao converter o parametro!"))
		return
	}

	db, err := db.Connect()
	if err != nil {
		w.Write([]byte("Erro ao conectar!"))
		return
	}
	defer db.Close()

	statement, err := db.Prepare("delete from users where id = ?")
	if err != nil {
		w.Write([]byte("Erro ao criar statement!"))
		return
	}
	defer statement.Close()

	if _, err := statement.Exec(ID); err != nil {
		w.Write([]byte("Erro ao executar statement!"))
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
