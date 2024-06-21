#!/bin/bash


cd cmd/deps/cmd/embeddings_layer

# Calculate the SHA256 hash of the requirements.txt file
hash=$(shasum -a 256 requirements.txt | awk '{ print $1 }')

# Check if requirements.lock exists
if [ -f "requirements.lock" ]; then
  read -r REQUIREMENTS_HASH < requirements.lock
  if [ "$hash" == "$REQUIREMENTS_HASH" ]; then
    echo "The requirements.txt file has not changed. Skipping."
    exit 0
  fi
fi

python -m venv venv
source venv/bin/activate

pip install -r requirements.txt || exit 1

runtime=$(python --version 2>&1 | cut -d ' ' -f 2 | cut -d '.' -f 2)

mkdir -p python/lib/python3.${runtime}/site-packages/ || exit 1

pip install --platform manylinux2014_x86_64 -t python/lib/python3.${runtime}/site-packages/ --only-binary=:all: -r requirements.txt || exit 1

rm -rf dist || exit 1
mkdir -p dist || exit 1
mv python dist/ || exit 1
# Store the hash in the requirements.lock file
echo $hash > requirements.lock