# Introduction

source ../golang/setup.source.me

./build.sh

./run.sh 5 5 && ./convert.sh && open images/*.png

./run.sh "hello kuba" && ./convert.sh && open images/*.png

alias view='open images/maze.original.png images/maze.enhanced.png'

