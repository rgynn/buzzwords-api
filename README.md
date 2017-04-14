# buzzwords-api : http server to generate plain text buzzwords

## Overview

![screenshot](https://raw.githubusercontent.com/hajhatten/buzzwords-api/master/readme/screnshot.png)

HTTP server to generate plain/text buzzwords from the [github.com/hajhatten/buzzwords](https://github.com/hajhatten/buzzwords) package.

```
/ # UI with buttons to generate buzzwords
|-- buzzword # generates a plain text response with a buzzword sentence
|-- suffix # generates a plain text response with a buzzword sentence including suffix
|-- verb # generates a plain text response with a buzzword sentence including verb
|-- verbsuffix # generates a plain text response with a buzzword sentence including verb and suffix
```

## Install

```
go get github.com/hajhatten/buzzwords-api
```

## Example

```
docker run -e PORT=80 -p 80:80 hajhatten/buzzwords-api:latest
```

## Author

Rikard Gynnerstedt

## License

MIT.
