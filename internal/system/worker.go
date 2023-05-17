package system

import (
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
		Setup() error
		Start() error
		Teardown() error
		Stop() error
	}
)

type (
	SimpleWorker struct {
		name     string
		log      log.Logger
		cfg      *config.Config
		didSetup bool
		didStart bool
	}
)

func NewWorker(name string, cfg *config.Config, log log.Logger) *SimpleWorker {
	name = GenName(name, "worker")

	return &SimpleWorker{
		name: name,
		cfg:  cfg,
		log:  log,
	}
}

func (sw SimpleWorker) Name() string {
	return sw.name
}

func (sw SimpleWorker) SetName(name string) {
	sw.name = name
}

func (sw SimpleWorker) Log() log.Logger {
	return sw.log
}

func (sw SimpleWorker) Cfg() *config.Config {
	return sw.Cfg()
}

func (sw SimpleWorker) Setup() error {
	sw.Log().Info("Setup")
	return nil
}

func (sw SimpleWorker) Start() error {
	sw.Log().Info("Start")
	return nil
}

func (sw SimpleWorker) Teardown() error {
	sw.Log().Info("Teardown")
	return nil
}

func (sw SimpleWorker) Stop() error {
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
