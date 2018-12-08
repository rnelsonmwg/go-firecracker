FROM go-firecracker/client:latest

ADD scripts/lint/lint.sh /lint.sh
RUN chmod o+x /lint.sh

ENTRYPOINT /lint.sh