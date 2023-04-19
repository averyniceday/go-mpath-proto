#!/usr/bin/env bash
SCRIPT_DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"
SCRIPT_DIR_PARENT="$( cd "$SCRIPT_DIR/.." && pwd )"
GENOME_NEXUS_API_DOCS=${1:-https://genomenexus.org/v2/api-docs}
SWAGGER_CODEGEN_CLI_JAR="$SCRIPT_DIR/swagger-codegen-cli.jar"

# only download the client if it is not already downloaded
if [ ! -f ${SWAGGER_CODEGEN_CLI_JAR} ]; then
    if which curl ; then
        curl 'https://repo1.maven.org/maven2/io/swagger/swagger-codegen-cli/2.3.1/swagger-codegen-cli-2.3.1.jar' --output ${SWAGGER_CODEGEN_CLI_JAR}
    else
        if which wget ; then
            wget 'http://central.maven.org/maven2/io/swagger/swagger-codegen-cli/2.3.1/swagger-codegen-cli-2.3.1.jar' -O ${SWAGGER_CODEGEN_CLI_JAR}
        else
            echo "Error : could not find curl or wget for downloading swagger-codegen-cli" 1>&2
            exit 1
        fi
    fi
fi

# remove the modules if available (otherwise some legacy methods in the test will not get removed)
rm -rf "${SCRIPT_DIR_PARENT}/genome-nexus-public-api"
rm -rf "${SCRIPT_DIR_PARENT}/genome-nexus-internal-api"

# generate java modules (see config json files for more details)
#java -jar ${SWAGGER_CODEGEN_CLI_JAR} config-help -l go
java -jar ${SWAGGER_CODEGEN_CLI_JAR} generate -l go -c genomeNexusPublicApiClientConfig.json -i ${GENOME_NEXUS_API_DOCS} -o "$SCRIPT_DIR_PARENT/genome-nexus-public-api" --api-package genome_nexus_public_api
java -jar ${SWAGGER_CODEGEN_CLI_JAR} generate -l go -c genomeNexusInternalApiClientConfig.json -i ${GENOME_NEXUS_API_DOCS}?group=internal -o "$SCRIPT_DIR_PARENT/genome-nexus-internal-api" --api-package genome_nexus_internal_api
