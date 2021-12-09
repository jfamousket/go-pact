# Example TODO App

Example TODO app built using `wails`, `go`, `react` and `go-pact`.

# Usage

This app uses a local pact server and wails to build a desktop TODO app. The todos are stored and managed by the pact local server, therefore you can change the pact server url in the `main.go` file to use any pact server of your choice.

### Start Pact Server

From the `example` directory run:

```
$ make start_pact
```

### Seed Pact Server

From the `example` directory run:

```
$ make pact_seed
```

### Start App

- If you want to run the development version of the app on your browser do the following:

  - From the `example` directory, run:

  ```
  $ wails serve
  ```

  - From the `example/frontend` directory run:

  ```
  $ yarn start
  ```

- If you want to run the built version as a desktop application do the following:
  - From the `example` directory run:
  ```
  $ wails build -p
  ```
  - Depending on your desktop environment, an app would be built to a `example/build` directory.

# Contributors

Built with :heart: by [jfamousket](https://twitter.com/jfamousket)
