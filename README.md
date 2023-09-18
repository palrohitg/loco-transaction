# :metal:	Loco Transaction Service 
- Transaction Service will store the different transaction  by user and Parent child relationships between transaction.



## High Level Design:
![LocoTransaction](https://github-production-user-asset-6210df.s3.amazonaws.com/40069230/268563842-67d59d56-4385-4f96-9e0b-a12e5d36bf3d.png?X-Amz-Algorithm=AWS4-HMAC-SHA256&X-Amz-Credential=AKIAIWNJYAX4CSVEH53A%2F20230918%2Fus-east-1%2Fs3%2Faws4_request&X-Amz-Date=20230918T055259Z&X-Amz-Expires=300&X-Amz-Signature=f3575dd9845c6095c66cacbe4b0fc57b01c0cb07fb5bd237bbba0870178c2f6d&X-Amz-SignedHeaders=host&actor_id=40069230&key_id=0&repo_id=294484791)

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