# CommandFTL

> An online text-based game that will connect you with different players through a network of web components representing universes and ships.

## Introduction

Before further reading, you might want to check the [problem description](https://github.com/Anglepi/CommandFTL/blob/main/docs/ProblemDescription.md) to know what this project is about or check [other documentations](https://github.com/Anglepi/CommandFTL/blob/main/docs/README.md) available about this repository.

The project, as you might have noticed, will be written in Go, and below you will find some choices I have made according to this language and the nature of the project.

## Current state of development

Aside from a few tweaks, this project has basically reached the [first milestone](https://github.com/Anglepi/CommandFTL/milestone/2) that I had set for it.

In the last few days I worked on the functionalities needed to be able to join the game, travel to a sector and shoot a different ship using just HTTP requests, and so far it is working pretty nice.

I also created unit test for all this new code and used Make as task manager to simplify tasks. Technical information about all this can be found [here](https://github.com/Anglepi/CommandFTL/blob/main/docs/TaskManager.md).

### How to try it

For now, you can run `go run main.go` and use your preferred software to perform POST request. Currently available requests are:

- POST: `/new/{shipName}` -> Returns a JSON message containing `{sectorName/shipName}` if succeeds. You will need this information for later.
- POST: `/action/ftl/{currentSector}/{shipName}/{destinationSector}` -> To move from your current sector to a different one.
- POST: `/action/shoot/{currentSector}/{shipName}/{targetName}` -> To attack a target
