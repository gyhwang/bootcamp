package main

type led [5]string

var zero = led{
	"███",
	"█ █",
	"█ █",
	"█ █",
	"███",
}

var one = led{
	"██ ",
	" █ ",
	" █ ",
	" █ ",
	"███",
}

var two = led{
	"███",
	"  █",
	"███",
	"█  ",
	"███",
}

var three = led{
	"███",
	"  █",
	"███",
	"  █",
	"███",
}

var four = led{
	"█ █",
	"█ █",
	"███",
	"  █",
	"  █",
}

var five = led{
	"███",
	"█  ",
	"███",
	"  █",
	"███",
}

var six = led{
	"███",
	"█  ",
	"███",
	"█ █",
	"███",
}

var seven = led{
	"███",
	"  █",
	"  █",
	"  █",
	"  █",
}

var eight = led{
	"███",
	"█ █",
	"███",
	"█ █",
	"███",
}

var nine = led{
	"███",
	"█ █",
	"███",
	"  █",
	"███",
}

var colon = led{
	" ",
	"░",
	" ",
	"░",
	" ",
}

var dot = led{
	" ",
	" ",
	" ",
	" ",
	"░",
}

var digits = [...]led{zero, one, two, three, four, five, six, seven, eight, nine}
