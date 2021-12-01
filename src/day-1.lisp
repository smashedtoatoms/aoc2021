(defpackage :aoc2021/day-1
  (:use :cl)
  (:export :get-depth-increases))

(in-package :aoc2021/day-1)

(defun run-in-window (list fn n)
  "Applies the provided function to a subset the size of n of the provided list."
  (loop for x on list
        while (nthcdr (1- n) x)
        collect (funcall fn (subseq x 0 n))))

(defun get-depth-increases (window-size raw-entries)
  "Returns the number of depth increases in the data set at the specified window
   size."
  (let ((measurements (run-in-window raw-entries (lambda (e) (reduce #'+ e)) window-size)))
    (count t (loop for e1 in measurements
                   for e2 in (rest measurements)
                   if (> e2 e1)
                     collect t
                   else
                     collect nil))))
