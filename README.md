# Advent of Code 2021

Using the [Advent of Code 2021](https://adventofcode.com/2021/) puzzles to become more proficient in [Go](https://golang.org/).

### Environment Variables

#### AOC_SESSION_TOKEN

Required for this application to automatically download your user-specific input files for each puzzle. You can find your session token in adventofcode.com's session cookie in your browser.

If not set, the application will not be able to automatically download puzzle inputs. You will have to download each manually and save it to the correct location, e.g. `./day01/input.text`.

#### AOC_LOG_LEVEL

Sets the [logrus](https://github.com/Sirupsen/logrus) log level. Valid values are `panic`, `fatal`, `error`, `warn`, `info`, `debug`, and `trace`. Defaults to `warn` if not set.

### Makefile

This repo's [`Makefile`](./Makefile) automates many useful operations.

#### make setup

Downloads and installs [Go 1.17.3](https://go.dev/dl/) and downloads all of this application's package dependencies.

#### make test

Runs all unit tests.

#### make build

Builds the `advent-of-code-2021` executable in the current directory.

#### make clean

Deletes all build artifacts from the current directory, i.e. the `advent-of-code-2021` executable.

#### make run

Builds the `advent-of-code-2021` executable and runs a single day's solution. Specify the day by setting the `DAY` parameter in the command line. Optionally, you can also set the `LOG_LEVEL` variable, otherwise it will default to `warn`.

```
make run DAY=1 [LOG_LEVEL=trace]
```

#### make run-all

Builds the `advent-of-code-2021` executable and runs all solutions for all days. Optionally, you can also set the `LOG_LEVEL` variable, otherwise it will default to `warn`.

```
make run-all [LOG_LEVEL=trace]
```

#### make install

Builds the `advent-of-code-2021` executable and installs it in the `$GOPATH/bin` directory.

#### make uninstall

Deletes the `advent-of-code-2021` executable from the `$GOPATH/bin` directory.