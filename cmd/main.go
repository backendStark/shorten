package main

import (
	"fmt"
	"net/http"
	"shorten/configs"
	"shorten/internal/auth"
	"shorten/internal/link"
	"shorten/internal/stat"
	"shorten/internal/user"
	"shorten/pkg/db"
	"shorten/pkg/event"
	"shorten/pkg/middleware"
)

const serverPortNumber = "8081"
const serverPort = ":" + serverPortNumber

func main() {
	conf := configs.LoadConfig()
	db := db.NewDb(conf)
	router := http.NewServeMux()
	eventBus := event.NewEventBus()

	// Repositories
	linkRepository := link.NewLinkRepository(db)
	userRepository := user.NewUserRepository(db)
	statRepository := stat.NewStatRepository(db)

	// Services
	authService := auth.NewAuthService(userRepository)
	statService := stat.NewAuthService(stat.StatServiceDeps{
		EventBus:       eventBus,
		StatRepository: statRepository,
	})

	// Handlers
	auth.NewAuthHandler(router, auth.AuthHandlerDeps{
		Config:      conf,
		AuthService: authService,
	})
	link.NewLinkHandler(router, link.LinkHandlerDeps{
		LinkRepository: linkRepository,
		Config:         conf,
		// StatRepository: statRepository,
		EventBus: eventBus,
	})
	stat.NewStatHandler(router, stat.StatHandlerDeps{
		StatRepository: statRepository,
		Config:         conf,
	})

	// Middlewares
	stack := middleware.Chain(
		middleware.CORS,
		middleware.Logging,
	)

	server := http.Server{
		Addr:    serverPort,
		Handler: stack(router),
	}

	go statService.AddClick()

	fmt.Println("Server start and listening port:", serverPortNumber)
	if err := server.ListenAndServe(); err != nil {
		panic(err)
	}
}
