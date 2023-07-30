# cron-expression-parser

This is a command line application that parses a cron string and 
expands each field to show the times at which it will run. 
The application considers the standard cron format with five time fields (minute, hour, day of month, month, and day of week) plus a command. 
It does not handle special time strings such as "@yearly". 
The input should be provided on a single line, passed as a single argument to the application.

## Usage

- To use the application, run the command:

``go run cmd/main.go``

- Provide the cron string as input. Eg : */15 2/3 1,2,4,6,15 * 1-5 /usr/bin/find

## Output

The output will be formatted as a table with the field name taking the first 14 columns and the times as a space-separated list following it.

Example output for the given input:

||                     |
| --- |---------------------|
| minutes | 0 15 30 45|
| hour | 2 5 8 11 14 17 20 23 |
| day of month  | 1 2 4 6 15|
| month | 1 2 3 4 5 6 7 8 9 10 11 12|
| day of week  | 1 2 3 4 5|
| command | /usr/bin/find|


Please note that the application does not handle special time strings such as "@yearly". It only considers the standard cron format.

