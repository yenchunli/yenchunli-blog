---
title: "Install Hugo on MacOS"
date: 2020-04-11T11:39:20+08:00
draft: true

tags: [ "Hugo" ]

summary: "`Hugo` is the fastest static website generator in the world. It is really suitable for software developers to generate their personal website. In this tutorial, we start to build up a website which used `Hugo`."
---



## 1. Install go

Go to the [Go Officail Webiste]("https://golang.org/dl/") and download `go`.


## 2. Install brew (OSX)

[Brew Official Website]("https://brew.sh/")
```bash
/bin/bash -c "$(curl -fsSL https://raw.githubusercontent.com/Homebrew/install/master/install.sh)"
```

## 3. Install Hugo
```bash
brew install hugo
```

## 4. Run the command to generate a site
```bash
hugo new site myblog
```

The directory will like:

```bash
myblog
  |- archetypes
  |- assets
  |- content
      |-posts
  |- data
  |- layout
  |- public
  |- static
  |- theme
```

## 5. Install theme

You can download themes from [here]("https://themes.gohugo.io/"). 

In this tutorial we use [hugo-theme-noteworthy]("https://github.com/kimcc/hugo-theme-noteworthy") for example.

```bash
cd <hugo-project-dir>
git clone https://github.com/kimcc/hugo-theme-noteworthy.git themes/noteworthy
```

There are two ways to utilize the theme.

1. Write `theme=noteworthy` to config.toml
2. Copy necessary data to project folder 

For example:
```bash
// Maybe you need to move more data
mv themes/noteworthy/layout layout
mv themes/noteworthy/assets asset
mv themes/noteworthy/static static
```

## 6. Set config.toml

Usually we can take a look from `themes/exampleSite/config.toml` for details.

## 7. Run the hot-load testing server

```bash
hugo server
// You can see your website on http://localhost:1313
```

## 8. Post articles

```bash
hugo new posts/hello.md
```










