#!/usr/bin/env bash

RUN_PATH="${CI_PROJECT_DIR:=/exam}"

$(dirname "$0")/overwrite-readonly.sh $RUN_PATH
$RUN_PATH/score.sh