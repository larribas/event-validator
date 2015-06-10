# Event Validator

The Event Validator is a web service aimed to define the information and structure of business events within a company.

## Motivation

Moder companies tend to have dozens of internal services and tools to support the workflow of different departments. Oftentimes, these services communicate with one another and react to meaningful events that happen elsewhere.

Implementing a successful event-driven communication model has some challenges:
* Event producers and consumers must agree on a certain format, structure and information to send with each event.
* Such agreement may be implicit or explicit. We have found the latter to be more robust an provide accountability.
* Should the requirements of a certain type of event change, all parties involved must coordinate the implementation of those changes (which can be painful).

The Event Validator deals with these problems in the most simple way:
* Any party can define the information and structure of an event type, via a web service (that could, potentially, be connected to an internal toolbox).
* Definitions use a certain format, unique for the whole company. Each company may choose the format they prefer (see _Technologies_). Usually, the validation format will be capable of automatically validating an event (e.g. as an XML Schema or a JSON Schema), reducing human mistakes.
* Definitions are immutable. Once created, they cannot be modified nor deleted. This behavior helps producer and consumer applications to reason about the correctness of their implementation.
* Definitions are versioned. Whenever a new definition for an event type is created, it is assigned an incremental numeric version (unique within that type). This avoids the need to coordinate the implementation of a requirement change. Applications can rely on the event version to remain backwards-compatible.


## HTTP API

The Event Validator exposes a straightforward Web API with the following endpoints:

__POST /validators/:type__
To create a new validator for a certain event type.
The request should be of type json, with the following parameters:
```json
{
    "rules": "string"
}
```
Where 'rules' is a string specifying the definition (by default, a JSON schema)

When successful, the endpoint will respond:
```json
{
    "version": "int"
}
```


__GET /validators/:type/versions__
To get the available versions for a certain event type.
If the type has at least one version, it will respond with:
```json
{
    "min_version": "int",
    "max_version": "int"
}
```


__GET /validators/:type/versions/:version__
To get the validation rules for a certain definition (type + version combination).
If the specified definition exists, it will respond with:
```json
{
    "rules": "string"
}
```


__Error format__
Should any of the previous endpoint result in a failure (due to a wrong request format, invalid parameters, or the nonexistence of the requested resource, the endpoint will respond with a 400 status code, and the following application/json structure:
```json
{
    "error_type": "string",
    "error_message": "string"
}
```



## Technologies

The default implementation uses JSONSchema-based validators, and Redis as a storage mechanism. However, it is very easy to plug in different validator formats or storage mechanisms, adapted to each company's particular technology stack. To do so:
* Implement the `domain.Repository` or `domain.FormatChecker` interfaces, within or without the `repositories` and `formats` folders.
* At `main.go`, get the `SetCurrentEnvironment()` function to return a `domain.Environment` with the desired components.


## Tests

Some of the implementations (e.g. Redis) may need certain infrastructure and environment variables to be set up. To execute such tests, a docker-based script is provided under `bin/test-with-docker`.