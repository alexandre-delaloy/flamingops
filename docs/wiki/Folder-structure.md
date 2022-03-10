# Folder structure

The structure of this project follows [these conventions](https://github.com/golang-standards/project-layout).

- `/.github`: Conventions, template, labels, ci, cd, settings
- `/cmd`: main files
- `/build`: Dockerfiles
- `/config`: Environment files
- `/deploy`: Docker Compose, K8s and Terraform files
- `/internal`: "private" code
- `/pkg`: "public" code
- `/scripts`: Makefile bash scripts for setup/install/start
- `/web`: UI / Webapp files
- `/worker`: Nodejs Lambda