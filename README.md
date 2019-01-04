# Pandoc in a Docker Container (Work in Progress)

# Build the Image

```sh
$ docker build --tag pandoc .
```

# Run the Container

```sh
$ docker run --rm -p "8080:8080" pandoc
```

Visit <http://localhost:8080/pandoc/help> or <http://localhost:8080/pandoc/version>.
