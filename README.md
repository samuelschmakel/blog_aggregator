# gator

gator is a command-line blog aggregator that allows users to save and browse feeds, follow or unfollow feeds created by others, and manage their own personalized collection of blog posts. Built with Go and PostgreSQL, gator provides a lightweight and efficient way to keep track of your favorite blogs.

## Features

- Save and browse your favorite blog feeds.
- Follow or unfollow other users’ feeds.
- Simple and intuitive CLI commands.
- Store configuration data in a JSON file.

## Prerequisites

Before installing gator, make sure you have the following dependencies installed on your system:

### 1. Go Programming Language

Download and install Go by following the instructions on the [official Go website](https://golang.org/dl/).

After installation, verify it by running:

```bash
$ go version
```

### 2. PostgreSQL Database

Download and install PostgreSQL from the [official PostgreSQL website](https://www.postgresql.org/download/).

After installation, start the PostgreSQL server and verify it by running:

```bash
$ psql --version
```

## Installation

Install gator using the `go install` command:

```bash
$ go install github.com/your-username/gator@latest
```

Make sure the Go binary directory (usually `$GOPATH/bin`) is added to your system’s PATH.

## Configuration

Create a configuration file for gator at `~/.gatorconfig.json` with the following structure:

```json
{
  "db_url": "postgres://postgres:postgres@localhost:5432/gator?sslmode=disable",
  "current_user_name": "sam"
}
```

Replace `sam` with the username you will create using the `gator register` command.

## Usage

Run gator by invoking its commands from the terminal. Below are the available commands and their descriptions:

### Commands

- \*\*register \*\***`<username>`**
  Register a new user with the specified username.

  Example:

  ```bash
  $ gator register john
  ```

- \*\*login \*\***`<username>`**
  Log in as an existing user.

  Example:

  ```bash
  $ gator login john
  ```

- **reset**
  Resets user data, removing all saved data.

- **users**
  Displays all registered users.

  Example:

  ```bash
  $ gator users
  ```

- **agg**
  Aggregates saved feeds and fetches the latest posts.

  Example:

  ```bash
  $ gator agg
  ```

- \*\*addfeed \*\***`<feed_url>`**
  Adds a new feed for the currently logged-in user.

  Example:

  ```bash
  $ gator addfeed https://example.com/feed
  ```

- **feeds**
  Displays all feeds saved by the current user.

  Example:

  ```bash
  $ gator feeds
  ```

- \*\*follow \*\***`<username>`**
  Follow another user's feeds.

  Example:

  ```bash
  $ gator follow jane
  ```

- **following**
  Displays all the feeds the current user is following.

  Example:

  ```bash
  $ gator following
  ```

- \*\*unfollow \*\***`<feed_id>`**
  Unfollows a specific feed by its ID.

  Example:

  ```bash
  $ gator unfollow 3
  ```

- **browse**
  Browse through the posts saved in your feeds.

  Example:

  ```bash
  $ gator browse
  ```

## Running gator

1. Register a new user using the `register` command.
2. Log in as the registered user.
3. Use the various commands to manage your feeds, follow others, and browse posts.

Enjoy using gator to stay update blogs!
