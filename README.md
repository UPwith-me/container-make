<div align="center">

⚡ Container-Make

The Developer Experience Tool for the Container Era





容器时代的极致开发体验工具

 English  |  中文文档 

</div>

<a id="-english"></a>

📖 English

Container-Make (cm) transforms your devcontainer.json into a powerful CLI artifact. It fuses the speed of Makefiles, the isolation of Docker, and the intelligence of VSCode DevContainers into a single, lethal binary.

✨ Why Container-Make?

🎯 Single Source of Truth
Your devcontainer.json defines the universe. No more maintaining separate Dockerfiles or Makefiles for local dev.

💎 Native Fidelity
vim, htop, and interactive shells work exactly as they do locally. We handle the complex TTY and signal forwarding for you.

🚀 BuildKit Powered
Leverages Docker BuildKit for aggressive caching. Your environment spins up in seconds, not minutes.

🛡️ Zero Pollution
Dependencies live in the container, not on your host OS. Keep your machine clean.

🛠️ Workflow

1. Install

Build from source or download the binary.

git clone [https://github.com/container-make/cm.git](https://github.com/container-make/cm.git)
cd cm && go build -o cm ./cmd/cm


2. Initialize

Generate shell aliases for a seamless experience.

./cm init
# Follow the on-screen instructions to update your .bashrc/.zshrc


3. Execute

Go to any project with a .devcontainer folder and run commands.

# Prepare the environment (Pre-build image)
cm prepare

# Run any command inside the container
cm run -- go build -o app main.go
cm run -- npm install
cm run -- python train.py

# Drop into an interactive shell
cm run -- /bin/bash


⚙️ Configuration

We support the standard devcontainer.json specification.

// .devcontainer/devcontainer.json
{
    "image": "mcp/firecrawl:latest",
    "forwardPorts": [8080],
    "containerEnv": {
        "APP_ENV": "development"
    },
    "postStartCommand": "echo 'Ready to code!'"
}


<a id="-chinese"></a>

🇨🇳 中文文档

Container-Make (cm) 将您的 devcontainer.json 转化为一个强大的命令行工具。它集成了 Makefile 的极致速度、Docker 的绝对隔离以及 DevContainers 的现代开发体验。

✨ 核心价值

🎯 单一真理来源
使用 devcontainer.json 定义整个开发宇宙。无需再为本地开发维护额外的 Dockerfile 或 Makefile。

💎 原生级保真
vim、htop 和交互式 Shell 的体验与宿主机完全一致。我们为您处理了复杂的 TTY 和信号转发。

🚀 BuildKit 驱动
利用 Docker BuildKit 的激进缓存策略。环境启动仅需秒级，而非分钟级。

🛡️ 零环境污染
所有依赖均活在容器内，保持宿主机纯净。告别 "it works on my machine"。

🛠️ 工作流

1. 安装

从源码编译或下载二进制文件。

git clone [https://github.com/container-make/cm.git](https://github.com/container-make/cm.git)
cd cm && go build -o cm ./cmd/cm


2. 初始化

生成 Shell 别名，获得无缝体验。

./cm init
# 按照屏幕提示更新您的 .bashrc 或 .zshrc


3. 执行

进入任何包含 .devcontainer 文件夹的项目即可执行。

# 准备环境 (预构建镜像)
cm prepare

# 在容器内运行任意命令
cm run -- go build -o app main.go
cm run -- npm install
cm run -- python train.py

# 进入交互式终端
cm run -- /bin/bash


⚙️ 配置指南

我们支持标准的 devcontainer.json 规范。

// .devcontainer/devcontainer.json
{
    // 基础镜像
    "image": "mcp/firecrawl:latest",

    // 端口自动转发 (映射到 localhost)
    "forwardPorts": [8080],

    // 注入环境变量
    "containerEnv": {
        "APP_ENV": "development"
    },

    // 生命周期钩子
    "postStartCommand": "echo '环境已就绪！'"
}


<div align="center">
<sub>Designed for the future of development.</sub>





<sub>MIT License &copy; 2025 Devin HE</sub>
</div>
