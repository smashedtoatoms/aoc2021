(in-package :asdf-user)

(defsystem "aoc2021"
  :version "0.0.1"
  :author ("Jason Legler")
  :maintainer ("Jason Legler")
  :license "MIT"
  :description "smashedtoatoms attempt at aoc 2021"
  :class :package-inferred-system
  :depends-on ("aoc2021/src/day-1")
  :in-order-to ((test-op (load-op "aoc2021/test/all")))
  :perform (test-op (o c) (symbol-call :aoc2021/test :test-all)))

(defsystem "aoc2021/test"
  :depends-on ("parachute"
               "aoc2021/test/all"))

(register-system-packages "aoc2021/src/day-1" '(:aoc2021/day-1))
(register-system-packages "aoc2021/test/all" '(:aoc2021/test))
