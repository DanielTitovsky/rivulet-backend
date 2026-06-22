package main

import (
	"context"
	"fmt"
	"log"
	"os/signal"
	"syscall"

	app_loger "github.com/DanielTitovsky/rivulet-backend.git/internal/app/loger"
	app_oauth "github.com/DanielTitovsky/rivulet-backend.git/internal/app/oauth"
	app_minIo_storage "github.com/DanielTitovsky/rivulet-backend.git/internal/app/repository/minIo/storage"
	app_postgres_pool "github.com/DanielTitovsky/rivulet-backend.git/internal/app/repository/postgres/pool"
	app_postgres_transaction "github.com/DanielTitovsky/rivulet-backend.git/internal/app/repository/postgres/transaction"
	app_http_server "github.com/DanielTitovsky/rivulet-backend.git/internal/app/transport/http/server"
	artist_minio_repository "github.com/DanielTitovsky/rivulet-backend.git/internal/features/artist/repository/minio"
	artist_postgres_repository "github.com/DanielTitovsky/rivulet-backend.git/internal/features/artist/repository/postgres"
	artist_service "github.com/DanielTitovsky/rivulet-backend.git/internal/features/artist/service"
	artist_transport_http "github.com/DanielTitovsky/rivulet-backend.git/internal/features/artist/transport"
	auth_postgres_repository "github.com/DanielTitovsky/rivulet-backend.git/internal/features/auth/repository"
	auth_service "github.com/DanielTitovsky/rivulet-backend.git/internal/features/auth/service"
	auth_transport_http "github.com/DanielTitovsky/rivulet-backend.git/internal/features/auth/transport/http"
	playlis_minio_repository "github.com/DanielTitovsky/rivulet-backend.git/internal/features/playlist/repository/minIo"
	playlist_postgres_repository "github.com/DanielTitovsky/rivulet-backend.git/internal/features/playlist/repository/postgres"
	playlist_service "github.com/DanielTitovsky/rivulet-backend.git/internal/features/playlist/service"
	playlist_transport_http "github.com/DanielTitovsky/rivulet-backend.git/internal/features/playlist/transport"
	tokens_postgres_repository "github.com/DanielTitovsky/rivulet-backend.git/internal/features/tokens/repository/postgres"
	tokens_service "github.com/DanielTitovsky/rivulet-backend.git/internal/features/tokens/service"
	token_transport_http "github.com/DanielTitovsky/rivulet-backend.git/internal/features/tokens/transport/http"
	tracks_minio_repository "github.com/DanielTitovsky/rivulet-backend.git/internal/features/tracks/repository/minIo"
	tracks_postgres_repository "github.com/DanielTitovsky/rivulet-backend.git/internal/features/tracks/repository/postgres"
	tracks_service "github.com/DanielTitovsky/rivulet-backend.git/internal/features/tracks/service"
	tracks_transport_http "github.com/DanielTitovsky/rivulet-backend.git/internal/features/tracks/transport/http"
	users_minio_repository "github.com/DanielTitovsky/rivulet-backend.git/internal/features/users/repository/minIo"
	users_postgres_repository "github.com/DanielTitovsky/rivulet-backend.git/internal/features/users/repository/postgres"
	users_service "github.com/DanielTitovsky/rivulet-backend.git/internal/features/users/service"
	users_transport_http "github.com/DanielTitovsky/rivulet-backend.git/internal/features/users/transport/http"
	"github.com/bytedance/gopkg/util/logger"
	"github.com/joho/godotenv"
	"go.uber.org/zap"
)

