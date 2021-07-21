#!/bin/sh

echo "$(<LICENSE)"

# Configure your fisicus app below:
SITEMAP=https://b-nova.com/sitemaps/sitemap.xml
FILTER=/home/content/*

./bin/fisicus serve --sitemap=$SITEMAP --filter=$FILTER
