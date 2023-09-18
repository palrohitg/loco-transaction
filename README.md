# :metal:	Loco Transaction Service 
- Transaction Service will store the different transaction  by user and Parent child relationships between transaction.



## High Level Design:
![LocoTransaction](https://github.com/palrohitg/Machine-Coding/assets/40069230/49cbbdb0-df8a-4f22-bbfe-fd82c12bce8e)

## :computer: Tech Stack

* [Golang](https://go.dev/)
* [Postgres](https://www.mysql.com/)
* [Gin](https://gin-gonic.com/)
* [Docker](https://www.docker.com/)



## :running_woman: Local Installation Guide :

1. Clone the repository by using Git Client:

         git clone https://github.com/palrohitg/logservice.git

2. To Setup and Run Application + DataBase + CRON:

        chmod +x application_container_start.sh
        ./application_container_start.sh

3. To Setup and Run Load Testing / Request:

        chmod +x test_container_start.sh
        ./test_container_start.sh


## Logs of Docker Shell
![LogService](https://github-production-user-asset-6210df.s3.amazonaws.com/40069230/250387944-decd7b32-12a0-4ffe-ae77-329ebcfb549e.png)
![LogService](https://github-production-user-asset-6210df.s3.amazonaws.com/40069230/250387972-2aabc407-d806-41a4-a024-d3b8a56069e7.png)

## 📜 LICENSE

[MIT](https://github.com/palrohitg/logservice.git) 