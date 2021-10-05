# CommandFTL

> An online text-based game that will connect you with different players through a network of webservices representing universes and ships.

## The problematic

Since it is just a game, I will try to help and relieve some of that Sunday afternoon boredom.

Text-based games are no longer mainstream, they are considered a classic or even an antique, but in this case it is online.
I am not pretending it is the only one of its kind, because it is not, but in this case the game will be built from a set of webservices that will represent different aspects of it.

Users will interact directly with a webservice that will represent their character or spaceship, and I know this is not the most comfortable way to play a game, so I will be providing with a very basic script that will do as CLI to the game (it is not part of it), allowing players to perform specified request by using simplified commands, and to display any information necessary such as feedback of their actions or changes in their environment.

Of course, nobody will be enforced to play with such interface. Users will be able to code their own interface however they want, they can even group different requests in order to perform complex tasks with just a single action.

It will be a completely server-side game, since every single calculation and logic will be done there and a client software is not part of this project. Of course, latency should not be an issue for a game of this nature, but its effect will be eventually analyzed during the developmen and testing of the game.

## Basic concept of the game

So, it is a spaceships multiplayer game. Players will be able to fight each other when they are close enough, and maybe even some secondary task such as scavenge the area in search of resources or upgrading your ships capabilities, but how will it work?

### The universe

Each universe will be divided in a set of named sectors that players can visit, and each sector may have some peculiarities such as resources, plasma storms or collapsing stars that will affect players in it.

### The spaceship

Players will start their game with their spaceship in a random sector of an universe. Each ship will have these properties:

- **Hull Hit Points.** If this number reachs 0, it is game over.
- **Enery Core.** It provides a specific amount of energy that can be freely distributed among the ship's systems.

Here are some examples of systems:

- **Weapons Systems.** Allows the attachment of a certain type of weapons and provides these with power from the Energy Core.
- **Shield System.** Uses energy to mitigate incoming damage.
- **Engine System.** Allows the ship to move. This includes FTL jumps to move to different sectors and maybe even to perform evasive manneuvers.
- **Life Support Systems.** Allows you to breathe within the ship.
- **Sensor Systems.** Provides information about what is in the sector.
- **Hacking Systems.** Allows you to affect the performance of enemy ship's systems.

There might be more systems in the future, and every single one of them will be upgradeable providing some extra benefits.

## References

- [Git and repository configuration (In Spanish, from another project).](https://github.com/Anglepi/Aura/blob/main/docs/configuracion_git.md)
- [FTL:Faster Than Light. A game I like and which from I took some inspiration.](https://subsetgames.com/ftl.html)
