#!/bin/sh

cat <<EOL
{
  "project_path": "$GITHUB_REPOSITORY",
  "commit_id": "$GITHUB_SHA"
}
EOL
