#!/usr/bin/env bash

#
# usage:
#
#       dos2unix   ./install-dev-dependencies.sh       \
#  &&   chmod +x   ./install-dev-dependencies.sh       \
#  &&   sudo       ./install-dev-dependencies.sh
#

# note that 'protoc' and 'thrift' must be installed independently because the
# packages provided by ubuntu are not up-to-date at least in regard to golang

apt-get    install                                     \
	          --assume-yes                             \
              make                                     \
              gnuplot                                  \
              montage                                  \
              graphicsmagick-imagemagick-compat
