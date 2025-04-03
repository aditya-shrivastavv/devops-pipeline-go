package main

import (
    "database/sql"
    "fmt"
    "log"
    "net/http"
    "os"

    "github.com/gin-gonic/gin"
    "github.com/prometheus/client_golang/prometheus"
    "github.com/prometheus/client_golang/prometheus/promhttp"
    _ "github.com/lib/pq" // this will help make postgres work with datbase/sql
)

// Define Prometheus metrics
var (
    addGoalCounter    = prometheus.NewCounter(prometheus.CounterOpts{
        Name: "add_goal_requests_total",
        Help: "Total number of add goal requests",
    })
    removeGoalCounter = prometheus.NewCounter(prometheus.CounterOpts{
        Name: "remove_goal_requests_total",
        Help: "Total number of remove goal requests",
    })
    httpRequestsCounter = prometheus.NewCounterVec(
        prometheus.CounterOpts{
            Name: "http_requests_total",
            Help: "Total number of HTTP requests",
        },
        []string{"path"},
    )
)

func init() {
    // Register Prometheus metrics
    prometheus.MustRegister(addGoalCounter)
    prometheus.MustRegister(removeGoalCounter)
    prometheus.MustRegister(httpRequestsCounter)
}

func createConnection() (*sql.DB, error) {
    connStr := fmt.Sprintf("user=%s password=%s host=%s port=%s dbname=%s sslmode=%s",
        os.Getenv("DB_USERNAME"),
        os.Getenv("DB_PASSWORD"),
        os.Getenv("DB_HOST"),
        os.Getenv("DB_PORT"),
        os.Getenv("DB_NAME"),
        os.Getenv("SSL"),
    )

    db, err := sql.Open("postgres", connStr)
    if err != nil {
        return nil, err
    }

    err = db.Ping()
    if err != nil {
        return nil, err
    }

    return db, nil
}
