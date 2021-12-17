# CommandFTL

> An online text-based game that will connect you with different players through a network of web components representing universes and ships.

## Introduction

Before further reading, you might want to check the [problem description](https://github.com/Anglepi/CommandFTL/blob/main/docs/ProblemDescription.md) to know what this project is about or check [other documentations](https://github.com/Anglepi/CommandFTL/blob/main/docs/README.md) available about this repository.

The project, as you might have noticed, will be written in Go, and below you will find some choices I have made according to this language and the nature of the project.

If you want to check the current state of the project, click [here](https://github.com/Anglepi/CommandFTL/blob/main/docs/StateOfDevelopment.md).

## Building REST API

This section contains the details of the REST API built for CommandFTL, such as the framework used and the different endpoints and how to use them.

### Toolkit used

I chose [gorilla/mux](https://github.com/gorilla/mux), a toolkit that extends the default Handler from Go to manage my endpoints and build the API.

The main reason behind this choice is that it just extends a bit the already existent functionality from Go, which allows to add more functionalities and frameworks in the future if required, or make any changes without massive effort.

Another reason is that it is simple, the toolkit is not hard to use and very straightforward. Having this very little complexity, and being just a small extension, makes it easy to work with in many aspects such as testing.

#### Alternatives

[Gin](https://github.com/gin-gonic/gin) is a popular framework that stands out for being minimalistic and very fast. It is a very promising option to consider, but I finally rejected it because the future of the project is not completely planned, and if I had to require extra utilities or perform any other changes regarding the API, this framework will make it too much work.

Another alternative to consider is [Martini](https://github.com/go-martini/martini). Unlike Gin, this one seems to work pretty well with other frameworks and utilities, being modular and scalable. The problem I found that made me reject this is that it seems to use its own router which is pretty slow, and the nature of CommandFTL does not allow slow request handling. Also, when I got to check their repository, I saw it is no longer maintained.

[Revel](https://revel.github.io/) is also pretty popular, and has a ton of features and third-party libraries that could come in handy in some situations. This makes the framework very useful when imported, but also very large, and with a lot of features I will never need to use. As I said in a previous stage of the project, I do not need unnecessary functions, I want a lightweight source code.

### Defined endpoints

| Endpoint                 | Method | Body                               | Response Status Code | Description                                                                                                           |
| ------------------------ | ------ | ---------------------------------- | -------------------- | --------------------------------------------------------------------------------------------------------------------- |
| /newGame                 | POST   | {"shipName": "myShipName"}         | 201, 400, 409        | Used to start a new game, creating a ship with the given name. Returns the newly created ship with the initial sector |
| /{sector}/{ship}/weapons | PUT    | {"target": "targetShipName"}       | 200, 400, 422        | Update your weapons systems giving it a target to shoot at. Basically, allows you to shoot at a given target          |
| /{sector}/{ship}/engines | PUT    | {"destination": "neighbourSector"} | 200, 400, 404, 422   | Attempts to use your engines to travel to the specified sector                                                        |

The "main" resource is the sector. Within each sector, we can find ships, and each ship will have systems to access. This is the reasoning behind the structuring of these endpoints.

## Including logs

I have used Go's native framework for logs. It's pretty versatile, and does not require additional modules to install, making it the most lightweight possible. To use it I just built a [small file](https://github.com/Anglepi/CommandFTL/blob/main/logger/logger.go) to encapsulate its functionality and that's it.

Although there are tools and frameworks for logging in Go, such as [logrus](https://github.com/sirupsen/logrus) and [glog](https://github.com/golang/glog), I do not feel necessary to use them, since right now and during the following states of development of the project, I do not require additional features or tools for my logging needs, adding them would just make everything more complex and heavyweight than necessary.

## Additional links of interest

- [Previous and extra documentation.](https://github.com/Anglepi/CommandFTL/blob/main/docs/README.md)
- [Task manager selection](https://github.com/Anglepi/CommandFTL/blob/main/docs/TaskManager.md)
- [Testing frameworks research](https://github.com/Anglepi/CommandFTL/blob/main/docs/TestingFramework.md)
- [Assertion libraries research](https://github.com/Anglepi/CommandFTL/blob/main/docs/AssertionLibrary.md)
