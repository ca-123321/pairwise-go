# pairwise_solver

## What it does(TODO: update)

This takes in a hard-coded set of arrays (`/particulars/solver_order3.py` etc) or uses
PP.json if created by `../main.go`, and permutes them such that every element appears in a column exactly once.

Then it outputs `arrangedPP.json` for further use, as well as prints the solution. We can also enumerate all possible solutions (see below)

## Usage

`python3 solver.py` will give one solution of the PP.json

Note: Sometimes this will reprint the single solution multiple times. Not sure
why (TODO), but it's not a problem at the moment.

`python3 solver.py enum` will enumerate all solutions of the PP.json to stdout,
but will still generate a single `arrangedPP.json` of the last found solution.

## Notes

Uncomment `Enumerate all results` and wait for 3 billion years. Enumerating the
results isn't important, but it's an interesting problem for the future, that
seems open. 

- Successfully reduced problem space by the (order+1)! solutions which are
just column permutations of the whole set of arrays

## Todo

- (very low on the list) Reduce the solution space by (order+1)! when enumerating all results, each time a solution is found by eliminating all possibilties of the form of shifting entire columns of the set.  ( D O N E )
- Move to a Go setting to avoid writing and reading
- In the meantime, if the `solver` flag is `true`:
- - PPmaker should output to a text file
- - The python solver should read input from a text file, perform solving, and
    export to a text file
- - `pairwise` should then pick up and do things to it (send it to `hexmaker`
    or whatever future card/tile creation
- - Probably all organized by a shell script/TUI/GUI

(TODO: Update this as json marshaling works) 

- In the meantime of having to use a python solver, marshal the solver output
  to feed back into pairwise

## Credits

thanks to rope321 on reddit, and also to the googles, I wrote almost
none of it

thanks to the or-tools mailing group

thanks to Mr Fred Rogers for teaching at a young age that doing your best is
good enough

thanks to YOU, just for being who you are.
