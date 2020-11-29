[![PkgGoDev](https://pkg.go.dev/badge/github.com/k0swe/lotw-qsl)](https://pkg.go.dev/github.com/k0swe/lotw-qsl)
[![Go Report Card](https://goreportcard.com/badge/github.com/k0swe/lotw-qsl)](https://goreportcard.com/report/github.com/k0swe/lotw-qsl)

# Go API client for Logbook of the World

A GoLang client library for the Logbook of the World query API. LoTW provides a
web service that accepts RESTful queries that report QSOs satisfying specified 
criteria:

  * accepted by LoTW after a specified date
  * confirmed by LoTW after a specified date
  * with a specified callsign
  * with an operator in a specified DXCC entity
  * in a specified mode
  * on a specified band
  * at a specified date and timeusing a specified station callsign
  
The API is documented
[here](https://lotw.arrl.org/lotw-help/developer-query-qsos-qsls/).

This client library was generated based on the OpenAPI specification in
the `api/openapi.yaml` file. However, the API itself is not well-described by
OpenAPI, so the generated library is supplemented with `wrapper.go`.

A simple application to demonstrate how to integrate the library is located in
`cmd/lotw-qsl/main.go`.
