FROM golang:1.20 as builder

COPY . /ooze
WORKDIR /ooze/mayhem

RUN go install github.com/dvyukov/go-fuzz/go-fuzz@latest github.com/dvyukov/go-fuzz/go-fuzz-build@latest
RUN go get github.com/dvyukov/go-fuzz/go-fuzz-dep
RUN go get github.com/AdaLogics/go-fuzz-headers
RUN apt update && apt install -y clang

RUN cd fuzz_ooze_gotextdiff && go-fuzz-build -libfuzzer -o fuzz_ooze_gotextdiff.a && \
    clang -fsanitize=fuzzer fuzz_ooze_gotextdiff.a -o fuzz_ooze_gotextdiff.libfuzzer

RUN cd fuzz_ooze_prettydiff && go-fuzz-build -libfuzzer -o fuzz_ooze_prettydiff.a && \
    clang -fsanitize=fuzzer fuzz_ooze_prettydiff.a -o fuzz_ooze_prettydiff.libfuzzer

RUN cd fuzz_ooze_options && go-fuzz-build -libfuzzer -o fuzz_ooze_options.a && \
    clang -fsanitize=fuzzer fuzz_ooze_options.a -o fuzz_ooze_options.libfuzzer

FROM debian:bookworm-slim
COPY --from=builder /ooze/mayhem/fuzz_ooze_gotextdiff/fuzz_ooze_gotextdiff.libfuzzer /
COPY --from=builder /ooze/mayhem/fuzz_ooze_prettydiff/fuzz_ooze_prettydiff.libfuzzer /
COPY --from=builder /ooze/mayhem/fuzz_ooze_options/fuzz_ooze_options.libfuzzer /