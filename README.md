# CommandFTL

> An online text-based game that will connect you with different players through a network of web components representing universes and ships.

## Project planning

I have planned a total of 5 different MVPs to be developed, each one of them being an increment of the previous one.
All of these MVPs are represented with a milestone each, which are the following:

1. [Creating communications template.](https://github.com/Anglepi/CommandFTL/milestone/1). The idea is to prepare the environment to work on the project with a set of defined endpoints that can be used to listed to request even if they do not process them at all.
1. [Most basic combat.](https://github.com/Anglepi/CommandFTL/milestone/2). This includes the capability of a player to join the game, keep their session and perform a very basic action of attack.
1. [Introducing Systems.](https://github.com/Anglepi/CommandFTL/milestone/3). A set of basic systems and their functionalities will be introduced into each players ships, expanding the commands kit and adding interest to battles. Also, some extra actions such as repairs will be added as well.
1. [Advanced Systems and Management.](https://github.com/Anglepi/CommandFTL/milestone/4). Extra and more advanced (or rather complex) systems will be added into the game. These systems and also new weapons will be obtained as a reward for defeating another player.
1. [World Environment.](https://github.com/Anglepi/CommandFTL/milestone/5). Includes several environmental hazzards that affect a sector and random events such as wreckages, distress calls and other NPC interactions.

Each of the above milestones include some extra information about them in the description.

## User Stories

The next set of User Stories has been created an associated to the first milestone they correspond to. Some of them will be moved between different milestones.

1. [[US1] As a player, I want to be able to join the game from anywhere with access to the internet.](https://github.com/Anglepi/CommandFTL/issues/1)
1. [[US2] As a player, during combat, I want to make use of my Weapon Systems to attack and destroy my target.](https://github.com/Anglepi/CommandFTL/issues/2)
1. [[US3] As a developer, in order to ease manual testing, I want a client interface that allows me to make requests easily.](https://github.com/Anglepi/CommandFTL/issues/3)
1. [[US4] As a player, to make the game more interesting, I want my ship to have different systems that increase my possible actions during the game.](https://github.com/Anglepi/CommandFTL/issues/4)
1. [[US5] As a developer, during every game for each player, I want to establish a session in order to correctly associate each player with their ship.](https://github.com/Anglepi/CommandFTL/issues/6)
1. [[US6] As a player, during or after my ship has been under attack, I want to be able to repair the damages.](https://github.com/Anglepi/CommandFTL/issues/7)

## Entities created

As a base for future development, I have created a few 'classes' with simple data structures and empty methods. These can be found at [cc.yaml](https://github.com/Anglepi/CommandFTL/blob/main/cc.yaml) file, indicating the path of each one of them in the 'entidad' key.

## Additional links of interest

- [Previous documentation.](https://github.com/Anglepi/CommandFTL/blob/main/docs/README.md)
