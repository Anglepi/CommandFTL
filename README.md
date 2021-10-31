# CommandFTL

> An online text-based game that will connect you with different players through a network of web components representing universes and ships.

## Introduction

Before further reading, you might want to check the [problem description](https://github.com/Anglepi/CommandFTL/blob/main/docs/ProblemDescription.md) to know what this project is about or check [other documentations](https://github.com/Anglepi/CommandFTL/blob/main/docs/README.md) available about this repository.

The project, as you might have noticed, will be written in Go, and below you will find some choices I have made according to this language and the nature of the project.

## Testing

During the next few days I will be focusing on including everything necessary to perform testing on this project in order to easily check code and functionalities as they are included to CommandFTL. Next you will find what choices I have made, and a link for each decision to a page where I state the reason of each choice and some alternatives that I have studied.

### Task manager

The task manager I will be using is **Make**. There are other options such as **Rake** and **GoGradle** and [here](https://github.com/Anglepi/CommandFTL/blob/main/docs/TaskManager.md) I explain why I took this choice.

### Testing framework

You probably already know that Go already has a testing framework, and that is the one I will be using although it is not the only option. Find out more [here](https://github.com/Anglepi/CommandFTL/blob/main/docs/TestingFramework.md).

### Assertion library

I found a lot of different assertion libraries available for Go, and ended up choosing [stretchr/testify/assert](https://pkg.go.dev/github.com/stretchr/testify/assert). Check out why and other options [here](https://github.com/Anglepi/CommandFTL/blob/main/docs/AssertionLibrary.md).

## Additional links of interest

- [Previous and extra documentation.](https://github.com/Anglepi/CommandFTL/blob/main/docs/README.md)
- [Task manager selection](https://github.com/Anglepi/CommandFTL/blob/main/docs/TaskManager.md)
- [Testing frameworks research](https://github.com/Anglepi/CommandFTL/blob/main/docs/TestingFramework.md)
- [Assertion libraries research](https://github.com/Anglepi/CommandFTL/blob/main/docs/AssertionLibrary.md)

