Support-go
======

Contents
-----
The purpose of this application is to update a light status page. This go app will consist of a get service, to fetch current light status, and put service to update the light status based on radio button clicks on an admin page.

Setup
-----
Run
```
$ vagrant up
```
to create a virtual box with ubuntu 14.04LTS which has mysql server and go 1.4.2 installed, as well as a go workspace.

Running
-----
Run
```
$ vagrant ssh
```
to connect to your provisioned box, once there cd into the top level director of this project and
Run
```
$ PORT=(pass your port here) go run *.go
```
This should display the _"listening..."_ text.
Visit your localhost with the port you chose to access this app
/statuses - will show current queue status
/statuses/?querystring - will show nothing in the browser but log the variable to the console window
