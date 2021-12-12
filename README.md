# Realtime VWAP Calculator

A realtime VWAP calculation engine. It uses coinbase websocket feed as its default provider for real time data.

## How to run

To run the app

`make run`

To run all the tests

`make test`

To build 

`make build`

## Project Structure

Project structure is a simple variation of post-adapter(hexagon) architecture.

### `/cmd`
Main applications for this project.
The directory name for each application  name of the executable.

### `/internal`
Private application and library code.
This is the code we don't want others importing in their applications or libraries.

### `/ws_client`
Contains Websocket framework codes, JSON schemas, protocol definitions

### `/core`
Business logic stores in this package

## Desing and Architecture

The Domain-Driven-Design approach was used for the services. The project's main focus is located under the core package. Also, the Port-Adapter architecture is used for implementing interfaces and their integrations with the domain.
The core package is dependent on interfaces, not technology-related structs. So if interfaces change from Websocket Client to MQ consumer, there will be no change for the domain layer.
The domain layer is agnostic to any technology, it is only related to business logic.

Before the compile process, dependencies creation and integration script generated by wired package. When a new interface is required, we only need to implement the new adapter and attach it.

Because of time constraints, I need to make some assumptions and simplifications. But they can be easily improvable because of design and architecture.

