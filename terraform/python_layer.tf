resource "aws_lambda_layer_version" "lambda_layer" {
  filename            = var.dependencies_layer_path
  layer_name          = "deps-layer"
  compatible_runtimes = ["python3.12"]
  source_code_hash    = filesha256(var.dependencies_layer_path)
}

resource "aws_lambda_layer_version" "langchain_lambda_layer" {
  filename            = var.dependencies_langchain_path
  layer_name          = "langchain-layer"
  compatible_runtimes = ["python3.12"]
  source_code_hash    = filesha256(var.dependencies_langchain_path)
}


variable "dependencies_layer_path" {
  type    = string
  default = "../cmd/deps/cmd/lambda_keep_py/dependencies.zip"
}

variable "dependencies_langchain_path" {
  type    = string
  default = "../cmd/deps/cmd/langchain_layer/dependencies.zip"
}
