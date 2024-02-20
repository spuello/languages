// swift-tools-version: 5.9
// The swift-tools-version declares the minimum version of Swift required to build this package.

import PackageDescription

let package = Package(
    name: "008_commandline_app",
    targets: [        // Targets are the basic building blocks of a package, defining a module or a test suite.
        .executableTarget(
            name: "008_commandline_app"),
    ]
)
