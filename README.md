# CommandFTL

> An online text-based game that will connect you with different players through a network of web components representing universes and ships.

## Introduction

Before further reading, you might want to check the [problem description](https://github.com/Anglepi/CommandFTL/blob/main/docs/ProblemDescription.md) to know what this project is about or check [other documentations](https://github.com/Anglepi/CommandFTL/blob/main/docs/README.md) available about this repository.

The project, as you might have noticed, will be written in Go, and below you will find some choices I have made according to this language and the nature of the project.

## Current state of development

Aside from a few tweaks, this project has basically reached the [first milestone]() that I had set for it.

In the last few days I worked on the functionalities needed to be able to join the game, travel to a sector and shoot a different ship using just HTTP requests, and so far it is working pretty nice.

I also created unit test for all this new code and used Make as task manager to simplify tasks. Technical information about all this can be found [here]().

### How to try it

For now, you can run `go run main.go` and use your preferred software to perform POST request. Currently available requests are:

- POST: `/new/{shipName}` -> Returns a JSON message containing `{sectorName/shipName}` if succeeds. You will need this information for later.
- POST: `/action/ftl/{currentSector}/{shipName}/{destinationSector}` -> To move from your current sector to a different one.
- POST: `/action/shoot/{currentSector}/{shipName}/{targetName}` -> To attack a target

## Using docker images for testing

I created a [Dockerfile](https://github.com/Anglepi/CommandFTL/blob/main/Dockerfile) as optimized as I could that contains everything necessary to run all the unit tests.

The base image I chose is `golang:alpine` since it includes just what I need and nothing more, it is a very small image perfect to optimize the size of my image.

It sets the workdir and perform the updates without using the cache and in the same `RUN` instruction, and finally sets the user to the newly created one which has no root permissions and adds the GOPATH which is necessary for it to work.

As you can see, the Dockerfile is really simple, the only thing worth mentioning is `ENV CGO_ENABLED 0`. This is to disable the feature that allows Go packages to call C code. Since it is not necessary for me, I disabled it to optmize everything a little more.

### Deploying the image

This image is currently deployed to both [Dockerhub](https://hub.docker.com/r/anglepi/commandftl) and [Github Container Registry](https://github.com/Anglepi/CommandFTL/pkgs/container/commandftl) (both links to my image at these domains). The process to do this is defined in a [github action](https://github.com/Anglepi/CommandFTL/blob/main/.github/workflows/deploy-test-image.yml) I created with the following workflow:

1. Download the repository to get the image
1. Log in at Dockerhub
1. Extracts the metadata of my repository
1. Build and push the image to Dockerhub
1. Log in at Github Container Registry
1. Build and push the image to GHCR

This image is not built with the code, that is the main reason I do not need to update it with every change in the repository. I specified this action to be launched every time I make changes in my Dockerfile or in this file itself.

## Additional links of interest

- [Previous and extra documentation.](https://github.com/Anglepi/CommandFTL/blob/main/docs/README.md)
- [Task manager selection](https://github.com/Anglepi/CommandFTL/blob/main/docs/TaskManager.md)
- [Testing frameworks research](https://github.com/Anglepi/CommandFTL/blob/main/docs/TestingFramework.md)
- [Assertion libraries research](https://github.com/Anglepi/CommandFTL/blob/main/docs/AssertionLibrary.md)
