# X-BOT

A Slack BOT who does anything.

## Commands

`@bot` is your BOT's name.

### `@bot c|contribute`

Post the URL to contribute.

### `@bot j|joke`

Post a joke with the [Joke API](https://github.com/15Dkatz/official_joke_api).

### `@bot l|list`

Post a list of commands.

### `@bot o|ojichat`

Post an ojichat with the [ojichat](https://github.com/greymd/ojichat).  
Valid arguments are the following.

- `@bot ojichat`
- `@bot ojichat @skmatz`

### `@bot p|pick`

Pick randomly.  
Valid arguments are the following.

- `@bot pick 3 a b c d e`

### `@bot ping`

Check to see if the BOT server is working properly.  
If it is working, the BOT posts the message `pong` to the same channel.

### `@bot t|timer`

Measure the time.  
Valid arguments are the following.

- `@bot timer 3`: Measure 3 seconds.
- `@bot timer 3 sec`: Measure 3 seconds.
- `@bot timer 3 min`: Measure 3 minutes.
- `@bot timer 3 sec Hello, world!`: Measure 3 seconds with the memo: "Hello, world!".

## Develop

Add a BOT with `app_mentions:read` and `chat:write` permissions to your Slack workspace and set .env.

```
docker-compose up --build
```

## Deploy

I recommend using [Cloud Run](https://cloud.google.com/run) provided by GCP.  
You just select the forked-repository, set two environment variables, and set an endpoint for the Slack event API.
