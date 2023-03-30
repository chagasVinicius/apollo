#!/usr/bin/env sh
old="go-starter"
new="$1"
dir="."

find $dir -type f -not -path '*/\.git/*' -not -name 'init.sh' -print0 | xargs -0 sed -i '' -e "s/$old/$new/g"
