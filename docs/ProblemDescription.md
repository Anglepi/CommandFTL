# CommandFTL

> An online text-based game that will connect you with different players through a network of web components representing universes and ships.

## The problematic

Since it is just a game, I will try to help and relieve some of that Sunday afternoon boredom.

Text-based games are no longer mainstream, they are considered a classic or even an antique, but in this case it is online.
I am not pretending it is the only one of its kind, because it is not, but in this case the game will be built from a set of web components that will represent different aspects of it.

Users will interact directly with a webs component that will represent their character or spaceship, and I know this is not the most comfortable way to play a game, so I will be providing with a very basic script that will do as CLI to the game (it is not part of it), allowing players to perform specified request by using simplified commands, and to display any information necessary such as feedback of their actions or changes in their environment.

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

## How will I build it?

Since it is a completely online game, all of its components will be developed in the cloud. As I said before, the only thing users will need to play this game is a way to build requests to the cloud.

I will create different components that will interact with each other in order to make everything works:

1.  The communication with the player. It is a text-based game, the player has to be able to send their commands and get information of what is going on in their surroundings.
1.  The result of their actions. This is highly related to the previous point. Actions have consequences and that is something you learn when you reach adulthood. Shooting a weapon, fixing your ship's systems, buying equipment and every other kind of possible interactions with your ship, other players or the environment will have consequences that will be processed in the cloud.
1.  The management of the environment. A game with just players and no world could probably do better. I am planning to add some random events such as asteroid rain, supernovas and other not so natural phenomena that will have effects on players.
    The world is not just disasters happening again and again, players could find with resources such as abandoned ships or NPCs in search of assistance that will reward them with weapons or additional score, maybe even credits if I am willing to implement some kind of store or trading between players.

Someone might think that, if this works, why don't everybody build games this way? One answer might be because latency has a great impact in the gameplay (Google Stadia failed because of this, among other reasons) and the infrastructure required to run these modern graphics and complex games is not affordable to build in the cloud, the servers would be too expensive.

So, why build this game entirely in the cloud? It has its advantages:

1. You do not need a powerful machine to run its logic and it does not even have graphics, it is a low computing cost game.
1. It is an online game, it makes sense that at least some of its components are on the web.
1. And again, its a text based game, people need more than a few miliseconds to correctly type a command. The game will be designed to give more importance to resource management and strategy than to low reaction times. It is still to be determined, but my guess is that latency will not be an issue.

## References

- [Git and repository configuration (In Spanish, from another project).](https://github.com/Anglepi/Aura/blob/main/docs/configuracion_git.md)
- [FTL:Faster Than Light. A game I like and which from I took some inspiration.](https://subsetgames.com/ftl.html)
