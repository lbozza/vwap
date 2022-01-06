# Description

VWAP is volume-weighted average price calculator, it works by fetching trades information from an exchange websocket in real-time by subscribing to one of its channels, treating it's information and then passing it's information through an open channel that will then call a calculator that will calculate the VWAP value, the max number of trades is set to 200, we call these 200 trades: datapoints.

## Design
In this project I choose to separate each part of the software by responsabilities using hexagonal architecture. 
The code is separated in the following folders:
 - entity 
 - infra
 - usecase

The entity layer is where our domain structures are stored, the response model of our data.

The infra layer is responsible for connecting to the exchange and fetching the trading information.

The usecase layer is where our service and calculator are located. 
 
### How all of this works?

When we start our application we initialize one routine for each crypto pair we want to fetch the trading information. Each routine will have it own websocket connection and it's own service and vwap calculator instance

Once we subscribe to the websocket channel and start listening to it we broadcast this information to an internal channel that is passed through our service, our service will loop over this channel calling our vwap calculator for each new trade listened coming from the websocket.

Then each by each trade is added to our calculator and the actual vwap value is calculated in real-time and printed on our console screen.


<img width="756" alt="Screen Shot 2022-01-06 at 11 45 48" src="https://user-images.githubusercontent.com/21343976/148400769-a314d617-4afc-4ca3-bace-207f2440b65d.png">

## Running

#### With docker

Build the application

````make docker-build````

Run the application

```make docker-run```

#### Locally with golang

Install dependencies

````make deps````

Build the application

```make build```

Run the application

```./vwap```


### Tests

The application is being tested only on the vwap calculator layer whereas this is where the main core business rule is located.

#### Running tests
To run the tests you must have golang 1.13+ installed on your computer

Running tests

`````make test-cov`````

The tests will be ran and two files will be generated, one called cover.txt and one called cover.html, you can open the cover.html on your browser to check the coverage of the code by the unit tests.
