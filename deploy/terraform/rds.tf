resource "aws_db_instance" "default" {
  allocated_storage    = 10
  engine               = "postgres"
  engine_version       = "13.6"
  instance_class       = "db.t3.micro"
  name                 = "flamingops_db"
  username             = var.database_user
  password             = var.database_password
  parameter_group_name = "default.mysql5.7"
  skip_final_snapshot  = true
}