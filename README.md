# One Hour One Life Name Picker

Small web app that suggests first and last names from the game's allowed names lists.

## Development

Composed of two parts: a static web page, and a [AWS lambda](https://aws.amazon.com/lambda/) backend written in [Go](https://golang.org/) that picks the names.

Backend program is compiled in a [Vagrant](https://www.vagrantup.com/) VM, as much for getting a zip file with unix permissions as anything. The Vagrant setup links the project directory into the VM's Go workspace.

### AWS

I've only set up Lambda once, but roughly:

Lambda function on Go 1.x, with an API Gateway attached.

API Gateway required CORS to be set up for the OPTIONS request. The program also returns the Access-Control-Allow-Origin header on each request.

## Credits

- Name lists from OHOL, prepared by Jason Rohrer with this notice:

> All US baby names for 2016 with more than 5 occurrences found here:
> 
> https://www.ssa.gov/oact/babynames/limits.html
> 
> 
> 
> All US last names from 2000 with more than 200 occurrences pulled from here:
> 
> https://www.census.gov/topics/population/genealogy/data/2000_surnames.html
