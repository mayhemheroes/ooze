FROM golang:1.20 as builder

COPY . /ooze
WORKDIR /ooze/mayhem

RUN go install github.com/dvyukov/go-fuzz/go-fuzz@latest github.com/dvyukov/go-fuzz/go-fuzz-build@latest
RUN go get github.com/dvyukov/go-fuzz/go-fuzz-dep
RUN go get github.com/AdaLogics/go-fuzz-headers
RUN apt update && apt install -y clang

RUN go-fuzz-build -libfuzzer -o fuzz_ooze.a && \
    clang -fsanitize=fuzzer fuzz_ooze.a -o fuzz_ooze.libfuzzer

FROM debian:bookworm-slim
COPY --from=builder /ooze/mayhem/fuzz_ooze.libfuzzer /