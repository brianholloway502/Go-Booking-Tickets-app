This is a ticket booking application built with Go. It is a CLI program that will do the following:
  - Ask the user for their first name
  - Ask the user for their last name
  - Ask the user for their email
  - Ask the user how many tickets they want to book for the coference. Ticket count is hardcoded to 50 inside the program.
      - Then perform the following calculations output:
          - Apply field validation checks to first and last name, email and ticket count. If error found, display error.
          - If no error found, display message stating tickets were purchased and sent to the user's email and deduct the ticket count from the ticket total which is a hardcoded value in the program
            NOTE: There is no back end to this program, so no actual tickets are being purchased and nothing is being saved to the database. Everything is being kept tracked internally during the program run.
          - After user successfully display the information for the user that purhased the ticket and start the program up again for the next user to reserve tickets. 
