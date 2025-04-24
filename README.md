Go Conference Booking System

This project is part of my practice and learning journey in Go programming. Through this project, I am covering various Go concepts to deepen my understanding and improve my programming skills.

Overview
This simple Go-based application simulates a conference ticket booking system. Users can input their personal information and the number of tickets they wish to book. The program validates the inputs, updates the ticket availability, and simulates sending a confirmation email using Go’s concurrency (goroutines).

Learning Highlights
Key Concepts Covered:
Local Setup
Installed Go and configured my editor to get started with Go development.

Writing Our First Program & Structure of a Go File
Created the initial program structure, including understanding the basic syntax of a Go file.

Variables & Constants in Go
Defined variables and constants for tracking conference details and user bookings.

Formatted Output
Used fmt.Printf for formatted output to display user and booking information clearly.

Data Types in Go
Learned about various data types (e.g., string, int, uint) and their usage in Go.

Getting User Input
Used fmt.Scan to take user input for booking tickets.

What is a Pointer?
Explored Go's pointer concept and how it affects memory management.

Book Ticket Logic
Implemented the logic for booking a ticket, including user validation and updating the remaining tickets.

Arrays & Slices
Worked with arrays and slices to store and manage booking data.

Loops in Go
Implemented loops to manage continuous user input until tickets are sold out.

Conditionals (if / else) and Boolean Data Type
Used if / else statements for conditional checks on user input (e.g., ticket number, email validity).

Validate User Input
Created a function to validate user input, ensuring the correct format for names, emails, and tickets.

Switch Statement
Incorporated the switch statement for handling different conditions in the program flow.

Encapsulate Logic with Functions
Organized the logic into functions to make the program more readable and maintainable.

Organise Code with Go Packages
Split the logic into multiple Go packages to better organize the code and adhere to Go's modular structure.

Scope Rules in Go
Learned about scope rules in Go and how variables are scoped within functions and packages.

Maps
Used maps to store and look up booking data efficiently.

Structs
Created structs to define user data (e.g., name, email, ticket number) and encapsulate related information.

Goroutines - Concurrency in Go
Practiced concurrency in Go by sending confirmation emails in parallel using goroutines.

How to Run
Install Go (if not already installed): Check Go Installation Guide

Clone the repository to your local machine.

Open the project in your preferred editor (e.g., VS Code).

Run the program using the following command:

go run main.go

main.go: The main program file that contains the logic for booking tickets and user interactions.

helper.go: Contains functions for validating user input.

Challenges Faced
Handling user input and ensuring valid data was a key challenge. I implemented various checks and validations for user details.

Understanding and implementing Go's concurrency with goroutines was another area where I expanded my knowledge, especially in sending emails.

Next Steps
I plan to add additional features, such as handling multiple conferences and integrating a database to store booking information.

Further improve the validation of user input and enhance error handling.

Conclusion
This project has been a valuable learning experience as I get hands-on practice with Go. It’s helping me understand fundamental programming concepts while also diving deeper into Go-specific features like goroutines, pointers, and struct types. I look forward to building more complex applications as I continue learning!
