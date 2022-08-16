# pairwise

## What it is 

Using finite projective plane to generate card (maybe video) games. 

Currently uses package hexmaker to create hextiles based on generated PPs. 

## Running

`go run main.go -flag=value`

For instance: 
`go run main.go -order=7 -show=array` to display the projective plane order 7.

`go run main.go -order=3 -show=text -arrange=solve` to run the solver and display the projective plane order 3, "solved" (ie every element appears in exactly one column using OR-Tools)

`go run main.go -help` for options.

```
Usage:
  -arrange string
      'shuffle', 'solve', 'test', 'order', or default ''
  -color
      Render in color
  -hexdeck
      Send the PP to hexmaker and make a set of images
  -order int
      Order of the projective plane (default 2)
  -show string
      'text', or 'array', (default no output)
```

The (commented-out) texmaker part both requires an external library, and seeks
to write a test file into etc/test.pdf -- shouldn't be a problem unless you
uncomment and try to run the texmaker portion.

## Notes

https://www.youtube.com/watch?v=k472YAVzoGU

## Fracmaker 

The (commented-out) portion(s) of fracmaker seek to create a set of cards with
equivalent fractions on each. Lucky gift from the universe, that there are 31
(n^2 + n + 1, for n=5) irreducible fractions from 1/2 to 9/10. 

# Next

- fix Fracmaker
- generalize geometries and alignment
- ~~figure out how to integrate a solver so as to not leave Go, or have the
  solver (python) write a file and have pairwise read the file~~
- ~~Check for deck/ subdirectory before trying to write card images to it, give
  error or attempt to create the directory~~
- Add a sheet option for tiling the hexagons for easy printing and cutting
  prototypes
- Work on generalizing from hexmaker to "make some kind of card shape"
- - Choose edge or corner alignment
- - Check and deal with requests (order 8, hexagon for instance)
- ~~Wrap it all in a shell script until a native Go solver is used~~
- - ~~Give that wrapper an ability to take in a textfile of parameters~~
- Introduce a text UI for growing branches of request possibilities
- - Check for incompatible requests, should be handled in the UI
