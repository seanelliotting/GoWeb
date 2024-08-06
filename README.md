## GoWeb

### this is educational project 

This project was created for teaching ___Golang___.  
It shows how to create a web server and HTML page.

First install Go on your computer.  
Check the version with the command.

```shell
$ go version
```

If you need to install of Go here are the links.

[Removing an old version of Go](https://go.dev/doc/manage-install)  
[Installing Go](https://go.dev/doc/install#tarball)


---

The project ___GoWeb___ has the following directory structure.

- GoWeb  
   - cmd
     - web
        - handlers.go
        - main.go
   - pkg
   - ui
     - html
     - static

---

### ___Catalog desciptions___

- cmd 
  - directory will contain the application-specific  
   code for the executable applications in the project.

- pkg
   - directory will contain the ancillary non-application-specific code used in the project. For example models database SQL.

- ui
   - directory will contain the user-interface assets used by the web application. Specivically, the ui/html directory will contain HTML templates, and the ui/static directory will contain static files (like CSS and images)

   ---

   To start the server you need to copy GoWeb directory to your home directory or another folder. Then go to the cmd/web directory and run the command:

   ``` shell
   $ cd $HOME/GoWeb
   $ go run ./cmd/web
   ```

   After this, the server will start on port 4000.
   
   Open to your browser and type in the path

``` 
   http://localhost:4000
```

A web page of the small application will be displayed on the screen.

---
## Attention! Important!
 When downloading a project from GitHub, the word "-main" will be added to the project name. The project name after downloading will be "GoWeb-main". The word "-main" must be removed from the name so that the project is simply called "GoWeb". This is necessary so that the route in the team matches
  ``` shell
  $ cd $HOME/GoWeb
  $ go run ./cmd/web
INFO	2024/07/24 16:00:51 Starting server on :4000
  ```
If the server starts there will be a line   
INFO	2024/07/24 16:00:51 Starting server on :4000
