# Fountains!

This is an example of web application for [Web and Software Application](http://gamificationlab.uniroma1.it/en/wasa/)
course.

See the [Fantastic Coffee (decaffeinated)](https://github.com/sapienzaapps/fantastic-coffee-decaffeinated) template for
instructions and project structure.

Note: this example uses SQLite, however you may implement a "naive" database using slices and maps.

## How to build container images

### Backend

```sh
$ docker build -t wasa-photos-backend:latest -f Dockerfile.backend .
```

### Frontend

```sh
$ docker build -t wasa-photos-frontend:latest -f Dockerfile.frontend .
```

## How to run container images

### Backend

```sh
$ docker run -it --rm -p 3000:3000 wasa-photos-backend:latest
```

### Frontend

```
$ docker run -it --rm -p 8081:80 wasa-photos-frontend:latest
```

## License

See [LICENSE](LICENSE).
