# Build & Run (Linux)

Prerequisites:

- Go (recommended >= 1.20)
- Linux (the project uses kernel BTF and bpftool)
- `bpftool` available in PATH and `/sys/kernel/btf/vmlinux` present

Quick build & run:

```bash
# build binary into ./bin/goserverps
make build

# run the built binary
make run

# development run (no build artifact)
make run-dev
```

Other useful targets:

- `make clean` — remove `bin/` and `./.cache`
- `make fmt` — format all Go files with `gofmt`
- `make vet` — run `go vet ./...`

Notes and troubleshooting:

- The program may call `bpftool` (via `BaseRun`) and read `/sys/kernel/btf/vmlinux`. Those steps require root privileges or proper capabilities in many environments.
- If `bpftool` is not present, run `sudo apt install bpftool` (or install the package provided by your distribution). On some distros `bpftool` is packaged alongside `linux-tools`.
- To inspect logs, run the binary directly (`./bin/goserverps`) or use `journalctl` if you install it as a service.
