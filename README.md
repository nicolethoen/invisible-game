# Invisible Game

An adventure game I'm building as a way to practice writing Go and Python.

The main goal is to create an "invisible game": an adventure whose world is
discovered through screen-reader output and interaction, with elements that
aren't apparent visually until the player explores them. I also want the game
to help engineers build first-hand experience with screen-reader tools and
controls. Many engineers are responsible for building accessible experiences
and testing screen-reader usability without having much opportunity to use a
screen reader themselves. The game aims to make that practice engaging and
approachable through play.

Development is starting with a command-line adventure. The game is written in
Go, level data is stored in JSON, and a Python script validates that data. A
dedicated UI and screen-reader-focused interactions will come later.

## Run locally

Run these commands from the project root.

Start the game (defaults to `level_1`):

```sh
go -C go run .
```

Choose a level by ID:

```sh
go -C go run . level_1
```

Validate the level data with Python:

```sh
python3 python/validate_level_data.py
```

The Python validator is the current data check; a separate automated test suite
hasn't been added yet.

## Practice approach

I'm using this project to practice writing Go and Python through small,
incremental changes.

## Project layout

- `go/` — Go game engine
- `python/` — Python level-data validator
- `level_catalog/` — level index and JSON level files
