# ------
# BUILD
# ------
FROM golang:1.18-alpine AS megacron-go-build

LABEL maintainer="github.com/taggart-comet"
LABEL name="megacron-go-build"
LABEL description="build stage for golang project"

WORKDIR /app/megacron

COPY ./../. .

RUN CGO_ENABLED=0 go build -o /runner -a -ldflags '-extldflags "-static"' ./runner/
RUN CGO_ENABLED=0 go build -o /initial -a -ldflags '-extldflags "-static"' ./initial/

# ------
# FOR PROD RUNNING
# ------
FROM scratch as megacron-go-run

LABEL maintainer="github.com/taggart-comet"
LABEL name="megacron-go-run"
LABEL description="production image for megacron with binary executables"

COPY --from=megacron-go-build /runner /runner
COPY --from=megacron-go-build /initial /initial

CMD ["/sidecar"]
