#!/bin/bash -ex

EXISTING_ISSUE=$( gh issue list --search "Release v${CANDIDATE_BETA}/v${CANDIDATE_STABLE}" --json url --jq '.[].url' --repo "${REPO}" )

if [ "${EXISTING_ISSUE}" != "" ]; then
    echo "Issue already exists: ${EXISTING_ISSUE}"
    exit 0
fi

gh issue create -a "${GITHUB_ACTOR}" --repo "${REPO}" --label release --title "Release v${CANDIDATE_BETA}/v${CANDIDATE_STABLE}" --body "Like #4522, but for v${CANDIDATE_BETA}/v${CANDIDATE_STABLE}"
**Performed by ocg release manager**

- [ ] Prepare stable  release v"${CANDIDATE_STABLE}"
- [ ] Prepare beta  release v"${CANDIDATE_BETA}"
- [ ] Tag and release stable  v"${CANDIDATE_STABLE}"
- [ ] Tag and release beta  v"${CANDIDATE_BETA}"
- [ ] Prepare ocg-releases v"${CANDIDATE_BETA}"
