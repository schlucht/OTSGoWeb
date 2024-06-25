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
TEXT="ADD - Upload Project"

build:
	@go build -o dist/${BINARY_NAME} ./cmd

run: build
	clear
	@env ./dist/${BINARY_NAME} --version ${NEW_VERSION}	

clean:
	@go clean
	@rm -rf dist

stop:
	@pkill -f ${BINARY_NAME}
	@echo "Backend stopped..."

restart: stop run	

git:	
	@git add . && git commit -m ${TEXT} && 

tagging:
	@git tag -a ${NEW_VERSION} -m "${TEXT} - ${NEW_VERSION}"
	@git push && git push --tags
	@echo ${NEW_VERSION} > ${VERSION_FILE}

gitVersion: git tagging
	@echo ${NEW_VERSION}

# @if [ -z ${TEXT} ]; then echo "test"; else echo ${TEXT}; fi
test:
	ifeq(${TEXT},)		
		echo ${TEXT}
	endif
	@echo ${TEXT}