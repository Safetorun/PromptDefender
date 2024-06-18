#!/bin/bash

# Loop over each Python module specified in the PYTHON_PACKAGES environment variable
echo 'Building Python dependencies...'
echo "$1"

 # Change directory to the current Python module
 python_module=$1
 cd ${python_module}

# Calculate the SHA256 hash of the requirements.txt file
hash=$(shasum -a 256 requirements.txt | awk '{ print $1 }')
echo "The hash of requirements.txt is: ${hash}"

# Check if requirements.lock exists
changed=false

if [ -f "requirements.lock" ]; then
  read -r REQUIREMENTS_HASH < requirements.lock
  if [ "$hash" == "$REQUIREMENTS_HASH" ]; then
    echo "The requirements.txt file has not changed. Skipping.";
    changed=false
  else
    changed=true
  fi
else
  changed=true
fi

echo "The requirements.txt file has/has not changed: ${changed}"

if [ "$changed" = true ]; then
  # Remove any existing dist directory and create a new one
   rm -rf dist
   mkdir -p dist

   # Create a new Python virtual environment
   python3 -m venv venv

   # Activate the virtual environment
   source venv/bin/activate

   # Get the Python runtime version
   runtime=$(python --version 2>&1 | cut -d ' ' -f 2 | cut -d '.' -f 2)

   # Create a new directory for the Python site-packages
   mkdir -p python/lib/python3.${runtime}/site-packages/


   # Install the Python packages specified in the requirements.txt file to the site-packages directory
   pip install  --index-url https://test.pypi.org/simple/ --extra-index-url https://pypi.org/simple --platform manylinux2014_x86_64 -t python/lib/python3.${runtime}/site-packages/ --only-binary=:all: -r requirements.txt

   # Create a zip file of the Python directory

   # Remove any existing dependencies in the deps directory and create a new one
   rm -rf ../deps/${python_module}/*
   mkdir -p ../deps/${python_module}

   # Move the dependencies.zip file to the deps directory
   mkdir -p ../deps/${python _module}/dist
   mv python ../deps/${python_module}/dist

   # Store the hash in the requirements.lock file
   echo $hash > requirements.lock
fi
# Remove the Python directory and the dist directory, then create a new dist directory
rm -rf python
rm -rf dist
mkdir ../dist

 # Copy all files to the dist directory
 cp -r * ../dist/

 # Remove the venv and dist directories from the dist directory
 rm -rf ../dist/venv
 rm -rf ../dist/dist
 mv ../dist .