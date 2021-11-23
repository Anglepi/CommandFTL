# CommandFTL

> An online text-based game that will connect you with different players through a network of web components representing universes and ships.

## Introduction

Before further reading, you might want to check the [problem description](https://github.com/Anglepi/CommandFTL/blob/main/docs/ProblemDescription.md) to know what this project is about or check [other documentations](https://github.com/Anglepi/CommandFTL/blob/main/docs/README.md) available about this repository.

The project, as you might have noticed, will be written in Go, and below you will find some choices I have made according to this language and the nature of the project.

If you want to check the current state of the project, click [here](https://github.com/Anglepi/CommandFTL/blob/main/docs/StateOfDevelopment.md).

## Including Continuous Integration to the repository

The last few days I have been working on including Continuous Integration for the repository, allowing to run the tests every time the code is updated, providing a way to see if the current code of the repository is working.

To do this I have included two different mechanisms, [Github Actions](https://github.com/Anglepi/CommandFTL/blob/main/.github/workflows/run-tests.yml) and [Travis CI](https://travis-ci.org/), which is configured in [this file](https://github.com/Anglepi/CommandFTL/blob/main/.travis.yml).

### Github Actions

I worked previously with github actions (to update the docker image, explained [here](https://github.com/Anglepi/CommandFTL/blob/main/docs/BuildingDockerImage.md)), so creating a github action is not new for me.

I created [this action](https://github.com/Anglepi/CommandFTL/blob/main/.github/workflows/run-tests.yml), which will run every time I update existent code or another github action, and every time a PR is made into master.

The workflow of the action is the following:

1. Set up the GOPATH environment variable which is required for it to work.
1. Change the working directory. The reason behind this is explained later on.
1. Perform a checkout of the repository in the new working directory.
1. Define the different go environments to do the tests, with different versions, from an old one to one of the newest.
1. Run the tests in the docker image by using `` -t -v `pwd`:/app/test anglepi/commandftl ``

I had to move from the default working directory since apparently there was a `go.mod` file already existing there, and it seems to be a common issue, so when trying to checkout the repository it [fails](https://github.com/Anglepi/CommandFTL/runs/4289541104?check_suite_focus=true). After doing some research, I found out that the best solution for this was moving to a different folder.

### Travis CI

Travis requires you to sign up for a free plan, which does not seem to be free at all since they ask for a payment method and charged me 0.92€ that has not been refunded to me so far.

Once this is done, following their [tutorial](https://github.com/Anglepi/CommandFTL/blob/main/.travis.yml), I linked my Github Account with Travis, and configured it so it has access to this repository.

The next step is to create a [`.travis.yml` file](https://github.com/Anglepi/CommandFTL/blob/main/.travis.yml), where I have specified the language, different versions of it and the script necessary to run the tests.

As you can see, the file is really simple, and it is more than enough for Travis to do its work.

There is also a way to use docker images for the testing, explained [here](https://docs.travis-ci.com/user/docker/), that seems pretty simple as well. A proposal of the file to use such configuration would be:

```
language: go

services:
  - docker

before_install:
- docker pull anglepi/commandftl

script:
- docker run -t -v `pwd`:/app/test anglepi/commandftl
```

Though I have to say I have not tried it.

Having this setup, you could find in the situation where you want to upload changes to your repository that does not affect your code, but adding the changes will run the tests and consume credits. To skip it, just add `[skip travis]` in your commit message.

## Additional links of interest

- [Previous and extra documentation.](https://github.com/Anglepi/CommandFTL/blob/main/docs/README.md)
- [Task manager selection](https://github.com/Anglepi/CommandFTL/blob/main/docs/TaskManager.md)
- [Testing frameworks research](https://github.com/Anglepi/CommandFTL/blob/main/docs/TestingFramework.md)
- [Assertion libraries research](https://github.com/Anglepi/CommandFTL/blob/main/docs/AssertionLibrary.md)
