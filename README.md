cxmate
======

cxmate is a RESTful network API proxy service for network algorithms. If you're interested in turning a network algorithm into a robust web service, cxmate can drastically reduce the investment of time and effort required by providing the following key features:

- **Streaming support for CX, an extensible aspect oriented network network interchange format**<br>
  CX supports steaming of arbitrarily large networks, and is well suited for encoding rich networks through the use of aspects. cxmate reads and writes streams of CX, allowing high throughput with lower memory consumption. Your service need not know the exact details of CX to take advantage of its power and flexibility. cxmate supports one-to-one, one-to-many, and many-to-many network algorithms. You decide how many networks cxmate will receive and send.
  
- **Work with native objects in your language of choice instead of HTTP request and responses**<br>
  cxmate provides an efficient translation between the CX interchange format and objects native to the proxied service. By the time cxmate calls your service, your code will receive a stream of easy to use element objects containing network pieces, algorithm parameters, and formatted errors to work with. cxmate only expects a stream of native objects in return. Never work with raw HTTP again.
  
- **A fully RESTful JSON HTTP interface managed by cxmate on behalf of your service**<br>
  Any service proxied by cxmate need not write any HTTP handlers, URL parsers, or deal with any of the boilerplate associated with creating a RESTful web service. Clients will have full access to this popular method of interfacing with your service through cxmate, allowing you to focus on writing and maintaining service logic instead of interfaces.
  
- **Algorithm parameters and error handling made easy**<br>
  When cxmate receives a request, query string parameters are automatically translated to key/value elements and streamed to your service like any other object. Any errors detected by cxmate while parsing the incoming network and parameters will also be turned into error objects your service can then decide to send back to the client, handle internally, or ignore.
  
- **Service insights via automated metrics gathering and logging**<br>
  cxmate exposes a plethora of useful statistics about itself and the proxied service via its RESTful HTTP API, allowing service authors to monitor the health and usage of their service over time.  
 
 cxmate is a subproject of Cytoscape and the Ideker Lab at the University of California, San Diego. cxmate greatly decreases the time bioinformaticians, computer scientists, and researchers from other disciplines spend writing code, allowing them to focus on their algorithms and providing biological value to research community. cxmate also decreases the time spent creating services for features used by tens of thousands of Cytoscape users every day.

Installation
------------

While we recommend eventually running cxmate and your service in Docker containers for maximum portability and deployability on the Cytoscape Cyberinfrastructure, we also precompile cxmate binaries for popular platforms for testing and development:

- Download a precompiled binary for your platform [here](https://github.com/ericsage/cxmate/releases)
- Run cxmate in a docker container with the [official Docker Hub image](https://hub.docker.com/r/ericsage/cxmate/)

Getting Started
---------------

cxmate works with a number of popular programming languages. Currently, our official tutorial walks through building a cxmate proxied service with Python. If you're interested in using cxmate with a different language, please [contact us via email](eric.david.sage@gmail.com)! We'd love to help support your use case and create a guide for your language.

- The official cxmate [Python tutorial](https://github.com/ericsage/cxmate/wiki/Python-tutorial)

Configuration
-------------
All configuration is done through environment variables that cxmate reads once on startup. Once started, cxmate will print out the read values for easy debugging. The supported configuration parameters are as follows:

```
//cxmate's address
LISTENING_ADDRESS        //Default 0.0.0.0
//cxmate's port
LISTENING_PORT          //Default 80
//The address of the service cxmate will proxy
//SERVICE_ADDRESS       //Default 127.0.0.1
//The port of the service cxmate will proxy
//SERVICE_PORT          //Default 8080
//Will cxmate receive more than one network on behalf of this service?
//RECEIVES_COLLECTION   //Default false
//ONLY ACTIVE IF RECIEVES_COLLECTION IS TRUE, how many networks will cxmate receive?
//EXPECTED_NUM_NETWORKS //Default 1
//What aspects from the CX should cxmate pass to the service?
//RECEIVES_ASPECTS     //No Default, MUST BE SET, should be a commma seperated list of aspect names
//Will cxmate send more than one network to the client?
//SENDS_COLLECTION     //Default false
//ONLY ACTIVE IF SENDS_COLLECTION IS TRUE, how many networks will cxmate send?
//SENDS_NUM_NETWORKS   //Default 1
//What aspects from the CX should cxmate forward to the client?
//SENDS_ASPECTS        //No Default, MUST BE SET, should be a comma seperated list of aspect names
```

Note that `RECEIVES_ASPECTS` and `SENDS_ASPECTS` must be set or the behaviour of cxmate will be unpredictable. You should consider what aspects your service and clients expect, and set them accordingly, which will allow cxmate to function properly.

cxmate currently supports the following aspects for sending and receiving:

```nodes edges nodeAttributes edgeAttributes networkAttributes cartesianLayout```

You do not need to declare errors or parameters as aspects, they are always available in every service.

Contributors
------------

We welcome all contributions via Github pull requests. We also encourage the filing of bugs and features requests via the Github [issue tracker](https://github.com/ericsage/cxmate/issues/new). For general questions please [send us an email](eric.david.sage@gmail.com).

License
-------

cxmate is MIT licensed and a product of the [Cytoscape Consortium](http://www.cytoscapeconsortium.org).

Please see the [License](https://github.com/ericsage/cxmate/blob/master/LICENSE) file for details.
