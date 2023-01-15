# University Admission Procedure

## About

It takes a lot of hard work to enroll in the university of your dreams. Although, we tend to dismiss how difficult it is for the university to handle the document volume. 
In this project, I deal with university applicants. I implement an algorithm to determine which applicants are going to be enrolled. 
At each stage, the algorithm has gradually become more complex and comprehensive!

## Stages

1. Calculate the mean score for a single applicant.
2. Set the mean score threshold and determine whether the applicants are going to be enrolled.
3. Deal with a whole list of applicants instead of just one! Find the lucky ones by comparing their GPA score.
4. Every potential student can apply for one particular department of their liking. The ranking takes place between applicants that chose the same department.
5. GPAs are no longer needed as there are final exams after all! Also, applicants can now apply to several departments.
6. More exams! Rank the applicants by the mean score of several final exams.
7. Now the applicants pass an additional exam for their department, and the best score will be chosen to determine the applicant's ranking: either the mean score or the score of the special exam.

## Learning Outcomes
- Practice loops and various mathematical operations. 
- Learn how to handle files and different types of collections such as lists (including nested lists) and dictionaries. 
- Put to use the sorting function and see how useful it can be.

## Objectives

1. Read an integer N from the input. This integer represents the maximum number of students for each department. 
2. Read the file  named applicant_list.txt. The fields in the file are separated by one space character: 
   - First name
   - Last name 
   - Physics exam score 
   - Chemistry exam score
   - Mathematics exam score
   - Computer science exam score
   - Special admission exam score
   - Department of the first priority
   - Department of the second priority
   - Department of the third priority
3. Choose the best score for a student in the ranking: either the mean score for the final exam(s) or the special exam's score. Each department makes the admission decision based on the mean score of a different set of exams:
   - physics and mathematics exams for the Physics department, 
   - chemistry exam for the Chemistry department, 
   - mathematics exam for the Mathematics department, 
   - computer science and mathematics exams for the Engineering Department, 
   - chemistry and physics exams for the Biotech department.
4. Applicants should be ranked by their exam score and, in case they have the same score, by their full name in alphabetic order.
5. Each department accepts N (maximum number of students for the department) best students from the department's ranking list. If there are fewer than N students on the department's list, all students from the list are accepted. 
6. The accepted students are removed from the general list of applicants and no longer participate in the ranking. 
7. The same procedure is repeated for the second and third priorities. If there are departments that accepted fewer than N students in the first stage of admission, these departments try to accept more students to fill all N student positions.
8. Output the admission lists to files. Create a file for each department, name it %department_name%.txt, for example, physics.txt. Write the names of the students accepted to the department and the student's best score, either the mean finals score or the special exam's score to the corresponding file (one student per line).