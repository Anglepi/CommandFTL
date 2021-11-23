# CommandFTL

> An online text-based game that will connect you with different players through a network of web components representing universes and ships.

## Introduction

Before further reading, you might want to check the [problem description](https://github.com/Anglepi/CommandFTL/blob/main/docs/ProblemDescription.md) to know what this project is about or check [other documentations](https://github.com/Anglepi/CommandFTL/blob/main/docs/README.md) available about this repository.

The project, as you might have noticed, will be written in Go, and below you will find some choices I have made according to this language and the nature of the project.

## Using docker images for testing

I created a [Dockerfile](https://github.com/Anglepi/CommandFTL/blob/main/Dockerfile) as optimized as I could that contains everything necessary to run all the unit tests.

I want this image to be fast when executing tests. When applying continuous integration mechanisms, most of the platforms that provide these services charges you in some way for execution time, so making smaller images that take less to build will save money.

One of the advantages of CommandFTL so far is that it does not require any additional components to install besides Go, so I can start searching by lightweight images.

There are quite a few [Go docker images](https://hub.docker.com/_/golang). Let's see some interesting tags:

- **Alpine**. This is the most promising one since it is really small, almost everything it has is only the very basic and necessary things to make it work.
- **Buster and Stretch**. These images are built on Debian, so they're probably larger and contain more things that I probably won't need. Of these two images, Buster has a newer version of Debian, so it is more interesting than the other.
- The rest of the version were discarded instantly the moment I found out about what they have. **Bullseye** uses Debian, but is a not yet stable version, **WindowsServerCore** has, you guessed it, Windows Server stuff that will have features I do not need at all, **NanoServer** is a lightweight version of Windows used for development, but again, I don't need Windows in any of its variants for my purpose.

Buster, as expected, is large, a total of 901MB. It already includes make and git, among a lot of other things, but I am looking for something lightweight:

```
FROM golang:buster
LABEL maintainer="Ángel Píñar <angle@correo.ugr.es>"
ENV CGO_ENABLED 0

WORKDIR /app/test

RUN apt-get update && \
    adduser commandftl

USER commandFTL

ENV GOPATH=/home/commandFTL/go

CMD make test
```

Out of curiosity, I tried to get to work the Nanoserver image in a Windows Environment, but I could not even make it work.

The [alpine version](https://github.com/Anglepi/CommandFTL/blob/main/Dockerfile) has a size of 329MB, looking so much better, and by far the best of the options for my case.

So this is the base image I chose, `golang:alpine`, since it is a very small image perfect to optimize the size of it, and I only have to install make and git to accomplish what I want.

It sets the workdir and performs the updates and installs without using the cache and in the same `RUN` instruction, and finally sets the user to the newly created one which has no root permissions and adds the GOPATH which is necessary for it to work.

As you can see, the Dockerfile is really simple, the only thing worth mentioning is `ENV CGO_ENABLED 0`. This is to disable the feature that allows Go packages to call C code. Since it is not necessary for me, I disabled it to optimize everything a little more.

### Deploying the image

This image is currently deployed to both [Dockerhub](https://hub.docker.com/r/anglepi/commandftl) and [Github Container Registry](https://github.com/Anglepi/CommandFTL/pkgs/container/commandftl) (both links to my image at these domains). The process to do this is defined in a [github action](https://github.com/Anglepi/CommandFTL/blob/main/.github/workflows/deploy-test-image.yml) I created with the following workflow:

1. Download the repository to get the image
1. Login at Dockerhub
1. Extracts the metadata of my repository
1. Build and push the image to Dockerhub
1. Login at Github Container Registry
1. Build and push the image to GHCR

This image is not built with the code, that is the main reason I do not need to update it with every change in the repository. I specified this action to be launched every time I make changes in my Dockerfile or in this file itself.

## Additional links of interest

- [Previous and extra documentation.](https://github.com/Anglepi/CommandFTL/blob/main/docs/README.md)
- [Task manager selection](https://github.com/Anglepi/CommandFTL/blob/main/docs/TaskManager.md)
- [Testing frameworks research](https://github.com/Anglepi/CommandFTL/blob/main/docs/TestingFramework.md)
- [Assertion libraries research](https://github.com/Anglepi/CommandFTL/blob/main/docs/AssertionLibrary.md)
