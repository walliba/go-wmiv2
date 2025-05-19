# Windows Management Infrastructure (MI) for Go

Go wrapper for Windows Management Infrastructure (WMIv2), built on syscall-level bindings to `mi.dll`

[MI Reference](https://learn.microsoft.com/en-us/previous-versions/windows/desktop/wmi_v2/windows-management-infrastructure)

> [!WARNING]
> This project is under active development and is not production-ready. APIs may change without notice.
Once the first stable release is ready, [Semantic Versioning](https://semver.org/) will be used to communicate breaking changes and compatibility guarantees.

## Goal

This package aims to provide an intuitive and idiomatic Go interface for interacting with WMI and CIM on Windows. Inspired by PowerShell's `Get-CimInstance` and related cmdlets, it simplifies access to system information by abstracting away the complexities of COM or MI APIs.

While [Microsoft's official WMI package](https://github.com/microsoft/wmi) offer comprehensive support and broad compatibility, this project focuses on performance, efficiency, and minimalism. It aims to provide a streamlined alternative for scenarios where speed and lower memory overhead are critical—such as system agents, monitoring tools, or automation frameworks.

## License

This project is licensed under the [MIT License](LICENSE).
