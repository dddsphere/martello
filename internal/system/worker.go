package system

import (
	"context"
	"fmt"
	"hash/fnv"
	"strings"
	"time"

	"github.com/dddsphere/martello/internal/config"
	"github.com/dddsphere/martello/internal/log"
)

type (
	Worker interface {
		Name() string
		Log() log.Logger
		Cfg() *config.Config
		Start(ctx context.Context) error
		Stop(ctx context.Context) error
	}
)

type (
	BaseWorker struct {
		name     string
		log      log.Logger
		cfg      *config.Config
		didSetup bool
		didStart bool
	}
)

func NewWorker(name string, cfg *config.Config, log log.Logger) *BaseWorker {
	name = GenName(name, "worker")

	return &BaseWorker{
		name: name,
		cfg:  cfg,
		log:  log,
	}
}

func (sw BaseWorker) Name() string {
	return sw.name
}

func (sw BaseWorker) SetName(name string) {
	sw.name = name
}

func (sw BaseWorker) Log() log.Logger {
	return sw.log
}

func (sw BaseWorker) Cfg() *config.Config {
	return sw.cfg
}

func (sw BaseWorker) Start(ctc context.Context) error {
	sw.Log().Info("Start")
	return nil
}

func (sw BaseWorker) Stop(ctx context.Context) error {
	sw.Log().Info("Stop")
	return nil
}

func GenName(name, defName string) string {
	if strings.Trim(name, " ") == "" {
		return fmt.Sprintf("%s-%s", defName, nameSufix())
	}
	return name
}

func nameSufix() string {
	digest := hash(time.Now().String())
	return digest[len(digest)-8:]
}

func hash(s string) string {
	h := fnv.New32a()
	h.Write([]byte(s))
	return fmt.Sprintf("%d", h.Sum32())
}
