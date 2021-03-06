/*
 * Logbook of the World Query API
 *
 * LoTW provides a web service that accepts RESTful queries that report QSOs   satisfying specified criteria:    * accepted by LoTW after a specified date   * confirmed by LoTW after a specified date   * with a specified callsign   * with an operator in a specified DXCC entity   * in a specified mode   * on a specified band   * at a specified date and timeusing a specified station callsign
 *
 * API version: 1.0
 * Contact: k0swe@arrl.net
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package lotw

// YesOrNo the model 'YesOrNo'
type YesOrNo string

// List of YesOrNo
const (
	YES YesOrNo = "yes"
	NO  YesOrNo = "no"
)
