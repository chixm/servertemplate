# ATAGO
 All in one
 Template
 Assembled
 GO lang

 - author Chixm

# Web Server Template
ATAGO is an Web Server Template with very often used libraries.
The ATAGO is assembled to quickly make an web page.

It contains

html layout
static file server (files like css and javascript files)
database
websocket
keyvalue store(redis integration)
react

So, making a copy of The ATAGO will easily brings you to start developing your web service. 

# requirements
Install Golang, Source code of Golang is written in version 1.14.(https://golang.org/)
Install Node.js, We require npm command to use React. Developed with npm version 6.12 (https://nodejs.org/en/download/)

# Build Server
Command
npm install 
in terminal. then
npm build
to create server binary.

this project is developed in Windows PC. So the scripts in package.json is basically work only on Windows. 

## To Run Server
To run server. Just run Executable named atago.exe in src/main directory.

## About Golang
Golang version 1.14 is used to develop.

## About React
React file in this project is inside src/view.
React file is compiled to normal javascript from Typescript and Webpack put them togather and put into resource deirectory. 

# Updates
2019 Created basic web template with Golang.
2020 Added React to Project.