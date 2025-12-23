# edgetts

一个使用 **Go 语言** 调用 **Microsoft Edge 神经网络语音（Edge TTS）** 的文本转语音（TTS）库。

`edgetts` 是对 Microsoft Edge 内部 TTS 服务的 **非官方 Go 实现**，无需浏览器即可将文本转换为高质量语音音频，适合后端服务、CLI 工具及自动化场景。

> ⚠️ 本项目为 **非官方实现**，与 Microsoft 无任何隶属或授权关系。  
> Edge TTS 属于未公开接口，可能随时变更或失效。

---

## ✨ 功能特性

- 使用与 Microsoft Edge 相同的高质量神经网络语音
- 支持多语言、多音色（如中英文）
- 支持 MP3 音频输出
- API 简洁，符合 Go 语言习惯
- 支持流式写入，便于大文本处理
- 适用于后端服务、批量生成、CLI 工具

---

## 📦 安装

```bash
go get github.com/Nilin1/edgetts
