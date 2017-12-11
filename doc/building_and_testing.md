![foreman_logo](https://raw.github.com/cybergarage/foreman-doc/master/img/icon.png)

# Building and Testing

`foreman-go` is a main package of Foreman which is written by Go, and it depends on the following packages.

- [`foreman-cc`](https://github.com/cybergarage/foreman-cc)
- [`go-graphite`](https://github.com/cybergarage/go-graphite)

To build `foreman-go`, you have to install [`foreman-cc`](https://github.com/cybergarage/foreman-cc) according to the following document at first.

- [`foreman-cc`](https://github.com/cybergarage/foreman-cc)
  -  [Building and Testing](https://github.com/cybergarage/foreman-cc/blob/master/doc/building_and_testing.md)

About [`go-graphite`](https://github.com/cybergarage/go-graphite), you don't need to install the package directly because it is installed automatically by the following step.

## Dependency Packages

### ANTLR

To build `foreman-go`, you have to install ANTLR v4 later.

- [Getting Started with ANTLR v4](https://github.com/antlr/antlr4/blob/master/doc/getting-started.md)
- [ANTLR4 Language Target, Runtime for Go](https://github.com/antlr/antlr4/blob/master/doc/go-target.md)

#### MacOSX

##### Using Homebrew, you can install the ANTLR easily as the following.

```
brew install antlr
```

## Setup

To setup an initial develop environment for `foreman-go` in your comupter, you have to run the following command at first.

```
git clone https://github.com/cybergarage/foreman-go.git
source setup
make test
```
