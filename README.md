# fisicus

**fisicus is a cli tool that allows seo analysts to filter a target sitemap with a desired pattern.**

# About

With our **_b-nova_** collection of microservices, we're offering a complete set of modules which enables our JAMstack.
fisicus is an important module as it reads and forwards a customized sitemap from a given website.

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

## Contact

If you want to contact us you can reach us at [hello@b-nova.com](hello@b-nova.com).

## License

<!--- If you're not sure which open license to use see https://choosealicense.com/--->

No license has been applied yet.