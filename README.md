# aoc2021

It's that time of year again, so here we go.

[Last year, I tried to do it in Rust, then switched to Common Lisp, and
then back to Rust before bailing out after a few days to tinker with
build stuff an refactors](https://github.com/smashedtoatoms/aoc2020) I
left frustrated with the Common Lisp ecosystem and frustrated with how
hard it is to stay focused on business problems in Rust.

Throughout 2021, I made some contributions to
[Alive](https://marketplace.visualstudio.com/items?itemName=rheller.alive),
[CLPM](https://clpm.dev), and [The Common Lisp
Cookbook](https://lispcookbook.github.io/cl-cookbook/vscode-alive.html).
In doing so, I learned enough about the Common Lisp ecosystem for me to
actually get something done this year... or so I thought.  I got 3 done
before I got so lost in Common Lisp for-loop macros that I bailed out to
Go so I could actually get something done in under an hour and not be a
zombie at work and life due to staying up until 3 working on AOC.  If I
could figure out how to efficiently use the loop macro, it seems
powerful, but it's challenging to get right.

I guess I will have to put some more time in on Lisp.  It is
legitimately different enough that I can't just force myself through it
like I can with most new languages I learn.

I got to day 12 this year before calling it.  I don't have enough time
in my day to do this, work, music, side-projects, etc.  I have decided
that this one isn't going to make the cut.  I feel appropriately guilty.

## How To Dev in Common Lisp

### Prerequisites

Be in the cl directory and have these installed however you prefer to
install them:

- [SBCL](http://www.sbcl.org)
- [CLPM](https://clpm.dev)
- [Lake](https://github.com/takagi/lake)

I used [SBCL](http://www.sbcl.org) installed via the
[asdf-sbcl](https://github.com/smashedtoatoms/asdf-sbcl) plugin for
[asdf-vm](https://asdf-vm.com). I used [CLPM](https://clpm.dev) as my
package manager and [Lake](https://github.com/takagi/lake) for my
minimal build automation needs.  [The
cookbook](https://lispcookbook.github.io/cl-cookbook/vscode-alive.html#configuring-vscode-alive-to-work-with-clpm)
specifies how to set up REPL dev in
[VSCode](https://code.visualstudio.com), and it works; however, I'm
using [Emacs](https://www.gnu.org/software/emacs/) with
[Sly](https://github.com/joaotavora/sly).

This setup is a 100% common lisp way to do repeatable builds with
pinnable dependencies, automatic installation and recording of dependencies,
automatic downloads, and all the other goodies a proper package manager should
do for me.  My asdf-sbcl installer gives me a version of SBCL that can do
image compression, threading, linking of c libs (like openssl) and crypto. I
likely won't need any of that for this, but it is nice knowing it is there.
I believe this project structure to be modern idiomatic lisp.

### Testing

I used [Parachute](https://shinmera.github.io/parachute/) for my testing
framework.

### Useful Commands (assuming all prerequisets are met)
- install dependencies and run tests
  ```sh
  lake
  ```
- install dependencies
  ```sh
  lake install-dependencies
  ```
- run tests
  ```sh
  lake test
  ```
- start up sly repl on port 4005 with project systems loaded
  ```sh
  lake repl
  ```
- start up slime repl on port 4005 with project systems loaded
  ```sh
  lake repl-slime
  ```

## How To Dev in Go

### Prerequisites

Be in the go directory and have these installed however you prefer to
install them:

- [Go](https://go.dev/doc/install)
- [Mage](https://github.com/magefile/mage)

I used [Go](https://go.dev/doc/install) installed via the
regular installer. I used [Mage](https://github.com/magefile/mage) for my
minimal build automation needs.

This setup is a 100% Go way to do repeatable builds with pinnable
dependencies, automatic installation and recording of dependencies,
automatic downloads, and all the other goodies a proper package manager
should do for me.  It is mostly idiomatic, though the use of `internal`
and `cmd` directories on a project like this is probably overkill and
the use of `internal` at all can be contentious.  I probably don't need
any of it and could just dump everything into main at the top.  I chose
to split it up to make it easier to keep straight for myself.

### Testing

I used the built-in go tools for my test.

### Useful Commands (assuming all prerequisets are met)
- install dependencies and run tests
  ```sh
  mage
  ```
- download dependencies
  ```sh
  go mod download
  ```
- run tests
  ```sh
  mage test
  ```
- run build
  ```sh
  lake build
  ```
- run binary (after build) for exercise 4 for specified test data.
  ```sh
  out/aoc2020 4 ./inputs/day4.txt
  ```
