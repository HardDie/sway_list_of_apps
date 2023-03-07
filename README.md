## sway_list_of_apps
Allows you to get a list of running applications in a compact form with class and appId attributes.

The usual way to get the class and appId is to run the command **swaymsg -t get_outputs**.
But in such a large json response, it's hard to find the right fields.
This utility allows you to get a list of applications in this view:
```
[
        {
                "name": "1",
                "apps": [
                        {
                                "name": "Waterfox",
                                "class": "Waterfox",
                                "appId": null
                        }
                ]
        },
        {
                "name": "2",
                "apps": [
                        {
                                "name": "zsh",
                                "class": "",
                                "appId": "Alacritty"
                        }
                ]
        }
]
```

## How to install
```
go install github.com/HardDie/sway_list_of_apps@latest
```

## How to run
```
export PATH=$PATH:$HOME/go/bin
sway_list_of_apps
```
