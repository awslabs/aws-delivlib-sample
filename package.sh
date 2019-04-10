#!/bin/bash
set -euo pipefail

# clean up
rm -fr dist
mkdir dist

# build.json and changelog
node ./create-build-manifest.js > dist/build.json
cp CHANGELOG.md dist/

# use jsii to produce packages in all supported languages under dist/
jsii-pacmak || echo "Not everything worked but that's okay"

# create docs
mkdir -p dist/docs
echo "<pre>$(cat package.json)</pre>" > dist/docs/index.html
echo "nojekyll!" > dist/docs/.nojekyll
