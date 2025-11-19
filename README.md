<!--

/ // __ / | / /  /   |  /  / | / / / __ \ /  |/  /   |  / //// /
_ / / / /  |/ / / / / /| |  / //  |/ / __/ / // // /|/ / /| | / ,<  / /

/ / // / /|  / / / / ___ |/ // /|  / // , // /  / / ___ |/ /| |/ /

//_// |/ // //  |/// |/// ||//  ///  |// |/**_/

>> ORCHESTRATE CHAOS. CONTAINERIZE LOGIC. <<
-->

<div align="center">

 English  |  中文文档 

</div>

<a id="-english-version"></a>

⚡ Container-Make

The missing link between your Code and the Kernel.

Container-Make (cm) transforms your devcontainer.json into a powerful CLI artifact. It fuses the speed of Makefiles, the isolation of Docker, and the intelligence of VSCode DevContainers into a single, lethal binary.

No more "it works on my machine". No more 500-character docker run commands.

/// The Philosophy

Single Source of Truth: Your devcontainer.json defines the universe.

Native Fidelity: vim, htop, and interactive shells work exactly as they do locally.

Zero Pollution: Dependencies live in the container, not on your host OS.

/// Quick Start

1. Installation

Download the latest binary from Releases or build from source:

git clone [https://github.com/container-make/cm.git](https://github.com/container-make/cm.git)
cd cm && go build -o cm ./cmd/cm


2. Initialize

Generate shell integration (aliases and shims) to make cm feel like part of your shell.

./cm init
# Follow the instructions to add the alias to your .zshrc or .bashrc


3. Execution

Navigate to any project with a .devcontainer folder and just run.

Prepare the environment (Build):

cm prepare


Run a command:

# Syntax: cm run -- <your-command>

cm run -- go build -o app main.go
cm run -- npm install
cm run -- python train_model.py


Enter interactive mode:

cm run -- /bin/bash


/// Configuration

We support the standard devcontainer.json specification.

// .devcontainer/devcontainer.json
{
    "image": "mcp/firecrawl:latest",
    "forwardPorts": [8080],
    "containerEnv": {
        "APP_ENV": "development"
    },
    "postStartCommand": "echo 'Environment Ready.'"
}


<a id="-chinese-version"></a>

⚡ Container-Make (中文版)

代码与内核之间的缺失环节。

Container-Make (cm) 将您的 devcontainer.json 转化为一个强大的命令行工具。它集成了 Makefile 的极致速度、Docker 的绝对隔离以及 DevContainers 的现代开发体验。

告别 “在我的机器上是好的”。告别 500 个字符长的 docker run 命令。

/// 核心哲学

单一真理来源: 使用标准 devcontainer.json 定义开发宇宙。

原生级保真度: vim、htop 和交互式 Shell 的体验与宿主机完全一致。

零环境污染: 所有依赖均活在容器内，保持宿主机纯净。

/// 快速开始

1. 安装

从 Releases 下载最新二进制文件，或源码编译：

git clone [https://github.com/container-make/cm.git](https://github.com/container-make/cm.git)
cd cm && go build -o cm ./cmd/cm


2. 初始化

生成 Shell 集成脚本（别名和垫片），让 cm 与您的终端融为一体。

./cm init
# 按照屏幕提示将 alias 添加到您的 .zshrc 或 .bashrc 中


3. 执行

进入任何包含 .devcontainer 文件夹的项目即可开始使用。

准备环境 (构建镜像):

cm prepare


运行命令:

# 语法: cm run -- <您的命令>

cm run -- go build -o app main.go
cm run -- npm install
cm run -- python train_model.py


进入交互模式:

cm run -- /bin/bash


/// 配置指南

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
    "postStartCommand": "echo '环境已就绪。'"
}


<div align="center">
<p>
<sub>Designed & Engineered by Devin He</sub>





<sub>MIT License</sub>
</p>
</div>
