#!/bin/bash
set -euo pipefail

# clean up
rm -fr dist
mkdir dist

# create build.json
node ./create-build-manifest.js > dist/build.json

# use jsii to produce packages in all supported languages under dist/
jsii-pacmak

# create docs
mkdir -p dist/docs
echo "<pre>$(cat package.json)</pre>" > dist/docs/index.html
echo "nojekyll!" > dist/docs/.nojekyll
