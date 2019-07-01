# lookout

Lookout is an upstream software respository watcher built for maintaining large collections of up-to-date applications.



## Installation

1. Clone this repository and run

2. `go build`

3. `go install`



## Usage

### Commands

| Command | Alias | Discription                                                                |
| ------- | ----- | -------------------------------------------------------------------------- |
| help    | ?     | General help message for other commands.                                   |
| add     | a     | Add an application entry to the database.                                  |
| search  | s     | Search for an application in the database and retrieve all available data. |
| run     | r     | Run a full update scan on all the application in the database.             |



### Examples

| Command | Example                                     |
| ------- | ------------------------------------------- |
| add     | lookout add [APP_ID] [VERSION] [SOURCE_URL] |
| search  | lookout search [APP_ID]                     |
| run     | lookout run                                 |



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
