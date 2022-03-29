# wowlaunch

Program that launches pre-bnet launcher versions of World of Warcraft and enters a chosen account's username and password

## Installation

1. Download the latest release [from the release page](https://github.com/KarazhanChessClub/wowlaunch/releases/latest). Be sure to pick the correct operating system

2. Extract the zip or tarball into a location of your choice

3. Edit `config.yaml`. Follow the instructions in the comments (lines that start with `#`).

4. If you're on Linux and want to do a proper install, do the following:

    ```bash
    # Move the executable to somewhere on your $PATH
    sudo mv wowlaunch /usr/bin/wowlaunch

    # Make a XDG compliant config directory
    mkdir -p /home/$USER/.config/wowlaunch

    # Move the config file
    mv config.yaml /home/$USER/.config/wowlaunch/
    ```

## Usage

Tell wowlaunch which account name to use and it will do the rest. Wowlaunch must be run from the command line and given the name of the account you wish to launch

```bash
# Using the above example config, launch the 'main' account
wowlaunch main

# Launch the 'second' account
wowlaunch second
```

## Building From Source

1. Clone the repo and build the software

    ```bash
    git clone https://github.com/KarazhanChessClub/wowlaunch
    cd wowlaunch
    go build
    ```

2. Copy the executable to somewhere useful

    ```bash
    # Local bin dir
    cp wowlaunch $HOME/.local/bin

    # there's always /usr/bin on Linux
    sudo cp wowlaunch /usr/bin
    ```

    Windows users can just put it in a directory of their choice
    ```cmd
    mkdir C:\wowlaunch
    cp wowlaunch C:\wowlaunch
    ```

3. Copy the config file to a valid directory and rename it to `config.yaml`:
    ```bash
    # XDG compliant config directory. Recommended for Linux users
    mkdir -p $HOME/.config/wowlaunch
    cp exampleconfig.yaml $HOME/.config/wowlaunch/config.yaml

    # The same directory as the wowlaunch executable. Recommended for Windows users
    cp exampleconfig C:\wowlaunch\config.yaml
    ```

4. Edit the config file, following the instructions in the comments

