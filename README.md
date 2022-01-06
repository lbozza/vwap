VWAP is volume-weighted average price calculator, it works by fetching trades information from a exchange websocket in real-time by subscribing to one of its channels, treating it's information and then passing it's information through an open channel that will then call a calculator that will calculate the VWAP value, the max number of trades is set to 200, we call these 200 trades: datapoints.

In this project I choose to separate each part of the software by responsabilities using hexagonal architecture, we have the infrastructure layer that is responsible for opening the websocket connection and then subscribing to a channel and listen to the trades that are being made in realtime, then this information is passed through a service that calls the vwap calculator, the calculator will receive a new datapoint that represents each trade, it will then sum this trade volume and price to a datapoint list that will calculate the vwap as each new datapoint is added to the list. Once the datapoint list reaches the size of 200 we will only replace the oldest datapoint with the newest trade and so on.

In this VWAP Calculator we are fetching trades from three crypto currency pairs which are "BTC-USD", "ETH-USD" and "ETH-BTC", each pair is ran on a separate routine, so each one has its own service and calculator instance.

