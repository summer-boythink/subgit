## subgit
subgit is a Command line tool for cloning a folder in a git repository

## Why write it
When I use git to clone repositories, I tend to want to keep a subfolder, not the whole thing.

For example, in the following repository, I only need one of the templates, not the entire repository

https://github.com/solidjs/templates

So I developed a little tool to make it easy for me to only get subfolders

Of course, there are many better repositories with similar features on github


## Install
```bash
go install github.com/summer-boythink/subgit
```
```bash
subgit --help
```

## Features
* Select a subfolder to leave behind the git clone
* Select a file to leave behind the git clone
* Rename a folder or file (todo)

## As example

```bash
subgit clone "https://github.com/summer-boythink/autoup.git" -d src/

subgit clone "https://github.com/summer-boythink/autoup.git" -f src/test.js

subgit --storedir /home/ clone "https://github.com/summer-boythink/autoup.git"
```