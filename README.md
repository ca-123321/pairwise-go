# pairwise

## What it is 

Using finite projective plane to generate card (maybe video) games. 

## Running

`go build main.go`

Usage of ./main:
  -format (default false)
      Format with commas for python/etc
  -order int (default 3)
      Order of the projective plane (default 3)
  -show (default false)
      Displays to stdout
  -shuffle (default false)
      Permutes cards randomly, use with fracmaker

Right now uses Go 1.17 but 1.18 just came out, I'm just
learning Go really, so I don't think there should be conflicts unless Go
insists you use 1.17 and you only have 1.18, idk. 

The (commented-out) texmaker part both requires an external library, and seeks
to write a test file into etc/test.pdf -- shouldn't be a problem unless you
uncomment and try to run the texmaker portion.

## Notes

There are flags in `main.go`, for printing and formatting. The formatting is
for outputting in a form easy to paste into the (currently hardcoded)
`pairwise_solver` repo (TODO: link)

## Fracmaker 

The (commented-out) portion(s) of fracmaker seek to create a set of cards with
equivalent fractions on each. Lucky gift from the universe, that there are 31
(n^2 + n + 1, for n=5) irreducible fractions from 1/2 to 9/10. 

# Next

- decide on TeX or SVG
- implement one, according to some printing process, export pdf
- art
- write a solver in Go instead of using pairwise_solver (TODO: link)


