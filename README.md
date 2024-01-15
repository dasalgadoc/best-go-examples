<h1 align="center">
  🚀 🐹 Best - Go - Examples
</h1>

<p align="center">
    <a href="#"><img src="https://img.shields.io/badge/technology-go-blue.svg" alt="Go"/></a>
</p>
<p align="center">
  This repository explore basic concepts of Go.
</p>

It's an evolution of [Go examples](https://github.com/dasalgadoc/go-examples)

## 🧲 Environment Setup

### 🛠️ Needed tools

1. Go 1.18 or higher
2. (Optional) [fzf](https://github.com/junegunn/fzf)

### 🏃🏻 Application execution

1. Make sure to download all Needed tools, fzf its optional.
2. Clone the repository
```bash
git clone https://github.com/dasalgadoc/best-go-examples.git
```
3. Build up go project
```bash
go mod download
go get .
```
4. If you have fzf installed on your computer you can use the shell __runner.sh__

So, you can run nany example by selecting from de menu:

Make sure to give execution permissions to the file
```bash
chmod +x ./runner.sh
```

```bash
source ./runner.sh
```

If you don't have fzf installed, you can manually access to any folder an run the application.

```bash
cd ./01-console
go run main.go
```

5. Enjoy! 😎

## 📌  Example list

| Name                                                           | Description                   | Related Examples                                                                                                                                                                                  |
|----------------------------------------------------------------|-------------------------------|---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|
| [01-console](./01-console)                                     | Hello World                   |                                                                                                                                                                                                   |
| [02-variables](./02-variables)                                 | Variables and zero values     |                                                                                                                                                                                                   |
| [03-operators](./03-operators)                                 | Basic operations              |                                                                                                                                                                                                   |
| [04-functions](./04-functions)                                 | Functions                     | [33-currying](./33-currying)                                                                                                                                                                      |
| [05-recursion](./05-recursion)                                 | Recursion                     |                                                                                                                                                                                                   |
| [06-loops](./06-loops)                                         | Loops                         |                                                                                                                                                                                                   |
| [07-conditional](./07-conditional)                             | Conditionals                  |                                                                                                                                                                                                   |
| [08-casting](./08-casting)                                     | Casting                       |                                                                                                                                                                                                   |
| [09-array_and_slice](./09-array_and_slice)                     | Arrays and Slices             |                                                                                                                                                                                                   |
| [10-maps](./10-maps)                                           | Maps                          |                                                                                                                                                                                                   |
| [11-structs](./11-structs)                                     | Structs                       | [12-struct_with_behaviour](./12-struct_with_behaviour)                                                                                                                                            |
| [12-struct_with_behaviour](./12-struct_with_behaviour)         | Structs II                    | [11-structs](./11-structs)<br/>[15-channels](./15-channels)                                                                                                                                       |
| [13-interfaces](./13-interfaces)                               | Interfaces and composition    |                                                                                                                                                                                                   |
| [14-go_routines](./14-go_routines)                             | Go routine                    | [15-channels](./15-channels)<br/>[30-go_routines_2](./30-go_routines_2)<br/>[31-go_routines_communication](./31-go_routines_communication)<br/>[32-workers-poke-api](./32-workers-poke-api)       |
| [15-channels](./15-channels)                                   | Channels                      | [14-go_routines](./14-go_routines)<br/>[30-go_routines_2](./30-go_routines_2)<br/>[31-go_routines_communication](./31-go_routines_communication)<br/>[32-workers-poke-api](./32-workers-poke-api) |
| [16-abstract_factory](./16-abstract_factory)                   | Abstract Factory pattern      |                                                                                                                                                                                                   |
| [17-dates](./17-dates)                                         | Dates handling                |                                                                                                                                                                                                   |
| [18-json](./18-json)                                           | Json and JsonPath             | [22-yaml](./22-yaml)                                                                                                                                                                              |
| [19-reflections](./19-reflections)                             | Reflections                   |                                                                                                                                                                                                   |
| [20-exercise-tags](./20-exercise-tags)                         | An algorithm example          |                                                                                                                                                                                                   |
| [21-pointers_generators](./21-pointers_generators)             | Memory directions             |                                                                                                                                                                                                   |
| [22-yaml](./22-yaml)                                           | Yaml                          | [18-json](./18-json)                                                                                                                                                                              |
| [23-value_objects](./23-value_objects)                         | Value Objects                 |                                                                                                                                                                                                   |
| [24-enum](./24-enum)                                           | Enum                          | [25-enum-json](./25-enum-json)                                                                                                                                                                    |
| [25-enum-json](./25-enum-json)                                 | Enum with json marshaller     | [24-enum](./24-enum)                                                                                                                                                                              |
| [26-consume-rest-api](./26-consume-rest-api)                   | Consume rest API native Go    | [32-workers-poke-api](./32-workers-poke-api)                                                                                                                                                      |
| [27-errors](./27-errors)                                       | Errors handling Go            |                                                                                                                                                                                                   |
| [28-profiling](./28-profiling)                                 | Profiling                     | [29-benchmark](./29-benchmark)                                                                                                                                                                    |
| [29-benchmark](./29-benchmark)                                 | Benchmarking                  | [28-profiling](./28-profiling)                                                                                                                                                                    |
| [30-go_routines_2](./30-go_routines_2)                         | Go Routines 2                 | [14-go_routines](./14-go_routines)<br/>[15-channels](./15-channels)<br/>[31-go_routines_communication](./31-go_routines_communication)<br/>[32-workers-poke-api](./32-workers-poke-api)           |
| [31-go_routines_communication](./31-go_routines_communication) | Routines communication        | [14-go_routines](./14-go_routines)<br/>[15-channels](./15-channels)<br/>[30-go_routines_2](./30-go_routines_2)<br/>[32-workers-poke-api](./32-workers-poke-api)                                   |
| [32-workers-poke-api](./32-workers-poke-api)                   | Rest consume with go routines | [14-go_routines](./14-go_routines)<br/>[15-channels](./15-channels)<br/>[30-go_routines_2](./30-go_routines_2)<br/>[31-go_routines_communication](./31-go_routines_communication)                 |
| [33-currying](./33-currying)                                   | Currying                      | [04-functions](./04-functions)                                                                                                                                                                    |
| [34-graphql](./34-graphql)                                     | Graphql                       |                                                                                                                                                                                                   |
| [35-linear_search](./35-linear_search)                         | Linear search                 | [36-binary_search](./36-binary_search)                                                                                                                                                            |
| [36-binary_search](./36-binary_search)                         | Binary search                 | [35-linear_search](./35-linear_search)                                                                                                                                                            |
