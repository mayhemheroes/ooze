package ooze

import (
	"os"
	"testing"

	"github.com/gtramontina/ooze/internal/cmdtestrunner"
	"github.com/gtramontina/ooze/internal/color"
	"github.com/gtramontina/ooze/internal/consolereporter"
	"github.com/gtramontina/ooze/internal/fsrepository"
	"github.com/gtramontina/ooze/internal/fstemporarydir"
	"github.com/gtramontina/ooze/internal/gotextdiff"
	"github.com/gtramontina/ooze/internal/ignoredrepository"
	"github.com/gtramontina/ooze/internal/iologger"
	"github.com/gtramontina/ooze/internal/laboratory"
	"github.com/gtramontina/ooze/internal/ooze"
	"github.com/gtramontina/ooze/internal/prettydiff"
	"github.com/gtramontina/ooze/internal/scorecalculator"
	"github.com/gtramontina/ooze/internal/testingtlaboratory"
	"github.com/gtramontina/ooze/internal/verboselaboratory"
	"github.com/gtramontina/ooze/internal/verbosereporter"
	"github.com/gtramontina/ooze/internal/verboserepository"
	"github.com/gtramontina/ooze/internal/verbosetemporarydir"
	"github.com/gtramontina/ooze/internal/verbosetestrunner"
	"github.com/gtramontina/ooze/internal/viruses"
	"github.com/gtramontina/ooze/internal/viruses/floatdecrement"
	"github.com/gtramontina/ooze/internal/viruses/floatincrement"
	"github.com/gtramontina/ooze/internal/viruses/integerdecrement"
	"github.com/gtramontina/ooze/internal/viruses/integerincrement"
	"github.com/gtramontina/ooze/internal/viruses/loopbreak"
)

var defaultOptions = Options{ //nolint:gochecknoglobals
	Repository:               fsrepository.New("."),
	TestRunner:               cmdtestrunner.New("go", "test", "-count=1", "./..."),
	TemporaryDir:             fstemporarydir.New("ooze-"),
	MinimumThreshold:         1.0,
	Parallel:                 false,
	IgnoreSourceFilesPattern: nil,
	Viruses: []viruses.Virus{
		floatdecrement.New(),
		floatincrement.New(),
		integerdecrement.New(),
		integerincrement.New(),
		loopbreak.New(),
	},
}

func banner(log ooze.Logger) {
	border := color.Yellow
	text := color.Cyan

	log.Logf(border("┏━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━┓"))
	log.Logf("%[1]s%[2]s%[1]s", border("┃"), text("┌──────┐  ┌──────┐  ┌──────┐  ┌──────┐"))
	log.Logf("%[1]s%[2]s%[1]s", border("┃"), text("│      │  │      │  ┌──────┘  ├─────  "))
	log.Logf("%[1]s%[2]s%[1]s", border("┃"), text("└──────┘  └──────┘  └──────┘  └──────┘"))
	log.Logf(border("┗━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━┛"))
	log.Logf("Running…")
}

func Release(t *testing.T, options ...Option) {
	t.Helper()

	opts := defaultOptions
	for _, option := range options {
		opts = option(opts)
	}

	var logger ooze.Logger = iologger.New(os.Stdout)

	var reporter ooze.Reporter = consolereporter.New(
		logger,
		prettydiff.New(gotextdiff.New()),
		scorecalculator.New(),
		opts.MinimumThreshold,
	)

	if opts.IgnoreSourceFilesPattern != nil {
		opts.Repository = ignoredrepository.New(opts.IgnoreSourceFilesPattern, opts.Repository)
	}

	if testing.Verbose() {
		opts.Repository = verboserepository.New(t, opts.Repository)
		opts.TemporaryDir = verbosetemporarydir.New(t, opts.TemporaryDir)
		opts.TestRunner = verbosetestrunner.New(t, opts.TestRunner)
		reporter = verbosereporter.New(t, reporter)
	}

	var lab ooze.Laboratory = laboratory.New(opts.TestRunner, opts.TemporaryDir)
	if testing.Verbose() {
		lab = verboselaboratory.New(t, lab)
	}

	t.Cleanup(func() {
		t.Helper()
		res := reporter.Summarize()
		if !res.IsOk() {
			t.FailNow()
		}
	})

	lab = testingtlaboratory.New(t, lab, opts.Parallel)

	banner(logger)

	ooze.New(opts.Repository, lab, reporter).Release(
		opts.Viruses...,
	)
}
