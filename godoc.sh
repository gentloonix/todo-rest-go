#!/bin/sh

godoc -http=localhost:8080 & wget -r -np -N -E -p -k http://localhost:8080/pkg/
