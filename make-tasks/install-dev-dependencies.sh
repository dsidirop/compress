#!/usr/bin/env bash

# usage:
#
#       dos2unix   ./install-dev-dependencies.sh       \
#  &&   chmod +x   ./install-dev-dependencies.sh       \
#  &&   sudo       ./install-dev-dependencies.sh
#

apt-get    install                                     \
	          --assume-yes                             \
              make                                     \
              gnuplot                                  \
              montage



