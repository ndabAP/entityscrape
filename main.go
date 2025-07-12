package main

import (
	"context"
	"embed"
	_ "embed"
	"log/slog"
	"os"

	"github.com/ndabAP/entityscrape/cases"
	"github.com/ndabAP/entityscrape/cases/nsops"

	"golang.org/x/sync/errgroup"
)

var (
	logger = slog.New(slog.NewTextHandler(os.Stderr, &slog.HandlerOptions{
		Level: slog.LevelDebug,
	}))

	ctx = context.Background()

	//go:embed corpus/*
	corpus embed.FS
)

func init() {
	slog.SetDefault(logger)

	cases.Corpus = corpus
	gcloudSvcAccountKey := os.Getenv("GCLOUD_SERVICE_ACCOUNT_KEY")
	if len(gcloudSvcAccountKey) == 0 {
		panic("missing GCLOUD_SERVICE_ACCOUNT_KEY")
	}
	cases.GoogleCloudSvcAccountKey = gcloudSvcAccountKey
}

func main() {
	g, ctx := errgroup.WithContext(ctx)

	g.Go(func() error {
		return nsops.Conduct(ctx)
	})

	if err := g.Wait(); err != nil {
		logger.Error(err.Error())
		os.Exit(1)
	}
}
