# CI/CD, release and container registry

This repository provides **commit writting** and **branch naming conventions**, **issues** and **pull request templates**, and a **custom issues labels preset**.

But also **CI/CD** and **release** using [GitHub Actions](https://github.com/features/actions), [GitHub Container Registry](https://github.com/features/packages) and [Docker](https://www.docker.com/).

And finally, a simple **RESTful API**, using [Golang](https://golang.org/), [Postgres](https://www.postgresql.org/) and [Adminer](https://www.adminer.org/), build with [Docker](https://www.docker.com/) and [Docker Compose](https://docs.docker.com/compose/), using a [Makefile](<https://en.wikipedia.org/wiki/Make_(software)>).

## CI
[![GO](https://github.com/blyndusk/flamingops/actions/workflows/ci.go.yml/badge.svg)](https://github.com/blyndusk/flamingops/actions/workflows/ci.go.yml)

The **CI** workflow is located at [.github/workflows/ci.go.yml](https://github.com/blyndusk/flamingops/.github/workflows/ci.go.yml). It's triggered a **each push** on **all branches**.

It consist of:

- **install Golang** on the VM
- get the dependancies **using the cache of other Actions run**
- **lint the files** to check or syntax errors
- **build** the application

## CD

[![DOCKER](https://github.com/blyndusk/flamingops/actions/workflows/cd.docker.yml/badge.svg)](https://github.com/blyndusk/flamingops/actions/workflows/cd.docker.yml)

The **CD** workflow is located at [.github/workflows/cd.docker.yml](https://github.com/blyndusk/flamingops/.github/workflows/cd.docker.yml). It's triggered a **each push** on **`main` branch**.

It consist of:

- **login** into the GitHub container registry (ghcr.io)
- **build and push** the Golang api using the **production Dockerfile** located at [build/pakage/sample-api/Dockerfile](bhttps://github.com/blyndusk/flamingops/build/pakage/sample-api/Dockerfile)

After that, you can check the **pushed container** at: `https://github.com/<username>?tab=packages&repo_name=<repository-name>`

> IMPORTANT: you need to **update the production Dockerfile** with your **username** AND **_repository name_**. Otherwise, there will be errors at push:

```bash
LABEL org.opencontainers.image.source = "https://github.com/<username>/<repository-name>"
```

## Release

[![RELEASE](https://github.com/blyndusk/flamingops/actions/workflows/md.release.yml/badge.svg)](https://github.com/blyndusk/flamingops/actions/workflows/md.release.yml)

The **release** workflow is located at [.github/workflows/md.release.yml](.github/workflows/md.release.yml). It's triggered **manually by user input** at: [Actions > RELEASE](https://github.com/blyndusk/flamingops/actions/workflows/md.release.yml).

> IMPORTANT: you need to set the **image tag** in the action input, for the action to be able to push the docker image and create a release **with a specific version**. The image tag is a [SemVer](https://en.wikipedia.org/wiki/Software_versioning) tag, e.g. `1.0.2`.

It consist of:

- check if the **environment match the branch**
- do the CD (docker) action again, but **with a specific image tag**
- create a release **with the same tag**, filled with the **generated changelog as closed issues since the last release**

After that, you can check the release at `https://github.com/<username>/<repository-name>/releases`.

## Labeler

[![LABELER](https://github.com/blyndusk/flamingops/actions/workflows/ci.labeler.yml/badge.svg)](https://github.com/blyndusk/flamingops/actions/workflows/ci.labeler.yml)

The **labeler** workflow consists in **assigning specific labels** on **pull requests** according to the files that have been modified in the **commits attached to this pull request**.