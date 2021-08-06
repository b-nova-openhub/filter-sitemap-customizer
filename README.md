![fisicus-header](https://docs.google.com/drawings/d/e/2PACX-1vS8TbBbZkyi33-PrvcPWexQqI60LiuWQ6l1vmro4n6umAC0ErwEOoyV01yRVKvDIrd-_bzPIe8PYdJC/pub?w=962&h=330)

# gofisicus  (go-filter-sitemap-customizer)

<!--- These are examples. See https://shields.io for others or to customize this set of shields. You might want to include dependencies, project status and licence info here --->
![GoDoc](https://godoc.org/github.com/go-git/go-git/v5?status.svg)
![GitHub repo size](https://img.shields.io/github/repo-size/b-nova-openhub/solr-page-exposer)
![GitHub contributors](https://img.shields.io/github/contributors/b-nova-openhub/solr-page-exposer)
![GitHub stars](https://img.shields.io/github/stars/b-nova-openhub/solr-page-exposer?style=social)
![GitHub forks](https://img.shields.io/github/forks/b-nova-openhub/solr-page-exposer?style=social)
![Twitter Follow](https://img.shields.io/twitter/follow/b__nova?style=social)

**fisicus is a cli tool that allows seo analysts to filter a target sitemap with a desired pattern.**

# About

With our **_b-nova_** collection of microservices, we're offering a complete set of modules which enables our JAMstack.
fisicus is an important module as it reads and forwards a customized sitemap from a given website.

## Prerequisites

Before you begin, ensure you have met the following requirements:
<!--- These are just example requirements. Add, duplicate or remove as required --->

* You have installed the latest version of `go1.16.5`
* You have a `Linux/Mac OS` machine with working knowledge of the underlying filesystem and Go build process.

# Installation

1. Download and extract fisicus to your web directory or get the latest development version with:

```
$ ❯  git clone https://github.com/b-nova-openhub/filter-sitemap-customizer
```

2. Check the run configuration in `fisicus.sh`. Adjust `SITEMAP` and `FILTER` if necessary.

3. Run `make init` inside the directory. This is going to resolve the required Go dependencies, build and run the
   application.
4. The fisicus web service should now be served under `:8080/sitemap`.
5. While the web service is running, you can overwrite the specified configs with the request params `sitemap`
   and `filter` within that request at any moment. **
   Example:** `:8080/sitemap?filter=/home/about/*&sitemap=https://yourwebsite.com/sitemaps/sitemap.xml`

# Build

You can build fisicus with the provided Makefile. In order to do so, execute following command in the projects root
directory.

```
$ ❯  make all
```

:beginner:: If you're building for the first time and jumped ahead of this README, replace `make all` with `make init`.

# Usage

```
fisicus – a service that applies pattern matching to a target sitemap

Usage:
fisicus [command]

Available Commands:
serve Serve serve service and its rest endpoints
completion generate the autocompletion script for the specified shell 
help Help about any command 

Flags:
-h, --help help for jamctl

Use "fisicus [command] --help" for more information about a command.
```

## Dockerfile

There is also a Dockerfile by which one can containerize the stapagen application. The port that is being exposed is the
default 8080.

## Deployment to K8s

{coming soon}

### Quick Deployment to DigitalOcean's Kubernetes

[![Deploy to DO](https://www.deploytodo.com/do-btn-blue.svg)](https://cloud.digitalocean.com/apps/new?repo=https://github.com/b-nova-openhub/solr-page-exposer/tree/main)

## Contributing to gofisicus

<!--- If your README is long or you have some specific process or steps you want contributors to follow, consider creating a separate CONTRIBUTING.md file--->
To contribute to gostapafor, follow these steps:

1. Fork this repository.
2. Create a branch: `git checkout -b <branch_name>`.
3. Make your changes and commit them: `git commit -m '<commit_message>'`
4. Push to the original branch: `git push origin <project_name>/<location>`
5. Create the pull request.

Alternatively see the GitHub documentation
on [creating a pull request](https://help.github.com/en/github/collaborating-with-issues-and-pull-requests/creating-a-pull-request)
.

## Contributors

Thanks to the following people who have contributed to this project:

* [@raffaelschneider](https://github.com/raffaelschneider) 💻📖🐛
* [@stefanwelsch](https://github.com/bnova-stefan) 💻🧑‍🏫
* [@tomtrapp](https://github.com/tomtrapp) 🤔👀
* [@rickyelfner](https://github.com/ricky-bnova) 💬🐛

You might want to consider using something like
the [All Contributors](https://github.com/all-contributors/all-contributors) specification and
its [emoji key](https://allcontributors.org/docs/en/emoji-key).

## Contact

If you want to contact us you can reach us at [info@b-nova.com](hello@b-nova.com).

## License

<!--- If you're not sure which open license to use see https://choosealicense.com/--->

No license has been applied yet.