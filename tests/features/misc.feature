# vim: ts=4 sw=4 sts=4 noet
Feature: Root command is informational only

	@smoke
	Scenario Outline: Can get help info by running termwg with inquisitive inputs
		When I run `termwg <INPUT>`

		Then it should pass with "Usage:"

		Examples:
			| INPUT  |
			|        |
			| -h     |
			| --help |
