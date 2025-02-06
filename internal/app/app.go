package app

import (
	"context"
	"finance-operations-service/internal/config"
	"finance-operations-service/internal/finance"
	"finance-operations-service/internal/finance/handlers"
	"finance-operations-service/internal/finance/repository/postgres"
	"finance-operations-service/internal/finance/service"
	"finance-operations-service/pkg/client/db"
	"finance-operations-service/pkg/client/db/pg"
	"finance-operations-service/pkg/client/db/transaction"
	"github.com/gin-gonic/gin"
	"log"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

type App struct {
	httpServer     *http.Server
	dbClient       db.Client
	financeService finance.Service
}

func NewApp(ctx context.Context) *App {
	err := config.Load(".env")
	if err != nil {
		log.Fatal(err)
	}

	pgConfig, err := config.NewPGConfig()
	if err != nil {
		log.Fatal(err)
	}

	dbClient, err := pg.New(ctx, pgConfig.DSN())
	if err != nil {
		log.Fatal(err)
	}

	err = dbClient.DB().Ping(ctx)
	if err != nil {
		log.Fatalf("ping error: %s", err.Error())
	}

	txManager := transaction.NewTransactionManager(dbClient.DB())

	financeRepository := postgres.NewFinanceRepository(dbClient)
	financeService := service.NewFinanceService(financeRepository, txManager)

	return &App{
		financeService: financeService,
		dbClient:       dbClient,
	}

}

func (a *App) Run(port string) error {
	router := gin.Default()
	router.Use(
		gin.Recovery(),
		gin.Logger(),
	)

	api := router.Group("/api")

	handlers.RegisterHTTPEndpoints(api, a.financeService)

	a.httpServer = &http.Server{
		Addr:         ":" + port,
		Handler:      router,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	go func() {
		if err := a.httpServer.ListenAndServe(); err != nil {
			log.Fatalf("failed to listen and serve: %+v", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)

	<-quit

	ctx, shutdown := context.WithTimeout(context.Background(), 5*time.Second)
	defer shutdown()

	err := a.dbClient.Close()
	if err != nil {
		slog.Error("close db client error", "err", err)
	}
	return a.httpServer.Shutdown(ctx)
}
