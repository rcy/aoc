;;; day-4.el --- day 4                               -*- lexical-binding: t; -*-

(defun day-4-part-1 (file)
  (with-temp-buffer
    (insert-file-contents file)
    (let ((phrases (split-string (buffer-string) "\n")))
      (seq-count (lambda (x) x)
                 (mapcar 'day-4-is-valid phrases)))))

(defun day-4-is-valid (phrase)
  (let* ((words (sort (split-string phrase) 'string-lessp))
         (len1 (length words))
         (len2 (length (delete-dups words))))
    (and (not (equal len1 0))
         (equal len1 len2))))

(day-4-part-1 "./day-4.input") ;; 325

(defun day-4-part-2 (file)
  (with-temp-buffer
    (insert-file-contents file)
    (let ((phrases (split-string (buffer-string) "\n")))
      (seq-count (lambda (x) x)
                 (mapcar 'day-4-part-2-is-valid phrases)))))

(defun sort-letters (word)
  (seq-sort '< word))

(defun day-4-part-2-is-valid (phrase)
  (let* ((words (sort
                 (mapcar 'sort-letters (split-string phrase))
                 'string-lessp))
         (len1 (length words))
         (len2 (length (delete-dups words))))
    (and (not (equal len1 0))
         (equal len1 len2))))

(day-4-part-2 "./day-4.input")

(provide 'day-4)
;;; day-4.el ends here
