resource "aws_iam_role" "lambda" {
name = "lambda-role"
assume_role_policy = <<EOF
{
"Version": "2012-10-17",
"Statement": [
    {
    "Action": "sts:AssumeRole",
    "Principal": {
        "Service": "lambda.amazonaws.com"
    },
    "Effect": "Allow",
    "Sid": ""
    }
]
}
EOF
}
resource "aws_iam_policy" "lambda" {
name        = "lambda-policy"
path        = "/"
description = "IAM policy lambda"
policy = <<EOF
{
"Version": "2012-10-17",
"Statement": [
    {
    "Effect": "Allow",
    "Action": [
        "ssm:GetParameter"
    ],
    "Resource": "arn:aws:ssm:*:*:parameter${local.this_lambdas_extension_data_path}"
    }
]
}
EOF
}
resource "aws_iam_role_policy_attachment" "lambda" {
role    = aws_iam_role.lambda.name
policy_arn = aws_iam_policy.lambda.arn
}
