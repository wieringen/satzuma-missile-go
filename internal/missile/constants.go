package missile

const (
	UP        = 0x02
	DOWN      = 0x01
	LEFT      = 0x08
	RIGHT     = 0x04
	LEFTUP    = LEFT + UP
	RIGHTUP   = RIGHT + UP
	LEFTDOWN  = LEFT + DOWN
	RIGHTDOWN = RIGHT + DOWN
	STOP      = 0x00
	FIRE      = 0x10
)
