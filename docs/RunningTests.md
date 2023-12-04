---
title: Running tests 
excerpt: Running the unit and integration tests in the project
category: 652f7b9ea6bf5e000bb8dc94
---

# Unit tests

To run all unit tests, simply run 

```shell
make test 
```

In the root directory of the project. This will run all unit tests in the subdirectory of the project.

When adding a new package, make sure that you add this to the makefile as well


# Integration tests

To run the integration tests, you first need to deploy - on any branch other than main run:

```make deploy```

This will deploy your branch into a clean, branch environment. and output the required variables into the terraform/terraform_output.json file


```make integration_test```

