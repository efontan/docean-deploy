# docean-deploy
Testing how to push docker image and deploy to DigitalOcean using Github Actions

![CI](https://github.com/efontan/docean-deploy/workflows/CI/badge.svg)

 Setup DigitalOcean droplet adding a new user for Github actions:

```
  sudo adduser new_username
  sudo usermod -aG sudo new_username
  sudo usermod -aG docker new_username
```
