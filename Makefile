.PHONY: setup release changelog tag

setup: ##Setup pre-commit hook
	pre-commit install

release: ##Create release
	./semtag ${version} -s ${scope}

changelog: ##Generate changelog
	git-chglog -o CHANGELOG.md

tag:
	./semtag get
