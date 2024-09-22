# Go-CRUD-API
Developing API's in GoLang using Gin Server and deployment using LocalStack to emulate the AWS cloud env

Pre-requisites
This development uses Golang to create a server and access MySQL.
Using Gin server, we would try to deploy the server into an ec2 instance using LocalStack.

Please follow the steps below to install the required software :
1. Visual Studio Code - IDE to develop the server and code in GoLang.
   Use the following link to download the IDE. (https://code.visualstudio.com/download)

2. Go Programming Language
   Use the link (https://go.dev/doc/install) to get the latest GoLang compiler.
   Make sure to update the env Path variable to the bin path for the downloaded installer (in case of windows).
   For Mac, follow the appropriate steps to add to the path variables.

3. Download the MySQL community server and MySQL workbench using the following links
   https://dev.mysql.com/downloads/mysql/8.0.html
   https://dev.mysql.com/downloads/workbench/

4. Make sure to have the docker desktop downloaded to enable LocalStack accessibility
   Download and set the docker desktop here https://www.docker.com/products/docker-desktop/

5. Set up LocalStack by following the steps here after signing up
   https://app.localstack.cloud/getting-started
   * Download the CLI
   * And try to install the docker extension of LoaclStack to ease the use of images.
