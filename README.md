# House

This tool aims to facilitate the integration of many Git repositories into one
source folder.

## Structure ##

A house is a folder containing at least a `src` folder with one's working
repositories. For example: suppose you work with the
`github.com/ishiikurisu/moneylog`, `github.com/ishiikurisu/moneyweb`, and
`5apps.com/pocket_monster/pikachu` repositories. Your house folder will look
like something in these lines:

```
/makes
    money.makefile # Containing build instructions for moneyweb
/src
    /github.com
        /ishiikurisu
            /moneylog
                # Files from this repo
            /moneyweb
                # Files from this repo
    /5apps.com
        /pocket_monster/pikachu
            # Files from this repo
```

This way, you can call house's tools from its source location. Every folder
like this can be a house. Each repo can also contain a `house.yml` file to
store its own configurations.

## Tools ##

To pull the current repository state from internet, call:
```
house load <repo/name>
```
This will call `git pull origin master` on its parameter folder. If no folder
is provided, house will run these calls on the current folder.

To upload data to internet, call:
```
house upload <repo/name> [(-m <message>)]
```
This will call `git add -A`, `git commit`, and `git push origin master` on the
given folder. If no folder is provided, house will run these calls on the
current folder. To add already append a message to the commit, use the `-m` flag
to insert a message

To build your project based on its configuration file, call:
```
house build <repo/name>
```
There is no default call for build therefore it must configured if you want to
use this feature.

To edit the current project, run:
```
house edit <repo/name>
```
This will call the editor that is set on the configuration file.

To execute a command from the project, run:
```
house execute <repo/name> [(-a <arguments>...)]
```
Each argument is a pair:

- The first item is the name of the argument preceded by an "at" (@) symbol
- The second item is the value of the argument. Strings must be between
  quotation marks.

## Configuration file ##

Every repository in a house can contain a file named `house.yml`. Exactly
`house.yml`. In fact this is a YAML file containing specifications for each
tool:

``` yaml
---
build:
  local: true
  commands:
  - make
execute:
  local: true
  commands:
  - ./bin/main @op @where
  - ./par/cmd @op $
edit:
  edit: atom
```

This snippet will configure your house to go the repository where this file is
and run `make try` when you run `house build <repo/name>`.
Or to start `atom .` when you run `house edit <repo/name>`.
Or to execute a command from the repo when you run
`house execute <repo/name> -a <arguments>...`.

Build and execute commands can be executed in parallel: just put a dollar sign
`$` at the end of the line to make it run on background.
