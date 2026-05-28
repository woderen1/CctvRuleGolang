# CctvRuleGolang

<p><span style="font-size: 3rem; font-weight: bold; color: red;">央视网视频下载辅助–LinkGet[原理:通过将加密伪链接替换成可用链接辅助下载]</span></p>

<b>缺点:视频质量目前只能低些</b>

---


  <h2>📖 项目描述</h2>
  <p>这是一个用 Go 语言编写的命令行工具，用于批量处理包含特定央视网m3u8 视频流链接的文本文件。它可以将加密链接（下载后会花屏），替换为可用链接。</p>

<p><b>规则灵感来源 [https://www.cnblogs.com/luomocn/articles/18016493] </b></p>

<p>类似于其他央视网视频链接获取项目多是针对 Windows 和 macOS 的预编译包，本项目提供了源代码，可在任何能编译 Go 程序的平台上使用。没有提供预编译包是因为我比较懒 (～￣▽￣)～（Actions也懒得办呢</p>
<h2>✨ 功能</h2>
   <ul>
   <li>读取一个文本文件（如包含m3u8链接的网页源码或播放列表）。</li>
   <li>使用正则表达式匹配特定格式的 CNTV m3u8 链接。</li>
   <li>将匹配到的链接域名从 <code>dh5wswx02.v.cntv.cn</code> 替换为 <code>hls.cntv.lxdns.com</code>。</li>
  <li>将路径从 <code>/asp/h5e/hls</code> 替换为 <code>/asp/hls</code>。</li>
   <li>移除链接中的所有查询参数（<code>?</code> 之后的部分）。</li>
   <li>将处理后的内容保存到一个新文件中，原文件保持不变。</li></ul>

<h2>🔨 编译教程</h2>
<p>在运行程序前，你需要先将其编译为可执行文件。</p>
        <ol><li><strong>安装 Go 环境</strong>
<p>访问 <a href="https://go.dev/dl/" target="_blank">Go 官方下载页面</a>，根据你的操作系统（Windows/macOS/Linux）下载并安装 Go。</p></li>
<li><strong>获取源代码</strong>
<p>将 <code>main.go</code> 文件保存到你的本地目录。</p></li>
<li><strong>打开终端/命令提示符</strong>
<p>进入存放 <code>main.go</code> 文件的目录。</p>
</li>
<li><strong>执行编译命令</strong>
<p>在终端中运行以下命令：</p>
<pre><code># 在 Windows 上编译，会生成 main.exe
go build -o main.exe main.go

# 在 macOS 或 Linux 上编译，会生成名为 main 的可执行文件
go build -o main main.go</code></pre>
<p>编译成功后，当前目录下会生成一个可执行文件（<code>main</code> 或 <code>main.exe</code>）。</p></li></ol>

<h2>🚀 使用方法</h2>
<p>输入的txt文本文件示例可见 [./exampleinput.txt] </p>
<p>编译完成后，在终端中使用以下命令格式运行程序：</p>
<pre><code>./main &lt;输入文件&gt; [输出文件]</code></pre>
<p><strong>参数说明：</strong></p>
<ul>
<li><code>&lt;输入文件&gt;</code>：<strong>必需。</strong> 你要处理的原始文本文件的路径（例如：<code>source.txt</code>）。</li>
<li><code>[输出文件]</code>：<strong>可选。</strong> 处理完成后新文件的保存路径。如果省略，程序会自动在输入文件名前加上 <code>modified_</code> 前缀（例如：输入 <code>list.txt</code>，则输出文件为 <code>modified_list.txt</code>）。</li></ul>
 <h3>使用示例：</h3>
        <pre><code># 示例1：处理 input.txt，输出到默认文件 modified_input.txt
./main input.txt

# 示例2：处理 mylist.txt，并指定输出到 processed_list.txt
./main mylist.txt processed_list.txt

# 在 Windows 上示例
main.exe urls.txt final_urls.txt</code></pre>
 

    
    
<p>程序运行后，会在终端显示处理摘要，包括总行数和成功修改的行数。</p>
<h2>⚠️ 注意事项</h2>
<ul>
<li>该工具仅处理匹配特定正则表达式模式的链接，其他内容会原样保留。</li>
<li>程序会创建一个新文件，<strong>不会修改或覆盖你的原始文件</strong>，请放心使用。</li>
<li>请确保你对输入文件有读取权限，对输出目录有写入权限。</li></ul><hr>
<p style="text-align: center; color: #7f8c8d; font-size: 0.9em;">Made with Go </p>
    </div>
</body>
</html>

