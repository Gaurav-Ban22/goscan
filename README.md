# GoScan

GoScan is a command-based tool for analysing go programs. 
<br>
I initially planned on making this tool to help explore the concept of an abstract syntax tree in one of my favorite languages.

However, this tool isn't only limited to that, despite being a test. It can find the exported functions, comments, modules, and more in directories, projects, and files.


## Installation

Coming soon! I will release a build and binary for Linux later on.

## Usage

```bash
> goscan [flags] [path]
```
<h3>Flags</h3>

Flags are optional.

```bash
> goscan --module [path]
```
This flag will scan the go files in an entire directory and its subdirectories.

Example when running
```bash
> <alias> --module .
```

![image](https://github.com/Gaurav-Ban22/goscan/assets/69488672/190d18c9-4c8c-408f-836e-420062199f84)

