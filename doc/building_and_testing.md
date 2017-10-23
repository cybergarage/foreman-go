![foreman_logo](https://raw.github.com/cybergarage/foreman-doc/master/img/icon.png)

# Building and Testing

`foreman-go` is a main package of Foreman which is written by Go, and it depends on the following packages.

- [`foreman-cc`](https://github.com/cybergarage/foreman-cc)
- [`go-graphite`](https://github.com/cybergarage/go-graphite)

To build `foreman-go`, you have to install [`foreman-cc`](https://github.com/cybergarage/foreman-cc) according to the following document at first.

- [`foreman-cc`](https://github.com/cybergarage/foreman-cc)
  -  [Building and Testing](https://github.com/cybergarage/foreman-cc/doc/building_and_testing.md)

About [`go-graphite`](https://github.com/cybergarage/go-graphite), you don't need to install the package directly because it is installed automatically by the following step.

## Setup

```
git clone https://github.com/cybergarage/foreman-go.git
source setup
make test
```