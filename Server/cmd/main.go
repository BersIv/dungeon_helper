package main

import (
	"dungeons_helper/db"
	"dungeons_helper/internal/account"
	"dungeons_helper/internal/alignment"
	"dungeons_helper/internal/character"
	"dungeons_helper/internal/class"
	"dungeons_helper/internal/lobby"
	"dungeons_helper/internal/races"
	"dungeons_helper/internal/skills"
	"dungeons_helper/internal/stats"
	"dungeons_helper/internal/subraces"
	"dungeons_helper/internal/websocket"
	"dungeons_helper/router"
	"dungeons_helper/util"
	"log"

	"github.com/joho/godotenv"
)

func main() {
	log.Printf("Server is starting...")
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Could not load env: %s", err)
	}
	dbConn, err := db.NewDatabase()
	if err != nil {
		log.Fatalf("Could not initialize database connection: %s", err)
	}

	jwtTokenGetter := util.JWTTokenGetter{}
	passwordHasher := util.BcryptPasswordHasher{}

	accountHandler := account.NewHandler(account.NewService(account.NewRepository(dbConn.GetDB()), passwordHasher), jwtTokenGetter, passwordHasher)
	alignmentHandler := alignment.NewHandler(alignment.NewService(alignment.NewRepository(dbConn.GetDB())), jwtTokenGetter)
	classHandler := class.NewHandler(class.NewService(class.NewRepository(dbConn.GetDB())), jwtTokenGetter)
	racesHandler := races.NewHandler(races.NewService(races.NewRepository(dbConn.GetDB())), jwtTokenGetter)
	subracesHandler := subraces.NewHandler(subraces.NewService(subraces.NewRepository(dbConn.GetDB())), jwtTokenGetter)
	statsHandler := stats.NewHandler(stats.NewService(stats.NewRepository(dbConn.GetDB())), jwtTokenGetter)
	skillHandler := skills.NewHandler(skills.NewService(skills.NewRepository(dbConn.GetDB())), jwtTokenGetter)
	characterHandler := character.NewHandler(character.NewService(character.NewRepository(dbConn.GetDB())), jwtTokenGetter)
	lobbyHandler := lobby.NewHandler(lobby.NewService(lobby.NewRepository(dbConn.GetDB())), jwtTokenGetter)

	hub := websocket.NewHub(dbConn.GetDB())
	go hub.Run()
	wsHandler := websocket.NewHandler(dbConn.GetDB(), hub, jwtTokenGetter)

	r := router.InitRouter(
		router.AccountRouter(accountHandler),
		router.AlignmentRouter(alignmentHandler),
		router.ClassRouter(classHandler),
		router.RacesRouter(racesHandler),
		router.SubracesRouter(subracesHandler),
		router.StatsRouter(statsHandler),
		router.SkillsRouter(skillHandler),
		router.CharacterRouter(characterHandler),
		router.LobbyRouter(lobbyHandler),
		router.WebsocketRouter(wsHandler),
	)

	log.Printf("Server started")

	if err := router.Start("0.0.0.0:5000", r); err != nil {
		log.Fatalf("Failed to start server: %s", err)
	}
}
