# Assignment
### Which is done
* All required functionalities are implemented
### How to use
* [STEP 1]
  * [go run main.go] for running up server
* [SETP 2]
  * [curl localhost:8080/view/get_users] for getting users
  * [curl localhost:8080/view/get_products] for getting products
  * [curl localhost:8080/buy/cat/lisa/50 --request "PATCH"] for buying a cat as lisa with 50 points
  * [curl localhost:8080/buy/cat/lisa/100 --request "PATCH"] for buying a cat as lisa with 100 points
  * [curl localhost:8080/activity/normal --request "PATCH"] set activity level to normal, for changing level of discount to 0.05 and point conversion ratio to 1.
  * [curl localhost:8080/activity/festival --request "PATCH"] set activity level to festival, for changing level of discount to 0.06 and point conversion ratio to 1.2
  * [curl localhost:8080/buy/cat/bruch/200 --request "PATCH"] for buying a cat as bruch with 200 points, therefore extra 90% discount activated
