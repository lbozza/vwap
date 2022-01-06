# Description

VWAP is volume-weighted average price calculator, it works by fetching trades information from an exchange websocket in real-time by subscribing to one of its channels, treating it's information and then passing it's information through an open channel that will then call a calculator that will calculate the VWAP value, the max number of trades is set to 200, we call these 200 trades: datapoints.

## How it works
In this project I choose to separate each part of the software by responsabilities using hexagonal architecture, we have the infrastructure layer that is responsible for opening the websocket connection and then subscribing to a channel and listening to the trades that are being made in realtime, then this information is passed through a service that calls the vwap calculator, the calculator will receive a new datapoint that represents each trade, it will then sum this trade volume and price to a datapoint list that will calculate the vwap as each new datapoint is added to the list. Once the datapoint list reaches the size of 200 we will only replace the oldest datapoint with the newest trade and so on.

In this VWAP Calculator we are fetching trades from three crypto currency pairs which are "BTC-USD", "ETH-USD" and "ETH-BTC", each pair is ran on a separate routine, so each one has its own service and calculator instance.


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
