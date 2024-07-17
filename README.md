# golang-continous-integration

## Setup Github Actions

- For using github actions we have to make a .github folder. In which we have to make a specific folder for github actions that is workflows. Inside workflows we difine a yaml file to describe the github about what actions we wish to perform.

- Github actions also runs on a docker container, so we have to specify on what type of machine should our actions work. For eg. For particular job we can write `runs-on: ubuntu-latest`

- We have to specify steps of executions for the github actions. It will be executed in order from top to bottom. 

- Every step has a name and uses field. uses field is used to specify which builtin github action to use for our particular purpose.

- run field can be used to execute particular command. For eg. `run: go test ./...`