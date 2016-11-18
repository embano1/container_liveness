FROM scratch
MAINTAINER Michael Gasch <michael_gasch@live.com>
ADD container_liveness /container_liveness
ENTRYPOINT ["/container_liveness"]
