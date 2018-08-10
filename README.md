# Beep 
The purpose of function Beep is to make it easy to play notification sounds.
It is intended to be useful, for example, if you are running a long analysis in the background and want to know when it is ready.

## Installation:
```
go get -u github.com/zquangu112z/beep-notify
```

## Usage:
```
import (
	beep "github.com/zquangu112z/beep-notify"
)

func main() {
	beep.Beep(1)
}
```

You can pass the id from 1 to 8 to the Beep() function