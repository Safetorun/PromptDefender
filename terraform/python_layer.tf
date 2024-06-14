resource "aws_lambda_layer_version" "lambda_layer_keep" {
  filename         = data.archive_file.dependencies.output_path
  layer_name       = "${terraform.workspace}-deps-layer"
  compatible_runtimes = ["python3.12"]
  source_code_hash = data.archive_file.dependencies.output_base64sha256
}

resource "aws_lambda_layer_version" "lambda_layer_wall" {
  filename         = data.archive_file.dependencies_wall.output_path
  layer_name       = "${terraform.workspace}-deps-layer-wall"
  compatible_runtimes = ["python3.12"]
  source_code_hash = data.archive_file.dependencies_wall.output_base64sha256
}


data "archive_file" "dependencies" {
  type        = "zip"
  source_dir  = var.dependencies_layer_path
  output_path = "dependencies.zip"
}

resource "aws_lambda_layer_version" "langchain_lambda_layer" {
  filename         = data.archive_file.dependencies_base.output_path
  layer_name       = "${terraform.workspace}-langchain-layer"
  compatible_runtimes = ["python3.12"]
  source_code_hash = data.archive_file.dependencies_base.output_base64sha256
}

data "archive_file" "dependencies_base" {
  type        = "zip"
  source_dir  = var.dependencies_langchain_path
  output_path = "dependencies_base.zip"
}

data "archive_file" "dependencies_wall" {
  type        = "zip"
  source_dir  = var.dependencies_layer_path_wall
  output_path = "dependencies_wall.zip"
}



variable "dependencies_layer_path" {
  type    = string
  default = "../cmd/deps/cmd/lambda_keep_py/dist"
}

variable "dependencies_layer_path_wall" {
  type    = string
  default = "../cmd/deps/cmd/lambda_wall_py/dist"
}

variable "dependencies_langchain_path" {
  type    = string
  default = "../cmd/deps/cmd/langchain_layer/dist"
}
