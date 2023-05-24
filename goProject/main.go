package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/jackc/pgx/v5"
)

var dbConn *pgx.Conn

func initDb() {
	dbUser := "postgres"
	dbPwd := "123456"
	db := "sys"
	dbPort := "5432"
	dbHost := "0.0.0.0"

	var err error
	url := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable",
		dbUser, dbPwd, dbHost, dbPort, db)
	dbConn, err = pgx.Connect(context.Background(), url)
	if err != nil {
		msg := fmt.Sprintf("Unable to connect to database: %v\n", err)
		fmt.Println(msg)
		os.Exit(-1)
		return
	}
	s := `CREATE TABLE IF NOT EXISTS users (
    userid SERIAL PRIMARY KEY,
    username VARCHAR(50) NOT NULL,
    password VARCHAR(50) NOT NULL
)`
	_, err = dbConn.Exec(context.Background(), s)
	if err != nil {
		msg := fmt.Sprintf("Unable to connect to database: %v\n", err)
		fmt.Println(msg)
		os.Exit(-1)
		return
	}

	fmt.Printf("connected to %s successfully\n", dbHost)
}

type User struct {
	Userid   int    `json:"userid"`
	Username string `json:"username"`
	Password string `json:"password"`
}

func createUser(w http.ResponseWriter, r *http.Request) {
	// 解析请求参数
	var user User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// 插入用户信息
	err = dbConn.QueryRow(context.Background(), "INSERT INTO users(username, password) VALUES($1, $2) RETURNING userid", user.Username, user.Password).Scan(&user.Userid)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// 返回新用户的 userid 值
	json.NewEncoder(w).Encode(map[string]int{"userid": user.Userid})
}

func login(w http.ResponseWriter, r *http.Request) {
	// 解析请求参数
	var user User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// 查询用户信息
	var userid int
	err = dbConn.QueryRow(context.Background(), "SELECT userid FROM users WHERE username=$1 AND password=$2", user.Username, user.Password).Scan(&userid)
	if err != nil {
		http.Error(w, "用户名或密码错误", http.StatusUnauthorized)
		return
	}

	// 返回用户信息
	json.NewEncoder(w).Encode(map[string]int{"userid": userid})
}
func getUsers(w http.ResponseWriter, r *http.Request) {
	// 查询所有用户信息
	rows, err := dbConn.Query(context.Background(), "SELECT userid, username, password FROM users")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	// 将查询结果编码为 JSON 格式，并写入到 HTTP 响应中
	var users []User
	for rows.Next() {
		var user User
		err := rows.Scan(&user.Userid, &user.Username, &user.Password)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		users = append(users, user)
	}
	err = rows.Err()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(users)
}
func deleteUser(w http.ResponseWriter, r *http.Request) {
	// 解析请求参数
	query := r.URL.Query()
	userid := query.Get("userid")
	var err error

	// 删除用户信息
	_, err = dbConn.Exec(context.Background(), "DELETE FROM users WHERE userid=$1", userid)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// 返回成功响应
	w.WriteHeader(http.StatusOK)
}
func getUser(w http.ResponseWriter, r *http.Request) {
	// 解析请求参数
	query := r.URL.Query()
	userid := query.Get("userid")
	var err error
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// 查询用户信息
	var user User
	err = dbConn.QueryRow(context.Background(), "SELECT userid, username, password FROM users WHERE userid=$1", userid).Scan(&user.Userid, &user.Username, &user.Password)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// 返回用户信息
	json.NewEncoder(w).Encode(user)
}

func updateUser(w http.ResponseWriter, r *http.Request) {
	// 解析请求参数
	var user User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	// 更新用户信息
	_, err = dbConn.Exec(context.Background(), "UPDATE users SET username=$1, password=$2 WHERE userid=$3", user.Username, user.Password, user.Userid)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// 返回成功响应
	w.WriteHeader(http.StatusOK)
}

func main() {
	initDb()

	http.HandleFunc("/api/reg", createUser)
	http.HandleFunc("/api/login", login)
	http.HandleFunc("/api/getUserList", getUsers)
	http.HandleFunc("/api/deleteUser", deleteUser)
	http.HandleFunc("/api/updateUser", updateUser)
	http.HandleFunc("/api/getUser", getUser)
	// http.HandleFunc("/api/login", login)
	// http.HandleFunc("/api/user", userMgr)
	http.ListenAndServe(":8081", nil)
}
