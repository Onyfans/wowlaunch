# wowlaunch

Program that launches pre-bnet launcher versions of World of Warcraft and enters a chosen account's username and password

## Installation

1. Clone the repo and build the software

    ```bash
    git clone https://github.com/KarazhanChessClub/wowlaunch
    cd wowlaunch
    go build
    ```

2. Copy the executable to somewhere on your `$PATH`

    ```bash
    # Local bin dir
    cp wowlaunch $HOME/.local/bin
    
    # there's always /usr/bin
    sudo cp wowlaunch /usr/bin
    ```

3. Copy the config file to a valid directory and rename it to `config.yaml`:
    ```bash
    # XDG compliant config directory. This is the recommended option
    mkdir -p $HOME/.config/wowlaunch
    cp exampleconfig.yaml $HOME/.config/wowlaunch/config.yaml
    
    # The same directory as the wowlaunch executable
    cp wowlaunch $HOME/.local/bin
    cp exampleconfig $HOME/.local/bin/config.yaml
    ```
   
4. Edit the config file:

    ```yaml
    ---
    # Configure the delay in seconds between starting the program and typing the 
    # username & password into the Wow client. Tweak to fit your PC's speed
    delay: 8
    # Path to Wow.exe
    wowpath: "/home/$USERNAME/.wine/drive_c/World of Warcraft/Wow.exe"
    # WINEPREFIX for Linux users. Windows users can leave this as the default.
    # Do not remove it or you will break wowlaunch!
    wineprefix: "/home/$USERNAME/.wine"
    
    # Accounts must be listed using the list syntax below
    accounts:
        # The name you'll give to wowlaunch at launch. Set this to whatever you want
       - name: 'main'
         username: 'accountone'
         password: 'somepassword'
    
       - name: 'second'
         username: 'accounttwo'
         password: 'someotherpassword'
    ```

## Usage

Tell wowlaunch which account name to use and it will do the rest

```bash
# Using the above example config
wowlaunch main

# Launch the second account
wowlaunch second
```