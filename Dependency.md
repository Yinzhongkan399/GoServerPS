# Dependency — 在 Ubuntu 24.04 上安装 & 运行说明 ✅

## 简介
本文件说明如何在 **Ubuntu 24.04 (x86_64)** 系统上准备环境、安装依赖并运行本项目（Go 与 BPF/BTF 相关工具）。内容尽量覆盖常见环境与可能遇到的问题。

---

## 目标与前提
- 目标系统：Ubuntu 24.04 LTS（x86_64）
- 建议内核：Linux kernel **5.10+**（越新越好，以确保 eBPF/BTF 功能兼容）
- 需要具备 sudo 权限

---

## 一、更新系统与基础工具 🔧
```bash
sudo apt update && sudo apt upgrade -y
sudo apt install -y build-essential git curl wget ca-certificates sudo
```

---

## 二、安装 Go（推荐使用官方二进制）🟦
建议安装 Go 1.20 或更高版本（请替换为最新稳定版本）：

```bash
# 从 https://go.dev/dl/ 下载并替换为最新版本
wget https://go.dev/dl/go1.20.7.linux-amd64.tar.gz
sudo rm -rf /usr/local/go && sudo tar -C /usr/local -xzf go1.20.7.linux-amd64.tar.gz
# 添加到 PATH（写入 ~/.profile 或 ~/.bashrc）
echo 'export PATH=$PATH:/usr/local/go/bin' >> ~/.profile
source ~/.profile
# 验证
go version
```

提示：也可以使用 apt 安装（packaged），但 apt 中的版本可能较旧。

---


## 四、eBPF / BTF / bpftool 相关依赖 🐧⚙️
本项目部分功能涉及 BTF/BPF 相关操作（读取 BTF 信息、使用 btftool / bpftool）。建议安装如下：

```bash
sudo apt install -y clang llvm libelf-dev libmnl-dev pkg-config libbpf-dev iproute2 iputils-ping
# 尝试安装 bpftool（若 apt 中无最新版，可从内核工具或源码编译）
sudo apt install -y bpftool
```

如果你的系统 apt 仓库没有 `libbpf-dev` 或 `bpftool` 的合适版本，建议从内核源码或 libbpf 仓库编译安装，或者使用官方的容器镜像来运行依赖较重/需要特权的工具（见下方 Docker 说明）。

---

## 五、可选：使用 Docker 运行 btftool / 依赖隔离 🐳
若不想在宿主机上安装多量工具，推荐使用特权容器来运行需要直接访问内核或模块的工具：

```bash
# 安装 docker（如尚未安装）
sudo apt install -y docker.io
# 以交互方式运行 Ubuntu 容器并挂载系统目录以便访问 BTF / bpftool
sudo docker run --rm -it --privileged -v /lib/modules:/lib/modules -v /sys:/sys ubuntu:24.04 /bin/bash
# 在容器内安装 bpftool/相关工具并运行 btftool 或其它脚本
```

---

## 六、克隆仓库并编译运行 🚀
假设仓库位于 GitHub（替换为实际仓库地址）：

```bash
git clone https://github.com/Yinzhongkan399/GoServerPS.git
cd GoServerPS
# 若使用 Go modules（推荐）
# go mod tidy
# 构建或运行
go build ./...
# 或直接运行（开发阶段）
go run ./main.go
```

启动后，默认监听的端口或者可访问的 URL 请参见 `main.go` 或与项目相关的 README 文档，使用 curl 或浏览器测试：

```bash
curl -v http://localhost:8080/your-endpoint
curl -v -X POST -H "Content-Type: application/json" -d '{"k": "v"}' http://localhost:8080/your-endpoint
```

---

## 七、常用调试与权限提示 ⚠️
- eBPF/BTF 操作常需更高权限，运行 bpftool、btftool 或加载内核对象可能需要 root 或 CAP_BPF/CAP_SYS_ADMIN。可用 `sudo` 或以 root 身份运行相应步骤。
- 若遇到 kernel headers / module 相关错误，确保安装了当前内核对应的 `linux-headers-$(uname -r)`。

```bash
sudo apt install -y linux-headers-$(uname -r)
```

- 如果 `bpftool` 报错找不到对象或 BTF，检查 `/sys/kernel/btf/vmlinux` 或使用 btftool 来获取 BTF 信息。

---

## 八、示例：创建 systemd 服务（可选） ⚙️
在生产环境中，你可能希望把服务作为 systemd 单元管理：

```ini
# /etc/systemd/system/goserverps.service
[Unit]
Description=GoServerPS Service
After=network.target

[Service]
Type=simple
User=youruser
WorkingDirectory=/home/youruser/GoServerPS
ExecStart=/home/youruser/GoServerPS/GoServerPS
Restart=on-failure

[Install]
WantedBy=multi-user.target
```

然后：
```bash
sudo systemctl daemon-reload
sudo systemctl enable --now goserverps
sudo journalctl -u goserverps -f
```

---

## 九、故障排查 & 参考资料 💡
- 查看日志：`journalctl`, `dmesg`（特别是 BPF 相关错误会有内核日志）
- 内核 BTF/ebpf 相关资料：Linux 内核文档、libbpf、bpftool 项目主页
- 如果需要，我可以为你：
  - 提供 `Dockerfile` 或 `docker-compose` 来简化部署

---

## 版权与备注
本说明旨在提供通用的安装与运行步骤，具体依赖请以仓库代码、脚本注释或 `README.md` 为准。若你希望我把这些步骤合并到仓库 `README.md` 或创建 `requirements.txt` / `Dockerfile`，告诉我我会继续实现。✅
