load("@rules_jvm_external//:defs.bzl", "maven_install")

DEPENDENCY_LIST = [
        "io.ktor:ktor-server-netty:1.3.2",
        "io.ktor:ktor-metrics:1.3.2",
        "io.ktor:ktor-server-core:1.3.2",
        "io.ktor:ktor-websockets:1.3.2",
        "io.ktor:ktor-jackson:1.3.2",
        "io.ktor:ktor-server-tests:1.3.2",
    
        "com.github.excitement-engineer:ktor-graphql:1.0.0",
    
        "org.slf4j:slf4j-api:2.0.0-alpha1",
        "ch.qos.logback:logback-classic:1.3.0-alpha5",
        "io.github.microutils:kotlin-logging:1.7.9",
        "io.sentry:sentry:1.7.30",
        "io.sentry:sentry-logback:1.7.30",
        "org.jetbrains.kotlinx:kotlinx-cli:0.2.1",
    
        "org.jetbrains.kotlin:kotlin-stdlib-jdk8:1.3.70"
]

MAVEN_REPOSITORIES = [
    "http://central.maven.org/maven2/",
    "https://jcenter.bintray.com",
    "https://kotlin.bintray.com/ktor",
    "https://kotlin.bintray.com/kotlinx",
    "https://dl.bintray.com/excitement-engineer/ktor-graphql"
]

def install_maven_dependencies(fetch_sources = True, **kwargs):
    maven_install(
        artifacts = DEPENDENCY_LIST,
        fetch_sources = fetch_sources,
        repositories = MAVEN_REPOSITORIES,
        **kwargs
    )