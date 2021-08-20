#!/bin/sh
sh -c "docker run -d \
                  -p 127.0.0.1:$INPUT_GENERAL_PORT:$INPUT_GENERAL_PORT \
                  -e INPUT_GENERAL_PORT=$INPUT_GENERAL_PORT \
                  -e INPUT_SAVE_PATH=$INPUT_SAVE_PATH \
                  -v $INPUT_SAVE_PATH:$INPUT_SAVE_PATH \
                  --restart always \
                  --name smuggler \
                  ghcr.io/autamus/smuggler:latest"