//Пправить респонсы чтобы они возвращали нужный тайп
//Возвращать в Routers не только список роутов но и список миделварей которые бы относились ко всем роутам.Так же добавить для возможность присваивать к каэдому роуту отдельный миделвеер
//Нужно на уровне app в repositoey создать подключение к MinIo чтобы генерировать ссылки для получения трека
//Нужно настроить нормальный возраст ошибок. Сейчас не возможно нормально понять где именно происходит ошибка и возвращается дял пользователя ерунда
//Дописать валидацию в домене для user и track а также вообщем добавить нормальную валидацию входных параметров
//Написать для transactionManager config
//Создать Makefile чтобы не дрочить постоянно env переменные

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

	txManager := app_postgres_transaction.NewTransactionManager(pool)

	defer pool.Close()

	storage, err := app_minIo_storage.NewMinioStorage(
		ctx,
		*app_minIo_storage.NewConfigMust(),
	)

	if err != nil {
		logger.Fatal("Failed to init minIo sto: %w", zap.Error(err))
	}

	logger.Debug("Initiazling features", zap.String("feature", "Users"))

	userRepo := users_postgres_repository.NewUsersRepository(pool)
	UserStorageRepo := users_minio_repository.NewUserRepository(storage)
	userServide := users_service.NewUserServise(userRepo, UserStorageRepo, *txManager)

	userTransportHttp := users_transport_http.NewUsersHttpHandler(userServide)
	userRoutes := userTransportHttp.Routers()

	logger.Debug("Initiazling features", zap.String("feature", "Tracks"))

	trackRepo := tracks_postgres_repository.NewTrackRepository(pool)
	trackStorageRepo := tracks_minio_repository.NewTrackRepository(storage)
	trackService := tracks_service.NewTrackServise(trackRepo, *txManager, trackStorageRepo)

	trackTransportHttp := tracks_transport_http.NewTrackHttpHandler(trackService)
	trackRouters := trackTransportHttp.Routers()

	logger.Debug("Initiazling features", zap.String("feature", "Token"))

	tokenRepo := tokens_postgres_repository.NewTokenRepository(pool)
	tokenService := tokens_service.NewTokensServise(tokenRepo, "JOPA", "POPA")

	tokenTransportHttp := token_transport_http.NewTokensHttpHandler(tokenService)
	tokenRouters := tokenTransportHttp.Routers()

	logger.Debug("Initiazling features", zap.String("feature", "Auth"))

	authRepository := auth_postgres_repository.NewAuthRepository(pool)
	authService := auth_service.NewAuthServise(userServide, tokenService, *txManager, authRepository)

	googleOAuthConfig := app_oauth.NewOAuthConfigMust()

	authTransportHttp := auth_transport_http.NewAuthHttpHandler(authService, googleOAuthConfig)
	authRouters := authTransportHttp.Routers()

	logger.Debug("Initiazling features", zap.String("feature", "palylistRouters"))

	palylistRepository := playlist_postgres_repository.NewUsersRepository(pool)
	playlistStorageRepository := playlis_minio_repository.NewPlaylistRepository(storage)
	palylistService := playlist_service.NewPlaylistService(palylistRepository, trackService, playlistStorageRepository)

	palylistTransportHttp := playlist_transport_http.NewPlaylistHttpHandler(palylistService)
	palylistRouters := palylistTransportHttp.Routers()

	logger.Debug("Initiazling features", zap.String("feature", "Artist"))

	artistRepository := artist_postgres_repository.NewArtistRepository(pool)
	artisttStorageRepository := artist_minio_repository.NewArtistStorageRepository(*storage)
	artistService := artist_service.NewArtistService(artistRepository, artisttStorageRepository, *txManager)

	artistTransportHttp := artist_transport_http.NewArtistHttpHandler(artistService)
	artistRouters := artistTransportHttp.Routers()

	logger.Debug("Initiazling HTTP server")

	httpServer := app_http_server.NewHttpServer(serverConfig, appLoger)
	apiVersionRoute := app_http_server.NewApiVersinRouter(app_http_server.ApiVersion1, &httpServer.ServerGin.RouterGroup)

	httpServer.RegisterRouters(apiVersionRoute, userRoutes...)
	httpServer.RegisterRouters(apiVersionRoute, trackRouters...)
	httpServer.RegisterRouters(apiVersionRoute, tokenRouters...)
	httpServer.RegisterRouters(apiVersionRoute, palylistRouters...)
	httpServer.RegisterRouters(apiVersionRoute, authRouters...)
	httpServer.RegisterRouters(apiVersionRoute, artistRouters...)

	httpServer.Run(ctx)
}
