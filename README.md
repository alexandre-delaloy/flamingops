<p align="center">
  <img src="https://github.com/blyndusk/flamingops/blob/main/assets/flamingo_v2.png?raw=true" width="256">
</p>

Flamingops
====
***Cloud Providers Aggregator***

<p align="center">
    <br/>
    <a href="https://github.com/blyndusk/flamingops/releases">
      <img src="https://img.shields.io/github/v/release/blyndusk/flamingops"/>
    </a>
    <a href="https://github.com/blyndusk/flamingops/commits/main">
      <img src="https://img.shields.io/github/release-date/blyndusk/flamingops"/>
    </a>
    <a href="https://github.com/blyndusk/flamingops/issues">
      <img src="https://img.shields.io/github/issues/blyndusk/flamingops"/>
    </a>
    <a href="https://github.com/blyndusk/flamingops/pulls">
      <img src="https://img.shields.io/github/issues-pr/blyndusk/flamingops"/>
    </a>
    <a href="https://github.com/blyndusk/flamingops/blob/main/LICENSE">
      <img src="https://img.shields.io/github/license/blyndusk/flamingops"/>
    </a>
</p>

<p align="center">
  <a href="https://github.com/blyndusk/flamingops/actions/workflows/ci.api.yml">
      <img src="https://github.com/blyndusk/flamingops/actions/workflows/ci.api.yml/badge.svg"/>
    </a>
     <a href="https://github.com/blyndusk/flamingops/actions/workflows/cd.api.yml">
      <img src="https://github.com/blyndusk/flamingops/actions/workflows/cd.api.yml/badge.svg"/>
    </a>
    <a href="https://github.com/blyndusk/flamingops/actions/workflows/ci.ui.yml">
      <img src="https://github.com/blyndusk/flamingops/actions/workflows/ci.ui.yml/badge.svg"/>
    </a>
     <a href="https://github.com/blyndusk/flamingops/actions/workflows/cd.ui.yml">
      <img src="https://github.com/blyndusk/flamingops/actions/workflows/cd.ui.yml/badge.svg"/>
    </a>
</p>

# Project Status
*This repo is a [HETIC school](https://www.hetic.net/) project and its purpose is purely educational.* 

*Feel free to fork the project, but be aware that development might slow down or stop completely at any time, and that we are currently not looking for maintainers or owner.*

# Table of Contents
- [Flamingops](#flamingops)
- [Project Status](#project-status)
- [Table of Contents](#table-of-contents)
- [Overview](#overview)
- [Project Demo](#project-demo)
- [Getting Started](#getting-started)
  - [Requirements](#requirements)
  - [Installation](#installation)
- [Documentation](#documentation)
- [Built With](#built-with)
  - [Ops](#ops)
  - [Back-end](#back-end)
  - [Front-end](#front-end)
- [Team Members](#team-members)
- [Acknowledgments](#acknowledgments)
- [License](#license)

# Overview
Flamingops is a Cloud-based application granting devops an interface to manage all their Cloud instances in a cross-cloud-provider way, hence making it easier to handle your web services consumption.

Current support includes:
* AWS instances
* ScaleWay instances

# Project Demo
*No demo available at the moment.*

# Getting Started
These instructions will get you a copy of the project up and running on your local machine for development and testing purposes. See [deployment](#deployment) section for notes on how to deploy the project on a live system.

## Requirements
* Docker
* Docker Compose
* Golang
* Node
* Make

## Installation
First, do this:
```bash
make init
```
Now with your browser of choice get to `127.0.0.1:3000` to access UI.

And get to `127.0.0.1:3333` to access API.

# Documentation
See the Flamingops [Developer Documentation](https://github.com/blyndusk/flamingops/wiki/) for information on classes and utility functions.


# Built With
## Ops
* [Kubernetes](https://kubernetes.io/) - Open-source system for automating deployment, scaling, and management of containerized applications
* [Terraform](https://www.terraform.io/) - Open-source infrastructure as code software tool
* [Docker](https://www.docker.com/) - PaaS product to deliver software in packages
* [GitHub Actions](https://github.com/features/actions) - Automate all your software workflows, now with world-class CI/CD. Build, test, and deploy your code right from GitHub
* [Go-yave](https://github.com/blyndusk/go-yave) - A Golang project-managed repository template
* [Amazon Web Services (AWS)](https://aws.amazon.com/) - On-demand cloud computing platforms

## Back-end
* [Golang](https://golang.org/) - Open source programming language
* [AWS SDK for Go](https://github.com/aws/aws-sdk-go) - Official AWS SDK for the Go programming language
* [sw sdk]() - Unofficial ScaleWay SDK for the Go programming language
* [Jwt-go](https://github.com/dgrijalva/jwt-go) - Golang implementation of JSON Web Tokens
* [gin-gonic](https://github.com/gin-gonic/gin) - HTTP web framework written in Go

## Front-end
* [Create React App](https://create-react-app.dev/) - Officially supported way to create single-page React applications

# Team Members
* **Alexandre Delaloy** - [blyndusk](https://github.com/blyndusk) - DevOps (SRE) / Repository manager
* **Lucas Lehot** - [lucaslehot](https://github.com/lucaslehot) - Lead Dev Back
* **Cyrille Banovsky** - [Ban0vsky](https://github.com/Ban0vsky) - Lead Dev Front
* **Florian Brunet** - [Floadas](https://github.com/Floadas) - Fullstack (Back-oriented)
* **Armand Benichou** - [ArmandBeni](https://github.com/ArmandBeni) - Dev Front
* **Quentin Maillard** - [Tichyus](https://github.com/Tichyus) - Fullstack (Back-oriented)
* **Corentin Boulanouar** - [Shawnuke](https://github.com/Shawnuke) - Dev Front / Documenter 

# Acknowledgments

Any resemblance of the Flamingops logo to an already existing logo (living or dead) would be purely coincidental (that's a lie).

# License
This project is licensed under the terms of the [MIT](./flamingops/LICENSE) license.