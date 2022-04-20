# Custom validate

[![CI](https://github.com/posteris/commons/actions/workflows/build.yml/badge.svg)](https://github.com/posteris/commons/actions/workflows/build.yml)
[![Go Reference](https://pkg.go.dev/badge/github.com/posteris/commons.svg)](https://pkg.go.dev/github.com/posteris/commons)
[![GitHub license](https://badgen.net/github/license/posteris/commons)](https://github.com/posteris/commons/blob/master/LICENSE)
[![GitHub release](https://img.shields.io/github/release/posteris/commons.svg)](https://GitHub.com/posteris/commons/releases/)


Project dedicated to stadardize some patterns like erros and validations that can be utils and shared by many others projects.

The next sessions was dedicated to show how to use each features.

## Validation

The validation model receives a model (struct) that contains the validation tag, in case of some validation fail the return will be a _ValidationError_ array, other else _nil_.

```go
import commom "github.com/posteris/commons/validation"

//define struct model
type User struct {
	Name    string `validate:"required,min=3,max=256"`
	Surname string `validate:"required,min=3,max=256"`
	Email   string `validate:"omitempty,email"`
}

//create a model based on struct
userModel := User{
    Surname: "da Silva"
    Email: "dasilva@gmail.com"
}

err := commom.ValidateModel(userModel)
```

## Errors

The error package has 2 possible erros that can be understand bellow.

### DefaultError

the default error is user to return to the client a json with just one error message.

```go
import commom "github.com/posteris/commons/errors"

err := commom.CreateDefaultError("some error message")
```
### ValidationError

The validationError have 2 fields  the are showed bellow:
* __Field:__ the validated field
* __Message:__ the error message.
```go
import commom "github.com/posteris/commons/errors"

err := commom.CreateValidationError("field-name", "some error message")
```

## Query Parameter

Set of functions to operate over query parameter

### IsAsyncRequest

check if the async parameter is sent as query parameter

```go
async := IsAsyncRequest(fiberCtx)

```

## Software Quality
[![Quality Gate Status](https://sonarcloud.io/api/project_badges/measure?project=posteris_commons&metric=alert_status)](https://sonarcloud.io/summary/new_code?id=posteris_commons)
[![Maintainability Rating](https://sonarcloud.io/api/project_badges/measure?project=posteris_commons&metric=sqale_rating)](https://sonarcloud.io/summary/new_code?id=posteris_commons)
[![Coverage](https://sonarcloud.io/api/project_badges/measure?project=posteris_commons&metric=coverage)](https://sonarcloud.io/summary/new_code?id=posteris_commons)
[![Security Rating](https://sonarcloud.io/api/project_badges/measure?project=posteris_commons&metric=security_rating)](https://sonarcloud.io/summary/new_code?id=posteris_commons)
[![Vulnerabilities](https://sonarcloud.io/api/project_badges/measure?project=posteris_commons&metric=vulnerabilities)](https://sonarcloud.io/summary/new_code?id=posteris_commons)


This lib use [Sonarcloud](https://sonarcloud.io/) to help understend the code quality and security.

In conjunction with [Sonarcloud](https://sonarcloud.io/), this lib uses [Horusec](https://horusec.io/) which blocks CI/CD in any vulnerability incidence


## Benchmark

Thinking in the software quality, the __benchmark regression__ was created. It's can be viewed at the link bellow.

[Performance Regeression](https://posteris.github.io/commons/dev/bench/)