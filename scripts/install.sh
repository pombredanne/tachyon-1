#!/bin/bash

echo 'Building new `tachyon` binary...'
go install

echo 'Installing the plugin...'
cf uninstall-plugin tachyon
cf install-plugin $GOPATH/bin/tachyon
