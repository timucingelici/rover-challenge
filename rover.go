package main

type rover struct {
	loc *point
	dir *direction
	min *point
	max *point
	locationHistory []point
	beaconLocations []point
}

type point struct {
	x int
	y int
}

type direction struct {
	current int
}

func (d *direction) turn(angle int) {
	d.current = (d.current + 360 + angle) % 360
}

func (d *direction) asVector() point {
	switch d.current {
		case 0:
			return point{x: 0, y: 1}
		case 90:
			return point{x: 1, y: 0}
		case 180:
			return point{x: 0, y: -1}
		case 270:
			return point{x: -1, y: 0}
	}

	return point{}
}

func (r *rover) execute(command string) {
	switch command {
	case "L":
		r.dir.turn(-90)
		break
	case "R":
		r.dir.turn(90)
		break
	case "M":
		r.move()
		r.updateMinMax()
		r.locationHistory = append(r.locationHistory, *r.loc)
		break
	case "P":
		r.beaconLocations = append(r.beaconLocations, *r.loc)
		break
	}
}

func (r *rover) move() {
	r.loc.x = r.loc.x + r.dir.asVector().x
	r.loc.y = r.loc.y + r.dir.asVector().y
}

func (r *rover) updateMinMax() {
	if r.loc.x < r.min.x {
		r.min.x = r.loc.x
	}

	if r.loc.y < r.min.y {
		r.min.y = r.loc.y
	}

	if r.loc.x > r.max.x {
		r.max.x = r.loc.x
	}

	if r.loc.y > r.max.y {
		r.max.y = r.loc.y
	}
}