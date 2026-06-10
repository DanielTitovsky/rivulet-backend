package main

import (
	"context"
	"fmt"
	"log"
	"os/signal"
	"syscall"

	app_loger "github.com/DanielTitovsky/rivulet-backend.git/internal/app/loger"
	app_postgres_pool "github.com/DanielTitovsky/rivulet-backend.git/internal/app/repository/postgres/pool"
	app_http_server "github.com/DanielTitovsky/rivulet-backend.git/internal/app/transport/http/server"
	users_postgres_repository "github.com/DanielTitovsky/rivulet-backend.git/internal/features/users/repository/postgres"
	users_service "github.com/DanielTitovsky/rivulet-backend.git/internal/features/users/service"
	users_transport_http "github.com/DanielTitovsky/rivulet-backend.git/internal/features/users/transport/http"
	"github.com/bytedance/gopkg/util/logger"
	"github.com/joho/godotenv"
	"go.uber.org/zap"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Ошибка загрузки .env файла")
	}

	ctx, cancel := signal.NotifyContext(
		context.Background(),
		syscall.SIGINT,
		syscall.SIGTERM,
	)
	defer cancel()

	loggerCongig, err := app_loger.NewConfig()
	serverConfig := app_http_server.NewConfigMust()

	if err != nil {
		panic(fmt.Errorf("Create logger config: %w", err))
	}

	appLoger, err := app_loger.NewLoger(loggerCongig)

	if err != nil {
		panic(fmt.Errorf("Create logger: %w", err))
	}

	defer appLoger.Close()

	logger.Debug("Initiazling postgres connectin pool")

	pool, err := app_postgres_pool.NewConnectinPool(
		ctx,
		*app_postgres_pool.NewConfigMust(),
	)

	if err != nil {
		logger.Fatal("Failed to init postgres connectin: %w", zap.Error(err))
	}

	defer pool.Close()

	logger.Debug("Initiazling features", zap.String("feature", "Users"))

	userRepo := users_postgres_repository.NewUsersRepository(pool)
	userServide := users_service.NewUserServise(userRepo)

	userTransportHttp := users_transport_http.NewUsersHttpHandler(userServide)
	userRoutes := userTransportHttp.Routers()

	logger.Debug("Initiazling HTTP server")

	httpServer := app_http_server.NewHttpServer(serverConfig, appLoger)
	apiVersionRoute := app_http_server.NewApiVersinRouter(app_http_server.ApiVersion1, &httpServer.ServerGin.RouterGroup)

	httpServer.RegisterRouters(apiVersionRoute, userRoutes...)

	httpServer.Run(ctx)
}
