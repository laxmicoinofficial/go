# Dakibot Service for the Rover Test Network

This calls out to orbit to submit the transaction

Orbit needs to be started with the following command line param: --dakibot-url="http://localhost:8004/"
This will forward any query params received against /dakibot to the dakibot instance.
The ideal setup for orbit is to proxy all requests to the /dakibot url to the dakibot service
