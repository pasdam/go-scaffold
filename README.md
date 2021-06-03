# go-scaffold

[![Go Report Card](https://goreportcard.com/badge/github.com/pasdam/go-scaffold)](https://goreportcard.com/report/github.com/pasdam/go-scaffold)
[![CI Status](https://github.com/pasdam/go-scaffold/workflows/Continuous%20integration/badge.svg)](https://github.com/pasdam/go-scaffold/actions)

Command line application that generates files/projects from a template.

The app will ask the values to replace and process the template files to
generate the output.

A template consists of a group of file, of which only the ones that end with
`.tpl` are processed, the rest (including the directory structure) are
preserved.

## Build and install

To build the executable:

```sh
go build -o go-scaffold .
```

To build and install it:

```sh
go install .
```

## Create a template

The first sting to do is to configure the variables to use for the files
generation. This is done by declaring those in a file called `prompts.yaml`
situated into the folder `<project_root>/.go-scaffold`.

The content is like:

```yaml
prompts:
  - name: name
    type: string
    default: My name
    message: Please enter your name
  - name: age
    type: int
    default: 30
    message: Please enter your age
```

The supported types are:

- `string`, for text values;
- `int`, for integer values;
- `bool`, for boolean (true/false) values.

At this point we can create the files to process. For instance lets say we want
to generate a markdown files using the above variables just create `TEST.md` in
the root of the template with the following content:

```md
# Test

Hello, I'm {{.name}}, and I'm {{.age}} years old.
```

The way we use the previously defined variables is:

```text
{{ .<variable> }}
```

so we have to wrap it with curly brackets and write a dot before the variable's
name. You can find more details about how to use template variables in the
[godoc](https://golang.org/pkg/text/template/).

The final template folder structure will be:

```text
.
├── .go-scaffold
│   └── prompts.yaml
└── TEST.md
```

An example can be found in [./examples/hello-world](./examples/hello-world).

## Run

### Command line

To run the program:

```sh
go-scaffold [arguments]
```

Arguments:

- `-o <path>`, `--output <path>`: path of the output dir, if not specified the
  template will be generated in place"; the default value is the current folder
  (working dir);
- `-r`, `--remove-source`: flag to indicate whether remove the template and
  config files, or not; this has effect only if the input and output folder are
  the same, as if they're different those files are not copied at all;
  default value is false;
- `-t <path>`, `--template <path>`: path of the template folder; the default
  value is the current folder.

For instance:

```sh
go-scaffold -t ./ -o /tmp/test-go-template
```

This will use the template in the current folder and generate the files in
`/tmp/test-go-template`.

Note that the application will not copy to the output directory the
`.go-scaffold` and `.git` folders.

## Experimental features

### Init script

After executing a template it's possible to run a script to initialize the newly
created project, for instance to checkout git submodules or run some other
custom commands.

The app will automatically execute the file
`<template_dir>/.go-scaffold/initScript`.

## TODO:

- fix readme and example
- post render script
- manifest structure
- load JSON data
- generate multiple files from a single one (i.e for ranges)
- openapi generator cmd
- validate manifest and values
- add custom data objects
- gitignore file
- template functions
- open files
- subcharts
- override chart template file(s)
