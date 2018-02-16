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

## Testing

### Unit Testing

For the unit test, Foreman conforms the following standard testing of Go language.

- [The Go Programming Language - Package testing](https://golang.org/pkg/testing/)

### Integration Testing

Foreman has an original testing tool, `foremantest`, to create the integration test case more easily than [Apache JMeter](http://jmeter.apache.org). 

```
foremantest a testing utility for Forman.

	NAME
	foremantest

	SYNOPSIS
	foremantest <scenario_file>

	DESCRIPTION
	foremantest is a a testing utility for Forman.

	RETURN VALUE
	  Return EXIT_SUCCESS or EXIT_FAILURE
```

#### Senario File Format

`foremantest` can execute a scenario file at once, and the following is the file format specification.

```
scenario_file = (scenario_lines)+

scenario_lines = (scenario_event_line | comment_line | blank_line)

scenario_event_line = query; http_status_code(; verify_json_path; verify_json_value)
query = FQL STRING
http_status_code = INTEGER
verify_json_path =  '/' json_path_id ('/' json_path_id)*
json_path_id =  (\[[0-9+]\] | [a-zA-Z0-9+]
verify_json_value = STRING | INTEGER | FLOAT

comment_line = ^#*
blank_line = (nothing)
```

The sample scenario file is bellow.

```
SET (m0, 1.0, 1514764800) INTO METRICS;200
SELECT (id) FROM METRICS;200;/[0];m0
SELECT (id) FROM METRICS WHERE id == m0;200;/[0];m0
SELECT * FROM METRICS WHERE id == m0 AND ts >= 1514764800 AND ts <= 1514765100;200;/[0]/datapoints/[0]/[0];1
```
