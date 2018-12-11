const pkg = require('./package.json');

console.log(JSON.stringify({
  name: pkg.name,
  version: pkg.version,
  commit: process.env.CODEBUILD_RESOLVED_SOURCE_VERSION
}, undefined, 2));


