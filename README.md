# Lookout

[![Go Report Card](https://goreportcard.com/badge/github.com/alecbcs/lookout)](https://goreportcard.com/report/github.com/alecbcs/lookout)

Lookout is an upstream software respository watcher built for maintaining large collections of up-to-date applications.

## Installation

#### Dependencies

- `GCC`

- `Golang`

#### Build

1. Clone this repository and run

2. `go build`

3. If you've added you're go bin to your system path you can also run `go install` 

## Usage

### Commands

| Command | Alias | Discription                                                                |
| ------- | ----- | -------------------------------------------------------------------------- |
| help    | ?     | General help message for other commands.                                   |
| add     | a     | Add an application entry to the database.                                  |
| search  | s     | Search for an application in the database and retrieve all available data. |
| run     | r     | Run a full update scan on all the application in the database.             |
| import  | i     | Import an application entry (and it's dependencies) to the database.       |

### Examples

#### Commands

| Command | Example                                     |
| ------- | ------------------------------------------- |
| add     | lookout add [APP_ID] [VERSION] [SOURCE_URL] |
| search  | lookout search [APP_ID]                     |
| run     | lookout run                                 |
| import  | lookout import [YAML FILE]                  |

#### YAML Application Import File

```yaml
name: cuppa
version: 1.1.0
source: https://github.com/DataDrake/cuppa/archive/v1.1.0.tar.gz
dependencies: 
    - golang
```

## License

Copyright 2019 Alec Scott <alecbcs@gmail.com>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
