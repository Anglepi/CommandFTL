# Assertion Library used

There are [plenty of assertion libraries](https://pkg.go.dev/search?q=assert&m=) available for Go, being the most popular one [stretchr/testify/assert](https://pkg.go.dev/github.com/stretchr/testify/assert)

It offers a wide selection of assertion types and functions, some of them seems to work with HTTP Requests which will come in real handy when testing a Cloud application such as CommandFTL.

It only needs to be imported, and needs to be use alongside with Go testing package which I will be using anyways, so it is not a problem.

## Alternatives

As I mentioned before, there are a lot of alternatives, but they are by far less popular and do not provide anything that makes them better than my choice in any aspect. In fact, I could not (as of right now) of anything that I could miss in `stretchr/testify/assert` or that could improve it.