package app.ryss.gateway

import ch.qos.logback.classic.Logger
import ch.qos.logback.classic.Level
import org.slf4j.event.Level as SLF4JLevel

fun main(args: Array<String>) {
    println("Hello World!")

    val rootLogger = LoggerFactory.getLogger(Logger.ROOT_LOGGER_NAME) as Logger
    rootLogger.level = Level.valueOf(logLevel)
}

// import app.ryss.gateway.core.GatewayService
// import io.ktor.server.engine.embeddedServer
// import io.ktor.server.netty.Netty
// import io.sentry.Sentry
// import kotlinx.cli.ArgParser
// import kotlinx.cli.ArgType
// import kotlinx.cli.default
// import org.slf4j.LoggerFactory
// import ch.qos.logback.classic.Logger
// import ch.qos.logback.classic.Level
// import org.slf4j.event.Level as SLF4JLevel

// fun main(args: Array<String>) {
//     val parser = ArgParser("gateway")
//     val debug by parser.option(ArgType.Boolean, shortName = "d", fullName = "debug", description = "Enables debug mode")
//         .default(false)
//     val logLevel by parser.option(
//         ArgType.Choice(SLF4JLevel.values().map(Any::toString)),
//         shortName = "ll",
//         fullName = "log-level",
//         description = "Sets the log level"
//     ).default("INFO")
//     val enableGraphiQL by parser.option(ArgType.Boolean, shortName = "giql", fullName = "graphiql", description = "Enabled graphiql").default(false)
//     parser.parse(args)

//     val rootLogger = LoggerFactory.getLogger(Logger.ROOT_LOGGER_NAME) as Logger
//     rootLogger.level = Level.valueOf(logLevel)

//     if (!debug) {
//         Sentry.init(
//             "${System.getenv("GATEWAY_SERVICE_SENTRY_DSN")}?stacktrace" +
//                     ".app.packages=app.ryss"
//         )
//     } else {
//         Sentry.init()
//     }

//     val application = GatewayService(enableGraphiQL, debug)
//     val port = System.getenv("GATEWAY_SERVICE_PORT")?.toInt() ?: 3500
//     embeddedServer(Netty, module = application::mainModule, port = port).start(wait = true)
// }