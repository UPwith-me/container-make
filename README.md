<div align="center">

<!-- TITLE & LOGO -->
<h1>
    <br>
    ⚡ CONTAINER-MAKE ⚡
    <br>
</h1>

<h3>The Developer Experience Tool for the Container Era</h3>
<h3>容器时代的极致开发体验工具</h3>

<p>
    <a href="https://golang.org"><img src="https://img.shields.io/badge/Built_with-Go_1.21+-00ADD8?style=for-the-badge&logo=go" alt="Go"></a>
    <a href="LICENSE"><img src="https://img.shields.io/badge/License-MIT-ff5252?style=for-the-badge" alt="License"></a>
    <a href="#"><img src="https://img.shields.io/badge/Platform-Windows_|_Linux_|_macOS-181717?style=for-the-badge&logo=linux" alt="Platform"></a>
</p>

<br>

<!-- INTRO -->
<p align="center" style="max-width: 600px; margin: auto;">
    <b>Container-Make (cm)</b> transforms your <code>devcontainer.json</code> into a powerful CLI.<br>
    It brings the <b>speed</b> of Makefiles, the <b>isolation</b> of Docker, and the <b>convenience</b> of modern tooling into one binary.
    <br><br>
    <b>Container-Make (cm)</b> 将您的 <code>devcontainer.json</code> 转化为强大的 CLI 工具。<br>
    它集 Makefile 的<b>速度</b>、Docker 的<b>隔离性</b>以及现代工具的<b>便捷</b>于一身。
</p>

<br>

<!-- DEMO / HERO -->
<pre align="left" style="background-color: #1e1e1e; color: #d4d4d4; padding: 20px; border-radius: 10px; border: 1px solid #333; box-shadow: 0 10px 30px rgba(0,0,0,0.5);">
<span style="color: #569cd6;">$</span> <span style="color: #4ec9b0;">cm</span> init
<span style="color: #6a9955;"># Shell integration configured.</span>

<span style="color: #569cd6;">$</span> <span style="color: #4ec9b0;">cm</span> prepare
<span style="color: #ce9178;">[+]</span> Building image... <span style="color: #6a9955;">Done (0.8s)</span>

<span style="color: #569cd6;">$</span> <span style="color: #4ec9b0;">cm</span> run -- <span style="color: #dcdcaa;">make</span> build
<span style="color: #ce9178;">[+]</span> Creating container...
<span style="color: #ce9178;">[+]</span> Mapping UID/GID...
<span style="color: #4ec9b0;">Build complete. Artifacts are in ./bin</span>
</pre>

</div>

<br>
<br>

<!-- FEATURES GRID -->
## ✨ Why Container-Make? / 核心价值

<div align="center">
<table>
  <tr>
    <td width="50%" valign="top">
      <h3>� Instant Environments</h3>
      <p>No more "it works on my machine". Define your environment once in <code>devcontainer.json</code> and run anywhere. <code>cm</code> handles the rest.</p>
      <br>
      <h3>� Seamless Networking</h3>
      <p>Need to access a database or web server? <code>forwardPorts</code> maps them to localhost instantly. No manual <code>docker run -p</code> needed.</p>
    </td>
    <td width="50%" valign="top">
      <h3>💎 Native Fidelity</h3>
      <p>We spent days perfecting TTY handling. Vim, htop, and interactive shells work exactly as they do on your host machine.</p>
      <br>
      <h3>⚡ BuildKit Performance</h3>
      <p>Leverages Docker BuildKit for aggressive caching. Your builds have never been this fast and reproducible.</p>
    </td>
  </tr>
</table>
</div>

<br>

<!-- USAGE -->
## 🛠️ Workflow / 工作流

### 1. Define / 定义
Create a standard `.devcontainer/devcontainer.json`.
创建标准的 `.devcontainer/devcontainer.json`。

```jsonc
{
  "image": "mcp/firecrawl:latest",
  "forwardPorts": [8080],
  "postStartCommand": "echo 'Ready to code!'"
}
```

### 2. Prepare / 准备
Pre-warm the environment (Optional but recommended for CI).
预热环境（可选，推荐用于 CI）。

```bash
cm prepare
```

### 3. Execute / 执行
Run any command inside the container.
在容器内执行任意命令。

```bash
cm run -- npm install
cm run -- go build
cm run -- python main.py
```

<br>

<!-- INSTALLATION -->
## 📦 Installation / 安装

```bash
# 1. Clone
git clone https://github.com/container-make/cm.git

# 2. Build
cd cm && go build -o cm.exe ./cmd/cm

# 3. Enjoy
./cm.exe init
```

<br>

<!-- FOOTER -->
<div align="center">
    <br>
    <p>
        <sub>Designed for the future of development.</sub>
        <br>
        <sub>面向未来的开发工具。</sub>
    </p>
    <br>
    <a href="#"><img src="https://img.shields.io/github/stars/container-make/cm?style=social" alt="GitHub Stars"></a>
</div>
