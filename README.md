# House

This tool aims to facilitate the integration of many Git repositories into one
source folder.

```
WARNING
```

This program is being rewritten!

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
like this can be a house. Each repo can also contain a `house.json` file to
store its own configurations and

## Tools ##

To pull the current repository state from internet, call:
```
house load <repo/name>
```
This will call `git pull origin master` on its parameter folder. If no folder
is provided, house will run these calls on the current folder.

To upload data to internet, call:
```
house upload <repo/name>
```
This will call `git add -A`, `git commit`, and `git push origin master` on the
given folder. If no folder is provided, house will run these calls on the
current folder.

To clone a repo to your repository, call:
```
house get <repo/name>
```

To build your project based on its configuration file, call:
```
house build <repo/name>
```
There is no default call for build therefore it must configured if you want to
use this feature.

To start your preferred text editor, call:
```
house edit <repo/name>
```
There is no default text editor to be called as it must configured beforehand.

## Configuration file ##

Every repository in a house can contain a file named `house.json`. Exactly
`house.json`. In fact this is a JSON file containing specifications for each
tool:

``` json
{
    "build": {
        "local": true,
        "command": "make try"
    },
    "edit": {
        "editor": "atom"
    }
}
```

This snippet will configure your house to go the repository where this file is
and run `make try` when you run `house build <repo/name>`; or to run
`atom <repo/name>` when you run `house edit <repo/name>`.
