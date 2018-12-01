# lunchideas

Slackbot for lunch ideas using Foursquare API

---

<p align="center">
  <img src="./screenshot.png" alt="Screenshot" />
</p>

This was built on MLH's Local Hack Day as a fun project to play around with the Foursquare API.

## How to deploy

First, set secrets in Heroku:

```
LOCATION => Shoreditch
CLIENT_ID => Foursquare's Client ID
SECRET => Foursquare's Client Secret
```

then push the image:

```sh
$ heroku container:push web -a lunchideas
```

and run it:

```sh
$ heroku container:release web -a lunchideas
```

## License

MIT
