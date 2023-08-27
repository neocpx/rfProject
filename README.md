# RapidFort Offline Project 
# Docker Container Build and Run Instructions

This repository contains my attempt of Rapid fort offline project.
This is a simple rest api server that returns the file information as json.
Follow the steps below to build and run the container using the provided script.

## Prerequisites

Before you begin, make sure you have the following installed:

- Docker: [Install Docker](https://docs.docker.com/get-docker/)
- Git: [Install Git](https://git-scm.com/book/en/v2/Getting-Started-Installing-Git)

## Getting Started

1. **Clone the Repository**: Clone this repository to your local machine using the following command:

   ```bash
   git clone https://github.com/neocpx/rfProject
   cd rfProject
   ```
2. **Build docker image**: Build the docker image on you local machine using the following command:

    ```bash
    docker bui<t_k€>Ã½ v:lua.cmp.utils.feedkeys.call.run(1)
    dld -t name .
    ```
3. **Run the container**: Run the built image using the command:

    ```bash
    docker run -p 8080:8080 name
    ```
4. **Run from Dockerhub**: Direct pull the latest image from dockerhub using command:
```bash
docker pull neocpx/rfproject
```

## To Test
1. **Run the container or project on local enviourment and run**
```bash
curl -X POST -F "file=@path/to/the/file" http://api-endpoint
```
