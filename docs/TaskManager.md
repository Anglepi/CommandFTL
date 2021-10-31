# Task Manager used

I will be using [make](https://www.gnu.org/software/make/) and here are some reasons to justify this decision:
- It is a simple tool and easy to use that you almost do not need to check out its documentation to know how to use it.
- It not require additional setup in the project, just the `Makefile` file.
- It is a versatile tool, it lets you do many different tasks with different tools.
- And finally, probably because of the previous reasons, it is one of the most popular (if not the most) task managers out there.

## Alternatives

When searching for a task manager, I found out that are a lot of them to choose from, but unfortunately most of them are barely known and not so straightforward to use. The main alternative to Make that I found is GoGradle, but there are other options such as [Rake](https://github.com/ruby/rake) or [Gofer](https://github.com/chuckpreslar/gofer) and many others that are generic task managers.

Again, the most popular ones used in Go projects seems to be Make and GoGradle, followed by a great distance by Rake an others.

### GoGradle

[GoGradle](https://github.com/gogradle/gogradle) is a Gradle plugin that was created to support Go language and its features. It includes a few [predefined tasks](https://github.com/gogradle/gogradle/blob/master/docs/tasks.md).

Although it has a nice documentation, and it uses Gradle which is also very popular, it requires a small setup to get it to work and has a more complex syntax than make, besides some functionalities that I will not be using. I chose make over this one mainly for its simplicity compared to GoGradle.