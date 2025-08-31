package main

import (
	"context"
	_ "embed"
	"log/slog"
	"os"
	"strconv"

	"github.com/ndabAP/entityscrape/cases"
	"github.com/ndabAP/entityscrape/cases/nsops"

	"golang.org/x/sync/errgroup"
)

var (
	logger = slog.New(slog.NewTextHandler(os.Stderr, &slog.HandlerOptions{
		Level: slog.LevelDebug,
	}))

	ctx = context.Background()
)

func init() {
	slog.SetDefault(logger)

	_, err := os.Stat("go.mod")
	if os.IsNotExist(err) {
		panic("must be executed in root")
	}

	cwd, err := os.Getwd()
	if err != nil {
		panic(err.Error())
	}
	cases.SetCorpusRootDir(cwd)

	gcloudSvcAccountKey := os.Getenv("GCLOUD_SERVICE_ACCOUNT_KEY")
	if len(gcloudSvcAccountKey) == 0 {
		panic("missing GCLOUD_SERVICE_ACCOUNT_KEY")
	}
	cases.GoogleCloudSvcAccountKey = gcloudSvcAccountKey

	sampleRate := os.Getenv("SAMPLE_RATE")
	if len(sampleRate) > 0 {
		s, err := strconv.ParseUint(sampleRate, 10, 64)
		if err != nil {
			panic(err.Error())
		}
		if s > 100 {
			panic("SAMPLE_RATE must be <= 100, or unset")
		}
		cases.SampleRate = s
	}
}

func main() {
	g, ctx := errgroup.WithContext(ctx)

	// g.Go(func() error {
	// 	// Recover for easier debugging.
	// 	defer func() {
	// 		if r := recover(); r != nil {
	// 			panic(r)
	// 		}
	// 	}()
	// 	return isopf.Conduct(ctx)
	// })
	g.Go(func() error {
		defer func() {
			if r := recover(); r != nil {
				panic(r)
			}
		}()
		return nsops.Conduct(ctx)
	})
	// g.Go(func() error {
	// 	defer func() {
	// 		if r := recover(); r != nil {
	// 			panic(r)
	// 		}
	// 	}()
	// 	return rvomg.Conduct(ctx)
	// })

	if err := g.Wait(); err != nil {
		logger.Error(err.Error())
		os.Exit(1)
	}
}
