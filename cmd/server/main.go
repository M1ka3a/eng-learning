package main

import (
	"backend-lab-go/internal/biz"
	"backend-lab-go/internal/dal"
	"backend-lab-go/internal/handler"
	"backend-lab-go/internal/worker"
	"context"
	"database/sql"
	"log"
	"net/http"

	_ "github.com/jackc/pgx/v5/stdlib"
)

func main() {
	ctx := context.Background()

	// 1. 初始化 PostgreSQL
	db, err := sql.Open(
		"pgx",
		"host=/tmp dbname=backend_lab sslmode=disable",
	)
	if err != nil {
		log.Fatal("Open database failed:", err)
	}
	defer db.Close()

	if err := db.Ping(); err != nil {
		log.Fatal("Connect database failed:", err)
	}

	log.Println("PostgreSQL connected successfully")

	// 2. 初始化 Stores，共享同一个数据库连接池
	categoryStore := dal.NewCategoryStore(db)
	jobStore := dal.NewJobStore(db)
	jobBiz := biz.NewJobBiz(jobStore)

	// 3. 初始化 Handlers
	categoryHandler := handler.NewCategoryHandler(categoryStore)
	jobHandler := handler.NewJobHandler(jobBiz)

	// 4. 注册 HTTP 路由
	http.HandleFunc("GET /categories", categoryHandler.Get)
	http.HandleFunc("POST /categories", categoryHandler.Create)

	http.HandleFunc("POST /jobs", jobHandler.Create)
	http.HandleFunc("GET /jobs/{id}", jobHandler.GetJobByID)

	// 5. 初始化并启动 Worker
	worker := worker.NewWorker(jobBiz)

	const workerCount = 3

	for i := 0; i < workerCount; i++ {
		go worker.Run(ctx)
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	go worker.Run(ctx)

	// 6. 启动 HTTP Server
	log.Println("HTTP server listening on :8080")

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Printf("HTTP server stopped: %v", err)
	}
}
