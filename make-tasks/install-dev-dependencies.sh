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
              openjdk-8-jdk                            \
              gradle                                   \
              graphicsmagick-imagemagick-compat

wget                                                                                                 \
           https://repo1.maven.org/maven2/org/apache/avro/avro-tools/1.11.0/avro-tools-1.11.0.jar    \
           -O ./arena/avro-tools.jar
