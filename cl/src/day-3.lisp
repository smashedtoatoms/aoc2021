(defpackage :aoc2021/day-3
  (:use :cl)
  (:export :calculate-power-usage
           :calculate-life-support-rating))

(in-package :aoc2021/day-3)

(defun find-most-common-bit (index bits)
  "Finds the most common bit in an array of bits and returns that bit."
  (if (>=
       (count #\1 bits :key (lambda (n) (aref n index)))
       (/ (length bits) 2))
      #\1
    #\0))

(defun calculate-power-usage (diagnostic-data)
  "Takes diagnostic data as a list of 5 bit arrays and from it derives the
   power usage.

   diagnostic-data example: '(\"00100\" \"11110\" \"10110\")"
  (let* ((data-length (length (car diagnostic-data)))
         (p (parse-integer
             (coerce (loop for i below data-length
                           collect (find-most-common-bit i diagnostic-data))
                     'string)
             :radix 2)))
    (* p (- (ash 1 data-length) 1 p))))

(defun rate (rating-type diagnostic-data)
  "Takes diagnostic data as a list of 5 bit arrays and calculates the rating
   for the specified type.

  rating-type example: :oxygen-generator-rating
  diagnostic-data example: '(\"00100\" \"11110\" \"10110\")"
  (assert (position rating-type '(:oxygen-generator-rating :co2-scrubber-rating)))
  (let ((test (case rating-type
                (:oxygen-generator-rating :test)
                (:co2-scrubber-rating :test-not))))
    (loop for i from 0 until (null (cdr diagnostic-data))
          do (setf diagnostic-data (remove (find-most-common-bit i diagnostic-data)
                                           diagnostic-data
                                           test
                                           #'eq
                                           :key (lambda (n) (aref n i)))))
    (parse-integer (car diagnostic-data) :radix 2)))

(defun calculate-life-support-rating (diagnostic-data)
  "Takes diagnostic data as a list of 5 bit arrays and from it derives the
   life-support rating"
  (* (rate :oxygen-generator-rating diagnostic-data) (rate :co2-scrubber-rating diagnostic-data)))
