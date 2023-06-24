GO=go
GOOPTS=

RM=rm -f

PROJECT_NAME=autobahn-info

INSTALL_PATH=/usr/local/bin

version:
	${GO} version

clean:
	${RM} ${PROJECT_NAME} 

all: clean version
	${GO} ${GOOPTS} build

install: all
	@[[ -d ${INSTALL_PATH} ]] && cp -v ${PROJECT_NAME}* ${INSTALL_PATH}/. || echo Directory ${INSTALL_PATH} does not exist, cancelling.
	
