# Pipdepgenerator

[![Build Status](https://travis-ci.org/dragtor/pipdepgenerator.svg?branch=master)](https://travis-ci.org/dragtor/pipdepgenerator)

Pipdepgenerator is tool to generate python project requirements.txt .

  - Fast build-in parser
  - Dynamic requirement.txt generator 

# New Features!

  - Generate requirement.txt for python project


### Build-Steps:


```sh
$ git clone https://github.com/dragtor/pipdepgenerator
$ cd pipdepgenerator
$ go build -o pipdepgenerator main.go
$ ./pipdepgenerator -p <project-dir> -r <project-dir>/requirements.txt
```

### Todos

 - integrate travic-ci
 - maintaining strict package version
 - pass configuration from config file

License
----

MIT
