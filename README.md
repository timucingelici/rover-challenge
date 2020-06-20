Specification
---------------

A squad of robotic rovers has landed by NASA on a plateau on
Mars.  This plateau, which is curiously rectangular, must be
navigated by the rovers so that it can plan laser beacons
that will be used for space navigation.

A rover's position is represented by a combination of an x and y
co-ordinates and a letter representing one of the four cardinal
compass points. The plateau is divided up into a grid to simplify
navigation. An example position might be 0, 0, N, which means the
rover is in the bottom left corner and facing North.

In order to control a rover, NASA sends a simple string of
letters. The possible letters are 'L', 'R', 'M' and 'P'. 'L' and 'R'
makes the rover spin 90 degrees left or right respectively, without
moving from its current spot.

'M' means move forward one grid point, and maintain the same heading.
Assume that the square directly North from (x, y) is (x, y+1).

'P' means Plant a laser beacon in the current position without
changing the position or direction

For example:
if you start at:
`{:pos [1 2] :dir :N}`

and you follow these commands
"LMLMLMLMMPM"

Your final position will be:
`{:pos [0 3] :dir :N}`
and one beacon placed at [0 2]

NASA developers use this simple language to write programs which
place laser beacons in special places.

Here is the plot.
-------------------
You noticed that one of the rovers was recently hacked
and you only manage to retrieve the last program which was executed.
The program is a long string of commands like the example above.
The Intelligence Bureau told you that they think the Interstellar
forces kidnapped one of the most prominent scientists and they are
forcing him to build a machine which will counter all Earth moves
in a conflict.
The identity of the missing scientist is unknown which it makes
impossible a rescue mission.
You suspect that the hacking of the rover is not accidental,
and that someone, maybe the kidnapped scientist, wants to send
you a message to let you know who the scientist is.
What you hope for is a name or an address to understand who
the missing scientist is.

Your Mission
--------------

Your mission is to use the retrieved program the rover to
understand who the scientist is. It won't be easy as all
transmissions are intercepted and under surveillance so it won't be
immediately obvious who he is.

Who is the missing scientist??

The survival and freedom of the Earth depends on you,
Good luck!
