resource "aws_lambda_function" "example" {
…
function_name = "lambda"
role = aws_iam_role.lambda.arn
tags = {
    "serverless.consul.hashicorp.com/v1alpha1/lambda/enabled" = "true"
}
variables = {
    environment = {
      CONSUL_MESH_GATEWAY_URI = var.mesh_gateway_http_addr
      CONSUL_SERVICE_UPSTREAMS = "static-server:2345:dc1"
      CONSUL_EXTENSION_DATA_PREFIX = "/lambda_extension_data"
    }
}
layers = [aws_lambda_layer_version.consul_lambda_extension.arn]
