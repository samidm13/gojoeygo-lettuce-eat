# Lettuce Eat: food ordering web app that allows group ordering

## Summary:
Welcome to the repository for our final engineering project at Makers Academy. Our team 'Go Joey Go' is made up of four apprentices. Our aim was to develop a web app where food can be ordered by individuals as part of a group order.

We believe that the key ingredients for our success as a team are sharing ideas and knowledge, working together in pairs and offering help to each other whenever required.

## A one-sentence description of the MVP interaction
    The order initiator can "Create Order" that selects the restaurant and generates a unique order token shared with whomever they want to be part of the group order.

## User stories implemented
1. As a user

   So I can order food as part of a group

   I can sign up with an encrypted password

2. As a user

   So I can order food as part of a group

   I can log in with an encrypted password


3. As a user

   So my data are securely stored

   I can log out from my account


4. As an order initiator

   So I can set up a group order

   I can "Create Order" that include the restaurant and generates an unique order token

5. As a User

   So that I can order food as part of a group

   I can login or signup with the valid order token

6. As a user

   So that I can add my order to the basket

   I can select menu dishes from the restaurant

## Technologies Used:
- Primary programming language: Go
- Front end languages: HTML, CSS
- Text editors: Atom / VS Code
- Web framework: Gin
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
