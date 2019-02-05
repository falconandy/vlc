# Golang wrapper for the VLC player (via TCP)

Golang wrapper to control the VLC Player via TCP commands.

```
p := NewPlayer(nil)
p.Start()
p.Play("/path/to/video")
time.Sleep(time.Second * 5)
p.Pause()
time.Sleep(time.Second * 3)
p.Pause()
time.Sleep(time.Second * 5)
p.Stop()
p.Shutdown()
```
