# delay-api

## Run Instructions:
* Run *make help* inside the *deployment* folder to see a list of available commands.
* Inside the *deployment* run *make* to setup the delay-engine


## Usage:
* Example usages: 
  * curl -X GET --url localhost:8080/5s
    * This will replicate a server response with a delay of 5 seconds.
  * curl -X GET --url localhost:8080/5ms
    * This will replicate a server response with a delay of 5 milliseconds.

