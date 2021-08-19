#!/bin/sh
sh -c "docker run -d \
                  -p $INPUT_GENERAL_PORT:$INPUT_GENERAL_PORT \
                  -e INPUT_GENERAL_PORT=$INPUT_GENERAL_PORT \
                  -e INPUT_SAVE_PATH=$INPUT_SAVE_PATH \
                  -v $INPUT_SAVE_PATH:$INPUT_SAVE_PATH \
                  --restart always \
                  ghcr.io/autamus/smuggler:latest"
