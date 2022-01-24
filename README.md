# pairwise

## What it is 

Using finite projective plane to generate card (maybe video) games. 

Currently uses package hexmaker to create hextiles based on generated PPs. 

## Running

`go run main.go -flag=value` All bool flags are false right now, default order
is 5.

For instance, `go run main.go -order=7 -show=true` to display the projective
plane order 7. 

`go run main.go -order=3 -json=true` then run the generic `solver.py` from
within its directory, to get an order 3 PP with array elements permuted such
that each element appears in each column exactly once

`go run main.go -help` for options.

```
Usage:
  -color
    (true) Print in color, (false) black and white with stroked hexagon outline
  -format
    (true) Format with commas for python/etc
  -hexdeck
    (true) Send the PP to hexmaker and make a set of images
  -json
    (true) Marshal the PP and write it to a json file
  -order 
    (int) Order of the projective plane (default 5)
  -show
    (true) Displays to stdout
  -shuffle 
    (true) Permutes cards randomly, use with fracmaker
  -solver 
      currently unimplemented, see pairwise_solver
```

Right now uses Go 1.17 but 1.18 just came out, I'm just
learning Go really, so I don't think there should be conflicts unless Go
insists you use 1.17 and you only have 1.18, idk. 

The (commented-out) texmaker part both requires an external library, and seeks
to write a test file into etc/test.pdf -- shouldn't be a problem unless you
uncomment and try to run the texmaker portion.

## Notes

https://www.youtube.com/watch?v=V_wQ7ByCSgw

## Fracmaker 

The (commented-out) portion(s) of fracmaker seek to create a set of cards with
equivalent fractions on each. Lucky gift from the universe, that there are 31
(n^2 + n + 1, for n=5) irreducible fractions from 1/2 to 9/10. 

# Next

- fix Fracmaker
- generalize geometries and alignment
- figure out how to integrate a solver so as to not leave Go, or have the
  solver (python) write a file and have pairwise read the file
- ~~Check for deck/ subdirectory before trying to write card images to it, give
  error or attempt to create the directory~~
- Add a sheet option for tiling the hexagons for easy printing and cutting
  prototypes
- Wrap it all in a shell script until a native Go solver is used
- - Give that wrapper an ability to take in a textfile of parameters
