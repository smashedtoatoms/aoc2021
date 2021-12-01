(defpackage :aoc2021/test
  (:use :cl :parachute)
  (:import-from :aoc2021/day-1 :get-depth-increases)
  (:export :test-all))

(in-package :aoc2021/test)

(define-test all)

(define-test "gets depth change for test data"
  :parent all
  (let* ((test-data '(199 200 208 210 200 207 240 269 260 263))
        (result (get-depth-increases 1 test-data)))
    (is = 7 result)))

(define-test "gets depth change for prod data"
  :parent all
  (let* ((test-data (mapcar #'parse-integer (uiop:read-file-lines #p"inputs/day1.txt")))
        (result (get-depth-increases 1 test-data)))
    (is = 1448 result)))

(define-test "gets shifting depth changes for test data"
  :parent all
  (let* ((test-data '(199 200 208 210 200 207 240 269 260 263))
        (result (get-depth-increases 3 test-data)))
    (is = 5 result)))

(define-test "gets shifting depth changes for prod data"
  :parent all
  (let* ((test-data (mapcar #'parse-integer (uiop:read-file-lines #p"inputs/day1.txt")))
        (result (get-depth-increases 3 test-data)))
    (is = 1471 result)))

(defun test-all ()
  (test 'all))
