#!/bin/bash

 python_module=$1
 cd ${python_module}

 rm requirements.lock
 rm -rf venv
 rm -rf dist
 cd ../..
 deps_directory=cmd/deps/${python_module}
 rm -rf $deps_directory

