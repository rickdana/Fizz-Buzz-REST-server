#  Fizz-Buzz REST server  
  
Fizz-Buzz REST server is a Golang base Rest API. The server exposes a service which Accepts five parameters : three integers int1, int2 and limit, and two strings str1 and str2 and   
Returns a list of strings with numbers from 1 to limit, where: all multiples of int1 are replaced by str1, all multiples of int2 are replaced by str2, all multiples of int1 and int2 are replaced by str1str2.  
  
## Installation  
  
#### Prequisit : Install a version of Golang. visit [How to install golang](https://golang.org/doc/install)

 - Clone the projet to your golang app folder 
 - Create server config file (see sample below)
 ```yml  
 # Server configurations  
server:  
  host: "localhost"  
  port: 8443  
  
dataSource:  
  fileStore: "${path_to_file}/fizz_buzz_request.txt"  
  
logger:  
  filePath: "${path_to_file}/fizz_buzz.log"  
  
auth:  
  username: "admin"  
  password: "changeit"
```  
 - Launch the server 
```bash  
go run main.go
```
 ## Usage
 The API offert 3 services:
 
 - **fizz-buzz** 
	 - end point: 
	 ```bash
	 POST: /api/v1/fizz-buzz
	 ```
	 - Input
		```yml  
		{
			"firstMultiple": 3,
			"SecondMultiple": 5,
			"limit": 15,
			"fizzWord": "Hello",
			"buzzWord": "World"
		}
		```  
	- output
		```yml  
		["1","2","Hello","4","World","Hello","7","8","Hello","World","11","Hello","13","14","HelloWorld"]
		```  
	- example
		```bash
		 curl --location --request POST 'https://localhost:8443/api/v1/fizz-buzz' \  
		  --header 'Authorization: Basic YWRtaW46Y2hhbmdlaXQ=' \  
		  --header 'Content-Type: application/json' \  
		  --data-raw '{  
		  "firstMultiple": 3,  
		  "SecondMultiple": 5,  
		  "limit": 15,  
		  "fizzWord": "Hello",  
		  "buzzWord": "World"  
		}'
		```
 
 - **statistics**
 statistics endpoint allowing users to know what the most frequent request has been
	 
	 - input
	    ```bash
	   No args
	    ```
	  - output
		  ```yml
		  [
		    {
		        "fizzBuzzDto": {
		            "firstMultiple": 4,
		            "secondMultiple": 9,
		            "limit": 100,
		            "fizzWord": "Hello",
		            "buzzWord": "Worl"
		        },
		        "hits": 7
		    },
		    {
		        "fizzBuzzDto": {
		            "firstMultiple": 7,
		            "secondMultiple": 9,
		            "limit": 100,
		            "fizzWord": "Hello",
		            "buzzWord": "Worl"
		        },
		        "hits": 1
		    },
		    {
		        "fizzBuzzDto": {
		            "firstMultiple": 3,
		            "secondMultiple": 5,
		            "limit": 15,
		            "fizzWord": "Fizz",
		            "buzzWord": "Buzz"
		        },
		        "hits": 1
		    }
		]
		  ```
	  - example
		  ```bash
		  curl --location --request GET 'https://localhost:8443/api/v1/statistics' \
		--header 'Authorization: Basic YWRtaW46Y2hhbmdlaXQ=' \
		--header 'Content-Type: application/json' \
		--data-raw '{
		    "firstMultiple": null,
		    "SecondMultiple": 5,
		    "limit": 15,
		    "fizzWord": "Fizz"
		}'
		  ```
- **health-check**
	Server heath check endpoint 
	- 	 input
		    ```bash
		   No args
		    ```
	- output
		```yml
		{
		    "up": "OK"
		}
		```
	- example
		```bash
		curl --location --request GET 'https://localhost:8443/api/v1/health-check' \
		--header 'Content-Type: application/json' \
		--data-raw '{
		    "firstMultiple": null,
		    "SecondMultiple": 5,
		    "limit": 15,
		    "fizzWord": "Fizz"
		}'
		```

## Contributing  
Pull requests are welcome. For major changes, please open an issue first to discuss what you would like to change.  
  
Please make sure to update tests as appropriate.  
  
## License  
[MIT](https://choosealicense.com/licenses/mit/)