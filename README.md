<div align="center">

# ⚡ C O N T A I N E R - M A K E ⚡

### The Missing Link Between Makefiles and Containers
### 连接 Makefile 与容器的缺失环节

[![Go Version](https://img.shields.io/badge/Go-1.21+-00ADD8?style=for-the-badge&logo=go)](https://golang.org)
[![License](https://img.shields.io/badge/License-MIT-ff5252?style=for-the-badge)](LICENSE)
[![PRs Welcome](https://img.shields.io/badge/PRs-welcome-brightgreen.svg?style=for-the-badge)](http://makeapullrequest.com)

<br/>

<p align="center">
  <b>Container-Make (cm)</b> is not just a tool; it's a <b>philosophy</b>. <br/>
  It bridges the gap between the raw power of local build tools and the pristine isolation of containers.
  <br/><br/>
  <b>Container-Make (cm)</b> 不仅仅是一个工具，它是一种<b>哲学</b>。<br/>
  它跨越了本地构建工具的原始力量与容器完美隔离之间的鸿沟。
</p>

<br/>

---

## 🔮 The Experience / 极致体验

</div>

<table align="center">
  <tr>
    <td align="center" width="33%">
      <h3>🚀<br/>Zero Config Start</h3>
      <p>Drop a <code>devcontainer.json</code> and go. No complex setups. Just pure productivity.</p>
      <p><i>零配置启动。只需一个配置文件，即刻开始高效工作。</i></p>
    </td>
    <td align="center" width="33%">
      <h3>💎<br/>Interactive Fidelity</h3>
      <p>Raw TTY mode, signal propagation, and resize handling. It feels exactly like your local shell.</p>
      <p><i>原生 TTY 模式，信号传递，自动调整大小。手感如本地 Shell 般丝滑。</i></p>
    </td>
    <td align="center" width="33%">
      <h3>⚡<br/>BuildKit Powered</h3>
      <p>Integrated with Docker BuildKit for blazing fast, cached image builds.</p>
      <p><i>集成 Docker BuildKit，带来闪电般的缓存构建速度。</i></p>
    </td>
  </tr>
  <tr>
    <td align="center" width="33%">
      <h3>🛡️<br/>Security First</h3>
      <p>Automatic UID/GID mapping. Say goodbye to <code>root</code> owned files in your workspace.</p>
      <p><i>自动 UID/GID 映射。彻底告别工作区中的 root 权限文件噩梦。</i></p>
    </td>
    <td align="center" width="33%">
      <h3>🔗<br/>Seamless Networking</h3>
      <p>Port forwarding support. Access your container's services from localhost instantly.</p>
      <p><i>端口转发支持。从 localhost 瞬间访问容器内的服务。</i></p>
    </td>
    <td align="center" width="33%">
      <h3>🧩<br/>Ecosystem Ready</h3>
      <p>Supports Lifecycle Hooks (`postCreate`, `postStart`) and DevContainer standards.</p>
      <p><i>支持生命周期钩子和 DevContainer 标准，融入庞大生态。</i></p>
    </td>
  </tr>
</table>

<div align="center">

---

## 🛠️ Installation / 安装指南

<br/>

```bash
# Clone the repository / 克隆仓库
git clone https://github.com/container-make/cm.git

# Build the binary / 构建二进制文件
cd cm && go build -o cm.exe ./cmd/cm

# Initialize shell integration / 初始化 Shell 集成
./cm.exe init
```

<br/>

---

## 💻 Usage / 使用方式

### 1. Prepare / 准备
Pre-warm your environment and build images.
预热环境，构建镜像。

```bash
./cm.exe prepare
```

### 2. Run / 运行
Execute commands in the container with native performance.
以原生性能在容器中执行命令。

```bash
# Run a single command / 运行单个命令
./cm.exe run -- make build

# Drop into a shell / 进入 Shell
./cm.exe run -- /bin/bash

# Expose ports / 暴露端口
# (Configured in devcontainer.json: "forwardPorts": [8080])
./cm.exe run -- python3 -m http.server 8080
```

<br/>

---

## ⚙️ Configuration / 配置艺术

`devcontainer.json`

```jsonc
{
  "image": "mcp/firecrawl:latest",
  // "build": { "dockerfile": "Dockerfile" },
  
  "forwardPorts": [8080, 3000],
  
  "postStartCommand": "echo '🚀 Environment Ready!'",
  
  "containerEnv": {
    "APP_ENV": "development"
  }
}
```

<br/>

---

## 🗺️ Roadmap / 未来蓝图

| Phase | Status | Description |
| :--- | :---: | :--- |
| **I. Genesis** | ✅ | Core Bootstrapping & Config Parsing |
| **II. Fidelity** | ✅ | TTY, Signals, Entrypoint Injection |
| **III. Velocity** | ✅ | BuildKit Integration & Caching |
| **IV. Ecosystem** | ✅ | Lifecycle Hooks & IDE Integration |
| **V. Connectivity** | ✅ | Advanced Networking (Port Forwarding) |

<br/>

<p align="center">
  <sub>Built with ❤️ by the Container-Make Team. Designed for Builders.</sub>
</p>

</div>
