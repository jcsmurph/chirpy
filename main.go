package main

import (
	"log"
	"net/http"
	"os"

	"github.com/jcsmurph/chirpy/internal/database"
	"github.com/joho/godotenv"

	"github.com/go-chi/chi/v5"
)

type apiConfig struct {
	fileserverHits int
	DB             *database.DB
    jwtSecret string
}

func main() {
	const port = "8080"
	const serverFilePath = "."

    godotenv.Load()
    jwtSecret := os.Getenv("JWT_SECRET")
    if jwtSecret == " "{
        log.Fatal("JWT_Secret env variable is not set")
    }
	db, err := database.NewDB("database.json")

	if err != nil {
		log.Fatal(err)
	}
	apiCfg := apiConfig{
		fileserverHits: 0,
		DB:             db,
        jwtSecret: jwtSecret,
	}
	router := chi.NewRouter()
	fsHandler := apiCfg.middlewareMetricsInc(http.StripPrefix("/app", http.FileServer(http.Dir(serverFilePath))))

    // Index Handlers
	router.Handle("/app", fsHandler)
	router.Handle("/app/*", fsHandler)

    // Metrics Handlers
	apiRouter := chi.NewRouter()
	apiRouter.Get("/healthz", handlerReadiness)
	apiRouter.Get("/reset", apiCfg.handlerReset)

    // Login Handlers
	apiRouter.Post("/login", apiCfg.handlerLogin)

    // Chirp Handlers
	apiRouter.Post("/chirps", apiCfg.handlerChirpsCreate)
	apiRouter.Get("/chirps", apiCfg.handlerChirpsRetrieve)
	apiRouter.Get("/chirps/{chirpID}", apiCfg.handlerChirpsGet)

    // User Handlers
	apiRouter.Post("/users", apiCfg.handlerUsersCreate)
	apiRouter.Put("/users", apiCfg.handlerUsersUpdate)

    // Token Handlers
	apiRouter.Post("/refresh", apiCfg.handlerRefreshToken)
	apiRouter.Post("/revoke", apiCfg.handlerRevokeToken)

	router.Mount("/api", apiRouter)

	adminRouter := chi.NewRouter()
	adminRouter.Get("/metrics", apiCfg.handlerMetrics)
	router.Mount("/admin", adminRouter)

	corsMux := middlewareCors(router)

	server := http.Server{
		Addr:    ":" + port,
		Handler: corsMux,
	}

	log.Printf("Serving on port: %s\n", port)
	log.Fatal(server.ListenAndServe())
}
