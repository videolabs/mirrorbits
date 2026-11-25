# syntax=docker/dockerfile:1
ARG GO_VERSION=1

FROM docker.io/library/golang:${GO_VERSION}-trixie AS build
WORKDIR /src
COPY go.mod go.sum ./
RUN go mod download
# .git is copied too: the Makefile takes the version from `git describe`
COPY . .
# PREFIX=/usr: the templates are looked up in /usr/share/mirrorbits
RUN CGO_ENABLED=0 make build PREFIX=/usr

# rsync is needed to scan the mirrors with an rsync:// URL. The GeoIP
# databases are not included: MaxMind requires an account to download them,
# and they must be updated regularly. Mount them in /usr/share/GeoIP (the
# default GeoipDatabasePath), e.g. from the ghcr.io/maxmind/geoipupdate image.
FROM docker.io/library/debian:trixie-slim
RUN apt-get update \
 && DEBIAN_FRONTEND=noninteractive apt-get install -y --no-install-recommends \
      ca-certificates rsync \
 && rm -rf /var/lib/apt/lists/* \
 && useradd --system --no-create-home --home-dir /nonexistent \
      --shell /usr/sbin/nologin mirrorbits \
 && mkdir -p /srv/repo /usr/share/GeoIP /var/log/mirrorbits \
 && chown mirrorbits:mirrorbits /var/log/mirrorbits
COPY --from=build /src/bin/mirrorbits /usr/bin/mirrorbits
COPY --from=build /src/templates/ /usr/share/mirrorbits/
COPY contrib/docker/mirrorbits.conf /etc/mirrorbits.conf
USER mirrorbits
EXPOSE 8080
ENTRYPOINT ["/usr/bin/mirrorbits"]
CMD ["daemon", "-config", "/etc/mirrorbits.conf"]
