# GoodBlast API

API for managing and conducting gaming tournaments.

## Table of Contents

- [Introduction](#introduction)
- [Getting Started](#getting-started)
  - [Prerequisites](#prerequisites)
  - [Installation](#installation)
- [Usage](#usage)
- [Project Structure](#project-structure)
- [Configuration](#configuration)
- [Deployment](#deployment)
- [Notes](#notes)

## Introduction

GoodBlast API is designed to simplify the management and execution of gaming tournaments. Allows users to create, enter tournaments and track their progress and view their rankings via the leaderboard.

## Getting Started

### Prerequisites

Before running the GoodBlast API, make sure you have the following installed:

- Go
- Redis

### Installation

Clone the repository and install the necessary dependencies:

```bash
git clone https://github.com/kerem-ozt/GoodBlast_API.git
cd GoodBlast_API
go run main.go
```

## Usage

The postman collection of the study is given in the project file. The Swagger document of the project is also available (http://13.51.6.67:3002/swagger/index.html#/), but due to limited time, not all endpoints are working yet. However, it can be controlled for the purpose of gaining insight.

## Project Struct

The project consists of 3 main parts: controllers, services and routes, which are frequently used. Services encapsulate the core logic of your application, controllers handle HTTP requests and responses, and routes define the structure and mapping of API.  There is also a model folder that contains the mongodb models and the request and response models for validation. And a separate middleware folder for middlewares and validations. Together, they contribute to a clean and modular architecture.

├── Dockerfile
├── LICENSE
├── README.md
├── controllers
│   ├── auth.go
│   ├── leaderBoard.go
│   ├── ping.go
│   ├── tournament.go
│   └── user.go
├── docker-compose.yaml
├── docs
│   ├── docs.go
│   ├── swagger.json
│   └── swagger.yaml
├── go.mod
├── go.sum
├── logs
│   └── access.log
├── main.go
├── middlewares
│   ├── auth.go
│   ├── cors.go
│   ├── logger.go
│   ├── recovery.go
│   └── validators
│       ├── auth.validator.go
│       ├── base.validator.go
│       ├── tournament.validator.go
│       └── user.validator.go
├── models
│   ├── config.go
│   ├── db
│   │   ├── country.go
│   │   ├── leaderBoard.go
│   │   ├── token.go
│   │   ├── tournament.go
│   │   └── user.go
│   ├── request.go
│   └── response.go
├── routes
│   ├── auth.go
│   ├── leaderBoard.go
│   ├── ping.go
│   ├── router.go
│   ├── tournament.go
│   └── user.go
└── services
    ├── config.service.go
    ├── leaderBoard.service.go
    ├── storage.service.go
    ├── token.service.go
    ├── tournament.service.go
    └── user.service.go

## Configuration

To create an env file containing configuration information, an example env file named .env.example was created in project.

## Deployment

An instance with a Linux operating system was created on an EC2 machine on AWS, and after installing Go and Redis on the machine, the source codes were transferred and run on the machine.

13.51.6.67 public ipv4 address is accessible.

## Notes

This is my first Go project, and I started by choosing a simple boilerplate to facilitate development. Since it's a REST-heavy application, I structured the project similarly to the file system I'm familiar with from Node.js. To speed up development, I chose MongoDB for the database. I implemented a basic Redis implementation. For deployment, I chose AWS to explore a different cloud provider. I created a Dockerfile to build the project image.

My work has been busier than I expected for a few days. So I couldn't allocate the planned time. As a result, I haven't completed the Swagger documentation and couldn't write detailed tests yet.

It was an instructive and enjoyable process for me :)