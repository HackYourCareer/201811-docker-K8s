#!/usr/bin/env bash

for d in step*; do

	pushd "$d"
	pwd
	rm -f maze
	rm -f maze-worker
	rm -rf images/*
	rm -rf generated/*
	popd

done
