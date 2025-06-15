# 保存原始环境变量
$originalGOOS = $env:GOOS
$originalGOARCH = $env:GOARCH


# 支持的构建目标列表
$targets = @(
    @{ id=1; os="windows"; arch="amd64"; dir="win_x86_64"; ext=".exe" },
    @{ id=2; os="linux";   arch="amd64"; dir="linux_x86_64"; ext="" },
    @{ id=3; os="darwin";  arch="amd64"; dir="macos_x86_64"; ext="" },
    @{ id=4; os="darwin";  arch="arm64"; dir="macos_arm64"; ext="" }
)

# 显示菜单
Write-Host "请选择你要构建的平台："
foreach ($target in $targets) {
    Write-Host "$($target.id). $($target.os) $($target.arch)"
}
Write-Host "5. 全部构建"

# 获取用户输入
[int]$choice = Read-Host "请输入数字 (1-5)"

# 验证选择
if ($choice -lt 1 -or $choice -gt 5) {
    Write-Host "❌ 无效的选择！"
    pause
    exit 1
}

# 创建基础输出目录
$outputBase = "release-manual"
if (-Not (Test-Path -Path $outputBase)) {
    New-Item -ItemType Directory -Path $outputBase | Out-Null
}

# 构建函数
function Build-Target($target) {
    $env:GOOS = $target.os
    $env:GOARCH = $target.arch

    $outputDir = Join-Path $outputBase $target.dir
    $exeExt = $target.ext

    Write-Output "Building for $($target.os)/$($target.arch) -> $outputDir"

    # 创建目录
    if (-Not (Test-Path -Path $outputDir)) {
        New-Item -ItemType Directory -Path $outputDir | Out-Null
    }

    # 编译
    go build -o "$outputDir\dayz-server$exeExt" main.go

    if ($LASTEXITCODE -ne 0) {
        Write-Error "Build failed for $($target.os)/$($target.arch)"
        return $false
    }

    # 复制配置文件
    if (Test-Path -Path "config.yaml") {
        Copy-Item -Path "config.yaml" -Destination "$outputDir\config.yaml" -Force
    } else {
        Write-Warning "config.yaml not found. Skipping copy."
    }

    return $true
}

# 执行构建
if ($choice -eq 5) {
    foreach ($target in $targets) {
        if (-Not (Build-Target $target)) {
            Write-Error "构建失败，停止执行。"
            pause
            exit 1
        }
    }
} else {
    $selected = $targets | Where-Object { $_.id -eq $choice }
    if (-Not (Build-Target $selected)) {
        Write-Error "构建失败。"
        pause
        exit 1
    }
}

# 还原环境变量为 初始 环境
$env:GOOS = $originalGOOS
$env:GOARCH = $originalGOARCH

Write-Host "✅ 构建完成！"
pause