# pairwise

## What it is 

Using finite projective plane to generate card (maybe video) games. 

Currently uses package hexmaker to create hextiles based on generated PPs. 

## Running

`go run main.go -flag=value`

For instance, `go run main.go -order=7 -show=true` to display the projective
plane order 7. 

```
Usage of ./main:
  -color (default true)
      Print in color, or if false, black and white with stroked hexagon outline
  -format (default false)
     Format with commas for python/etc
  -order int (default 5)
      Order of the projective plane (default 3)
  -show (default false)
      Displays to stdout
  -shuffle (default false)
      Permutes cards randomly, use with fracmaker
  -solver (default false) 
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

