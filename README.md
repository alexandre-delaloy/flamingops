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
  - [Deployment](#deployment)
- [Documentation](#documentation)
- [Known Issues](#known-issues)
- [Built With](#built-with)
  - [Other](#other)
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

## Installation
> TODO: Write local installation instructions.

First, do this:
```bash
docker run --rm -it -v ${PWD}:${PWD} -w ${PWD} koddr/cgapp:latest [COMMAND]
```
then this:
```
# even more commands
```

Congratulations. You can now access to ...

Now with your browser of choice get to `127.0.0.1/signup` to create an account. 

Then go to `127.0.0.1/login` to enter your credentials and configure your account.

## Deployment
> TODO: Write deployment instructions.

# Documentation
See the Flamingops [Developer Documentation](https://github.com/blyndusk/flamingops/wiki/) for information on classes and utility functions.

# Known Issues
- 

# Built With
The devops aspect of this project was made possible thanks to:
* [Terraform](https://www.terraform.io/) - Open-source infrastructure as code software tool
* [Docker](https://www.docker.com/) - PaaS product to deliver software in packages

The back-end part of this application is powered by:
* [Golang](https://golang.org/) - Open source programming language

Hat tip to the following Golang dependencies:
* [AWS SDK for Go](https://github.com/aws/aws-sdk-go) - Official AWS SDK for the Go programming language
* [sw sdk]() - Unofficial ScaleWay SDK for the Go programming language
* [Jwt-go](https://github.com/dgrijalva/jwt-go) - Golang implementation of JSON Web Tokens
* [gin-gonic](https://github.com/gin-gonic/gin) - HTTP web framework written in Go

The front-end part of this application is powered by:
* [Create React App](https://create-react-app.dev/) - Officially supported way to create single-page React applications

## Other
* [Amazon Web Services (AWS)](https://aws.amazon.com/) - On-demand cloud computing platforms

# Team Members
* **Alexandre Delaloy** - [blyndusk](https://github.com/blyndusk)
* **Lucas Lehot** - [lucaslehot](https://github.com/lucaslehot)
* **Cyrille Banovsky** - [Ban0vsky](https://github.com/Ban0vsky)
* **Florian Brunet** - [Floadas](https://github.com/Floadas)
* **Armand Benichou** - [ArmandBeni](https://github.com/ArmandBeni)
* **Quentin Maillard** - [Tichyus](https://github.com/Tichyus)
* **Corentin Boulanouar** - [Shawnuke](https://github.com/Shawnuke)

# Acknowledgments

Any resemblance of the Flamingops logo to an already existing logo (living or dead) would be purely coincidental (that's a lie).

# License
This project is licensed under the terms of the [MIT](./flamingops/LICENSE) license.