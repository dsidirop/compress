SHELL := /bin/bash

include ./make-tasks/utils.*.mk
include ./make-tasks/tasks.*.mk

.DEFAULT_GOAL := benchmark
