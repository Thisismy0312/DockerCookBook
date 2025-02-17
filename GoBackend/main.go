package main

import (
    "database/sql"
    "fmt"
    "log"
    "net/http"
    "os"

    "golang.org/x/crypto/bcrypt"
    _ "github.com/lib/pq"        // PostgreSQL driver
)

func getEnvOrDefault(key, defaultValue string) string {
    value := os.Getenv(key)
    if value == "" {
        return defaultValue
    }
    return value
}

func initDB() *sql.DB {
    host := getEnvOrDefault("DB_HOST", "localhost")
    port := getEnvOrDefault("DB_PORT", "5432")
    user := getEnvOrDefault("DB_USER", "admin")
    password := getEnvOrDefault("DB_PASSWORD", "111111")
    dbname := getEnvOrDefault("DB_NAME", "userdb")

    psqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
        host, port, user, password, dbname)

    db, err := sql.Open("postgres", psqlInfo)
    if err != nil {
        log.Fatalf("Error opening database: %v", err)
    }

    err = db.Ping()
    if err != nil {
        log.Fatalf("Error connecting to the database: %v", err)
    }
    fmt.Println("Successfully connected to the database!")
    return db
}

var db *sql.DB

func init() {
    db = initDB()
}

func main() {
    http.HandleFunc("/register", HandleRegisterRoute)
    fmt.Println("Server is running on port 8080...")
    log.Fatal(http.ListenAndServe(":8080", nil))
}

func HandleRegisterRoute(w http.ResponseWriter, r *http.Request) {
    if r.Method != http.MethodPost {
        http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
        return
    }

    // 从请求头中获取参数
    userName := r.Header.Get("user_name")
    userPwd := r.Header.Get("user_pwd")

    if userName == "" || userPwd == "" {
        http.Error(w, "Missing user_name or user_pwd header", http.StatusBadRequest)
        return
    }

    // 对密码进行加密（可选）
    hashedPassword, err := HashPassword(userPwd)
    if err != nil {
        http.Error(w, "Error hashing password", http.StatusInternalServerError)
        return
    }

    // 插入数据到数据库
    sqlStatement := `
    INSERT INTO users (user_name, user_pwd)
    VALUES ($1, $2)
    RETURNING id`
    id := 0
    err = db.QueryRow(sqlStatement, userName, hashedPassword).Scan(&id)
    if err != nil {
        if err == sql.ErrNoRows {
            http.Error(w, "No rows returned", http.StatusInternalServerError)
        } else {
            http.Error(w, "Error inserting data into database", http.StatusInternalServerError)
        }
        return
    }

    fmt.Fprintf(w, "User registered with ID: %d", id)
}

// HashPassword 使用 bcrypt 对密码进行哈希
func HashPassword(password string) (string, error) {
    bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
    return string(bytes), err
}
