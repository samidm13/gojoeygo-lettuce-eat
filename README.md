# Lettuce Eat: food ordering web app that allows group ordering

## Summary:
Welcome to the repository for our final engineering project at Makers Academy. Our team 'Go Joey Go' is made up of four apprentices. Our aim was to develop a web app where food can be ordered by individuals as part of a group order.

We believe that the key ingredients for our success as a team are sharing ideas and knowledge, working together in pairs and offering help to each other whenever required.

## A one-sentence description of the MVP interaction
    The order initiator can "Create Order" that selects the restaurant and generates a unique order token shared with whomever they want to be part of the group order.

## User stories implemented

1. As a new user

   So I can set up a food group order 

   I can sign up with an encrypted password and no token
   

2. As an existing user

   So I can set up a food group order 

   I can log in with an encrypted password and no token
   

3. As a new user

   So I can order food as part of a group

   I can sign up with an encrypted password and a valid not expired token
   

4. As an existing user

   So I can order food as part of a group

   I can log in with an encrypted password and a valid not expired token
   

5. As a new user

   So I cannot join a non-existent food group order

   I cannot sign up with an invalid or expired token
   
   
6. As an existing user

   So I cannot join a non-existent food group order

   I cannot log in with an invalid or expired token
   

7. As an existing user

   So I cannot access other users' confidential data 

   I cannot log in with incorrect email or password 
   

8. As a new user

   So I have a unique account  

   I cannot signup with the same email more than once 


9. As a user

   So my data are securely stored

   I can log out from my account
   
   
10. As an order initiator

    So I can set up a group order

    I can "Create Order" that includes the restaurant, delivery address, time expiry and generates a
   unique order token
   

11. As an order initiator
 
    So I can join a group order

    I can "Add dishes" to the order I've just set up
   

12. As an order initiator

    So I can join a group order

    I can "Add dishes" to the order I've just set up
   

13. As an order initiator

    So that I can have an overall image of each order

    I can view for a specific order that I set up: the users, the dishes ordered per user, the total and the time when the order  will be placed 
   
   
14. As an order initiator

    So that I can have an overall image of each order

    I can view for a specific order that I set up: the users, the dishes ordered per user, the total and the time when the order  will be placed 
   
   
15.  As a user 

     So I can join a group order

     I can "Add dishes" to the order identified with the token I've just signup/logged in
    
    
16.  As a user 

     So I know what I've ordered

     I can view the dishes ordered , the total and the time when the order will be placed 
    
   
## Technologies Used:
- UI - HTML, CSS and JavaScript
- Backend - Golang & Gin Gonic for web REST framework
- Text editors: Atom / VS Code
- Testing framework: Go’s inbuilt testing functionality
- Database: PostgreSQL, PQ, TablePlus, SQL
- Version control: GitHub

## Database Setup:
When working with this repo for the first time you must set up the database on your machine. We assume here that Postgres is installed.
To set up the database:
1. From the command line run psql.
2. Execute the commands stored in db/migrations.

## How to install and run the app:
   1. Make sure you have installed 'golang' on you computer.
    - [Go Website](http://www.golang.org/dl/)
   2. clone the project repository in the default golang path which is $HOME/go/src
   3. Run the 'dep ensure' command from inside the project folder to make sure all the dependencies are installed.
   4. Set up the databased as mentioned above.
   5. Run the app by calling ./app from the command-line.

## How to run tests:
The following commands can be run from command-line:
- go test

You can output print additional information about test function using verbose -v command:
- go test -v

Command fo checking test coverage:
- go test -cover

Command to generate detailed test coverage report:
- go test -coverprofile=coverage.out

Command for generating the coverage profile in HTML file:
- go tool cover -html=coverage.out

## Process followed to learn Go language:
 - Learned golang by translating the projects we did in Ruby language.
 - Used the resources like 'Codecademy' and 'A Tour of Go'.
 - Tried to learn by TDD approach.
 - Learned by trial and error.
 - Followed debugging approach.
 - Reaching out for support.
 - Kept the thing as simple as we could understand.
 - We used pairing technique to learn and implement new feature.
 - Learned by sharing knowledge and resources.
