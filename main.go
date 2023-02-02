package main

import (
	"context"
	"fmt"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/seantcanavan/lambda_jwt_router/lambda_router"
	"github.com/seantcanavan/learnosity_auth_poc/learnosity"
	"log"
	"net/http"
	"os"
	"time"
)

var router *lambda_router.Router

func init() {
	// This globally sets your local time for GoLang only to UTC so this guarantees
	// that every date that we save to the database is saved in UTC
	time.Local = time.UTC

	router = lambda_router.NewRouter("/api", loggerMiddleware)

	// BuildLearnosityRequest endpoints
	router.Route("POST", "/learnosity/auth", learnosity.AuthLambda)

	learnosity.ConsumerKey = os.Getenv("LEARNOSITY_CONSUMER_KEY")
	learnosity.Domain = os.Getenv("LEARNOSITY_DOMAIN")
	learnosity.ConsumerSecret = os.Getenv("LEARNOSITY_CONSUMER_SECRET")
}

func main() {
	environment := os.Getenv("STAGE")
	if environment == "staging" || environment == "production" {
		lambda.Start(router.Handler)
	} else {
		port := os.Getenv("PORT")
		if port == "" {
			port = "8080"
		}
		fmt.Println(fmt.Sprintf("Ready to listen and serve on port %s", port))
		err := http.ListenAndServe(":"+port, http.HandlerFunc(router.ServeHTTP))
		if err != nil {
			log.Fatalf("%+v", err)
		}
	}
}

func loggerMiddleware(next lambda_router.Handler) lambda_router.Handler {
	return func(ctx context.Context, req events.APIGatewayProxyRequest) (
		res events.APIGatewayProxyResponse,
		err error,
	) {
		// [LEVEL] [METHOD PATH] [CODE] EXTRA
		format := "[%s] [%s %s] [%d] %s"
		level := "INF"
		var code int
		var extra string

		res, err = next(ctx, req)
		if err != nil {
			level = "ERR"
			code = http.StatusInternalServerError
			extra = " " + err.Error()
		} else {
			code = res.StatusCode
			if code >= 400 {
				level = "ERR"
			}
		}

		log.Printf(format, level, req.HTTPMethod, req.Path, code, extra)

		return res, err
	}
}
