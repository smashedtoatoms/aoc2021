(defpackage :aoc2021/day-2
  (:use :cl)
  (:export :calculate-course-simple :calculate-course-complex :txt-to-direction))

(in-package :aoc2021/day-2)

(defun txt-to-direction (raw)
  "Returns a direction amount config list.

   For example: (list \"foward\" 10) will be returned when given a raw string
   such as \"forward 10\""
  (let* ((parts (uiop:split-string raw :separator " "))
       (direction (first parts))
       (amount (parse-integer (second parts))))
    (list direction amount)))

(defun calculate-course-simple (directions)
  "Takes a list of direction-amount lists and calculates the simple course.

   directions example: '((\"foward\" 10) (\"down\" 2) (\"up\" 5))"
  (loop for (direction amount) in directions
        when (string= direction "forward")
          sum amount into horizontal-position
        when (string= direction "down")
          sum amount into depth
        when (string= direction "up")
          sum (* -1 amount) into depth
        finally (return (* horizontal-position depth))))

(defun calculate-course-complex (directions)
  "Takes a list of direction-amount lists and calculates the complext course.

   directions example: '((\"foward\" 10) (\"down\" 2) (\"up\" 5))
   "
  (loop for (direction amount) in directions
        sum 0 into aim
        sum 0 into depth
        sum 0 into horizontal-position
        when (string= direction "forward")
          sum amount into horizontal-position
        when (string= direction "forward")
          sum (* amount (or aim 0)) into depth
        when (string= direction "down")
          sum amount into aim
        when (string= direction "up")
          sum (* -1 amount) into aim
        finally (return (* horizontal-position depth))))
