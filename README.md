<p align="center">
  <img src="logos/mountain.png">
</p>
<h1 align="center">
  Lookout
</h1>

[![Go Report Card](https://goreportcard.com/badge/github.com/alecbcs/lookout)](https://goreportcard.com/report/github.com/alecbcs/lookout)

Lookout is an upstream software respository watcher built for maintaining large collections of up-to-date applications.

## Backstory

As a software maintainer, it can be a difficult and time-consuming task to keep software up-to-date. Most of us achieve this by turning on email notifications or periodically remembering to check a project's releases. However, manually checking releases isn't fool proof and email notifications can quickly become overwhelming. As a result, I wrote Lookout, a simple command line tool to store package information and to help automate the process of checking for upstream project updates.

I would like to say a big thank you to DataDrake for writing CUPPA, the upstream polling assistant library that made the development of Lookout possible.

## Installation

| :exclamation: | It is **HIGHLY recommended** that you [generate a Github Personal Access Key](https://help.github.com/en/articles/creating-a-personal-access-token-for-the-command-line#creating-a-token) and [place it in Lookout's Config](https://github.com/alecbcs/lookout#configuration) if you are using any Github repositories. Otherwise you will likely get a `Not Found` error when adding an application to Lookout. |
| ------------- | ----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |

#### Download Prebuilt Binaries
1. Go to Lookout [Releases](https://github.com/alecbcs/lookout/releases) -->
2. Copy the link to your corresponding OS and Architecture.
3. Run `sudo curl -L "PATH-TO-RELEASE" -o /usr/local/bin/lookout`
4. Run `sudo chmod a+x /usr/local/bin/lookout`
5. (Optional) Run `sudo ln -s /usr/local/bin/ait /usr/bin/lookout`

#### Development Dependencies

- `GCC`

- `Golang`

- If you'd like to use lookout as an installed program you must add your go directory to your [PATH](https://golang.org/doc/install#install).

#### Build

1. With go installed simply run `go get github.com/alecbcs/lookout` 

or

1. Clone this repository and run

2. `go build` (This will build `lookout` into a binary you can add to your `bin`.)

3. If you've added your go bin to your system PATH you can also run `go install`

## Usage

### Commands

| Command           | Alias | Discription                                                                |
| ----------------- | ----- | -------------------------------------------------------------------------- |
| help              | ?     | General help message for other commands.                                   |
| add               | a     | Add an application entry to the database.                                  |
| add-dependency    | ad    | Add an entry, dependency relationship to the database.                     |
| import            | im    | Import an application entry (and it's dependencies) to the database.       |
| info              | in    | Search for an application in the database and retrieve all available data. |
| list              | ls    | List all of the applications in the database.                              |
| remove            | rm    | Remove an entry from the database.                                         |
| remove-dependency | rd    | Remove an entry, dependency relationship from the database.                |
| run               | r     | Run a full update scan on all the application in the database.             |
| upgrade           | up    | Set an entry to the latest version possible.                               |

### Examples

#### Commands

| Command           | Example                                                                          |
| ----------------- | -------------------------------------------------------------------------------- |
| add               | lookout add cuppa 1.1.0 https://github.com/DataDrake/cuppa/archive/v1.1.0.tar.gz |
| add-dependency    | lookout add-dependency cuppa golang                                              |
| import            | lookout import example.yml                                                       |
| info              | lookout info cuppa                                                               |
| list              | lookout list cuppa                                                               |
| remove            | lookout remove cuppa                                                             |
| remove-dependency | lookout remove-dependency cuppa golang                                           |
| run               | lookout run                                                                      |
| upgrade           | lookout upgrade cuppa                                                            |

#### YAML Application Import File

```yaml
name: cuppa
version: 1.1.0
source: https://github.com/DataDrake/cuppa/archive/v1.1.0.tar.gz
dependencies: 
    - golang
    - something
    - something-else
```
## Demo
[![asciicast](https://asciinema.org/a/391933.svg)](https://asciinema.org/a/391933?speed=2)

## Configuration

Lookout's default configuration file is located at `$HOME/.config/lookout/lookout.config`

#### Example Config

```toml
[General]
  Version = "0.0.1"

[Database]
  Path = "$HOME/.config/lookout/apps.db"

[Github]
  Key = "GITHUB-API-KEY"
```

#### Github Config

Github limits the number of requests per day for unauthenticated clients. If you are getting a `Not Found` error when trying add a Github entry to lookout, you'll need to create a Github personal access key and add it to the lookout config. To get a Github key [please follow the documentation on cuppa's Github page](https://github.com/DataDrake/cuppa#github-personal-access-keys).



#### Thanks to [ManyPixels](https://www.manypixels.co/) for the amazing artwork above.



## License

Copyright 2019-2021 Alec Scott <hi@alecbcs.com>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
