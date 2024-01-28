# Alcohol Consumption Tracker

Technical interview question

## Running the web-app

## Description

This question requires you to practically create a small system. It is recommended that the system
is built in a containerised fashion, ideally with each service in a container and a docker-compose file
to configure and launch them all in concert. Alternatively, they could be implemented as Kubernetes
deployments for minikube on your local machine. Non-dockerized applications communicating
over your computer’s local network would also fulfill the requirements of the question, but are not
advised, as the lack of containerisation does not guarantee the configuration will work on a
different computer to your own.

You may use any frameworks and any languages.

### Task

Create a system to keep track of bar patrons’ alcohol consumption. The system must include a frontend, a
backend including an API, and a database. Via the UI, you should be able to create and remove patrons. You
may add a drink to the patron’s tally of drinks for the night by using the following API/database
(https://www.thecocktaildb.com/api.php) to generate the list and retrieve information about the drinks. Use
this API’s information to keep track of how much alcohol the patron has consumed. Use a mapping between
common alcohol types, their amounts and ABV consumed by the patron to do this.
Using a formula you devise, when the alcohol saturation of a given user is too high, the patron’s ID should be
coloured red (on a page displaying all the patron’s IDs, for example, in a grid). A patron’s saturation should be able to be fetched at any time from the API. Their saturation should be proportional to the alcohol consumed, inversely proportional to their body mass (a property of the patron), and inversely proportional to time since the individual drink was consumed, in a decaying manner.