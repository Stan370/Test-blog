# Test-blog
A simple RESTful blog backend system.


## ⌨️ Local Development

You can use GitHub Codespaces for online development:

[![][codespaces-shield]][codespaces-link]

Or clone it for local development:

```fish
$ git clone https://github.com/Stan370/Test-blog.git
$ cd Test-blog
```

### how to start locally

Ask the owner of this repo to get a copy of local config file (config.json), place it in /config directory.
When you have problem get packets from dependency, please check your GOPROXY.

#### Start a MySQL container in Docker

- `docker pull mysql:5.7`

    - For Apple Silicon chips: `docker pull --platform linux/amd64 mysql:5.7`

- `docker run --name Blog -d -p 3306:3306 -e MYSQL_ROOT_PASSWORD=public mysql:5.7`
```
#### Start 'blog' service

- `./run.sh`         in Linux env
