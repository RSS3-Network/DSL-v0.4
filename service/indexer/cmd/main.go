package main

import (
	"net/http"
	"net/http/pprof"
	"time"

	"github.com/naturalselectionlabs/pregod/common/protocol"
	"github.com/naturalselectionlabs/pregod/common/utils/loggerx"
	"github.com/naturalselectionlabs/pregod/service/indexer/internal/config"
	"github.com/naturalselectionlabs/pregod/service/indexer/internal/server"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"go.uber.org/zap"
)

var rootCommand = cobra.Command{
	Use:           "hub",
	SilenceUsage:  true,
	SilenceErrors: true,
	Version:       protocol.Version,
}

func init() {
	logrus.SetReportCaller(true)
	logrus.SetFormatter(&logrus.TextFormatter{
		FullTimestamp:   true,
		TimestampFormat: time.RFC3339,
	})
}

func main() {
	config.Initialize()

	if err := loggerx.Initialize(string(config.ConfigIndexer.Mode)); err != nil {
		logrus.Fatalln(err)
	}

	srv := server.New(&config.ConfigIndexer)

	// start pprof http server
	go func() {
		server := http.NewServeMux()
		server.HandleFunc("/debug/pprof/", http.HandlerFunc(pprof.Index))
		server.HandleFunc("/debug/pprof/cmdline", http.HandlerFunc(pprof.Cmdline))
		server.HandleFunc("/debug/pprof/profile", http.HandlerFunc(pprof.Profile))
		server.HandleFunc("/debug/pprof/symbol", http.HandlerFunc(pprof.Symbol))
		server.HandleFunc("/debug/pprof/trace", http.HandlerFunc(pprof.Trace))
		logrus.Fatal(http.ListenAndServe("localhost:6060", server))
	}()

	rootCommand.RunE = func(cmd *cobra.Command, args []string) error {
		return srv.Run()
	}

	if err := rootCommand.Execute(); err != nil {
		loggerx.Global().Fatal("indexer execution failed", zap.Error(err))
	}
}
