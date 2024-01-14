package main

import (
	"dungeons_helper/db"
	"dungeons_helper/internal/account"
	"dungeons_helper/internal/alignment"
	"dungeons_helper/internal/class"
	"dungeons_helper/router"
	"log"
)

func main() {
	log.Printf("Server is starting...")
	dbConn, err := db.NewDatabase()
	if err != nil {
		log.Fatalf("Could not initialize database connection: %s", err)
	}

	accountHandler := account.NewHandler(account.NewService(account.NewRepository(dbConn.GetDB())))
	alignmentHandler := alignment.NewHandler(alignment.NewService(alignment.NewRepository(dbConn.GetDB())))
	classHandler := class.NewHandler(class.NewService(class.NewRepository(dbConn.GetDB())))

	r := router.InitRouter(
		router.AccountRouter(accountHandler),
		router.AlignmentRouter(alignmentHandler),
		router.ClassRouter(classHandler),
	)

	log.Printf("Server started")

	if err := router.Start("0.0.0.0:5000", r); err != nil {
		log.Fatalf("Failed to start server: %s", err)
	}
}
