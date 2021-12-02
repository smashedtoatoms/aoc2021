(defpackage :aoc2021/test
  (:use :cl :parachute)
  (:import-from
   :aoc2021/day-1
   :get-depth-increases)
  (:import-from
   :aoc2021/day-2
   :calculate-course-simple
   :calculate-course-complex
   :txt-to-direction)
  (:export :test-all))

(in-package :aoc2021/test)

(define-test all)

;; Day 1
(define-test day1 :parent all)

(define-test "gets depth change for test data"
  :parent day1
  (let* ((test-data '(199 200 208 210 200 207 240 269 260 263))
        (result (get-depth-increases 1 test-data)))
    (is = 7 result)))

(define-test "gets depth change for prod data"
  :parent day1
  (let* ((test-data (mapcar #'parse-integer
                            (uiop:read-file-lines #p"inputs/day1.txt")))
        (result (get-depth-increases 1 test-data)))
    (is = 1448 result)))

(define-test "gets shifting depth changes for test data"
  :parent day1
  (let* ((test-data '(199 200 208 210 200 207 240 269 260 263))
        (result (get-depth-increases 3 test-data)))
    (is = 5 result)))

(define-test "gets shifting depth changes for prod data"
  :parent day1
  (let* ((test-data (mapcar #'parse-integer
                            (uiop:read-file-lines #p"inputs/day1.txt")))
        (result (get-depth-increases 3 test-data)))
    (is = 1471 result)))

;; Day 2
(define-test day2 :parent all)

(define-test "calculates course for test data"
  :parent day2
  (let* ((test-data '(("forward" 5) ("down" 5) ("forward" 8) ("up" 3) ("down" 8) ("forward" 2)))
         (result (calculate-course-simple test-data)))
    (is equal 150 result)))

(define-test "calculates course for prod data"
  :parent day2
  (let* ((test-data (mapcar #'txt-to-direction
                            (uiop:read-file-lines #p"inputs/day2.txt")))
         (result (calculate-course-simple test-data)))
    (is equal 1635930 result)))

(define-test "calculates more complicated course for test data"
  :parent day2
  (let* ((test-data '(("forward" 5) ("down" 5) ("forward" 8) ("up" 3) ("down" 8) ("forward" 2)))
         (result (calculate-course-complex test-data)))
    (is equal 900 result)))

(define-test "calculates more complicated course for prod data"
  :parent day2
  (let* ((test-data (mapcar #'txt-to-direction
                            (uiop:read-file-lines #p"inputs/day2.txt")))
         (result (calculate-course-complex test-data)))
    (is equal 1781819478 result)))

(defun test-all ()
  (test 'all))
