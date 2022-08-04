# render-cli
CLI for https://render.com

# Use
The Render CLI expects that you have `RENDER_API_TOKEN` set in the environment.

```shell
$ export RENDER_API_TOKEN=<token>

$ render services list
srv-c9hpanirrk0d1bjgn48g   Web App

$ render services get srv-c9hpanirrk0d1bjgn47g
ID:         srv-c9hpanirrk0d1bjgn48g
Name:       Web App
Slug:       web-app
Created:    Sat, 23 Apr 2022 05:57:51 UTC
Updated:    Thu, 28 Jul 2022 00:53:56 UTC

AutoDeploy: yes
    Repo:   https://github.com/...
    Branch: main
```
