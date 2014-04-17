package sti

import (
	"github.com/kdar/factorlog"
	"os"
)

const debug_log_fmt = `%{Color "red" "ERROR"}%{Color "yellow" "WARN"}%{Color "green" "INFO"}%{Color "cyan" "DEBUG"}%{Color "blue" "TRACE"}[%{Date} %{Time}] [%{SEVERITY}:%{File}:%{Line}] %{Message}%{Color "reset"}`

const normal_log_fmt = `%{Color "red" "ERROR"}%{Color "yellow" "WARN"}%{Color "green" "INFO"}%{Color "cyan" "DEBUG"}%{Color "blue" "TRACE"}[%{Date} %{Time}] [%{SEVERITY}] %{Message}%{Color "reset"}`

var log = factorlog.New(os.Stdout, factorlog.NewStdFormatter(normal_log_fmt))

func SetLogSeverity(debug bool) {
	if debug {
		log = factorlog.New(os.Stdout, factorlog.NewStdFormatter(debug_log_fmt))
		log.SetMinMaxSeverity(factorlog.TRACE, factorlog.ERROR)
	} else {
		log.SetSeverities(factorlog.INFO | factorlog.WARN | factorlog.ERROR)
	}
}

func Log() *factorlog.FactorLog {
	return log
}

type StiError int

const (
	ErrDockerConnectionFailed StiError = iota
	ErrNoSuchBaseImage
	ErrNoSuchRuntimeImage
	ErrPullImageFailed
	ErrSaveArtifactsFailed
	ErrCreateDockerfileFailed
	ErrCreateContainerFailed
	ErrInvalidBuildMethod
	ErrBuildFailed
	ErrCommitContainerFailed
)

func (s StiError) Error() string {
	switch s {
	case ErrDockerConnectionFailed:
		return "Couldn't connect to docker."
	case ErrNoSuchBaseImage:
		return "Couldn't find base image"
	case ErrNoSuchRuntimeImage:
		return "Couldn't find runtime image"
	case ErrPullImageFailed:
		return "Couldn't pull image"
	case ErrSaveArtifactsFailed:
		return "Error saving artifacts for incremental build"
	case ErrCreateDockerfileFailed:
		return "Error creating Dockerfile"
	case ErrCreateContainerFailed:
		return "Error creating container"
	case ErrInvalidBuildMethod:
		return "Invalid build method - valid methods are: run,build"
	case ErrBuildFailed:
		return "Running /usr/bin/prepare in base image failed"
	case ErrCommitContainerFailed:
		return "Failed to commit built container"
	default:
		return "Unknown error"
	}
}
