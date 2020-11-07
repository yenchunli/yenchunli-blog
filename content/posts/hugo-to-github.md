---
title: "Push hugo to Github Page"
date: 2020-04-12T11:39:20+08:00
draft: false

tags: [ "Hugo", "Github Page" ]

summary: "We push our hugo website to Github Page."
---

## Push to Github Pages

**!!!: Remember to stop hot-loading test server before pushing to Github.**

There are two different types of `Github Page`.

1. User/Organization Page
2. Project Page

For `User/Organization Page`, we need to create two github repository. One for hugo, the other one is for static files. For `Project Page`, we need to add `publicDir=docs` to `config.toml`.

Here we demo how to create `User/Organization Page`.

### 1. Create two git repository

1. `412988937/412988937.github.io`    --> For static files
2. `412988937/412988937-blog`        --> For hugo files

### 2. Create `deploy.sh` at **hugo-project-dir**

```bash
#!/bin/sh

# If a command fails then the deploy stops
set -e

printf "\033[0;32mDeploying updates to GitHub...\033[0m\n"

# Build the project.
hugo # if using a theme, replace with `hugo -t <YOURTHEME>`

# Go To Public folder
cd public

# Add changes to git.
git add .

# Commit changes.
msg="rebuilding site $(date)"
if [ -n "$*" ]; then
	msg="$*"
fi
git commit -m "$msg"

# Push source and build repos.
git push origin master
```

### 2. Add public folder as Git Submodule and root folder

In `deploy.sh`, it run `hugo` to generate static files to `public/`. Then it will push whole files to `412988937.github.io.git`. Finally we can see our website in [https://412988937.github.io](https://412988937.github.io)

```bash
cd <hugo-project-dir>
git init
git submodule add -b master https://github.com/412988937/412988937.github.io.git public
```

Note: we can add `hugo-project-dir` to git respository for tracking.

### 3. Publish to Github Page when we create new things

```bash
sh deploy.sh
```