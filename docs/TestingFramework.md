# Testing Framework used

Since this project is built in Go, it makes sense to use its native test framework, and it seems to work pretty well. After some research, I have decided to use this one, but there are some alternatives for you to use in case you do not like this.

To use it, you create as many `*__test.go` files as you want, that contain functions with names starting by `Test` and have a testing object as a parameter. An example could be `func TestSomethingHappens(t *testing.T)`. When running `go test`, it identifies all those files and functions, executes them and prints out the result.

## Alternatives

### Goblin

[Goblin](https://github.com/franela/goblin) is a framework whose intention is to provide Mocha-like BDD testing. It requires you to add it to your test dependencies and you're good to go.

I have not chosen Goblin just because Go native frameworks seem more than enough for me and I do not want to include many non-essential stuff in the project.

### Ginkgo

[Ginkgo](https://github.com/onsi/ginkgo) is another Behavior Driven Development testing framework, which is often used with [Gomega](https://github.com/onsi/gomega), and has its own CLI that you need to install in order to use it.

Again, it is not that the framework is bad, but CommandFTL is not a very big project, and although I have some experience in BDD Testing, I do not think these last two alternatives are a must to have in a project like CommandFTL.