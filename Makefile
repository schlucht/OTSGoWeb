BINARY_NAME=liam

VERSION_FILE=version.txt
OLD_VERSION=$(shell cat ${VERSION_FILE})
VERSION=${OLD_VERSION}
VER=$(subst ., ,${VERSION})
MAJOR=$(word 1,${VER})
MINOR=$(word 2,${VER})
PATCH=$(word 3,${VER})
NEWPATCH=$(shell expr ${PATCH} + 1)
NEW_VERSION=${MAJOR}.${MINOR}.${NEWPATCH}

build:
	@go build -o dist/${BINARY_NAME} ./cmd


run: build
	clear
	@env ./dist/${BINARY_NAME} --version ${NEW_VERSION}
	@echo ${NEW_VERSION} > ${VERSION_FILE}	



clean:
	@go clean
	@rm -rf dist

git:	
	@git add . && git commit -m "bump version" && git tag -a ${NEW_VERSION} -m "${NEW_VERSION}"
