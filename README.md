# docean-deploy
POC: push docker image and deploy to DigitalOcean using Github Actions

![.github/workflows/ci.yml](https://github.com/efontan/docean-deploy/workflows/.github/workflows/ci.yml/badge.svg)

 Setup DigitalOcean droplet adding a new user for Github actions:

```
  sudo adduser new_username
  sudo usermod -aG sudo new_username
  sudo usermod -aG docker new_username
```