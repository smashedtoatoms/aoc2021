# aoc2021

It's that time of year again, so here we go. I look forward to completing 5
exercises before bailing because I rat-holed on trying to figure out some
build-system quirk.

[Last year, I tried to do it in Rust, then switched to Common Lisp, and then
back to Rust.](https://github.com/smashedtoatoms/aoc2020) I left disgusted
with the Common Lisp ecosystem. I also realized that Rust makes it really easy
for me to focus on the wrong things. I will blow hours trying to pass
something by reference to save myself 4 bytes of memory because I know it
should be doable instead of focusing on the actual business problem. I learned
that in Rust, I prefer writing whatever seems more readable. Since it's all
safe, why do a hard-to-read reduce when a for loop with mutations will do?  It
wasn't really fun though.  For me, Rust is not an ideal language to use as
tool to explore a problem in.  It's great for coding up a solution to a
problem I already know how to solve; but, for me, it's not a good tool for
thinking through a problem.  I might be doing it wrong.

Anyway, throughout the year, I made some contributions to
[Alive](https://marketplace.visualstudio.com/items?itemName=rheller.alive),
[CLPM](https://clpm.dev), and [The Common Lisp
Cookbook](https://lispcookbook.github.io/cl-cookbook/vscode-alive.html). In
doing so, I learned enough about the ecosystem for me to actually get
something done this year.  Something is wrong with me.  I may write something
about this quirk of mine.  I love complaining, but I like fixing the shit I
complain about more.  I almost always inadvertently end up solving the things
I complain about over time without really realizing I am doing it.  I am drawn
to others who do the same.

I have no idea why I am drawn to Common Lisp. I have no good reason to do
anything in it. It doesn't help my job. It doesn't make me more hirable. I
don't build anything that generates revenue with it. It's just a fun tool to
use for thinking. It's turtles all the way down. It's probably a total waste
of time.  I am oddly drawn to it against my better judgement.  On with the
exercises.

## How To Dev

### Prerequisites

These need to be installed however you prefer to install them:

- [SBCL](http://www.sbcl.org)
- [CLPM](https://clpm.dev)
- [Lake](https://github.com/takagi/lake)

I am using [SBCL](http://www.sbcl.org) installed via the
[asdf-sbcl](https://github.com/smashedtoatoms/asdf-sbcl) plugin for
[asdf-vm](https://asdf-vm.com). I am using [CLPM](https://clpm.dev) as my
package manager and [Lake](https://github.com/takagi/lake) for my minimal
build automation needs.  [The
cookbook](https://lispcookbook.github.io/cl-cookbook/vscode-alive.html#configuring-vscode-alive-to-work-with-clpm)
specifies how to set up REPL dev in [VSCode](https://code.visualstudio.com),
and it works; however, I'm using [Emacs](https://www.gnu.org/software/emacs/)
with [Sly](https://github.com/joaotavora/sly).

This setup gives me a 100% common lisp way to do repeatable builds with
pinnable dependencies, automatic installation and recording of dependencies,
automatic downloads, and all the other goodies a proper package manager should
do for me.  My asdf-sbcl installer gives me a version of SBCL that can do
image compression, threading, linking of c libs (like openssl) and crypto. I
likely won't need any of that for this, but it is nice knowing it is there.

### Testing

I'm using [Parachute](https://shinmera.github.io/parachute/) for my testing
framework.  I have never used it.  We'll see how it files with repl dev.
Assuming you have the prerequisites installed, you can use the following
commands.

### Useful Commands
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
