# integer-spiral

I think this is another example of the "interview question" genre.  I don't really know where it originated.  A buddy of mine mentioned this problem to me, and my curiosity took over.  I believe the original question was phrased something like:

"Given an integer size _x_ defining the length of the side of a square, write something to print out integers in a spiral starting from the outside the square of said length and working your way in."  Something like that.

The first draft in my brain was more procedural, i.e., start moving in a direction, then turn when you hit the boundary.  And repeat until you're done.

But ultimately, the physics/math part of my brain protested that there must be a function that determines the value of any given element in the 2-D array of size _x_ given the coordinates of the element.  (And there is at least one.)

So the way I broke it down was by "layer", then "leg", then "offset".

Layer essentially defines one trip around the square.  Remember, we're starting on the outside and spiraling inward toward the center.
Leg is one of the four segments of a layer.  I mean, it's a square, right?  :)
Offset defines the distance from the beginning of a leg to the location of the current element.

Using this schema, I started to manually calculate some things to see what pattern the sequence formed.

Have a look at the .pdf; those are basically my notes.

Once I determined the pattern, there's not a lot left to the implementation.  This was still in my early days of Golang, so the code's a bit rough around the edges.

But it gets the job done.  :)

Extending this code into functions to, say, print the value of one element of the array given the element's coordinates, is left as an exercise for the reader.
