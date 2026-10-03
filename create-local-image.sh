set -xeuo pipefail

docker build -t rmecha/enpitsu-npsn-server:latest .
docker save -o enpitsu-npsn-server.tar rmecha/enpitsu-npsn-server:latest
