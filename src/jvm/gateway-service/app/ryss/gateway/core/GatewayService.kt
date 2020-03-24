package app.ryss.gateway.core

import app.ryss.gateway.entites.RequestError
import app.ryss.gateway.gql.GraphQL
import app.ryss.gateway.ratelimits.MemoryRateLimiter
import app.ryss.gateway.ratelimits.RateLimiter
import com.fasterxml.jackson.databind.SerializationFeature
import io.ktor.application.Application
import io.ktor.application.ApplicationCallPipeline
import io.ktor.application.call
import io.ktor.application.install
import io.ktor.features.*
import io.ktor.http.HttpHeaders
import io.ktor.http.HttpMethod
import io.ktor.http.HttpStatusCode
import io.ktor.http.cio.websocket.pingPeriod
import io.ktor.http.cio.websocket.timeout
import io.ktor.jackson.jackson
import io.ktor.response.respond
import io.ktor.routing.get
import io.ktor.routing.routing
import io.ktor.websocket.WebSockets
import mu.KotlinLogging
import java.time.Duration

private val requestLogger = KotlinLogging.logger("Requests")

/**
 * Main class of gateway service.
 * @param enableGraphiQL whether graphiql should be enabled or not
 * @param debug whether debug mode is enabled or not
 */
class GatewayService(enableGraphiQL: Boolean, debug: Boolean) {

    /**
     * Rate-limiter used for rate-limiting requests.
     */
    val rateLimiter: RateLimiter = if (debug) MemoryRateLimiter() else null!!
    private val graphQL = GraphQL(enableGraphiQL, this)

    /**
     * Main KTor module.
     */
    fun mainModule(application: Application) {
        application.apply {
            install(Compression) {
                gzip {
                    priority = 1.0
                }
                deflate {
                    priority = 10.0
                    minimumSize(1024) // condition
                }
            }

            install(CORS) {
                method(HttpMethod.Options)
                method(HttpMethod.Put)
                method(HttpMethod.Delete)
                method(HttpMethod.Patch)
                header(HttpHeaders.Authorization)
                header("MyCustomHeader")
                allowCredentials = true
                anyHost() // @TODO: Don't do this in production if possible. Try to limit it.
            }

            install(WebSockets) {
                pingPeriod = Duration.ofSeconds(15)
                timeout = Duration.ofSeconds(15)
                maxFrameSize = Long.MAX_VALUE
                masking = false
            }

            install(ContentNegotiation) {
                jackson {
                    enable(SerializationFeature.INDENT_OUTPUT)
                }
            }

            install(StatusPages) {
                exception<Throwable> { cause ->
                    requestLogger.error(cause) { "There was an error handling an HTTP request" }
                    call.respond(
                        HttpStatusCode.InternalServerError,
                        RequestError(cause::class.simpleName, cause.message, HttpStatusCode.InternalServerError.value)
                    )
                }
            }

            routing {
                intercept(ApplicationCallPipeline.Call) {
                    rateLimiter.intercept(this)
                }
                get("/") {
                    call.respond(mapOf("stock" to "broaud"))
                }

                graphQL(this)
            }
        }
    }
}
