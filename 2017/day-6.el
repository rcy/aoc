;;; day-6.el ---                                     -*- lexical-binding: t; -*-

(defun day-6 (vec)
  (let (seen '())
    (while (not (seq-contains seen vec))
      (push (seq-copy vec) seen)
      (let* ((max (seq-max vec))
             (pos (seq-position vec max)))
        (aset vec pos 0)
        (dotimes (i max)
          (let ((i (% (+ 1 i pos) (length vec))))
            (aset vec i (1+ (elt vec i)))))))
    (cons (length seen)                   ; part 1
          (1+ (seq-position seen vec))))) ; part 2

; (day-6 [0 5 10 0 11 14 13 4 11 8 8 7 1 4 12 11])
