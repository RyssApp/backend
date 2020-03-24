load("@rules_java//java:defs_bzl", "java_binary", "java_library", "java_plugin")

package(default_visibility = ["//visibility:private"])

load("//kotlin:kotlin_bzl", "kt_jvm_library")

kt_jvm_library(
    name = "gateway_service_lib",
    srcs = glob(["src/**"]),
    deps = [
                "@maven//:io_ktor_ktor_server_netty",
                "@maven//:io_ktor_ktor_metrics",
                "@maven//:io_ktor_ktor_server_core",
                "@maven//:io_ktor_ktor_websockets",
                "@maven//:io_ktor_ktor_jackson",
                "@maven//:io_ktor_ktor_server_tests",
            
                "@maven//:com_github_excitement_engineer_ktor_graphql",
            
                "@maven//:org_slf4j_slf4j_api",
                "@maven//:ch_qos_logback_logback_classic",
                "@maven//:io_github_microutils_kotlin_logging",
                "@maven//:io_sentry_sentry",
                "@maven//:io_sentry_sentry_logback",
                "@maven//:org_jetbrains_kotlinx_kotlinx_cli",
            
                "@maven//:org_jetbrains_kotlin_kotlin_stdlib_jdk8"
    ],
)

java_binary(
    name = "gateway-service",
    main_class = "app.ryss.gateway.LauncherKt",
    visibility = ["//visibility:public"],
    runtime_deps = [":gateway_service_lib"],
)