package main

import (
	"yb-get/cmd"
)

/*
	TODO

Immediate:
- Add verbose flag that changes logging config to show more columns
- Add debug flag that shows debug log level
- Change stacktrace level to FATAL instead of ERROR
- Clean up output

More:
- Change selectors to show more than just UUIDs: also include file names, universe names, etc
- Add option to specify customerUUID and/or universeUUID at cli
- Add option to skip API call at CLI (will go to placeholder message about feature coming soon)

- Change select list structure:
  - Go out and collect all customer/universe data first
  - Create large list and pass it into the menu once
  - Make it searchable
*/

func main() {
	cmd.Execute()
}
