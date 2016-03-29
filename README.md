# slaidsciou

Set a slideshow of images as wallpaper, on GNOME Shell.

This thing uses a dirty hack, it basically sets the wallpaper using `gsettings` every interval.

It'' ugly but it works.

## Configuration

You can find an example configuration file [here](https://github.com/gsora/slaidsciou/blob/master/slaidsciou.conf).

It's JSON, put it in `~/.config/slaidsciou.conf` and you're good to go.

|Option|Needed|Type|Description|
|------|------|----|-----------|
|`wallpaper-path`|yes|`string`|A folder containing all the images you want as wallpaper|
|`randomize`|no|`boolean`|If true, pictures will be randomized|
|`delay`|yes|`integer`|Delay between pictures in minutes|

## Installation

It's written in Go, so it's just a matter of
```
go get github.com/gsora/slaidsciou
```

## Usage

If run without any argument, `slaidsciou` will quietly run in background.

You can run it with the `-foreground` switch for debug purposes, but who cares?

To start it with GNOME Shell, you have to copy `slaidsciou.desktop` to `~/.local/share/applications`, modifying the `Exec` directive according to where your `$GOPATH/bin` is.

After that, just add it to the autostart list using GNOME Tweak Tool.

## Why in Go?

Why not?
