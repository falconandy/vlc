# VLC TCP commands

## Linux (VLC 4.0.0-dev)
```
+----[ CLI commands ]
| add XYZ  . . . . . . . . . . . . . . . . . . . . add XYZ to playlist
| enqueue XYZ  . . . . . . . . . . . . . . . . . queue XYZ to playlist
| playlist . . . . . . . . . . . . .  show items currently in playlist
| search [string]  . .  search for items in playlist (or reset search)
| delete [X] . . . . . . . . . . . . . . . . delete item X in playlist
| move [X][Y]  . . . . . . . . . . . . move item X in playlist after Y
| sort key . . . . . . . . . . . . . . . . . . . . . sort the playlist
| sd [sd]  . . . . . . . . . . . . . show services discovery or toggle
| play . . . . . . . . . . . . . . . . . . . . . . . . . . play stream
| stop . . . . . . . . . . . . . . . . . . . . . . . . . . stop stream
| next . . . . . . . . . . . . . . . . . . . . . .  next playlist item
| prev . . . . . . . . . . . . . . . . . . . .  previous playlist item
| goto, gotoitem . . . . . . . . . . . . . . . . .  goto item at index
| repeat [on|off]  . . . . . . . . . . . . . .  toggle playlist repeat
| loop [on|off]  . . . . . . . . . . . . . . . .  toggle playlist loop
| random [on|off]  . . . . . . . . . . . . . .  toggle playlist random
| clear  . . . . . . . . . . . . . . . . . . . . .  clear the playlist
| status . . . . . . . . . . . . . . . . . . . current playlist status
| title [X]  . . . . . . . . . . . . . . set/get title in current item
| title_n  . . . . . . . . . . . . . . . .  next title in current item
| title_p  . . . . . . . . . . . . . .  previous title in current item
| chapter [X]  . . . . . . . . . . . . set/get chapter in current item
| chapter_n  . . . . . . . . . . . . . .  next chapter in current item
| chapter_p  . . . . . . . . . . . .  previous chapter in current item
|
| seek X . . . . . . . . . . . seek in seconds, for instance `seek 12'
| pause  . . . . . . . . . . . . . . . . . . . . . . . .  toggle pause
| fastforward  . . . . . . . . . . . . . . . . . . set to maximum rate
| rewind . . . . . . . . . . . . . . . . . . . . . set to minimum rate
| faster . . . . . . . . . . . . . . . . . .  faster playing of stream
| slower . . . . . . . . . . . . . . . . . .  slower playing of stream
| normal . . . . . . . . . . . . . . . . . .  normal playing of stream
| rate [playback rate] . . . . . . . . . .  set playback rate to value
| frame  . . . . . . . . . . . . . . . . . . . . . play frame by frame
| fullscreen, f, F [on|off]  . . . . . . . . . . . . toggle fullscreen
| info [X] . .  information about the current stream (or specified id)
| stats  . . . . . . . . . . . . . . . .  show statistical information
| get_time . . . . . . . . .  seconds elapsed since stream's beginning
| is_playing . . . . . . . . . . . .  1 if a stream plays, 0 otherwise
| get_title  . . . . . . . . . . . . . the title of the current stream
| get_length . . . . . . . . . . . .  the length of the current stream
|
| volume [X] . . . . . . . . . . . . . . . . . .  set/get audio volume
| volup [X]  . . . . . . . . . . . . . . .  raise audio volume X steps
| voldown [X]  . . . . . . . . . . . . . .  lower audio volume X steps
| achan [X]  . . . . . . . . . . . .  set/get stereo audio output mode
| atrack [X] . . . . . . . . . . . . . . . . . . . set/get audio track
| vtrack [X] . . . . . . . . . . . . . . . . . . . set/get video track
| vratio [X] . . . . . . . . . . . . . . .  set/get video aspect ratio
| vcrop, crop [X]  . . . . . . . . . . . . . . . .  set/get video crop
| vzoom, zoom [X]  . . . . . . . . . . . . . . . .  set/get video zoom
| vdeinterlace [X] . . . . . . . . . . . . . set/get video deinterlace
| vdeinterlace_mode [X]  . . . . . . .  set/get video deinterlace mode
| snapshot . . . . . . . . . . . . . . . . . . . . take video snapshot
| strack [X] . . . . . . . . . . . . . . . . .  set/get subtitle track
| hotkey, key [hotkey name]  . . . . . . . . . . simulate hotkey press
|
| vlm  . . . . . . . . . . . . . . . . . . . . . . . . .  load the VLM
| set [var [value]]  . . . . . . . . . . . . . . . . . set/get env var
| save_env . . . . . . . . . . . .  save env vars (for future clients)
| alias [cmd]  . . . . . . . . . . . . . . . . set/get command aliases
| description  . . . . . . . . . . . . . . . . .  describe this module
| license  . . . . . . . . . . . . . . . . print VLC's license message
| help, ? [pattern]  . . . . . . . . . . . . . . . . .  a help message
| longhelp [pattern] . . . . . . . . . . . . . . a longer help message
| lock . . . . . . . . . . . . . . . . . . . .  lock the telnet prompt
| logout . . . . . . . . . . . . . .  exit (if in a socket connection)
| quit . . . . . . . .  quit VLC (or logout if in a socket connection)
| shutdown . . . . . . . . . . . . . . . . . . . . . . .  shutdown VLC
+----[ end of help ]
```

## Windows (VLC 3.0.6)
```
+----[ Remote control commands ]
|
| add XYZ  . . . . . . . . . . . . add XYZ to playlist
| enqueue XYZ  . . . . . . . . . queue XYZ to playlist
| playlist . . . . .  show items currently in playlist
| play . . . . . . . . . . . . . . . . . . play stream
| stop . . . . . . . . . . . . . . . . . . stop stream
| next . . . . . . . . . . . . . .  next playlist item
| prev . . . . . . . . . . . .  previous playlist item
| goto . . . . . . . . . . . . . .  goto item at index
| repeat [on|off] . . . .  toggle playlist item repeat
| loop [on|off] . . . . . . . . . toggle playlist loop
| random [on|off] . . . . . . .  toggle random jumping
| clear . . . . . . . . . . . . . . clear the playlist
| status . . . . . . . . . . . current playlist status
| title [X]  . . . . . . set/get title in current item
| title_n  . . . . . . . .  next title in current item
| title_p  . . . . . .  previous title in current item
| chapter [X]  . . . . set/get chapter in current item
| chapter_n  . . . . . .  next chapter in current item
| chapter_p  . . . .  previous chapter in current item
|
| seek X . . . seek in seconds, for instance `seek 12'
| pause  . . . . . . . . . . . . . . . .  toggle pause
| fastforward  . . . . . . . .  .  set to maximum rate
| rewind  . . . . . . . . . . . .  set to minimum rate
| faster . . . . . . . . . .  faster playing of stream
| slower . . . . . . . . . .  slower playing of stream
| normal . . . . . . . . . .  normal playing of stream
| frame. . . . . . . . . .  play frame by frame
| f [on|off] . . . . . . . . . . . . toggle fullscreen
| info . . . . .  information about the current stream
| stats  . . . . . . . .  show statistical information
| get_time . . seconds elapsed since stream's beginning
| is_playing . . . .  1 if a stream plays, 0 otherwise
| get_title . . . . .  the title of the current stream
| get_length . . . .  the length of the current stream
|
| volume [X] . . . . . . . . . .  set/get audio volume
| volup [X]  . . . . . . .  raise audio volume X steps
| voldown [X]  . . . . . .  lower audio volume X steps
| adev [device]  . . . . . . . .  set/get audio device
| achan [X]. . . . . . . . . .  set/get audio channels
| atrack [X] . . . . . . . . . . . set/get audio track
| vtrack [X] . . . . . . . . . . . set/get video track
| vratio [X]  . . . . . . . set/get video aspect ratio
| vcrop [X]  . . . . . . . . . . .  set/get video crop
| vzoom [X]  . . . . . . . . . . .  set/get video zoom
| snapshot . . . . . . . . . . . . take video snapshot
| strack [X] . . . . . . . . .  set/get subtitle track
| key [hotkey name] . . . . . .  simulate hotkey press
|
| help . . . . . . . . . . . . . . . this help message
| logout . . . . . . .  exit (if in socket connection)
| quit . . . . . . . . . . . . . . . . . . .  quit vlc
|
+----[ end of help ]
```