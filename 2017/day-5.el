;;; day-5.el --- aoc2017 day 5                       -*- lexical-binding: t; -*-

;; Copyright (C) 2017  Ryan Yeske

;; Author: Ryan Yeske <rcyeske@gmail.com>
;; Keywords: 

;; This program is free software; you can redistribute it and/or modify
;; it under the terms of the GNU General Public License as published by
;; the Free Software Foundation, either version 3 of the License, or
;; (at your option) any later version.

;; This program is distributed in the hope that it will be useful,
;; but WITHOUT ANY WARRANTY; without even the implied warranty of
;; MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
;; GNU General Public License for more details.

;; You should have received a copy of the GNU General Public License
;; along with this program.  If not, see <http://www.gnu.org/licenses/>.

;;; Commentary:

;; 

;;; Code:

(defun aoc-d5p1-instructions (file)
  (with-temp-buffer
    (insert-file-contents file)
    (seq-into (mapcar 'string-to-number (split-string (buffer-string)))
              'vector)))

(defun aoc-d5p1 (file)
  (let ((instructions (aoc-d5p1-instructions file))
        (i 0)
        (c 0))
    (while (and (>= i 0) (< i (length instructions)))
      (setq c (1+ c))
      (let ((o (elt instructions i)))
        (aset instructions i (1+ o))
        (setq i (+ i o))))
    c))

(defun aoc-d5p2 (file)
  (let ((instructions (aoc-d5p1-instructions file))
        (i 0)
        (c 0))
      (while (and (>= i 0) (< i (length instructions)))
        (setq c (1+ c))
        (let ((o (elt instructions i)))
          (aset instructions i
                (if (>= o 3)
                    (1- o)
                  (1+ o)))
          (setq i (+ i o))))
      c))

;; (aoc-d5p1 "day-5.input") ; 339351
;; (aoc-d5p2 "day-5.input") ; 24315397

(provide 'day-5)
;;; day-5.el ends here
