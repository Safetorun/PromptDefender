#!/bin/bash

# Get the Python runtime version
runtime=$(python --version 2>&1 | cut -d ' ' -f 2 | cut -d '.' -f 2)

# Create a new directory for the Python site-packages
mkdir -p cmd/deps/cmd/lambda_wall_py/dist/python/lib/python3.${runtime}/site-packages/prompt_defender_client || exit 1

cp -r cmd/client/prompt_defender_client/* cmd/deps/cmd/lambda_wall_py/dist/python/lib/python3.${runtime}/site-packages/prompt_defender_client || exit 1