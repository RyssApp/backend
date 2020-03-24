load("@io_bazel_rules_kotlin//kotlin:kotlin.bzl", "kt_jvm_library")

kt_jvm_library(
    name = "gateway_service_lib",
    srcs = glob(["**/*.kt"]),
    deps = [
        "@maven//:ch_qos_logback_logback_classic",
        "@maven//:com_github_excitement_engineer_ktor_graphql",
        "@maven//:io_github_microutils_kotlin_logging",
        "@maven//:io_ktor_ktor_jackson",
        "@maven//:io_ktor_ktor_metrics",
        "@maven//:io_ktor_ktor_server_core",
        "@maven//:io_ktor_ktor_server_netty",
        "@maven//:io_ktor_ktor_server_tests",
        "@maven//:io_ktor_ktor_websockets",
        "@maven//:io_sentry_sentry",
        "@maven//:io_sentry_sentry_logback",
        "@maven//:org_jetbrains_kotlin_kotlin_stdlib_jdk8",
        "@maven//:org_jetbrains_kotlinx_kotlinx_cli",
        "@maven//:org_slf4j_slf4j_api",
    ],
)

java_binary(
    name = "gateway-service",
    main_class = "app.ryss.gateway.LauncherKt",
    runtime_deps = [":gateway_service_lib"],
)
