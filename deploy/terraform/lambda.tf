data "archive_file" "flamingops-lambda" {
  type = "zip"

  source_dir  = "${path.module}/worker"
  output_path = "${path.module}/worker.zip"
}

resource "aws_s3_object" "flamingops-lambda" {
  bucket = aws_s3_bucket.lambda_bucket.id

  key    = "worker.zip"
  source = data.archive_file.flamingops-lambda.output_path

  etag = filemd5(data.archive_file.flamingops-lambda.output_path)
}

resource "aws_lambda_function" "main-function" {
  function_name = "Main"

  s3_bucket = aws_s3_bucket.lambda_bucket.id
  s3_key    = aws_s3_object.lambda_main-function.key

  runtime = "nodejs14.x"
  handler = "main.handler"

  source_code_hash = data.archive_file.lambda_main-function.output_base64sha256

  role = aws_iam_role.lambda_exec.arn
}

resource "aws_cloudwatch_log_group" "main-function" {
  name = "/aws/lambda/${aws_lambda_function.main-function.function_name}"

  retention_in_days = 30
}

resource "aws_iam_role" "lambda_exec" {
  name = "serverless_lambda"

  assume_role_policy = jsonencode({
    Version = "2012-10-17"
    Statement = [{
      Action = "sts:AssumeRole"
      Effect = "Allow"
      Sid    = ""
      Principal = {
        Service = "lambda.amazonaws.com"
      }
      }
    ]
  })
}

resource "aws_iam_role_policy_attachment" "lambda_policy" {
  role       = aws_iam_role.lambda_exec.name
  policy_arn = "arn:aws:iam::aws:policy/service-role/AWSLambdaBasicExecutionRole"
}
