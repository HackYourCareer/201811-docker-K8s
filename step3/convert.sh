#!/usr/bin/env sh

set -e

FONT=FreeMono-Bold

if [ -z "${WORK_DIR}" ]; then
    WORK_DIR="./images"
fi

INPUT_FILE="${WORK_DIR}/maze.txt"
OUTPUT_FILE1="${WORK_DIR}/maze.original.png"
OUTPUT_FILE2="${WORK_DIR}/maze.enhanced.png"

#Convert from TXT to PNG
convert -background white -fill black -font "${FONT}" -pointsize 5 label:@"${INPUT_FILE}" "${OUTPUT_FILE1}"
#Enhance for printing
convert "${OUTPUT_FILE1}" -morphology Erode Octagon "${OUTPUT_FILE2}"
